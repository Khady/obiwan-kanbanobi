from flask import Flask
from flask import g
from flask.ext.sqlalchemy import SQLAlchemy
import sys
from config import IPSERVER, PORTSERVER

app = Flask(__name__)
app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///:memory:'
db = SQLAlchemy(app)
app.config.from_object('config')

from api import Api
a = Api(IPSERVER, PORTSERVER)
#a = Api(sys.argv[1], int(sys.argv[2]))
a.start()
#a.createColumns(1,"1",1)
from app import views
