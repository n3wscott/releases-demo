#!/usr/bin/env bash

# Copyright 2021 The CloudEvents Authors
# SPDX-License-Identifier: Apache-2.0

set -o errexit
set -o nounset
set -o pipefail

if [[ -z "$VERSION" ]]; then
    echo "Must provide VERSION in environment" 1>&2
    exit 1
fi

MODULES=(
  "subcomponent"
  "samples"
#  "protocol/amqp"
#  "protocol/stan"
#  "protocol/nats"
#  "protocol/nats_jetstream"
#  "protocol/pubsub"
#  "protocol/kafka_sarama"
#  "protocol/ws"
#  "observability/opencensus"
#  "sql"
#  "binding/format/protobuf"
)

# It is intended that this file is run locally. For a full release tag, confirm the version is correct, and then:
#   ./hack/tag-tag-release.sh --tag --push

CREATE_TAGS=0 # default is a dry run
PUSH_TAGS=0   # Assumes `upstream` is the remote name for sdk-go.
SAMPLES=0

# Pick one:
REMOTE="origin"   # if checked out directly
#REMOTE="upstream" # if checked out with a fork

REPOINT=(
  "github.com/n3wscott/releases-demo/v2"
#  "github.com/cloudevents/sdk-go/v2"
)
REPOINT_ALL=(
#  "github.com/cloudevents/sdk-go/protocol/amqp/v2"
#  "github.com/cloudevents/sdk-go/protocol/stan/v2"
#  "github.com/cloudevents/sdk-go/protocol/nats/v2"
#  "github.com/cloudevents/sdk-go/protocol/nats_jetstream/v2"
#  "github.com/cloudevents/sdk-go/protocol/pubsub/v2"
#  "github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2"
#  "github.com/cloudevents/sdk-go/protocol/ws/v2"
#  "github.com/cloudevents/sdk-go/observability/opencensus/v2"
#  "github.com/cloudevents/sdk-go/observability/opentelemetry/v2"
#  "github.com/cloudevents/sdk-go/sql/v2"
#  "github.com/cloudevents/sdk-go/binding/format/protobuf/v2"
#  "github.com/cloudevents/sdk-go/v2"                       # NOTE: this needs to be last.
  "github.com/n3wscott/releases-demo/subcomponent/v2"
  "github.com/n3wscott/releases-demo/v2"
)


# Loop through arguments and process them
for arg in "$@"
do
    case $arg in
        -t|--tag)
        CREATE_TAGS=1
        shift
        ;;
        -p|--push)
        PUSH_TAGS=1
        shift
        ;;
        # --samples is used to repoint the dep used for samples to the newly released submodules
        --samples)
        SAMPLES=1
        REPOINT=REPOINT_ALL
        shift
        ;;
    esac
done

echo --- All Modules ---
for gomodule in $(find . | grep "go\.mod" | awk '{gsub(/\/go.mod/,""); print $0}' | grep -v "./v2/test" | grep -v "./test")
do
  echo "  $gomodule"

  if [[ $gomodule == "./v2" ]]
  then
    echo "    skipping main module"
    continue
  fi

  if [ "$SAMPLES" -eq "1" ]; then
    if [[ $gomodule != "./samples"* ]]
    then
      echo "    skipping non-sample module"
      continue
    fi
  else
    if [[ $gomodule == "./samples"* ]]
    then
      echo "    skipping sample module"
      continue
    fi
  fi

  pushd $gomodule > /dev/null
  for repoint in "${REPOINT[@]}"; do
    if grep -Fq "$repoint" go.mod
    then
      tag="$VERSION"
      echo "    repointing dep on $repoint@$tag"
      go get -d $repoint@$tag
      go mod edit -dropreplace $repoint
    fi
    go mod tidy
  done
  popd > /dev/null

done

if [ "$SAMPLES" -eq "1" ]; then
  echo "Done."
  exit 0
fi

echo --- Tagging ---

for i in "${MODULES[@]}"; do
    tag=""
    if [ "$i" = "" ]; then
      tag="$VERSION"
    else
      tag="$i/$VERSION"
    fi
    if [ "$CREATE_TAGS" -eq "1" ]; then
      echo "  tagging with $tag"
      git tag $tag
    fi
    if [ "$PUSH_TAGS" -eq "1" ]; then
      echo "  pushing $tag to $REMOTE"
      git push $REMOTE $tag
    fi
done