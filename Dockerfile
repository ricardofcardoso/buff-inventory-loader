# syntax=docker/dockerfile:1

FROM golang:1.21.5 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /buff-inventory-loader ./main

FROM alpine:latest AS build-release-stage

WORKDIR /

COPY --from=build-stage /buff-inventory-loader /buff-inventory-loader

ENTRYPOINT ["/buff-inventory-loader"]