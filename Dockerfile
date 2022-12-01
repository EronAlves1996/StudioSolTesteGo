FROM golang:latest as build

ENV APP_EXEC go_teste

WORKDIR $GOPATH/src/go_teste_app
COPY . .

RUN go mod tidy
RUN go build -o /$APP_EXEC

FROM alpine:latest

RUN apk add libc6-compat
ENV APP_EXEC go_teste

WORKDIR /
COPY --from=build /$APP_EXEC /$APP_EXEC
 
EXPOSE 8080
 
CMD /$APP_EXEC