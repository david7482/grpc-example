export GOPATH = $(PWD)
export PATH := $(PATH):$(GOPATH)/bin

ifdef RACE
GOOPT += -race
endif

ifdef STATIC
GOOPT += -a -installsuffix cgo
endif

.PHONY: clean glide proto main

all: glide proto main

glide:
	@if [ ! -d $(PWD)/src/vendor ]; then \
		echo "Install packages via glide install ..."; \
		cd src; glide install; \
	fi

proto:
	GOOGLE_API_PATH="src/vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis"; \
	PB_FILES=`find pb -name "*.proto"`; \
	protoc -I/usr/local/include -I. -Isrc -I$$GOOGLE_API_PATH --go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. $$PB_FILES; \
	protoc -I/usr/local/include -I. -Isrc -I$$GOOGLE_API_PATH --grpc-gateway_out=logtostderr=true:. $$PB_FILES; \
	protoc -I/usr/local/include -I. -Isrc -I$$GOOGLE_API_PATH --swagger_out=logtostderr=true:. $$PB_FILES;
	mkdir -p src/example; find pb \( -name "*.go" -or -name "*.json" \) -exec mv {} src/example \;

main:
	go build -o bin/example $(GOOPT) src/main.go

clean:
	rm -f pb/*.go pb/*.json
	rm -rf bin pkg
	rm -rf src/vendor src/example