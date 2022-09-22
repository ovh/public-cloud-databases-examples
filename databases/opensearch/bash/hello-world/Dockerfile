FROM archlinux

COPY version.sh /version.sh
RUN pacman -Sy unzip jq --noconfirm \
    && curl -fL https://artifacts.opensearch.org/opensearch-clients/opensearch-cli/opensearch-cli-1.1.0-linux-x64.zip -o /tmp/opensearch-cli.zip \
    && unzip -d /bin /tmp/opensearch-cli \
    && whereis opensearch-cli \
    && chmod +x /version.sh

USER nobody

ENTRYPOINT ["/version.sh"]
