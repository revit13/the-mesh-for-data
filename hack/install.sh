#!/usr/bin/env bash

set -e

: ${KUBE_NAMESPACE:=m4d-system}
: ${WITHOUT_VAULT:=false}
: ${WITHOUT_EGERIA:=false}
: ${WITHOUT_OPA:=false}

source secret-provider/deploy/vault-util.sh

kubectl create ns $KUBE_NAMESPACE || true
kubectl config set-context --current --namespace=$KUBE_NAMESPACE

kubectl apply -f manager/config/prod/deployment_configmap.yaml
make cluster-prepare
kubectl create secret generic user-vault-unseal-keys --from-literal=vault-root=$(kubectl get secrets vault-unseal-keys -o jsonpath={.data.vault-root} | base64 --decode) || true
# Install third party components
$WITHOUT_EGERIA || make -C third_party/egeria deploy
$WITHOUT_OPA || make -C third_party/opa deploy

# Waiting for the vault deployment to become ready
make -C third_party/vault deploy-wait

# Perform a port-forward to communicate with Vault
port_forward

make -C secret-provider configure-vault

# Install the manager, the connectors and the secret-provider
make deploy

make install

# Kill the port-forward
kill -9 %%
