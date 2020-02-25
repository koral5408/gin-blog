FROM scratch

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/gin-blog
COPY . $GOPATH/gin-blog

EXPOSE 8000
ENTRYPOINT ["./gin-blog"]