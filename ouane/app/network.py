import socket
import OpenSSL
import sys
import struct
import select

class Network( ):
    def __init__(self, host, port):
        #self.context = OpenSSL.SSL.Context(OpenSSL.SSL.TLSv1_METHOD)
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
        self.readSize = 0

    def getReadedStack(self):
        return self.readStack

    def getReadedMessage(self):
        return self.readStack.pop()

    def setWriteStack(self, msg):
        self.writeStack.append(msg)

    def send(self):
        for l in self.writeStack:
            self.connection.send(struct.pack("!I", len(l)))
            self.connection.send(l)
        self.writeStack = []

    def recv(self):
        if self.readSize == 0:
            self.readSize = struct.unpack("!I", self.connection.recv(4))[0]
        else:
            self.readStack.append(self.connection.recv(self.readSize))

    def run(self):
        inputs = [self.s]
        outputs = [self.s]
        readable, writable, exceptional = select.select(inputs, [], inputs, 1)
        #if len(writable) != 0:
        self.send()
        if len(readable) != 0:
            self.recv()
