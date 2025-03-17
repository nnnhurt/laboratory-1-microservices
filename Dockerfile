# STEP-1
# build app from source

FROM golang:1.24.1-alpine3.21 AS builder

ENV GOCACHE=/root/.cache/go-build

WORKDIR /myapp

COPY ./cmd ./cmd

COPY ./vendor ./vendor
COPY ./go.mod ./go.sum ./

COPY ./internal ./internal

RUN --mount=type=cache,target="/root/.cache/go-build" \
   go build -o app ./cmd/main.go


# STEP-2
# make container

FROM alpine:3.21

WORKDIR /mysuperapp

COPY --from=builder /myapp/app ./

EXPOSE 7080

CMD [ "/mysuperapp/app" ]