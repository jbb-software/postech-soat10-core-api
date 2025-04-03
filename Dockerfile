FROM golang:1.23

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o main ./app/cmd/main.go
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init --parseInternal --parseDependency -g ./app/cmd/main.go -o ./app/docs

CMD ["./main"]
