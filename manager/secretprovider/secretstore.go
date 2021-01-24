package secretprovider

import (
	"fmt"

	"github.com/go-logr/logr"
	"github.com/ibm/the-mesh-for-data/manager/secretprovider/backend"
	smmeta "github.com/ibm/the-mesh-for-data/manager/secretprovider/backend/apis/meta/v1"
	smv1alpha1 "github.com/ibm/the-mesh-for-data/manager/secretprovider/backend/apis/secretmanager/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const secretStoreDummyName = "dummyStore"
const secretStoreDummyNamespace = "m4d-system"
const externalSecretDummyName = "dummySecret"

var secretStore backend.Backend

func SecretStoreSetUp() error {
	log := provider.Log
	//ctx := context.Background()

	//FIXME: take params from config map
	s := &smv1alpha1.ClusterSecretStore{
		TypeMeta: metav1.TypeMeta{
			Kind:       "SecretStore",
			APIVersion: "secret-manager.itscontained.io/v1alpha1",
		},
		ObjectMeta: metav1.ObjectMeta{Name: secretStoreDummyName, Namespace: secretStoreDummyNamespace},
		Spec: smv1alpha1.SecretStoreSpec{
			Vault: &smv1alpha1.VaultStore{
				Server: "http://vault.m4d-system:8200",
				Path:   "external",
				Auth: smv1alpha1.VaultAuth{
					TokenSecretRef: &smmeta.SecretKeySelector{
						Key: "token",
						LocalObjectReference: smmeta.LocalObjectReference{
							Name: "my-vault-token",
						},
					},
				},
			},
		},
	}
	var err error = nil
	err = secretStore.GetBackendClient(s, provider.Log, provider.Client, nil, secretStoreDummyNamespace)
	log.Info(fmt.Sprintf("SecretStoreSetUp!!!\n")) // Walk and print out all routes

	return err
}

func GetSecret(credentials_location string, log logr.Logger) (string, error) {
	var remoreRef []smv1alpha1.RemoteReference
	remoreRef[0] = smv1alpha1.RemoteReference{Name: "{\"ServerName\":\"cocoMDS3\",\"AssetGuid\":\"f679bbf8-0b28-4f83-a146-788ee6e8b4b7\"}"}

	extSecret := smv1alpha1.ExternalSecret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ExternalSecret",
			APIVersion: "secret-manager.itscontained.io/v1alpha1",
		},
		ObjectMeta: metav1.ObjectMeta{Name: externalSecretDummyName, Namespace: secretStoreDummyNamespace},
		Spec: smv1alpha1.ExternalSecretSpec{
			StoreRef: smv1alpha1.ObjectReference{Name: secretStoreDummyName},
			DataFrom: remoreRef,
		},
	}
	//log := provider.Log
	return credentials, err := secretStore.GetSecretFromSecretStore(credentials_location, extSecret, log)
}
