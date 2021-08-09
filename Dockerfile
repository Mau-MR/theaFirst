FROM golang:1.16-alpine AS build

RUN apk add --no-cache git
RUN apk add build-base

WORKDIR /thea

COPY . .

RUN go mod download

RUN go test -cover -race ./...

RUN go build -o ./temp/server cmd/server/main.go

FROM alpine:3.9 AS bin

COPY --from=build /thea/temp/server /app/server

ENTRYPOINT ["/app/server","-port=8080"]

EXPOSE 8080