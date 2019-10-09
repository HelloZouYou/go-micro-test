FROM alpine

# 设置时区
ARG TZ=PRC
ENV TZ ${TZ}
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY payment-web /payment-web
COPY conf /conf

ENTRYPOINT [ "/payment-web" ]

CMD [ "--registry_address=etcd:2379 --register_interval=5 --register_ttl=10" ]

EXPOSE 8090
