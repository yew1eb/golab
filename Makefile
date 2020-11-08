export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE=on
LDFLAGS := -s -w

all: fmt build

build: 
	go build

fmt:
	go fmt ./...

test: gotest

gotest:
	sh test.sh
	#go test -v --cover $(go list ./... | grep -v /third-party/)
	#go test -v --cover ./cmd/...
	#go test -v --cover ./client/...
	#go test -v --cover ./server/...
	#go test -v --cover ./pkg/...

ci:
	#go test -count=1 -p=1 -v ./tests/...

e2e:
	#./hack/run-e2e.sh

alltest: gotest ci e2e
	
clean:
	rm -f ./target