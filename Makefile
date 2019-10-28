init:
	if ! hash dep 2>/dev/null; then go get -u github.com/golang/dep/cmd/dep; fi
	if ! hash sql-migrate 2>/dev/null; then go get -v github.com/rubenv/sql-migrate/...; fi
	if ! hash mockgen 2>/dev/null; then go get github.com/golang/mock/mockgen; fi
	if ! hash swagger 2>/dev/null; then go get -u github.com/go-swagger/go-swagger/cmd/swagger; fi
	dep ensure

run: db-migrate
	dep ensure
	go run main.go -c=config/config.yaml

test: db-migrate
	dep ensure
	go test -v ./test/

spec:
	swagger generate spec -m -o ./spec/api.json

spec-ui: spec
	swagger serve -F=swagger ./spec/api.json

db-migrate:
	sql-migrate up -config=store/store.yaml
