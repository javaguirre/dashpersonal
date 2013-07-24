#!/bin/bash
export GOROOT=/usr/lib/go
export GOPATH=/var/lib/jenkins/go
export PATH=$PATH:$GOPATH/bin
revel test dashpersonal dev
