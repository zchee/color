SHELL = /usr/bin/env bash

GO_TEST ?= go test
GO_TAGS ?= benchmark
GO_BENCH_FLAGS ?= -benchtime=2s
GO_BENCH_FUNCS ?= .
GO_BENCH_CPUS ?= 1,4,12
GO_BENCH_COUNT ?= 10
GO_BENCH_OUTPUT ?= ../new.txt
GO_BENCH_WORKING_DIRECTORY ?= ./benchmarks

define target
@printf "\\x1b[1;32m$(patsubst ,$@,$(1))\\x1b[0m\\n"
endef

.PHONY: test
test:
	@${GO_TEST} -v -race ./...

.PHONY: bench/base
bench/base:
	$(call target,${TARGET})
	@pushd ${GO_BENCH_WORKING_DIRECTORY} > /dev/null 2>&1; go test -v -tags=${GO_TAGS} -cpu ${GO_BENCH_CPUS} -count ${GO_BENCH_COUNT} -run='^$$' -bench=${GO_BENCH_FUNCS} ${GO_BENCH_FLAGS} . | tee ${GO_BENCH_OUTPUT}

.PHONY: bench
bench: TARGET=bench
bench: bench/base

.PHONY: bench/fatih
bench/fatih: GO_TAGS=benchmark_fatih
bench/fatih: GO_BENCH_OUTPUT=../old.txt
bench/fatih: TARGET=bench/fatih
bench/fatih: bench/base

.PHONY: benchstat
benchstat: clean bench
	@benchstat benchmarks/old.golden.txt $(shell echo ${GO_BENCH_OUTPUT} | cut -d/ -f2)

.PHONY: benchstat/new
benchstat/new: clean
	@${MAKE} --silent bench/fatih
	@${MAKE} --silent bench
	@benchstat old.txt new.txt

.PHONY: bench/cpu
bench/cpu: GO_BENCH_FLAGS+=-cpuprofile=cpu.pprof
bench/cpu: clean bnech

.PHONY: bench/mem
bench/mem: GO_BENCH_FLAGS+=-memprofile=mem.pprof
bench/mem: clean bench

.PHONY: bnech/mutex
bnech/mutex: GO_BENCH_FLAGS+=-mutexprofile=mutex.pprof
bnech/mutex: clean bench

.PHONY: bnech/block
bnech/block: GO_BENCH_FLAGS+=-blockprofile=block.pprof
bnech/block: clean bench

.PHONY: bnech/trace
bnech/trace: GO_BENCH_FLAGS+=-trace=trace.out
bnech/trace: clean bench

.PHONY: clean
clean:
	@$(RM) *.txt *.prof *.pprof
