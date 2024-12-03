FROM ubuntu:latest

ENV URL=https://www.google.com

RUN apt update \
    && apt install -y curl

COPY curl.sh .

ENTRYPOINT ["./curl.sh"]
