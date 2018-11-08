FROM golang:1.11 as builder

ENV CGO_ENABLED=0

WORKDIR /go/src/github.com/hortonworks/superset-proxy
COPY . .

RUN go build

###

FROM nginx:stable-alpine

COPY --from=builder /go/src/github.com/hortonworks/superset-proxy/superset-proxy /
COPY nginx.conf.tmpl /etc/nginx/

ENTRYPOINT ["/superset-proxy"]
CMD ["/usr/sbin/nginx", "-g", "daemon off;"]