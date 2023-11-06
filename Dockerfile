FROM golang:latest
WORKDIR /app
ENV GOPATH=/

COPY ./ ./

RUN apt-get update && \
    apt-get -y install postgresql-client

EXPOSE 8080

# build go app
RUN go mod download
RUN go build -o url-shortener ./cmd/url-shortener/main.go

CMD ["./url-shortener", "-db-host", "postgres-for-test-ozon", "-db-port", "5432"]