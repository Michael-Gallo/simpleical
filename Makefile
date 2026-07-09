.PHONY: test-slow test lint fmt vet bench bench-profile bench-long bench-comparative pre-commit install-hooks

test-slow:
	go test ./... --race --count 1

test:
	go test --count 1 ./...

lint:
	golangci-lint run

fmt:
	go fmt ./...

vet:
	go vet ./...

bench:
	cd benchmarks && go test -bench=BenchmarkAllScenarios -benchmem

bench-profile:
	cd benchmarks && go test -bench=BenchmarkAllScenarios -benchmem -cpuprofile=cpu.prof -memprofile=mem.prof

# Writes host-specific raw output under benchmarks/artifacts/ (gitignored).
bench-long:
	mkdir -p benchmarks/artifacts
	cd benchmarks && go test -bench=BenchmarkAllScenarios -benchmem -count=10 > artifacts/results.txt

bench-comparative:
	mkdir -p benchmarks/artifacts
	cd benchmarks && go test -bench=BenchmarkComparativeAll -benchmem -count=10 > artifacts/results_comparative.txt


pre-commit:
	PRE_COMMIT_HOME="$${PRE_COMMIT_HOME:-$(CURDIR)/.precommit-cache}" XDG_CACHE_HOME="$${XDG_CACHE_HOME:-$(CURDIR)/.cache}" pre-commit run --all-files

install-hooks:
	chmod +x .githooks/pre-commit
	git config core.hooksPath .githooks

cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
