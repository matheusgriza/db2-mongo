FROM golang:1.25.3 AS build

WORKDIR /app

COPY . /app

RUN go build -o /server ./cmd/api

FROM gcr.io/distroless/base-debian12

WORKDIR /

COPY --from=build /server /server

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/server" ]

