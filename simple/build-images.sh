#!/bin/bash

# Using docker hub for now, the local thing doesnt work for me yet

docker build . -f Dockerfile.client -t oneeyedsunday/k8_playground:client

docker build . -f Dockerfile.server -t oneeyedsunday/k8_playground:server

docker push oneeyedsunday/k8_playground:client && docker push oneeyedsunday/k8_playground:server

