#!/usr/bin/python

from flask import Flask
from api import Api
import sys
app = Flask(__name__)

@app.route("/")
def main():
    return "Hello World!"

if __name__ == "__main__":
    f = Api(sys.argv[1], int(sys.argv[2]))
#    app.run()
