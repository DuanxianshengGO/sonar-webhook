FROM gaia-e2-01-registry.cn-shanghai.cr.aliyuncs.com/middleware/ubuntu:22.04


ENV LANG C.UTF-8

ENV TZ=Asia/Shanghai
RUN apt update
RUN apt-get install -y ca-certificates

ADD sonar-webhook /usr/bin/

CMD ["/usr/bin/sonar-webhook"]