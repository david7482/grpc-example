#!/bin/bash

checkout()
{
	if ! wget -nd -nH 'https://github.com/google/protobuf/releases/download/v3.0.0/protobuf-cpp-3.0.0.tar.gz'; then
		return 1
	fi

	tar zxf protobuf-cpp-3.0.0.tar.gz
	cd protobuf-3.0.0
}

configure()
{
	if ! ./configure --prefix="/usr/local"; then
		return 1
	fi
}

build()
{
	if ! make -j4; then
		return 1
	fi
}

install()
{
	if ! make install; then
		return 1
	fi
}

cd /tmp

checkout || exit 1
configure || exit 1
build || exit 1
install || exit 1

cd /tmp
rm -rf protobuf-3.0.0
