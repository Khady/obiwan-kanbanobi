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
        msg.projects.name = ""
        msg.projects.content = ""
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

    def getCardsByColumnID(self, author_id, session_id, columns_id, project_id):
        msg = Msg()
        msg.target = message_pb2.COLUMNS
        msg.command = message_pb2.GETBOARD
        msg.author_id = author_id
        msg.session_id = session_id
        msg.columns.id = columns_id
        msg.columns.project_id = project_id
        msg.columns.name = ""
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
        readstr = " ".join([str(r) for r in project.read])
        adminstr = " ".join([str(r) for r in project.admins_id])
        p = Projects.query.get(project.id)
#        print p
        if (p == None):
            p = Projects(project.id, project.name, adminstr, project.content, readstr)
            db.session.add(p)
        else:
            p.name = project.name
            p.admins_id = adminstr
            p.content = project.content
            p.read = readstr
        db.session.commit()

    def addNewColumnInDB(self, column):
        script = " ".join(str(column.scripts_ids))
        writestr = " ".join(str(column.write))
        c = Columns.query.get(column.id)
 #       print c
        if (c == None):
            c = Columns(column.id, column.name, column.project_id, column.desc, ' '.join(column.tags), 0, writestr)
            db.session.add(c)
        else:
            c.name = column.name
            c.content = column.desc
            c.tags = ' '.join(column.tags)
            c.scripts_id = 0
            c.write = writestr
        db.session.commit()


    def addNewCardInDB(self, card):
        script = " ".join(str(card.scripts_ids))
        writestr = " ".join(str(card.write))
        c = Cards.query.get(card.id)
 #       print c
        if (c == None):
            c = Cards(card.id, card.name, card.desc, card.column_id, card.project_id, ' '.join(card.tags), card.user_id, 0, writestr)
            db.session.add(c)
        else:
            c.name = card.name
            c.content = card.desc
            c.column_id = card.column_id
            c.tags = ' '.join(card.tags)
            c.user_id = card.user_id
            c.scripts_id = 0
            c.write = writestr
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
                    # db.session.add(c)
                    # db.session.commit()
                    red.publish('ouane', u'CARDS')
                if (msg.target == message_pb2.COLUMNS):
                    # db.session.add(c)
                    # db.session.commit()
                    # red.publish('ouane', u'COLUMNS')
                    if msg.command == message_pb2.ERROR:
                        print 'ERROR COLUMNS',
                        print msg.error.error_id
                    else:
                        # c = Columns(msg.columns.id, msg.columns.name, msg.columns.desc, msg.columns.project_id, msg.columns.tags,
                        #             msg.columns.scripts_ids, msg.columns.write)
                        columns = msg.columns
                        self.addNewColumnInDB(columns)
                        print "COLUMNS",
                        print [msg.columns.id, msg.columns.name, msg.columns.desc, msg.columns.project_id, msg.columns.tags,
                               msg.columns.scripts_ids, msg.columns.write]
                        for cards in msg.columns.ColumnCards:
                            self.addNewCardInDB(cards)
                            print "CARD",
                            print [cards.id, cards.name, cards.desc, cards.column_id, cards.project_id, cards.tags,
                                   cards.user_id, cards.scripts_ids, cards.write]
                if (msg.target == message_pb2.IDENT):
                    if msg.command == message_pb2.ERROR:
                        print 'ERROR IDENT',
                        print msg.error.error_id
                    else:
                        user = {"author_id": msg.author_id, "session_id": msg.session_id}
                        self.getUserById(msg.author_id, msg.session_id, msg.author_id)
                        self.getAllProjetList(msg.author_id, msg.session_id, msg.author_id)
                        self.userLogin[msg.ident.login] = user
                        #red.publish('ouane', u'IDENT')
                if (msg.target == message_pb2.PROJECTS):
                    if msg.command == message_pb2.ERROR:
                        print 'ERROR PROJECT',
                        print msg.error.error_id
                    else:
                        self.addNewProjectInDB(msg.projects)
                        project = msg.projects
                        dictproject = {'id' : project.id, 'name' : project.name, 'content' : project.content, 'read' : ' '.join([str(r) for r in project.read]), 'admins_id' : ' '.join([str(r) for r in project.admins_id])}
                        dictproject['type'] = 'project'
                        red.publish('ouane', json.dumps(dictproject))
                        for columns in msg.projects.projectColumns:
                            self.addNewColumnInDB(columns)
                            # dictcolumns = {'id' : column.id, 'name' : column.name, 'content' : column.content, 'read' : ' '.join(column.read), 'admins_id' : ' '.join(column.admins_id)}
                            print columns
                            self.getCardsByColumnID(msg.author_id, msg.session_id, columns.id, msg.projects.id)
                            #dictcolumns['type'] = 'column'
                            #red.publish('ouane', json.dumps(dictcolumns))
                if (msg.target == message_pb2.ERROR):
                    red.publish('ouane', u'ERROR')
                    print "ERROR"
                if (msg.target == message_pb2.USERS):
                    for project in msg.users.userProject:
                        self.addNewProjectInDB(project)
                        dictproject = {'id' : project.id, 'name' : project.name, 'content' : project.content, 'read' : ' '.join([str(r) for r in project.read]), 'admins_id' : ' '.join([str(r) for r in project.admins_id])}
                        dictproject['type'] = 'project'
                        red.publish('ouane', json.dumps(dictproject))
                        self.getColumnsByProjectId(msg.author_id, msg.session_id, project.id)
