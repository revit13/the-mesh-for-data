// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package secretprovider

import (
	"flag"
	"os"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kubectl/pkg/scheme"
)

const (
	// tokenReviewGroup is the group with which the TokenReview resource is associated
	tokenReviewGroup string = "authentication.k8s.io"

	// tokenReviewVersion version of the TokenReview resource's implementation
	tokenReviewVersion string = "v1"
)

// getClientConfig returns config for accessing kubernetes
func getClientConfig(kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	return rest.InClusterConfig()
}

// newForConfig configures the client
func newForConfig(cfg *rest.Config) (*clientset.Clientset, error) {
	config := *cfg
	config.ContentConfig.GroupVersion = &schema.GroupVersion{Group: tokenReviewGroup, Version: tokenReviewVersion}
	config.APIPath = "/apis"
	config.ContentType = runtime.ContentTypeJSON
	config.UserAgent = rest.DefaultKubernetesUserAgent()
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	client, err := clientset.NewForConfig(&config)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func K8sAuthClientInit() (*clientset.Clientset, error) {
	// For local testing set KUBECONFIG to $HOME/.kube/config
	// It is unset for deployment

	kubeconfigArg := ""
	if kubeconfigpath := os.Getenv("KUBECONFIG"); kubeconfigpath != "" {
		kubeconf := flag.String("kubeconf", kubeconfigpath, "Path to a kube config. Only required if out-of-cluster.")
		flag.Parse()
		kubeconfigArg = *kubeconf
	}

	config, err := getClientConfig(kubeconfigArg)

	if err != nil {
		return nil, err
	}

	cs, err := newForConfig(config)
	return cs, err
}
