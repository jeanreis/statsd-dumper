ARG ARCH_IMAGE
ARG ARCH_GO

FROM ${ARCH_IMAGE}/golang:1.19-alpine as base

# RUN adduser \
#   --disabled-password \
#   --gecos "" \
#   --home "/nonexistent" \
#   --shell "/sbin/nologin" \
#   --no-create-home \
#   foo

WORKDIR $GOPATH/src/statsd-dumper.reis.codes

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=${ARCH} go build -ldflags="-s -w" -o /main .

FROM scratch

COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group

COPY --from=base /main .

USER nobody:nobody

ENV STATSD_PORT 8125

EXPOSE ${STATSD_PORT}/udp

CMD ["./main"]