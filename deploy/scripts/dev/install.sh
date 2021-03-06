#!/usr/bin/env sh
helm install draco ./deploy/helm/draco \
  --values ./deploy/helm/values/dev/values.yaml \
  --set timestamp="$(date)" \
  --namespace dev
