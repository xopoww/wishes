# syntax=docker/dockerfile:1

FROM alpine:latest as clone

WORKDIR /repo
COPY .git/        ./.git/
COPY api/         ./api/
COPY cmd/         ./cmd/
COPY internal/    ./internal/
COPY restapi/     ./restapi/
COPY scripts/     ./scripts/
COPY go.mod       ./
COPY go.sum       ./
COPY Makefile     ./


FROM golang:1.19-alpine as go-base

RUN apk update && apk upgrade && apk add --update alpine-sdk && \
    apk add --no-cache bash git make


FROM go-base as build

WORKDIR /build
COPY --from=clone /repo/ ./

RUN make build


FROM go-base as sql-migrate

WORKDIR /db
COPY --from=clone /repo/scripts/                ./scripts
COPY --from=clone /repo/internal/repository/sqlite/migrations/ ./internal/repository/sqlite/migrations/

RUN go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN bash ./scripts/db/create.sh ./db.sqlite3


FROM alpine:latest

RUN adduser -D myuser

WORKDIR /app
COPY --from=build /build/bin/wishes-server ./
COPY --from=sql-migrate --chown=myuser /db/db.sqlite3 /data/db.sqlite3

USER myuser

EXPOSE 8080

ENV HOST=0.0.0.0 PORT=8080 WISHES_DBS=/data/db.sqlite3
CMD [ "./wishes-server" ]