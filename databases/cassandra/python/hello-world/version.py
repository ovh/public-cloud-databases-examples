#!/usr/bin/python3
import sys
from cassandra.cluster import Cluster
from cassandra.auth import PlainTextAuthProvider

def connect():
    if len(sys.argv) < 4:
            raise Exception("missing parameters, should be: Host, Port, Username, Password")

    auth_provider = PlainTextAuthProvider(username=sys.argv[3], password=sys.argv[4])
    cluster = Cluster([sys.argv[1]], ssl_options=dict(ca_certs="/certificates/ca.certificate.pem"), port=sys.argv[2], auth_provider=auth_provider)
    session = cluster.connect()
    try:
        session
        rows = session.execute("SELECT cql_version FROM system.local")
        print(rows[0])
    except:
        print("failure")
    finally:
        cluster.shutdown()

if __name__ == '__main__':
    connect()
