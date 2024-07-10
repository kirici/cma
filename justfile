set dotenv-required := true

alias c := clean
alias b := build
alias r := run
alias t := test-all

# clean binary, log file, db file
clean:
    rm -f main ${LOG_FILE} ${DB_URL}
    sqlite3 ${DB_URL} < init.sql

# build main
build: clean
    go build -o main cmd/api/main.go

run: build
    # go run cmd/api/main.go
    ./main

# test everything
test-all: build
    go test ./tests -v

# live reload, requires github.com/air-verse/air
watch:
    air
