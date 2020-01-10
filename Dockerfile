FROM scratch
LABEL maintainer="Robby.Lee<kunlin_lee@live.com>"
COPY ./conf /simple-go-api/conf
COPY ./dist/simple-go-api /simple-go-api/app
EXPOSE 8001
WORKDIR /simple-go-api
ENTRYPOINT ["./app"]