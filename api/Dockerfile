
FROM ubuntu:20.04

WORKDIR /app

COPY ./main /app

RUN mkdir /data

EXPOSE 10001

CMD ["./main", "-f", "/data/config/api.yaml"]
