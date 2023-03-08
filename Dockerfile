FROM golang:1.20-alpine as builder

RUN mkdir /build

WORKDIR /build

ADD . .

RUN go build -o ./health-check .

FROM valeriansaliou/sonic:v1.4.0

ENV PORT 1491
ENV PASSWORD SecretPassword
ENV HEALTH_CHECK_HOST localhost

COPY --from=builder /build/health-check /health-check

ADD sonic.cfg /etc/sonic.cfg

CMD [ "sonic", "-c", "/etc/sonic.cfg" ]
