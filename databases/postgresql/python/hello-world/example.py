#!/usr/bin/python3
import psycopg2
import sys

def connect():
    conn = None
    try:        
        if len(sys.argv) < 2:
            raise Exception("URI is missing in the environment.")

        conn = psycopg2.connect(sys.argv[1])

        cur = conn.cursor()
        cur.execute('SELECT version()')
        db_version = cur.fetchone()
        print(db_version)
        cur.close()

    except (Exception, psycopg2.DatabaseError) as error:
        print(error)
    finally:
        if conn is not None:
            conn.close()

if __name__ == '__main__':
    connect()
