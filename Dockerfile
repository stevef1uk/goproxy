FROM scratch
EXPOSE 8080
ENTRYPOINT ["/goproxy"]
COPY ./bin/ /