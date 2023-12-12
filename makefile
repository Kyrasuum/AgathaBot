srcs = $(wildcard pkg/**/*.go) $(wildcard internal/**/*) cmd/main.go

%:
	@cd .. && make --no-print-directory $@

.PHONY: build
build: $(srcs)
	@go build cmd/main.go

.PHONY: test
test: $(srcs)
	@go test ./... --cover
