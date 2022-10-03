#!/usr/bin/python3
import sys
from urllib.parse import urlparse
from opensearchpy import OpenSearch

def connect(uri):
  u = urlparse(uri)

  # Create the client with SSL/TLS enabled.
  client = OpenSearch(
      hosts = [{'host': u.hostname, 'port': u.port}],
      http_compress = True, # enables gzip compression for request bodies
      http_auth = (u.username, u.password),
      use_ssl = True,
      verify_certs = True
  )

  info = client.nodes.info()
  for node in info.get('nodes', {}).values():
    print(f"{node['name']}: {node['version']}")


if __name__ == '__main__':
  if len(sys.argv) < 2:
      raise Exception("URI is missing")
  connect(sys.argv[1])
