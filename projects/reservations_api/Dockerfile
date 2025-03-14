# Build
FROM golang:1.24.1 AS build

WORKDIR /migrate_tool

# install the migration tool
RUN curl -L -o migrate.tar.gz https://github.com/golang-migrate/migrate/releases/download/v4.18.2/migrate.linux-amd64.tar.gz && \
    tar -xvzf migrate.tar.gz && \
    mv ./migrate /migrate

WORKDIR /service
ADD . /service

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# pre downloads the dependencies for testing
RUN go mod download

RUN go build -o /http-service ./cmd

CMD [ "/http-service" ]

# Test
FROM build AS test

# Production
FROM alpine:latest AS production

RUN apk --no-cache add ca-certificates

# copy app
COPY --from=build /http-service ./
RUN chmod +x ./http-service

# copy migrations
COPY --from=build /migrations ./

ENTRYPOINT [ "./http-service" ]
EXPOSE 8080