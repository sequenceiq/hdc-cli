#!/usr/bin/env bash

load ../utils/commands
load ../utils/resources

DELAY=$(($SECONDS+2100))

@test "PRECONDITION: Create new OpenStack V2 credential" {
  run remove-stuck-credential "${OS_CREDENTIAL_NAME}"
  echo "$output" >&2

  OUTPUT=$(create-credential-openstack-v2 $OS_ARGS_V2 2>&1 | tail -n 2 | head -n 1)

  echo "${OUTPUT}" >&2

  [[ "${OUTPUT}" == *"credential created: ${OS_CREDENTIAL_NAME}"* ]]
  [[ "${OUTPUT}" != *"error"* ]]
}

@test "Create new OpenStack cluster" {
  run remove-stuck-cluster "${OS_CLUSTER_NAME}"

  OUTPUT=$(create-cluster --name "${OS_CLUSTER_NAME}" --cli-input-json $OS_INPUT_JSON_FILE 2>&1 | tail -n 2 | head -n 1)

  echo "${OUTPUT}" >&2

  [[ "${OUTPUT}" == *"stack created: ${OS_CLUSTER_NAME}"* ]]
  [[ "${OUTPUT}" != *"error"* ]]
}

@test "Wait for new cluster is created" {
  run wait-cluster-status $DELAY "${OS_CLUSTER_NAME}" "AVAILABLE"

  echo "$output" >&2

  [ $status -eq 0 ]
  [ "$output" = "true" ]
}

@test "Change Ambari password for previously created cluster" {
  run cluster-is-status "${OS_CLUSTER_NAME}" "AVAILABLE"
  if [[ "$output" != "true" ]]; then
    skip "Cluster is NOT created yet!"
  fi

  OUTPUT=$(change-ambari-password --name "${OS_CLUSTER_NAME}" --old-password admin --new-password 4321 --ambari-user admin 2>&1 | awk '{printf "%s",$0} END {print ""}' | grep -o '{.*}' | jq ' . |  [to_entries[].key] == ["oldPassword","password","userName"]')

  [[ "${OUTPUT}" ==  true ]]
}

@test "Previously created cluster should be listed" {
  run cluster-is-status "${OS_CLUSTER_NAME}" "AVAILABLE"
  if [[ "$output" != "true" ]]; then
    skip "Cluster is NOT created yet!"
  fi

  for OUTPUT in $(list-clusters | jq ' .[] | [to_entries[].key] == ["Name","Description","CloudPlatform","StackStatus","ClusterStatus"]');
  do
    [[ "$OUTPUT" == "true" ]]
  done
}

@test "Previously created cluster should be described" {
  run cluster-is-status "${OS_CLUSTER_NAME}" "AVAILABLE"
  if [[ "$output" != "true" ]]; then
    skip "Cluster is NOT created yet!"
  fi

  OUTPUT=$(describe-cluster --name "${OS_CLUSTER_NAME}" | jq '. "name" == "${OS_CLUSTER_NAME}"')

  [[ "$OUTPUT" == "true" ]]
}

@test "Previously created cluster can be stop" {
  run cluster-is-status "${OS_CLUSTER_NAME}" "AVAILABLE"
  if [[ "$output" != "true" ]]; then
    skip "Cluster is NOT created yet!"
  fi

  OUTPUT=$(stop-cluster --name "${OS_CLUSTER_NAME}")

  echo $OUTPUT >&2

  [[ "${OUTPUT}" != *"error"* ]]
}

@test "Previously created cluster should be stopped" {
  run wait-cluster-status $DELAY "${OS_CLUSTER_NAME}" "STOPPED"

  echo "$output" >&2

  [ $status -eq 0 ]
  [ "$output" = "true" ]
}

@test "Previously created cluster can be start" {
  run cluster-is-status "${OS_CLUSTER_NAME}" "STOPPED"
  if [[ "$output" != "true" ]]; then
    skip "Cluster has not been stopped"
  fi

  OUTPUT=$(start-cluster --name "${OS_CLUSTER_NAME}")

  echo $OUTPUT >&2

  [[ "${OUTPUT}" != *"error"* ]]
}

@test "Previously created cluster should be started" {
  run wait-cluster-status $DELAY "${OS_CLUSTER_NAME}" "AVAILABLE"

  echo "$output" >&2

  [ $status -eq 0 ]
  [ "$output" = "true" ]
}

