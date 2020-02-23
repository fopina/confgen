FROM nginx:1.17-alpine

RUN wget https://github.com/fopina/confgen/releases/latest/download/confgen_linux_amd64 -O /usr/local/bin/confgen
RUN chmod a+x /usr/local/bin/confgen

ADD docker-entrypoint.sh /docker-entrypoint.sh

ADD nginx.conf.tpl /etc/nginx/conf.d/default.conf.tpl

# set a default value
ENV PROXIED_URL https://www.google.com

ENTRYPOINT [ "/docker-entrypoint.sh" ]
CMD [ "nginx", "-g", "daemon off;" ]
