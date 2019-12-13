FROM xushikuan/alpine-build:1.0

ENV TIME_ZONE=Asia/Singapore
RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone
WORKDIR /app
RUN mkdir logs
COPY main .
ENTRYPOINT ./main