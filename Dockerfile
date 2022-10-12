FROM alpine
WORKDIR /root
COPY ./hello /root/
RUN apk add bash
RUN chmod +x hello
EXPOSE 8000
CMD [ "./hello" ]