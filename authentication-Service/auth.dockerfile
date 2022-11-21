FROM golang:alpine AS builder

RUN mkdir /app

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o auth main.go



FROM alpine

WORKDIR /app

COPY --from=builder /app/auth .

EXPOSE 8081

CMD [ "/app/auth" ]