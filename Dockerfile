FROM scratch
EXPOSE 5000
ENTRYPOINT ["/goproxy" "--port:5000", "--host=0.0.0.0"]
COPY ./bin/ /
