FROM debian

RUN apt-get update && apt-get install -y curl

COPY m3db-bash-hello-world.sh /m3db-bash-hello-world.sh
RUN chmod +x /m3db-bash-hello-world.sh

USER nobody
ENV HOME /tmp

ENTRYPOINT ["/m3db-bash-hello-world.sh"]
