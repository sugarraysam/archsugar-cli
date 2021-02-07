#!/bin/bash

set -e

SONAR_PROJECT=archsugar
SONAR_TOKEN=35a337d1bc15759ef33ae21802554a3351f22617
SONAR_LOCAL_URL="http://localhost:9000"
SONAR_HEALTHZ="${SONAR_LOCAL_URL}/api/system/health"
SONAR_CREDS="admin:admin"

DOCKER_BRIDGE="${SONAR_PROJECT}"
DOCKER_VOLUME="${SONAR_PROJECT}_sonarqube_data"

###
### Functions
###
function createDockerBridge() {
    if [ -z "$(docker network ls --filter name=${DOCKER_BRIDGE} --format '{{.ID}}')" ]; then
        docker network create ${DOCKER_BRIDGE}
    fi
}

function startServer() {
    if [ -z "$(docker ps -a --filter name=sonarqube --format '{{.ID}}')" ]; then
        docker run -d --name sonarqube --network ${DOCKER_BRIDGE} \
            -v "${DOCKER_VOLUME}:/opt/sonarqube/data" -p 9000:9000 sonarqube
    else
        docker start sonarqube
    fi
}

function waitOnServerStart() {
    spin='-\|/'
    i=0
    echo "Waiting on ${SONAR_HEALTHZ}..."
    until [[ "$(curl -s -u ${SONAR_CREDS} ${SONAR_HEALTHZ} | jq '.health')" =~ GREEN ]]; do
        i=$(((i + 1) % 4))
        printf "\r%s" "${spin:$i:1}"
        sleep .5
    done
    echo "Sonarqube server is ready!"
}

function sonarScan() {
    docker run --rm --network ${DOCKER_BRIDGE} --name sonarcli \
        -e SONAR_HOST_URL=http://sonarqube:9000 -v "${PWD}:/usr/src" \
        sonarsource/sonar-scanner-cli \
        -Dsonar.host.url=http://sonarqube:9000 \
        -Dsonar.login="${SONAR_TOKEN}"
}

###
### Execution
###
createDockerBridge
startServer
waitOnServerStart
sonarScan

echo "[SUCCESS] View analysis results ==> ${SONAR_LOCAL_URL}, default login creds ${SONAR_CREDS}"
