#!/usr/bin/env sh
helm install draco ./deploy/helm/draco \
  --values ./deploy/helm/values/prod/values.yaml \
  --namespace default
