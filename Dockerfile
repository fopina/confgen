FROM scratch

ARG BINARY=confgen
COPY ${BINARY} /confgen

ENTRYPOINT ["/confgen"]
