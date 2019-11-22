
FROM golang:latest
EXPOSE 3000 27017
RUN mkdir /app
ADD . /app/
WORKDIR /app
ENV GOPATH /app
RUN cd /app ; go install backend
CMD ["/app/bin/backend"]
