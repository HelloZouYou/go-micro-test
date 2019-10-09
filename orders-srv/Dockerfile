FROM alpine

# 设置时区
ARG TZ=PRC
ENV TZ ${TZ}
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY orders-srv /orders-srv
COPY conf /conf

ENTRYPOINT [ "/orders-srv" ]

CMD [ "--registry_address=etcd:2379 --register_interval=5 --register_ttl=10" ]
