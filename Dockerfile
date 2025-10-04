FROM golang:latest

COPY . /dnd-agent
WORKDIR /dnd-agent

RUN go mod download
RUN go build -o bin/ cmd/main.go
EXPOSE 8081
CMD ["/dnd-agent/bin/main"]