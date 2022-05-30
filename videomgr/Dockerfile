
FROM ubuntu:20.04

WORKDIR /app

COPY ./main /app

RUN mkdir /data

CMD ["./main", "-f", "/data/config/videomanager.yaml"]
