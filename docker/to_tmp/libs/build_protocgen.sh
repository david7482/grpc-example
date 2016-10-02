#!/bin/bash

checkout()
{
	return 0
}

configure()
{
	return 0
}

build()
{
	export GOPATH=`pwd`
	export PATH=$PATH:/usr/local/go/bin

	if ! go get github.com/golang/protobuf/protoc-gen-go; then
		return 1
	fi
	if ! go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway; then
		return 1
	fi
	if ! go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger; then
		return 1
	fi
}

install()
{
	if ! cp $GOPATH/bin/protoc-gen-go /usr/local/bin; then
		return 1
	fi
	if ! cp $GOPATH/bin/protoc-gen-grpc-gateway /usr/local/bin; then
		return 1
	fi
	if ! cp $GOPATH/bin/protoc-gen-swagger /usr/local/bin; then
		return 1
	fi
}

cd /tmp

checkout || exit 1
configure || exit 1
build || exit 1
install || exit 1

cd /tmp
rm -rf src
