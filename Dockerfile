# syntax=docker/dockerfile:1
FROM golang:1.18-alpine as build

RUN apk --no-cache add make

WORKDIR /app
COPY Makefile .
COPY go.mod .
COPY go.sum .
RUN make deps

COPY . .
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=build /app/WeatherFetcher ./
COPY config ./config/

EXPOSE 8000
ENV GIN_MODE=release
ENTRYPOINT ["./WeatherFetcher",  "-e", "production"]
