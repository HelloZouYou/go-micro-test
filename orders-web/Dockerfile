FROM alpine

# 设置时区
ARG TZ=PRC
ENV TZ ${TZ}
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY orders-web /orders-web
COPY conf /conf

ENTRYPOINT [ "/orders-web" ]

EXPOSE 8091
