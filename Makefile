.PHONY: test-slow test lint fmt vet bench bench-profile bench-long bench-comparative pre-commit

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

bench-long:
	cd benchmarks && go test -bench=BenchmarkAllScenarios -benchmem  -count 10 > results.txt

bench-comparative:
	cd benchmarks && go test -bench=BenchmarkComparativeAll -benchmem -count 10 > results_comparative.txt


pre-commit:
	pre-commit run --all-files

cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
