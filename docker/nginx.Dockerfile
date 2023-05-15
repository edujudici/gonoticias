FROM nginx:latest
LABEL maintainer="Eduardo Judici <edujudici@gmail.com>"
COPY /docker/config/nginx.conf /etc/nginx/nginx.conf
EXPOSE 80 443
ENTRYPOINT ["nginx"]

# Parametros extras para o entrypoint
CMD ["-g", "daemon off;"]