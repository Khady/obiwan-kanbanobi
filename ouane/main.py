#!/usr/bin/python

from flask import g
# from flask.ext.sqlalchemy import SQLAlchemy
import sys
from app import app

# app = Flask(__name__)
# app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite://test.db'
# db = SQLAlchemy(app) 

if __name__ == "__main__":
    # f.createColumns(0, "aa", 0, 0, "lol")
    # c = Cards(1, "test", "salut je suis un test", 1, 1, "jerox;test;", 1, 0, 1)
    # co = Columns(4, "test", 1, "test contenu", "jerox;test;", 1, 1)
    # u = Users(12, "test", True, "motdepasse", "contact@email.com", True)
    # p = Projects(10, "test", 1337, "salut je suis un test", "dieu")
    # com = Comments(42, "test", "je suis dieu", 1, 24)
    # m = Metadata(10, 42, 1337, "data", "dieu")
    # db.session.add(c)
    # db.session.add(co)
    # db.session.add(u)
    # db.session.add(p)
    # db.session.add(com)
    # db.session.add(m)
    # db.session.commit()
    # print c
    # print co
    # print u
    # print p
    # print com
    # print m
    app.run(debug=True)
    #n = Network("127.0.0.1", 9658)
    # n.start()
