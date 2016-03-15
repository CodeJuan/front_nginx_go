FROM golang:1.6


# ssh
#RUN apt-get update
#RUN apt-get install -y openssh-server
#RUN mkdir -p /var/run/sshd
#RUN echo "root:123456" | chpasswd
#EXPOSE 22
#ENTRYPOINT /usr/sbin/sshd -D

# go
ENV CUR_DIR /go/src/github.com/front_go
ENV GOPATH $CUR_DIR/Godeps/_workspace:$GOPATH

WORKDIR $CUR_DIR
COPY . $CUR_DIR
RUN go build
EXPOSE 80
ENTRYPOINT ["./front_go"]
