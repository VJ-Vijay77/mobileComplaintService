FROM golang:alpine AS builder

RUN mkdir /app

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o customer main.go



FROM alpine

WORKDIR /app

COPY --from=builder /app/customer .

EXPOSE 8080

CMD [ "/app/customer" ]