FROM resin/rpi-raspbian
ENV LAST_UPDATED=28-01-2017
RUN apt-get update -qq
RUN apt-get install -yqq apt-utils
RUN apt-get install -yqq nginx
RUN mkdir -p /tmp/nginx/
EXPOSE 80
CMD cp /tmp/nginx/nginx.conf /etc/nginx/nginx.conf && service nginx start
