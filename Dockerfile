# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY echo.go ./

RUN go build echo.go

ENTRYPOINT ["./echo"]