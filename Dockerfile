FROM golang:1.16-alpine
RUN mkdir/app
ADD . /app
WORKDIR /app
RUN go clean -modcache
RUN go build -o main
EXPOSE 8000
CMD ["/app/main"]