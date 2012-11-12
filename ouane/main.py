#!/usr/bin/python

from flask import Flask
from api import Api
from db import Db
import sys

app = Flask(__name__)

@app.route("/")
def main():
    return "Hello World!"

if __name__ == "__main__":
    #f = Api(sys.argv[1], int(sys.argv[2]))
    db = Db()
    db.addUser(1, "tata")
    db.addUser(2, "tata")
    db.addUser(1, "toto")
    db.addProjectsById(1, "UBER")
    db.addColumnsById(1, "LOL")
    print db.getAllUser()
    print db.getAllProjectsById()
    print db.getAllColumnsById()
    #app.run()
