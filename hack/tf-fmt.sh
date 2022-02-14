#!/bin/sh

# in prow, already in container, so no 'podman run'
if [ "$IS_CONTAINER" != "" ]; then
  if [ "${#N}" -eq 0 ]; then
    set -- -list -check -diff -write=false -recursive data/data/
  fi
  set -x
  terraform fmt "${@}"
else
  ENGINE="podman"
  if [ "$(uname)" = "Darwin" ]; then
    ENGINE="docker"
  fi
  "$ENGINE" run --rm \
    --env IS_CONTAINER=TRUE \
    --volume "${PWD}:${PWD}:z" \
    --workdir "${PWD}" \
    quay.io/coreos/terraform-alpine:v0.12.0-rc1 \
    ./hack/tf-fmt.sh
fi
