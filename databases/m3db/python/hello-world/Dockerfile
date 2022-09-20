FROM python

WORKDIR /app

RUN pip3 install requests-toolbelt

COPY m3db-python-hello-world.py .

USER nobody

ENTRYPOINT ["python3", "/app/m3db-python-hello-world.py"]
