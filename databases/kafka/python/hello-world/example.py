#!/usr/bin/python3
from kafka import KafkaAdminClient
import sys

def connect():
    conn = None
    try:        
        if len(sys.argv) < 2:
            raise Exception("usage: {} <URI>".format(sys.argv[0]))
        
        client = KafkaAdminClient(bootstrap_servers=sys.argv[1],
                    security_protocol="SSL",
                    ssl_cafile="/certificates/ca.certificate.pem",
                    ssl_certfile="/certificates/access.certificate.pem",
                    ssl_keyfile="/certificates/access.key")
        
        topics = client.list_topics()
        topics.sort()

        if (len(topics) == 0):
            print("No topics found on the cluster.")
        else:
            print("Found the following topics on the cluster:")
            for topic in topics:
                print("- {}".format(topic))

    except (Exception) as error:
        print(error)
    finally:
        if client is not None:
            client.close()

if __name__ == '__main__':
    connect()
