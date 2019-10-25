FROM evildecay/etcdkeeper

# 设置时区
ARG TZ=PRC
ENV TZ ${TZ}
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# 设置环境变量
# 可配置 https://goproxy.io、https://goproxy.cn
ARG GOPROXY=https://goproxy.cn
ENV GO111MODULE on