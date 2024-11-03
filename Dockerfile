FROM ubuntu:latest
LABEL authors="jscatena"

ENTRYPOINT ["top", "-b"]