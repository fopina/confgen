server {
    listen       80;
    server_name  localhost;

    location / {
        proxy_pass {{ "PROXIED_URL" | env }};
    }
}
