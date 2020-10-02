FROM golang:alpine3.12

COPY . .

RUN go build src/main/*.go

ENTRYPOINT [ "./main" ]