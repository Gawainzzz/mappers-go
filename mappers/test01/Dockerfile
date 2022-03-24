
FROM ubuntu:16.04

RUN mkdir -p kubeedge

COPY ./bin/test01 kubeedge/
COPY ./config.yaml kubeedge/

WORKDIR kubeedge

ENTRYPOINT ["/kubeedge/test01", "--v", "5"]
