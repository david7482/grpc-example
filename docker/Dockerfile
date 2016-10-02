FROM ubuntu:16.04

MAINTAINER David Chou david74.chou@gmail.com

COPY to_tmp /tmp

RUN apt-get update
RUN apt-get -y install \
	    wget \
	    g++ \
	    cmake \
	    git

# Build protobuf
RUN cd /tmp/libs; ./build_protobuf.sh

# Install the latest Go binary
RUN cd /tmp; \
    wget -nd -nH 'https://storage.googleapis.com/golang/go1.7.1.linux-amd64.tar.gz'; \
    tar zxf go1.7.1.linux-amd64.tar.gz -C /usr/local;

# Install after installing protoc go plugin
RUN cd /tmp/libs; ./build_protocgen.sh

# Install Glide, the package manager for Go
RUN cd /tmp; \
    wget -nd -nH 'https://github.com/Masterminds/glide/releases/download/v0.11.1/glide-v0.11.1-linux-amd64.tar.gz'; \
    tar zxf glide-v0.11.1-linux-amd64.tar.gz; \
    mv linux-amd64/glide /usr/local/bin

# Append to bashrc
RUN cat /tmp/bash/bashrc.append >> /root/.bashrc

# Cleanup temp files
RUN rm -rf /tmp/*

# Update ld.so.cache
RUN ldconfig