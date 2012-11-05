# Echo client program
import socket
import OpenSSL
import sys
import struct

context = OpenSSL.SSL.Context(OpenSSL.SSL.TLSv1_METHOD)
HOST = sys.argv[1]              # The remote host
PORT = int(sys.argv[2])         # The same port as used by the server
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
#s.settimeout(5)
connection = OpenSSL.SSL.Connection(context,s)
connection.connect((HOST, PORT))

#connection.setblocking(1)

# Set the timeout using the setsockopt
# tv = struct.pack('ii', int(6), int(0))
# connection.setsockopt(socket.SOL_SOCKET, socket.SO_RCVTIMEO, tv)

print "Connected to " , connection.getpeername()
print "Sate " , connection.state_string()

while True:
    try:
        connection.do_handshake()
        break
    except OpenSSL.SSL.WantReadError:
        print "Exception"
        pass
print "Sate " , connection.state_string()

connection.send("bobo\n")

while True:
    try:
        # data = sys.stdin.readline()
        recvstr = connection.recv(1024)
        print "l'autre moche >"+recvstr
        connection.send("bobo\n")
    except OpenSSL.SSL.WantReadError:
        print "Exception"
        pass



