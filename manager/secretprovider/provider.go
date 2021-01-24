// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package secretprovider

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/go-logr/logr"
	"github.com/ibm/the-mesh-for-data/manager/controllers/utils"
	clientset "k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type secretProvider struct {
	client.Client
	clientset.Clientset // For TokenReview requests
	Log                 logr.Logger
}

var provider secretProvider

// Routes are the REST endpoints for getting a secret
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		middleware.SetHeader("Access-Control-Allow-Origin", "*"), // Allow any client to access these APIs
		middleware.SetHeader("Access-Control-Allow-Methods", "GET, OPTIONS"),
		middleware.SetHeader("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"),
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		middleware.Logger,          // Log API request calls
		middleware.RedirectSlashes, // Redirect slashes to no slash URL versions
		middleware.Recoverer,       // Recover from panics without crashing server
	)

	router.Route("/v1/m4d", func(r chi.Router) {
		r.Mount("/secret", handleSecret())
	})

	return router
}

func SetupSecretProvider(client client.Client) {
	setupLog := ctrl.Log.WithName("setup SetupSecretProvider")

	cs, err := K8sAuthClientInit()
	if err != nil || cs == nil {
		setupLog.Error(err, "Failed getting kubernetes client for authentication!")
		os.Exit(1)
	}

	provider = secretProvider{client, *cs, ctrl.Log.WithName("Secretprovider")}

	// Setup secret store
	err = SecretStoreSetUp()
	if err != nil {
		setupLog.Error(err, "Failed setup secret store!")
		os.Exit(1)
	}
}

func FireUpSecretProviderServer() {
	// REST APIs provided
	router := Routes()

	log := provider.Log

	// Print out Secret-provider APIs
	log.Info(fmt.Sprintf("Server listening on port %s", utils.GetSecretProviderPort()))
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Info(fmt.Sprintf("%s %s\n", method, route)) // Walk and print out all routes
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Error(err, fmt.Sprintf("error when printing rounts: %s", err.Error()))
		os.Exit(1)
	}

	// TODO: Should be https server
	//err := http.ListenAndServe(fmt.Sprintf(":%s", utils.GetSecretProviderPort()), router)
	err := http.ListenAndServe(fmt.Sprintf(":8083"), router)
	log.Error(err, "SecretProvider Server exited")
	os.Exit(1)
}
