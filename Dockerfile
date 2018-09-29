FROM golang:1.11-stretch as builder
WORKDIR /root
COPY helloworld.go /root/
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /root/helloworld /root/helloworld.go

FROM scratch
WORKDIR /root
COPY --from=builder /root/helloworld /root/
COPY helloworld.html logo.png /root/
CMD ["/root/helloworld"]

EXPOSE 9000
