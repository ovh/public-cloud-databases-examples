FROM confluentinc/cp-kafkacat

USER root

RUN mkdir /app/
COPY kafka-bash-hello-world.sh /app/kafka-bash-hello-world.sh
RUN chmod +x /app/kafka-bash-hello-world.sh

USER nobody

ENTRYPOINT ["/app/kafka-bash-hello-world.sh"]
