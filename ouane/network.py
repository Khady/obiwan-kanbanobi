import socket
import OpenSSL
import sys
import struct

class Network:
    def __init__(self, host, port):
        self.context = OpenSSL.SSL.Context(OpenSSL.SSL.TLSv1_METHOD)
        self.s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
#s.settimeout(5)
        self.connection = OpenSSL.SSL.Connection(self.context, self.s)
        self.connection.connect((host, port))
        while True:
            try:
                self.connection.do_handshake()
                break
            except OpenSSL.SSL.WantReadError:
                print "Exception"
                pass
    def send(self, msg):
        self.connection.send(msg)

    def recv(self, nb):
        return self.connection.recv(nb)
