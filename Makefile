.PHONY: test

test:
	@echo "==> Benchmark"
	@go test -bench=. -benchmem -run=none
