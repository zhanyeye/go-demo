FROM alpine
WORKDIR /root/
ADD main /root/
EXPOSE 8000
ENTRYPOINT ["./main"]