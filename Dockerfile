FROM golang:latest
WORKDIR /usr/
COPY go.mod ./
RUN go mod download
COPY ./ .
CMD [ "go", "run", "main.go" ]
