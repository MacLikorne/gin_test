FROM golang:1.17.1-buster as build

ADD ./cmd /src/cmd
ADD ./pkg /src/pkg
ADD ./go.mod /src
WORKDIR /src/cmd
RUN go get && go build -o /gin_test.bin main.go

FROM debian:buster-slim as run

RUN apt-get update && apt-get install -y ca-certificates && apt-get clean
COPY --from=build /gin_test.bin /usr/bin/gin_test
CMD ["/usr/bin/gin_test"]