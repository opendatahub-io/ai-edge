.PHONY: cli-test cli-build cli-run

GO=go

cli-test:
	${GO} test ./cmd/... ./... -v

cli-build:
	${GO} build -o odh ./cmd/main.go

cli-run:
	${GO} run ./cmd/main.go

cli-clean:
	rm -f odh
