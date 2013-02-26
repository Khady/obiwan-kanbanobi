from network import Network
from message_pb2 import Msg
import message_pb2
import threading
from dbUtils import Cards, Columns, Users, Projects, Comments, Metadata
from app import app, db

class Api(threading.Thread):
    def __init__(self, host, port):
        threading.Thread.__init__(self)
        self.network = Network(host, port)
        db.create_all()

    def getAllProjetList(self, author_id, session_id):
        msg = Msg()
        msg.target = message_pb2.PROJECTS
        msg.command = message_pb2.GET
        msg.author_id = author_id
        msg.session_id = session_id
        self.network.setWriteStack(msg.SerializeToString())

    def getColumnsByProjectId(self, author_id, session_id, project_id):
        msg = Msg()
        msg.target = message_pb2.COLUMNS
        msg.command = message_pb2.GET
        msg.author_id = author_id
        msg.session_id = session_id
        msg.columns.project_id = project_id
        self.network.setWriteStack(msg.SerializeToString())

    def getAllColumns(self, author_id, session_id):
        msg = Msg()
        msg.target = message_pb2.COLUMNS
        msg.command = message_pb2.GET
        msg.author_id = author_id
        msg.session_id = session_id
        self.network.setWriteStack(msg.SerializeToString())

    def getAllCards(self, author_id, session_id):
        msg = Msg()
        msg.target = message_pb2.CARDS
        msg.command = message_pb2.GET
        msg.author_id = author_id
        msg.session_id = session_id
        self.network.setWriteStack(msg.SerializeToString())

    def getCardsByProjectID(self, author_id, session_id, project_id):
        msg = Msg()
        msg.target = message_pb2.CARDS
        msg.command = message_pb2.GET
        msg.author_id = author_id
        msg.session_id = session_id
        msg.cards.project_id = project_id
        self.network.setWriteStack(msg.SerializeToString())

    def getAllUsers(self, author_id, session_id):
        msg = Msg()
        msg.target = message_pb2.USERS
        msg.command = message_pb2.GET
        msg.session_id = session_id
        msg.author_id = author_id
        self.network.setWriteStack(msg.SerializeToString())

    def getUserById(self, author_id, session_id, user_id):
        msg = Msg()
        msg.target = message_pb2.USERS
        msg.command = message_pb2.GET
        msg.author_id = author_id
        msg.session_id = session_id
        msg.users.id = user_id
        self.network.setWriteStack(msg.SerializeToString())
        
    def createColumns(self, author_id, session_id, project_id, id = 0, name = "", desc = "", tags = [], scripts_ids = [], write = []):
        msg = Msg()
        msg.target = message_pb2.COLUMNS
        msg.command = message_pb2.CREATE
        msg.author_id = author_id
        msg.session_id = session_id
        msg.columns.project_id = project_id
        msg.columns.id = id
        msg.columns.name = name
        for elem in write:
            msg.columns.write.add(elem)
        self.network.setWriteStack(msg.SerializeToString())

    def sendLogin(self, login, password):
        msg = Msg()
        msg.author_id = 0
        msg.session_id = ""
        msg.target = message_pb2.IDENT
        msg.command = message_pb2.CONNECT
        msg.ident.login = login
        msg.ident.password = password
        self.network.setWriteStack(msg.SerializeToString())

    def run(self):
        while 1:
            self.network.run()
            if len(self.network.getReadedStack()) != 0:
                msg = Msg()
                data = self.network.getReadedMessage()
                if (data != ""):
                    print ">>>>>>" + data
                msg.ParseFromString(data)
                if (msg.target == message_pb2.CARDS):
                    c = Card(msg.cards.id, msg.cards.name, msg.cards.column_id, msg.cards.project_id, msg.cards.tags,
                             msg.cards.user_id, msg.cards.scripts_id, msg.cards.write)
                    db.session.add(c)
                    db.session.commit()
                if (msg.target == message_pb2.COLUMNS):
                    c = Columns(msg.columns.id, msg.cards.name, msg.columns.column_id, msg.columns.project_id, msg.columns.tags,
                                msg.columns.scripts_id, msg.columns.write)
                    db.session.add(c)
                    db.session.commit()
                if (msg.target == message_pb2.IDENT):
                    print msg.command
