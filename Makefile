run:
	go run ./cmd/pz5-db

build:
	go build -o pz5-db.exe ./cmd/pz5-db

test:
	go test ./... -v
