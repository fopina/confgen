ARG BINARY=confgen

FROM scratch

COPY ${BINARY} /

ENTRYPOINT ["/confgen"]
