FROM resin/rpi-raspbian:jessie-20170111
LABEL maintainer "github.com/westlad"
ENV LAST_UPDATED=1-2-2017
RUN apt-get update && apt-get install -y wget
RUN wget -O /opt/go.tar.gz https://storage.googleapis.com/golang/go1.7.5.linux-armv6l.tar.gz
RUN tar -C /usr/local -xzf /opt/go.tar.gz
RUN useradd -ms /bin/bash pi
RUN groupadd -g 997 gpio
RUN usermod -aG gpio pi
#USER pi
ENV GOPATH=/home/pi/go
ENV PATH=$PATH:/usr/local/go/bin
ENTRYPOINT go install github.com/westlad/temperature && /home/pi/go/bin/temperature
