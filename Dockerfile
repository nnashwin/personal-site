FROM golang:1.14.4-alpine3.11
RUN apk --no-cache add gcc g++ make git ca-certificates
WORKDIR /usr/app
COPY . /usr/app
RUN GOOS=linux go build -ldflags="-s -w" -o /usr/bin/test /usr/app/main.go
COPY static /usr/app/static
COPY templates /usr/app/templates
CMD ["/usr/bin/test", "-sd", "/usr/app/static", "-td", "/usr/app/templates"]
