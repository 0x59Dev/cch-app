FROM nginx:latest

RUN apt-get update \
  && apt-get install --no-install-recommends --no-install-suggests -y \
    m4 \
  && rm -rf /var/lib/apt/lists/*

COPY default.conf.m4 /etc/nginx/conf.d/default.conf.m4
COPY start.sh /start.sh

CMD ["/start.sh"]
