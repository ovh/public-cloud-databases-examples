FROM mongo

COPY version.sh /version.sh
RUN chmod +x /version.sh

USER nobody

ENTRYPOINT ["/version.sh"]
