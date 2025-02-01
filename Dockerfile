## Multistage build
FROM golang:1.23 AS build
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /src
COPY . .
RUN go mod download
RUN go build -o /app

## Multistage deploy
FROM gcr.io/distroless/static-debian12

WORKDIR /
COPY --from=build /src/configs /configs
COPY --from=build /app /app

ENTRYPOINT ["/app"]
