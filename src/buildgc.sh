#!/bin/bash
# assumes cmd/dist/dist is built with bootstrapping go compiler
# assumes following environment:
#GOROOT_BOOTSTRAP=/usr/local/go/
#GOROOT=/Users/jaymz/golang/
#GOPATH=/Users/jaymz/go
# assumes ~/golang/src/ is pwd.
rm -f cmd/internal/gc/y.go cmd/internal/gc/y.output && pushd cmd/internal/gc/ && ../../yacc/yacc go.y ; popd ; cmd/dist/dist bootstrap
