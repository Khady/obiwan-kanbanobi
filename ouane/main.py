#!/usr/bin/python

from app import app

if __name__ == "__main__":
    app.run(debug=True)
    # print 'Listening on http://127.0.0.1:%s and on port 10843 (flash policy server)' % PORT
    # SocketIOServer(('', PORT), app, resource="socket.io").serve_forever()
