.PHONY: test
test:
	cd palantir/controllers; go test -count=1 -v . 

.PHONY: setup-test
setup-test:
	docker compose up --build -d
	cd examples/churn; pip install -r requirements.txt; python 01_churn.py
