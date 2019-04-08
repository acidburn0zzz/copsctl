# copsctl

## Introduction

copsctl - the Conplement AG Kubernetes developer tooling

[![Build Status](https://cpgithub.visualstudio.com/GitHubPipelines/_apis/build/status/conplementAG.copsctl?branchName=master)](https://cpgithub.visualstudio.com/GitHubPipelines/_build/latest?definitionId=9&branchName=master)

## Contribution

### Howto get the source-code

`go get github.com/conplementAG/copsctl` will install the package into your `$GOPATH`.

### Howto build the project

1. Resolve all dependencies

`cd $GOPATH/src/github.com/conplementAG/copsctl`
`dep ensure -v`

2. Build the tool

`cd $GOPATH/src/github.com/conplementAG/copsctl/cmd/copsctl`
`go build .`