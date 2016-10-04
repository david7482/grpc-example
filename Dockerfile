FROM ubuntu:16.04

MAINTAINER David Chou david74.chou@gmail.com

RUN apt-get update
RUN apt-get -y install \
	    wget \
	    g++ \
	    cmake \
	    git

# Build protobuf
RUN cd /tmp; \
	wget -nd -nH 'https://github.com/google/protobuf/releases/download/v3.0.0/protobuf-cpp-3.0.0.tar.gz'; \
	tar zxf protobuf-cpp-3.0.0.tar.gz; cd protobuf-3.0.0; \
	./configure --prefix="/usr/local"; make -j4; make install;


# Install the latest Go binary
RUN cd /tmp; \
    wget -nd -nH 'https://storage.googleapis.com/golang/go1.7.1.linux-amd64.tar.gz'; \
    tar zxf go1.7.1.linux-amd64.tar.gz -C /usr/local;

# Install after installing protoc go plugin
RUN cd /tmp; \
	export GOPATH=`pwd`; export PATH=$PATH:/usr/local/go/bin; \
	go get github.com/golang/protobuf/protoc-gen-go; cp $GOPATH/bin/protoc-gen-go /usr/local/bin; \
	go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway; cp $GOPATH/bin/protoc-gen-grpc-gateway /usr/local/bin; \
	go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger; cp $GOPATH/bin/protoc-gen-swagger /usr/local/bin;

# Install Glide, the package manager for Go
RUN cd /tmp; \
    wget -nd -nH 'https://github.com/Masterminds/glide/releases/download/v0.11.1/glide-v0.11.1-linux-amd64.tar.gz'; \
    tar zxf glide-v0.11.1-linux-amd64.tar.gz; \
    mv linux-amd64/glide /usr/local/bin;

# Add /usr/local/go/bin to .bashrc
RUN echo "export PATH=\$PATH:/usr/local/go/bin" >> /root/.bashrc

# Cleanup temp files
RUN rm -rf /tmp/*

# Update ld.so.cache
RUN ldconfig