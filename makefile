.PHONY: test
test:
	cd palantir/controllers; go test -v .

.PHONY: setup-test
setup-test:
	docker compose up --build -d