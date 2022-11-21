FROM golang:alpine AS builder

RUN mkdir /app

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o associate main.go



FROM alpine

WORKDIR /app

COPY --from=builder /app/associate .

EXPOSE 8082

CMD [ "/app/associate" ]