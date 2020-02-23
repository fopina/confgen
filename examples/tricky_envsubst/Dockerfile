FROM nginx:1.17-alpine

ADD docker-entrypoint.sh /docker-entrypoint.sh

ADD nginx.conf.1.tpl /etc/nginx/conf.d/default.conf.1.tpl
ADD nginx.conf.2.tpl /etc/nginx/conf.d/default.conf.2.tpl
ADD nginx.conf.3.tpl /etc/nginx/conf.d/default.conf.3.tpl

# set a default value
ENV PROXIED_URL https://www.google.com

ENTRYPOINT [ "/docker-entrypoint.sh" ]
CMD [ "nginx", "-g", "daemon off;" ]
