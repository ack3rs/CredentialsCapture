#!/bin/sh

GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models
swagger serve -F=swagger swagger.yaml


