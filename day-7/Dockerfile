# step 1: build executable
# https://docs.docker.com/language/golang/build-images/
FROM golang:1.19-alpine as builder

RUN apk update && apk add --no-cache git && apk add --no-cache bash && apk add build-base

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /main ./cmd/main.go


# step 2: deploy
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /main /main
EXPOSE 8080

USER nonroot:nonroot

CMD ["./main"]