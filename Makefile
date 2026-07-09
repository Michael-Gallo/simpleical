.PHONY: test-slow test lint fmt fmt-check vet bench bench-profile bench-long bench-comparative check install-hooks

test-slow:
	go test ./... --race --count 1

test:
	go test --count 1 ./...

lint:
	golangci-lint run

fmt:
	golangci-lint fmt

fmt-check:
	golangci-lint fmt --diff

vet:
	go vet ./...

bench:
	cd benchmarks && go test -bench=BenchmarkAllScenarios -benchmem

bench-profile:
	cd benchmarks && go test -bench=BenchmarkAllScenarios -benchmem -cpuprofile=cpu.prof -memprofile=mem.prof

bench-long:
	cd benchmarks && go test -bench=BenchmarkAllScenarios -benchmem  -count 10 > results.txt

bench-comparative:
	cd benchmarks && go test -bench=BenchmarkComparativeAll -benchmem -count 10 > results_comparative.txt

check: fmt-check vet lint test-slow

install-hooks:
	git config core.hooksPath .githooks

cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
