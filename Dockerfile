FROM scratch
COPY confgen /
ENTRYPOINT ["/confgen"]
