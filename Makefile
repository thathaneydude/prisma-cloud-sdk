require-%:
	@ if [ "$(shell command -v ${*} 2> /dev/null)" = "" ]; then \
		echo "[$*] not found"; \
		exit 1; \
	fi

lint-project:
	@echo ">> validating code format" \
	@fmtRes=$$(gofmt -l .); \
	if [! -z $${fmtRes}]; then \
		echo "format check failed"; echo "$${fmtRes}"; \
		exit 1;
	fi

##############
TAGS = "containers_image_openpgp"

.PHONY: unit-test
unit-test: require-go
	go test -v ./... -tags ${TAGS}

.PHONY: lint
lint: require-go lint-project

.PHONY: fmt
fmt:
	@gofmt -l -w $(shell find . type -f -name '*.go' -not -path "./vendor/*")
	@echo "Formatted all project go files"