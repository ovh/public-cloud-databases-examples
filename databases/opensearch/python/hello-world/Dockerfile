FROM python

WORKDIR /app

RUN pip install opensearch-py

COPY version.py .

USER nobody

ENTRYPOINT ["python3", "/app/version.py"]
