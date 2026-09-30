FROM golang:1.27 AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /auth ./cmd/auth

FROM gcr.io/distroless/static
COPY --from=build /auth /auth
EXPOSE 8081
ENTRYPOINT ["/auth"]
