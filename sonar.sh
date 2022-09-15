#!/bin/bash

#golangci-lint run --out-format checkstyle ./... >report/report.xml
go test -v ./... -coverprofile=report/coverage.out -json >report/ut.out
#golint ./... > report/golint.out


sonar-scanner \
  -Dsonar.projectKey=go-common-module \
  -Dsonar.sources=. \
  -Dsonar.host.url=http://127.0.0.1:9000/sonar \
  -Dsonar.login=5c867d63901a6c8b4c6d9c65d23e0d9d35357102
