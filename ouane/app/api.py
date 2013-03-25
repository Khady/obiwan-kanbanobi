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
        self.userLogin = {}

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
        msg.users.name = ""
        msg.users.admin = False
        self.network.setWriteStack(msg.SerializeToString())

    def getUserById(self, author_id, session_id, user_id):
        msg = Msg()
        msg.target = message_pb2.USERS
        msg.command = message_pb2.GET
        msg.author_id = author_id
        msg.session_id = session_id
        msg.users.id = user_id
        msg.users.name = ""
        msg.users.admin = False
        self.network.setWriteStack(msg.SerializeToString())
        
    def createColumns(self, author_id, session_id, project_id, name = "", desc = "", tags = [], scripts_ids = [], write = []):
        msg = Msg()
        msg.target = message_pb2.COLUMNS
        msg.command = message_pb2.CREATE
        msg.author_id = author_id
        msg.session_id = session_id
        msg.columns.project_id = project_id
        msg.columns.id = 0
        msg.columns.name = name
        for elem in write:
            msg.columns.write.add(elem)
        self.network.setWriteStack(msg.SerializeToString())

    def createProject(self, author_id, session_id, name = "", content = "", read = []):
        msg = Msg()
        msg.target = message_pb2.PROJECTS
        msg.command = message_pb2.CREATE
        msg.author_id = author_id
        msg.session_id = session_id
        msg.projects.admins_id.append(author_id)
        msg.projects.name = name
        msg.projects.content = content
        msg.projects.id = 0
        for elem in read:
            msg.projects.read.add(elem)
        self.network.setWriteStack(msg.SerializeToString())

    def createUser(self, author_id, session_id, login, email, password, admin = False):
        msg = Msg()
        msg.target = message_pb2.PROJECTS
        msg.command = message_pb2.CREATE
        msg.author_id = author_id
        msg.session_id = session_id
        msg.users.id = 0
        msg.users.name = login
        msg.users.password = password
        msg.users.mail = email
        msg.users.admin = admin
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

    def checkIfConnected(self, name):
        if name in self.userLogin:
            return True
        return False

    def getUserConnectionData(self, name):
        return self.userLogin[name]

    def run(self):
        while 1:
            self.network.run()
            if len(self.network.getReadedStack()) != 0:
                msg = Msg()
                data = self.network.getReadedMessage()
                # if (data != ""):
                #     print ">>>>>>" + data
                msg.ParseFromString(data)
                if (msg.target == message_pb2.CARDS):
                    c = Card(msg.cards.id, msg.cards.name, msg.cards.column_id, msg.cards.project_id, msg.cards.tags,
                             msg.cards.user_id, msg.cards.scripts_id, msg.cards.write)
                    db.session.add(c)
                    db.session.commit()
                if (msg.target == message_pb2.COLUMNS):
                    c = Columns(msg.columns.id, msg.columns.name, msg.columns.column_id, msg.columns.project_id, msg.columns.tags,
                                msg.columns.scripts_id, msg.columns.write)
                    db.session.add(c)
                    db.session.commit()
                if (msg.target == message_pb2.IDENT):
                    # print msg.target
                    # print msg.command
                    # print msg.author_id
                    # print msg.session_id
                    # print msg.ident.login
                    user = {"author_id": msg.author_id, "session_id": msg.session_id}
                    self.getUserById(msg.author_id, msg.session_id, msg.author_id)
                    self.userLogin[msg.ident.login] = user
                if (msg.target == message_pb2.PROJECTS):
                    p = Columns(msg.projects.id, msg.projects.name, msg.projects.admin_id, msg.projects.content, msg.projects.read)
                    db.session.add(c)
                    db.session.commit()
                if (msg.target == message_pb2.ERROR):
                    print "ERROR"
                if (msg.target == message_pb2.USERS):
                    print "USERS"
