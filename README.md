# Benchmarks

This repository contains basic benchmarks for time.ParseDuration vs the proposed
time.ParseDurationBytes.

There is an optional benchmark for testing time.Duration marhsal support
proposed in [cl24473](https://go-review.googlesource.com/#/c/24473/)

There are two version of the BenchmarkParseDurationByteInput benchmark,
one for benchmarking master and the other for benchmarking
[cl24842](https://go-review.googlesource.com/#/c/24842/)

## Running Benchmarks

```bash
go get github.com/Shelnutt2/go-benchmark-ParseDuration
cd $GOPATH/src/github.com/Shelnutt2/go-benchmark-ParseDuration
go test -bench=BenchmarkParseDuration
```
