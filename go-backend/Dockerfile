FROM golang:latest

# Зависимости
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY ./ ./

RUN go build -o go-app cmd/main.go


CMD ["./go-app"]

