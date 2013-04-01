import json
from network import Network
from message_pb2 import Msg
import message_pb2
import threading
from app import app, db, red
from dbUtils import Cards, Columns, Users, Projects, Comments, Metadata

class Api(threading.Thread):
    def __init__(self, host, port):
        threading.Thread.__init__(self)
        self.network = Network(host, port)
        db.create_all()
        self.userLogin = {}

    def getAllProjetList(self, author_id, session_id, user_id):
        msg = Msg()
        msg.target = message_pb2.USERS
        msg.command = message_pb2.GETBOARD
        msg.author_id = author_id
        msg.session_id = session_id
        msg.users.id = user_id
        msg.users.name = ""
        msg.users.admin = False
        self.network.setWriteStack(msg.SerializeToString())

    def getProjectById(self, author_id, session_id, project_id):
        msg = Msg()
        msg.target = message_pb2.PROJECTS
        msg.command = message_pb2.GET
        msg.author_id = author_id
        msg.session_id = session_id
        msg.projects.id = project_id
        msg.projects.content = ""
        msg.projects.name = ""
        self.network.setWriteStack(msg.SerializeToString())

    def getColumnsByProjectId(self, author_id, session_id, project_id):
        msg = Msg()
        msg.target = message_pb2.PROJECTS
        msg.command = message_pb2.GETBOARD
        msg.author_id = author_id
        msg.session_id = session_id
        msg.projects.id = project_id
        self.network.setWriteStack(msg.SerializeToString())

    # def getAllColumns(self, author_id, session_id):
    #     msg = Msg()
    #     msg.target = message_pb2.COLUMNS
    #     msg.command = message_pb2.GET
    #     msg.author_id = author_id
    #     msg.session_id = session_id
    #     self.network.setWriteStack(msg.SerializeToString())

    # def getAllCards(self, author_id, session_id):
    #     msg = Msg()
    #     msg.target = message_pb2.CARDS
    #     msg.command = message_pb2.GET
    #     msg.author_id = author_id
    #     msg.session_id = session_id
    #     self.network.setWriteStack(msg.SerializeToString())

    def getCardsByColumnID(self, author_id, session_id, project_id):
        msg = Msg()
        msg.target = message_pb2.COLUMNS
        msg.command = message_pb2.GETBOARD
        msg.author_id = author_id
        msg.session_id = session_id
        msg.columns.id = project_id
        self.network.setWriteStack(msg.SerializeToString())

    # def getAllUsers(self, author_id, session_id):
    #     msg = Msg()
    #     msg.target = message_pb2.USERS
    #     msg.command = message_pb2.GET
    #     msg.session_id = session_id
    #     msg.author_id = author_id
    #     msg.users.name = ""
    #     msg.users.admin = False
    #     self.network.setWriteStack(msg.SerializeToString())

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
        msg.target = message_pb2.USERS
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

    def addNewProjectInDB(self, project):
        readstr = " ".join(project.read)
        adminstr = " ".join(project.admins_id)
        p = Projects.query.get(project.id)
        print p
        if (p == None):
            p = Projects(project.id, project.name, adminstr, project.content, readstr)
            db.session.add(p)
        else:
            p.name = project.name
            p.admins_id = adminstr
            p.content = project.content
            p.read = readstr
        db.session.commit()


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
                    red.publish('ouane', u'CARDS')
                if (msg.target == message_pb2.COLUMNS):
                    c = Columns(msg.columns.id, msg.columns.name, msg.columns.column_id, msg.columns.project_id, msg.columns.tags,
                                msg.columns.scripts_id, msg.columns.write)
                    db.session.add(c)
                    db.session.commit()
                    red.publish('ouane', u'COLUMNS')
                if (msg.target == message_pb2.IDENT):
                    # print msg.target
                    # print msg.command
                    # print msg.author_id
                    # print msg.session_id
                    # print msg.ident.login
                    user = {"author_id": msg.author_id, "session_id": msg.session_id}
                    self.getUserById(msg.author_id, msg.session_id, msg.author_id)
                    self.getAllProjetList(msg.author_id, msg.session_id, msg.author_id)
                    self.userLogin[msg.ident.login] = user
                    red.publish('ouane', u'IDENT')
                if (msg.target == message_pb2.PROJECTS):
                    if (msg.command == message_pb2.GET):
                        self.addNewProjectInDB(msg.projects)
                        project = msg.projects
                        dictproject = {'id' : project.id, 'name' : project.name, 'content' : project.content, 'read' : ' '.join(project.read), 'admins_id' : ' '.join(project.admins_id)}              
                        dictproject['type'] = 'project'
                        red.publish('ouane', json.dumps(dictproject))
                        print "PROJECTS"
                if (msg.target == message_pb2.ERROR):
                    red.publish('ouane', u'ERROR')
                    print "ERROR"
                if (msg.target == message_pb2.USERS):
                    red.publish('ouane', u'USERS')
                    print "USERS"
                    for project in msg.users.userProject:
                        self.addNewProjectInDB(project)
                        dictproject = {'id' : project.id, 'name' : project.name, 'content' : project.content, 'read' : ' '.join(project.read), 'admins_id' : ' '.join(project.admins_id)}              
                        dictproject['type'] = 'project'
                        print dictproject
                        red.publish('ouane', json.dumps(dictproject))
                        print project.id
                        print project.name
                        print project.content
                        print project.read
                        print project.admins_id
                        # self.getProjectById(msg.author_id, msg.session_id, project.id)
