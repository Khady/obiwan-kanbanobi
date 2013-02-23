from flask import Flask
from flask import g
from flask.ext.sqlalchemy import SQLAlchemy
import sys

app = Flask(__name__)
app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///:memory:'
db = SQLAlchemy(app)
app.config.from_object('config')

from api import Api
a = Api(sys.argv[1], int(sys.argv[2]))
a.start()

from app import views
