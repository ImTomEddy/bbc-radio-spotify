FROM golang:1.15-buster as build

EXPOSE 8081

WORKDIR /go/src/github.com/imtomeddy/bbc-radio-spotify
ADD . .

RUN go get -d -v ./...

RUN go build -o /go/bin/app

FROM gcr.io/distroless/base-debian10
COPY --from=build /go/bin/app /

CMD ["/app"]