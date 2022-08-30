FROM python

WORKDIR /app

RUN pip3 install kafka-python

COPY example.py .

USER nobody

ENTRYPOINT ["python3", "/app/example.py"]
