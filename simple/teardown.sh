#!/bin/bash
kubectl delete deployment/client-deployment deployment/server-deployment
kubectl delete service/client-service service/server-service
