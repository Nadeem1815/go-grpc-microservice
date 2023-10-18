FROM golang:1.21.0-alpine AS build-stage
WORKDIR /app  
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -v -o ./build/api ./cmd
# Define environment variables
ENV PORT=:3000
ENV AUTH_SVC_URL=localhost:50051
ENV PRODUCT_SVC_URL=localhost:50052
ENV ORDER_SVC_URL=localhost:50053

FROM alpine:3.14 AS prod
WORKDIR /app
COPY --from=build-stage /app/build/api api
CMD [ "/app/api" ]
