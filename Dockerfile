FROM alpine:latest
ADD ./test /test
CMD ["./test"]
