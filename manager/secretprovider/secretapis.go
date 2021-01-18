// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package secretprovider

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

/*func GetSecretFromBackend(w http.ResponseWriter, r *http.Request) {
	provider.Log.Info(fmt.Sprintf("In GetSecretFromBackend"))
	datasetCreds := `{"access_key": "x", "secret_key": "x"}"`
	render.JSON(w, r, datasetCreds) // Return the M4DApplication as json
}*/

// requestSecretInfo contains the information needed for a module to retrieve a token.
type requestSecretInfo struct {
	SecretName string `json:"secret_name"`
	Token      string `json:"token"` //service account token
}

func GetSecretFromBackend(w http.ResponseWriter, r *http.Request) {
	log := provider.Log

	var requestInfo requestSecretInfo
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	secret := r.URL.Query().Get("secret")

	token := r.Header.Get("Authorization")

	if secret == "" {
		err := errors.New("Url Param 'secret' is missing")

		suberr := render.Render(w, r, ErrInvalidRequest(err))
		if suberr != nil {
			log.Error(err, fmt.Sprintf(err.Error()))
		}
		return
	}

	fmt.Println("TOKENxxx")
	fmt.Println(requestInfo.Token)
	log.Info(fmt.Sprintf(fmt.Sprintf("In GetSecretFromBackend token %s\n secret_name: %s\n", token, secret)))
	/*
		genTokenReview := &auth.TokenReview{
			TypeMeta: metav1.TypeMeta{
				Kind:       "TokenReview",
				APIVersion: "authentication.k8s.io/v1",
			},
			ObjectMeta: metav1.ObjectMeta{Name: "TokenReview1", Namespace: "default"},
			Spec:       auth.TokenReviewSpec{Token: string(token)},
		}
		trev, err := authClient.AuthenticationV1().TokenReviews().Create(ctx, genTokenReview, metav1.CreateOptions{})
		if err != nil {
			log.Info(fmt.Sprintf(fmt.Sprintf("TokenReview Authentication failed for token %s", token)))
			render.Status(r, http.StatusForbidden)
			return
		}
		log.Info(fmt.Sprintf("Username from TokenReviews: %s", trev.Status.User.Username))
		datasetCreds := `{"access_key": "x", "secret_key": "x"}"`
		render.Status(r, http.StatusOK)
		render.JSON(w, r, datasetCreds)*/
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
