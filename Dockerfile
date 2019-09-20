FROM library/golang AS build

MAINTAINER tailinzhang1993@gmail.com

ENV APP_DIR /go/src/fabric-orderer-seek
RUN mkdir -p $APP_DIR
WORKDIR $APP_DIR
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o fabric-orderer-seek .

# Create a minimized Docker mirror
FROM scratch AS prod

COPY --from=build /go/src/fabric-orderer-seek/fabric-orderer-seek /fabric-orderer-seek
EXPOSE 8080
CMD ["/fabric-orderer-seek"]
