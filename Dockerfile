FROM alpine:3.6 as alpine
RUN apk add -U --no-cache ca-certificates

FROM scratch

ENV GODEBUG=netdns=go

COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

LABEL org.label-schema.version=latest
LABEL org.label-schema.vcs-url="https://github.com/josmo/drone-netapp-snapshot.git"
LABEL org.label-schema.name="Drone netapp-snapshot"
LABEL org.label-schema.vendor="Josmo"

ADD release/linux/amd64/drone-netapp-snapshot /bin/
ENTRYPOINT ["/bin/drone-netapp-snapshot"]
