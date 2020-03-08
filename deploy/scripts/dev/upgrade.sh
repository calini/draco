#!/usr/bin/env sh
helm upgrade draco ./deploy/helm/draco \
  --values ./deploy/helm/values/dev/values.yaml \
  --set timestamp="$(date)" \
  --namespace dev