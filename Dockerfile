FROM golang:1.16.8-alpine AS builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build

ARG http_port=1323
ENTRYPOINT /build

ENV PORT $http_port
ENV DB_HOST 172.17.0.2
ENV DB_PORT 3306
ENV DB_USER user2
ENV DB_PASS Password_123
ENV DB_DATABASE oderValidator

EXPOSE $http_port

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/delivery-validation /app/
WORKDIR /app
CMD ["./delivery-validation"]