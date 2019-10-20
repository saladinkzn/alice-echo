FROM golang:1.10.4 as build
WORKDIR /go/src/alice-echo

COPY alice alice
COPY handler handler

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -a --installsuffix cgo -o alice-echo .

FROM alpine:3.8

WORKDIR /root

EXPOSE 8080

RUN apk --update --no-cache add ca-certificates

COPY --from=build /go/src/alice-echo/alice-echo .

COPY docker-entrypoint.sh .
RUN chmod +x docker-entrypoint.sh

CMD ["./docker-entrypoint.sh"]