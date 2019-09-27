FROM alpine

# 设置时区
ARG TZ=PRC
ENV TZ ${TZ}
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY user-srv /user-srv
COPY conf /conf

ENTRYPOINT [ "/user-srv" ]

CMD [ "--registry_address=registry:8500 --register_interval=5 --register_ttl=10" ]

EXPOSE 8080
