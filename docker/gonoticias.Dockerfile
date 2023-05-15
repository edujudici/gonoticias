FROM golang

LABEL maintainer="Eduardo Judici <edujudici@gmail.com>"

WORKDIR /app/src/gonoticias
ENV GOPATH=/app
COPY . /app/src/gonoticias
RUN go build main.go
ENTRYPOINT ["./main"]
EXPOSE 8080