FROM python

WORKDIR /app

RUN pip3 install mysql-connector-python

COPY version.py .

USER nobody

ENTRYPOINT ["python3", "/app/version.py"]
