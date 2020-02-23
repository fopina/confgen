FROM nginx:1.17-alpine

ADD docker-entrypoint.sh /docker-entrypoint.sh

ADD nginx.conf.tpl /etc/nginx/conf.d/default.conf.tpl

# set a default value
ENV PROXIED_URL https://www.google.com

ENTRYPOINT [ "/docker-entrypoint.sh" ]
CMD [ "nginx", "-g", "daemon off;" ]
