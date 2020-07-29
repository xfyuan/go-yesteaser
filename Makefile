module := $(shell go list -m)

build: # build
	CGO_ENABLED=0 go build -v \
	-ldflags "-X '$(module)/cmd.BUILD_TIME=`date`' \
			  -X '$(module)/cmd.GO_VERSION=`go version`'"

lint:
	CGO_ENABLED=0 golangci-lint run

test:
	YESTEA_ENV=test CGO_ENABLED=0 ginkgo ./...

SpecSteps = lint test
spec: $(SpecSteps)
