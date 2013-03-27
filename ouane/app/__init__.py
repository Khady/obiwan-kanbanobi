from flask import Flask
from flask import g
from flask.ext.sqlalchemy import SQLAlchemy
from config import IPSERVER, PORTSERVER
import datetime
import redis

red = redis.StrictRedis()
app = Flask(__name__)
app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///:memory:'
db = SQLAlchemy(app)
app.config.from_object('config')

def event_stream():
    pubsub = red.pubsub()
    pubsub.subscribe('ouane')
    # TODO: handle client disconnection.
    for message in pubsub.listen():
        print message
        yield 'data: %s\n\n' % message['data']

from api import Api
a = Api(IPSERVER, PORTSERVER)
#a = Api(sys.argv[1], int(sys.argv[2]))
a.start()
#a.createColumns(1,"1",1)
from app import views
