#from main import app, db
from flask import Flask
from flask.ext.sqlalchemy import SQLAlchemy

# app = Flask(__name__)
# app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///:memory:'
# db = SQLAlchemy(app)
# app.config.from_object('config')

from app import app

class Cards(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.Text)
    content = db.Column(db.Text)
    column_id = db.Column(db.Integer)
    project_id = db.Column(db.Integer)
    tags = db.Column(db.Text)
    user_id = db.Column(db.Integer)
    scripts_id = db.Column(db.Integer)
    write = db.Column(db.Text)

    def __init__(self, id, name, content, column_id, project_id, tags, user_id, scripts_id, write):
        self.id = id
        self.name = name
        self.content = content
        self.column_id = column_id
        self.project_id = project_id
        self.tags = tags
        self.user_id = user_id
        self.scripts_id = scripts_id
        self.write = write

    def __repr__(self):
        return '<Cards id %d name %r>' % (self.id, self.name)

class Columns(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.Text)
    project_id = db.Column(db.Integer)
    content = db.Column(db.Text)
    tags = db.Column(db.Text)
    scripts_id = db.Column(db.Integer)
    write = db.Column(db.Text)

    def __init__(self, id, name, project_id, content, tags, scripts_id, write):
        self.id = id
        self.name = name
        self.content = content
        self.project_id = project_id
        self.tags = tags
        self.scripts_id = scripts_id
        self.write = write

    def __repr__(self):
        return '<Columns id %d name %r>' % (self.id, self.name)

class Users(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.Text)
    admin = db.Column(db.Boolean)
    password = db.Column(db.Text)
    mail = db.Column(db.Text)
    active = db.Column(db.Boolean)

    def __init__(self, id, name, admin, password, mail, active):
        self.id = id
        self.name = name
        self.admin = admin
        self.password = password
        self.mail = mail
        self.active = active

    def __repr__(self):
        return '<Users id %d name %r>' % (self.id, self.name)

class Projects(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.Text)
    admins_id = db.Column(db.Integer)
    content = db.Column(db.Text)
    read = db.Column(db.Text)

    def __init__(self, id, name, admins_id, content, read):
        self.id = id
        self.name = name
        self.content = content
        self.admins_id = admins_id
        self.read = read

    def __repr__(self):
        return '<Projects id %d name %r>' % (self.id, self.name)

class Comments(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.Text)
    content = db.Column(db.Text)
    cards_id = db.Column(db.Integer)
    author_id = db.Column(db.Integer)

    def __init__(self, id, name, content, cards_id, author_id):
        self.id = id
        self.name = name
        self.content = content
        self.cards_id = cards_id
        self.author_id = author_id

    def __repr__(self):
        return '<Comments id %d name %r>' % (self.id, self.name)

class Metadata(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    object_type = db.Column(db.Integer)
    object_id = db.Column(db.Integer)
    data_key = db.Column(db.Text)
    data_value = db.Column(db.Text)

    def __init__(self, id, object_type, object_id, data_key, data_value):
        self.id = id
        self.object_type = object_type
        self.object_id = object_id
        self.data_key = data_key
        self.data_value = data_value

    def __repr__(self):
        return '<Metadata id %d name %r>' % (self.id, self.data_key)
