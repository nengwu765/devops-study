FROM ubuntu
ENV VERSION=0.1
LABEL author="765" type="httpserver"
COPY bin/amd64/httpserver /opt/case/httpserver
EXPOSE 5555
ENTRYPOINT ["/opt/case/httpserver"]

