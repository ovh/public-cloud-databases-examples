FROM postgres

RUN mkdir /entrypoint
COPY example.sh /entrypoint/example
RUN chmod +x /entrypoint/example

USER nobody

ENTRYPOINT ["/entrypoint/example"]