FROM golang:1.18

WORKDIR /go/src/github.com/Shambou/todolist

COPY . .

RUN ["go", "install", "github.com/githubnemo/CompileDaemon@latest"]
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64
ENTRYPOINT ["/go/bin/CompileDaemon", "-exclude-dir", ".git", "-exclude-dir", ".idea", "-directory", ".", "-build", "go build -o app ./cmd/server/main.go", "-command", "./app", "-verbose"]
