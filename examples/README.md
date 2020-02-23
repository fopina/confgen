The `simple_` examples scenario is where you want an image that `proxypass` to a configurable URL

```
docker run -e PROXIED_URL=http://example.com ...
```

Straightforward and fits `envsubst` perfectly. `confgen` works just as easily but no real advantage and requires this fat binary.


For a scenario where you need logic control such as a simple conditional or loop, `envsubst` can't help.

The `tricky_` examples scenario is the same `proxypass` image with the added option to enable HTTPS by passing a certificate and its key with a volume

```
# without HTTPS
docker run -e PROXIED_URL=http://example.com ...

# with HTTPS
docker run -e PROXIED_URL=http://example.com \
           -v /path/to/cert_and_key/:/etc/nginx/cert_and_key/ \
           -e SSL_CERT=/etc/nginx/cert_and_key/cert.crt \
           -e SSL_CERT_KEY=/etc/nginx/cert_and_key/cert.key \
           ...
```

The `ssl ...` block in nginx.conf should only be present if there is any `SSL_CERT` defined.

There are a few hacky ways to accomplish this with `envsubst`. My favorite (before `confgen` of course) is splitting the configuration into blocks and adding the include logic to the docker-entrypoint. This makes it easy to re-use blocks in case of loop logic.

But using `confgen` that doesn't matter as there's only one condition check added to the original template making it a lot easier to maintain at the mere cost of a 2M file added to your image.
