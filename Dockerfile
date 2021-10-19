FROM golang:1.17.2-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . /app

RUN go build -o /KafkaCallApis

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /KafkaCallApis /KafkaCallApis

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/KafkaCallApis"]