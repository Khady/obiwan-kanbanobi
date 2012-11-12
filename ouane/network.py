import socket
import OpenSSL
import sys
import struct

class Network:
    def __init__(self, host, port):
        self.context = OpenSSL.SSL.Context(OpenSSL.SSL.TLSv1_METHOD)
        self.s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
#s.settimeout(5)
        self.connection = OpenSSL.SSL.Connection(context,s)
        self.connection.connect((host, port))
        while True:
            try:
                connection.do_handshake()
                break
            except OpenSSL.SSL.WantReadError:
                print "Exception"
                pass
    def send(self, msg):
        connection.send(msg)

    def recv(self, nb):
        return connection.recv(nb)
