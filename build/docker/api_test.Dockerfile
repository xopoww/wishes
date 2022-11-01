# syntax=docker/dockerfile:1

FROM alpine:latest as clone

WORKDIR /repo
COPY test/api_test/ ./


FROM python:3.10-alpine

WORKDIR /test
COPY --from=clone /repo/ .

RUN adduser -D myuser
USER myuser

RUN pip install -r requirements.txt

ENTRYPOINT [ "python", "-m", "pytest", "-p", "no:cacheprovider", "--color=yes" ]
CMD [ "--collect-only" ]
