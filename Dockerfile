FROM golang:1.14.8-alpine

ENV GO111MODULE=on

WORKDIR /build

COPY ./go.mod .

COPY ./go.sum .

COPY . .

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .

EXPOSE 3000

CMD ["/dist/main"]