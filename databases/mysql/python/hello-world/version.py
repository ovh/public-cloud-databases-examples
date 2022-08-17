#!/usr/bin/python3

import sys
import mysql.connector
from mysql.connector.constants import ClientFlag
from urllib.parse import urlparse

def connect():
    try:        
        if len(sys.argv) < 2:
            raise Exception("URI is missing")

        pu = urlparse(sys.argv[1])
        with mysql.connector.connect(user=pu.username, password=pu.password, host=pu.hostname, port=pu.port) as connection:
            with connection.cursor(buffered=False) as cursor:
                cursor.execute("select version();")
                print(cursor.fetchone()[0])

    except Exception as error:
        print(error)

if __name__ == '__main__':
    connect()
