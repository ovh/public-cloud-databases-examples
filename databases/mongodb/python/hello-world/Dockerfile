FROM python

WORKDIR /app

RUN pip3 install pymongo[srv]

COPY version.py .

USER nobody

ENTRYPOINT ["python3", "/app/version.py"]
