export GOPATH = $(PWD)
export PATH := $(PATH):$(GOPATH)/bin:/usr/local/go/bin

ifdef RACE
GOOPT += -race
endif

ifdef STATIC
export CGO_ENABLED = 0
GOOPT += -a -installsuffix cgo
endif

.PHONY: clean glide proto example1

all: glide proto example1

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
	mkdir -p src/pb; find pb \( -name "*.go" -or -name "*.json" \) -exec mv {} src/pb \;

example1:
	go build -o bin/example1 $(GOOPT) src/example1/main.go

clean:
	rm -f pb/*.go pb/*.json
	rm -rf bin pkg
	rm -rf src/vendor src/example