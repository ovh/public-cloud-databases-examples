FROM python

WORKDIR /app

RUN pip3 install psycopg2

COPY example.py .

USER nobody

ENTRYPOINT ["python3", "/app/example.py"]
