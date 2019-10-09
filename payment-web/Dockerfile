FROM alpine

# 设置时区
ARG TZ=PRC
ENV TZ ${TZ}
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY payment-web /payment-web
COPY conf /conf

ENTRYPOINT [ "/payment-web" ]

EXPOSE 8090
