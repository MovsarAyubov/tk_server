FROM golang:1.21.0

WORKDIR /app

COPY go.mod go.sum ./

# Install the package
RUN go mod download

COPY . ./

RUN go build -o /cmd/apiserver cmd/main.go

EXPOSE ${SERVER_PORT}

CMD ["/cmd/apiserver"]