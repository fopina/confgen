server {
    listen       80;
    server_name  localhost;

    {{ if ne ("SSL_CERT"|env) "" }}
    ssl on;
    ssl_certificate {{ "SSL_CERT" | env }};
    ssl_certificate_key {{ "SSL_CERT_KEY" | env }};
    {{ end }}

    location / {
        proxy_pass {{ "PROXIED_URL" | env }};
    }
}
