FROM golang:alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go build -o ./godo && mv ./godo /usr/bin/godo

FROM alpine

COPY --from=builder /usr/bin/godo /usr/local/bin/godo

ENTRYPOINT [ "godo" ]
