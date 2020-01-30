#!/usr/bin/env sh
helm upgrade draco ./deploy/helm/draco \
  --values=./deploy/helm/values/prod/values.yaml \
  --set=spec.temdraco.metadata.labels.date=`date +'%s'` \
  --namespace=default
