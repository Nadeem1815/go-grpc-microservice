FROM golang:1.21.0-alpine AS build-stage  
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -v -o ./build/api ./cmd

FROM alpine:3.14 AS prod   
WORKDIR /app
COPY --from=build-stage /app/build/api /app/api
CMD [ "/app/api" ]