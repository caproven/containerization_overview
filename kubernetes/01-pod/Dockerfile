FROM golang:1.24-alpine AS build

WORKDIR /build

COPY go.mod .
RUN go mod download

COPY . .
RUN go build -o httpserver .

FROM alpine

USER 1001

WORKDIR /app
COPY --from=build /build/httpserver .

CMD ["./httpserver"]
