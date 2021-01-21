// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package secretprovider

import (
	"context"
	b64 "encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	auth "k8s.io/api/authentication/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/rand"
)

const tokenReviewNamespace = "default"
const serviceAccountPrefix = "system:serviceaccount:"

/*func GetSecretFromBackend(w http.ResponseWriter, r *http.Request) {
	log := provider.Log
	ctx := context.Background()

	log.Info(fmt.Sprintf("REVIT   GetSecretFromBackend Failed: %s", module))
	datasetCreds := `{"access_key": "x", "secret_key": "x"}"`
	render.JSON(w, r, datasetCreds) // Return the M4DApplication as json
}*/

// Check if the sevice account token is authorized to access the secret.
func isAuthorized(token string) error {
	log := provider.Log
	ctx := context.Background()
	random := rand.String(5)
	randomTokenReviewName := "m4d-token-review-" + random

	log.Info(fmt.Sprintf(fmt.Sprintf("In isAuthorized created tokenReview %s\n", randomTokenReviewName)))

	authClient := provider.Clientset
	genTokenReview := &auth.TokenReview{
		TypeMeta: metav1.TypeMeta{
			Kind:       "TokenReview",
			APIVersion: "authentication.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{Name: randomTokenReviewName, Namespace: tokenReviewNamespace},
		Spec:       auth.TokenReviewSpec{Token: token},
	}
	trev, err := authClient.AuthenticationV1().TokenReviews().Create(ctx, genTokenReview, metav1.CreateOptions{})

	if err != nil {
		return err
	}
	// Check the status of the TokenReviews
	if !trev.Status.Authenticated {
		err := errors.New("Authorization Token is not Authenticated by TokenReview")
		return err
	}

	module := trev.Status.User.Username
	// Check that the service account name is from the expected namespace
	// FIXME: Fix namespace
	if !strings.HasPrefix(module, serviceAccountPrefix+"m4d-system") {
		return err
	}
	return nil
}

func isValidRequest(secret string, tokenBase64Encoded string) error {
	if secret == "" {
		err := errors.New("Url Param 'secret' is missing")
		return err
	}

	if tokenBase64Encoded == "" {
		err := errors.New("Authorization Token is missing from header")
		return err
	}
	return nil
}

func GetSecretFromBackend(w http.ResponseWriter, r *http.Request) {
	log := provider.Log

	secret := r.URL.Query().Get("secret")
	tokenBase64Encoded := r.Header.Get("Authorization")

	err := isValidRequest(secret, tokenBase64Encoded)
	if err != nil {
		suberr := render.Render(w, r, ErrInvalidRequest(err))
		if suberr != nil {
			log.Error(err, fmt.Sprintf("GetSecretFromBackend Failed: %s", err.Error()))
		}
		return
	}
	token, _ := b64.URLEncoding.DecodeString(tokenBase64Encoded)

	err = isAuthorized(string(token))
	if err != nil {
		suberr := render.Render(w, r, ErrForbiddenRequest(err))
		if suberr != nil {
			log.Error(err, fmt.Sprintf("GetSecretFromBackend Failed: %s", err.Error()))
		}
		return
	}

	datasetCreds := `{"access_key": "x", "secret_key": "x"}"`
	render.Status(r, http.StatusOK)
	render.JSON(w, r, datasetCreds)
}

// handleSecret is a list of the REST APIs supported by the secret provider
func handleSecret() *chi.Mux {

	router := chi.NewRouter()
	router.Get("/", GetSecretFromBackend) // Returns the M4DApplication CRD including its status
	router.Options("/*", getSecretOptions)
	return router
}

// getSecretOptions returns an OK status, but more importantly its header is set to indicate
// that future POST, PUT and DELETE calls are allowed as per the header values set when the router was initiated in main.go
func getSecretOptions(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
}
