.PHONY: cli-test cli-build cli-run

GO=go

GENERATOR=./hack/openapi-generator-cli.sh
GENERATOR_CONFIG=./hack/openapi-generator-config.yaml
GENERATOR_INPUT_SPEC=./api/model-registry-openapi.yaml
GENERATOR_OUTPUT_DIR=./pkg/generated/model_registry_client


cli-test:
	${GO} test -v ./cmd/... ./... ${GENERATOR_OUTPUT_DIR}/...

cli-build:
	${GO} build -o odh ./cmd/main.go

cli-run:
	${GO} run ./cmd/main.go

cli-clean:
	rm -f odh

cli-generate-clients:
	@which mvn > /dev/null 2>&1 || (echo "mvn (Maven) is not installed, please install it to proceed" && exit 1)
	@which java > /dev/null 2>&1 || (echo "java is not installed, please install it to proceed" && exit 1)
	${GENERATOR} generate -i ${GENERATOR_INPUT_SPEC} -g go -o ${GENERATOR_OUTPUT_DIR} -c ${GENERATOR_CONFIG}
