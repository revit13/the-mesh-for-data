// Copyright 2021 IBM Corp.
// SPDX-License-Identifier: Apache-2.0
package utils

import (
	"fmt"
	"net/url"
)

// The Vault plugin to use to retrive the dataset credential
const vaultPluginPath = "kubernetes-secrets-reader"

// VaultSecretPath returns the path to Vault secret that holds the dataset credential.
// The path contains the kubernetes-secrets-reader plugin name and the secret name and namespace as parameters,
// for example, for secret name my-secret and namespace default it will be of the form:
// "/v1/kubernetes-secrets-reader/my-secret?namespace=default"
func VaultSecretPath(secretName string, secretNamespace string) string {
	pluginPath := "/v1/" + vaultPluginPath + "/"
	// Construct the path to the secret in Vault that holds the dataset credentials
	secretParam := fmt.Sprintf("%s?namespace=%s", secretName, secretNamespace)
	return fmt.Sprintf("%s%s", pluginPath, url.PathEscape(secretParam))
}
