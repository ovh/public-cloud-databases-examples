#!/usr/bin/python3

import requests
import sys

def connect():
    try:        
        if len(sys.argv) < 2:
            raise Exception("URI is missing")

        res = requests.get('{}/api/v1/labels'.format(sys.argv[1]))
        print(res.content)

    except Exception as error:
        print(error)

if __name__ == '__main__':
    connect()
