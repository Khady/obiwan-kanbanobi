from network import Network
from message_pb2 import Msg
import message_pb2

class Api:
    def __init__(self, host, port):
        self.network = Network(host, port)
        self.network.start()

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
        
    def sendLogin(self, login, password):
        msg = Msg()
        msg.author_id = 0
        msg.session_id = ""
        msg.target = message_pb2.IDENT
        msg.command = message_pb2.CONNECT
        # msg.target = message_pb2._TARGET.values_by_name["IDENT"]
        # msg.command = message_pb2._CMD.values_by_name["CONNECT"]
        msg.ident.login = login
        msg.ident.password = password
        self.network.setWriteStack(msg.SerializeToString())
