.PHONY: build
buid: 
	go build -v ./cmd/apiserver
.DEFAULT_GOAL: build