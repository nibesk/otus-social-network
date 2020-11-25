#!/usr/bin/env bash

go get -d -v && \
go install -v && \
go build -o ../build

../build server 2>&1 | tee -i logs/log
