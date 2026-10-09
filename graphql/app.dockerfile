FROM golang:1.26-alpine AS build

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY account account
COPY catalog catalog
COPY order order
COPY graphql graphql
COPY telemetry telemetry
COPY lifecycle lifecycle

RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app ./graphql

FROM alpine:3.22

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=build /go/bin/app ./app

EXPOSE 8080

CMD ["./app"]