set -e
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -buildmode=pie -ldflags "-w -s" -o ./main.run
