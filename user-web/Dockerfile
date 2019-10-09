FROM alpine

# 设置时区
ARG TZ=PRC
ENV TZ ${TZ}
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY user-web /user-web
COPY conf /conf

ENTRYPOINT [ "/user-web" ]

EXPOSE 8088