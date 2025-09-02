# build app
FROM golang:1.22.1 AS main-env
RUN mkdir /app
ADD . /app/
ADD cmd/ /app/cmd/
WORKDIR /app
RUN go build -ldflags="-s -w" -o run ./cmd

# run artifacts
FROM scratch
WORKDIR /app
COPY --from=main-env /app/run /app
EXPOSE 8080

CMD ["/app/run"]
