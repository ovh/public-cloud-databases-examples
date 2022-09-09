#!/usr/bin/env python3

import redis
import sys
import json

def hello_redis():
    try:
        if len(sys.argv) < 2:
            raise Exception("URI is missing")

        # connect to Redis
        r = redis.Redis.from_url(sys.argv[1])

        # Read the key
        msg = r.info()

        # Parse the infos
        #infos = json.loads(msg)

        # Print the key
        #print(infos.redis_version)
        print(msg["redis_version"])

    except Exception as e:
        print(e)


if __name__ == '__main__':
    hello_redis()
