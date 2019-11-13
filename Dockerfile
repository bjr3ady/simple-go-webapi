FROM scratch
LABEL maintainer="Robby.Lee<kunlin_lee@live.com>"
COPY ./conf /school-board/conf
COPY ./dist/school-board-rest-api /school-board/app
EXPOSE 8001
WORKDIR /school-board
ENTRYPOINT ["./app"]