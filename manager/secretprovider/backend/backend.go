package backend

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	smv1alpha1 "github.com/ibm/the-mesh-for-data/manager/secretprovider/backend/apis/secretmanager/v1alpha1"
	"github.com/ibm/the-mesh-for-data/manager/secretprovider/backend/internal/store"
	storebase "github.com/ibm/the-mesh-for-data/manager/secretprovider/backend/internal/store/base"
	"github.com/ibm/the-mesh-for-data/manager/secretprovider/backend/util/merge"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ExternalSecretReconciler reconciles a ExternalSecret object
type Backend struct {
	client.Client
	Log logr.Logger
	//Scheme *runtime.Scheme
	//Clock  clock.Clock

	storeFactory store.Factory
	Reader       client.Reader
}

const secretStoreDummyName = "dummyStore"
const secretStoreDummyNamespace = "default"

var storeClient store.Client

func (b *Backend) GetBackendClient(store smv1alpha1.GenericStore, log logr.Logger,
	kubeClient client.Client, kubeReader client.Reader, namespace string) error {

	b.storeFactory = &storebase.Default{}

	ctx := context.Background()
	var err error
	storeClient, err = b.storeFactory.New(ctx, log, store, kubeClient, kubeReader, namespace)

	if err != nil {
		log.Error(err, fmt.Sprintf("Error in GetBackendClient %s", err.Error()))
		return err
	}
	return nil
}

func (b *Backend) getSecret(ctx context.Context, storeClient store.Client, extSecret *smv1alpha1.ExternalSecret, log logr.Logger) (map[string][]byte, error) {
	secretDataMap := make(map[string][]byte)
	for _, remoteRef := range extSecret.Spec.DataFrom {
		secretMap, err := storeClient.GetSecretMap(ctx, remoteRef)
		if err != nil {
			log.Error(err, fmt.Sprintf("Error in getSecret %s", err.Error()))
			return nil, err
		}
		secretDataMap = merge.Merge(secretDataMap, secretMap)
	}

	for _, secretRef := range extSecret.Spec.Data {
		secretData, err := storeClient.GetSecret(ctx, secretRef.RemoteRef)
		if err != nil {
			log.Error(err, fmt.Sprintf("Error in getSecret %s", err.Error()))
			return nil, err
		}
		secretDataMap[secretRef.SecretKey] = secretData
	}

	return secretDataMap, nil
}

func (b *Backend) GetSecretFromSecretStore(credentials_location string, extSecret smv1alpha1.ExternalSecret, log logr.Logger) (map[string][]string, error) {
	ctx := context.Background()
	return b.getSecret(ctx, storeClient, &extSecret, log)

}
