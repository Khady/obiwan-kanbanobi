import socket
import OpenSSL
import sys
import struct
import select
import threading

class Network(threading.Thread):
    def __init__(self, host, port):
        #self.context = OpenSSL.SSL.Context(OpenSSL.SSL.TLSv1_METHOD)
        threading.Thread.__init__(self)
        self.s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
#s.settimeout(5)
        #self.connection = OpenSSL.SSL.Connection(self.context, self.s)
        #self.connection.connect((host, port))
        # while True:
        #     try:
        #         self.connection.do_handshake()
        #         break
        #     except OpenSSL.SSL.WantReadError:
        #         print "Exception"
        #         pass
        self.s.connect((host, port))
        self.connection = self.s
        self.readStack = []
        self.writeStack = []
        self.readSize = 4096

    def getReadedStack():
        return self.readStack

    def setWriteStack(self, msg):
        self.writeStack.append(msg)

    def send(self):
        for l in self.writeStack:
            print len(l)
            self.connection.send(struct.pack("!I", len(l)))
            self.connection.send(l)
        self.writeStack = []

    def recv(self):
        if self.readSize == 0:
            self.readSize = struct.unpack("!I", self.connection.recv(4))[0]
            print self.readSize
        else:
            self.readStack.append(self.connection.recv(self.readSize))

    def run(self):
        while 1:
            inputs = [self.s]
            outputs = [self.s]
            readable, writable, exceptional = select.select(inputs, outputs, inputs)
            if len(writable) != 0:
                self.send()
            if len(readable) != 0:
                self.recv()
