FROM mj509.flymks.com:9090/mj15r/alpine:3.15.4
RUN \
  apk update && \
  apk add --no-cache nano sudo bash wget curl git tree grep && \
  wget https://busybox.net/downloads/binaries/1.28.1-defconfig-multiarch/busybox-x86_64 && \
  chmod +x busybox-x86_64 && \
  mv busybox-x86_64 bin/busybox1.28 && \
  mkdir -p /opt/www && echo "let me go" > /opt/www/index.html

ENTRYPOINT ["/bin/busybox1.28"]
CMD ["httpd", "-f", "-h", "/opt/www"]
