#!/usr/bin/python3

import sys
from pymongo import MongoClient
from pymongo.errors import ConnectionFailure

def connect():
    try:        
        if len(sys.argv) < 2:
            raise Exception("URI is missing")

        client = MongoClient(sys.argv[1])
        server_info = client.server_info()
        print(server_info["version"])


    except (Exception, ConnectionFailure) as error:
        print(error)

if __name__ == '__main__':
    connect()
