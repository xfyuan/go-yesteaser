module := $(shell go list -m)

build: # build
	CGO_ENABLED=0 go build -mod=vendor \
	-ldflags "-X '$(module)/cmd.BUILD_TIME=`date`' \
			  -X '$(module)/cmd.GO_VERSION=`go version`'"
