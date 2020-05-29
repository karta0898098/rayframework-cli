package main

const TemplateDockerfile = `
FROM golang:1.14.2-alpine3.11 AS builder-env

WORKDIR /app

COPY go.mod /
COPY go.sum /

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app -tags=jsoniter .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache bash
RUN apk add tzdata

COPY --from=builder-env /app /app

RUN cd /app && ls

WORKDIR /app

EXPOSE 8080

ENTRYPOINT [ "./app" ]
`

const TemplateDockerfileWait = `
FROM golang:1.14.2-alpine3.11 AS builder-env

WORKDIR /app

COPY go.mod /
COPY go.sum /

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app -tags=jsoniter .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache bash
RUN apk add tzdata

COPY --from=builder-env /app /app

RUN cd /app && ls

WORKDIR /app

EXPOSE 8080

CMD ["./wait-for-it.sh","database:3306","-t","60","--","./app"]
`
