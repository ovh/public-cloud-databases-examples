FROM archlinux

RUN pacman -Sy redis --noconfirm
COPY version.sh /version.sh
RUN chmod +x /version.sh

USER nobody
ENV HOME /tmp

ENTRYPOINT ["/version.sh"]
