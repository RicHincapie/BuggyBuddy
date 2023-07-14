ARG GO_VERSION=1.20.2

FROM golang:${GO_VERSION}-alpine as BUILD

WORKDIR /src

COPY ./ .

RUN go build -o app main.go

FROM scratch as FINAL

LABEL maintainer="Ric Hincapie"

EXPOSE 8000

COPY --from=BUILD /src/app /app

ENTRYPOINT [ "/app" ]
