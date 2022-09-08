FROM cassandra

ENV SSL_CERTFILE=/certificates/ca.certificate.pem
COPY version.sh /version.sh
USER root
RUN chmod +x /version.sh

USER nobody


ENTRYPOINT ["/version.sh"]