@test "Previously created cluster can be upscale" {
  run cluster-is-status "${OS_CLUSTER_NAME}" "AVAILABLE"
  if [[ "$output" != "true" ]]; then
    skip "Cluster is NOT created yet!"
  fi

  CHECK_RESULT=$(describe-cluster --name "${OS_CLUSTER_NAME}" | jq ' ."instanceGroups" | .[] | select(."group"=="compute") | . "nodeCount"')
  INSTANCE_COUNT_DESIRED=$(($CHECK_RESULT + 1))

  echo $INSTANCE_COUNT_DESIRED > /tmp/clitestutil

  run scale-cluster --name "${OS_CLUSTER_NAME}" --group-name compute --desired-node-count $INSTANCE_COUNT_DESIRED
  [ $status -eq 0 ]
}

@test "Previously created cluster should be upscaled" {
  run cluster-is-status "${OS_CLUSTER_NAME}" "AVAILABLE"
  if [[ "$output" != "true" ]]; then
    skip "Cluster is NOT created yet!"
  fi

  INSTANCE_COUNT_DESIRED=`cat /tmp/clitestutil`
  echo $INSTANCE_COUNT_DESIRED

  run wait-stack-cluster-status $DELAY "${OS_CLUSTER_NAME}" AVAILABLE
  status_available=$status

  run node-count-are-equal $OS_CLUSTER_NAME

  INSTANCE_COUNT_CURRENT=$(describe-cluster --name "${OS_CLUSTER_NAME}" | jq ' ."instanceGroups" | .[] | select(."group"=="compute") | . "nodeCount" ')
  echo $INSTANCE_COUNT_CURRENT

  STATUS_REASON=`cat /tmp/status_reason`
  echo $STATUS_REASON
  [ $status_available = 0 ] && [ $status = 0 ] && [ $INSTANCE_COUNT_DESIRED = $INSTANCE_COUNT_CURRENT ]
}

@test "Previously created cluster can be downscale" {
  run cluster-is-status "${OS_CLUSTER_NAME}" "AVAILABLE"
  if [[ "$output" != "true" ]]; then
    skip "Cluster is NOT created yet!"
  fi

  CHECK_RESULT=$( describe-cluster --name "${OS_CLUSTER_NAME}" | jq ' ."instanceGroups" | .[] | select(."group"=="compute") | . "nodeCount" ')
  INSTANCE_COUNT_DESIRED=$(($CHECK_RESULT - 1))

  echo $INSTANCE_COUNT_DESIRED > /tmp/clitestutil

  run scale-cluster --name "${OS_CLUSTER_NAME}" --group-name compute --desired-node-count $INSTANCE_COUNT_DESIRED
  echo "$output" >&2
  [ $status -eq 0 ]
}

@test "Previously created cluster should be downscaled" {
  run cluster-is-status "${OS_CLUSTER_NAME}" "AVAILABLE"
  if [[ "$output" != "true" ]]; then
    skip "Cluster is NOT created yet!"
  fi

  INSTANCE_COUNT_DESIRED=`cat /tmp/clitestutil`
  echo $INSTANCE_COUNT_DESIRED

  run wait-stack-cluster-status $DELAY "${OS_CLUSTER_NAME}" AVAILABLE
  status_available=$status

  run node-count-are-equal "${OS_CLUSTER_NAME}"

  INSTANCE_COUNT_CURRENT=$(describe-cluster --name "${OS_CLUSTER_NAME}" | jq ' ."instanceGroups" | .[] | select(."group"=="compute") | . "nodeCount" ')
  echo $INSTANCE_COUNT_CURRENT

  STATUS_REASON=`cat /tmp/status_reason`
  echo $STATUS_REASON

  [ $status_available = 0 ] && [ $status = 0 ] && [ $INSTANCE_COUNT_DESIRED = $INSTANCE_COUNT_CURRENT ]
}

@test "Generate reinstall template" {
  OUTPUT=$(generate-reinstall-template --name "${OS_CLUSTER_NAME}" --blueprint-name "${BLUEPRINT_NAME}" | jq .blueprintName -r)

  echo "${OUTPUT}" >&2

  [[ "${OUTPUT}" == "${BLUEPRINT_NAME}" ]]
}

@test "TEARDOWN: Delete cluster then credential" {
  OUTPUT=$(delete-cluster --name "${OS_CLUSTER_NAME}" --wait)
  echo "${OUTPUT}" >&2

  run wait-cluster-delete $DELAY "${OS_CLUSTER_NAME}"
  [ $status = 0 ]

  CHECK_RESULT=$( delete-credential --name "${OS_CREDENTIAL_NAME}" )
  echo $CHECK_RESULT >&2
}

@test "TEARDOWN: Delete temporary files" {
  CHECK_RESULT=$( rm -f /tmp/status_reason )
  echo $CHECK_RESULT >&2

  CHECK_RESULT=$( rm -f /tmp/clitestutil )
  echo $CHECK_RESULT >&2
}