from network import Network
from message_pb2 import message

class Api:
    def __init__(self, host, port):
        self.network = Network(host, port)

    def getAllProjetList(self, author_id, session_id):
        msg = message()
        msg.target = message.TARGET.PROJECTS
        msg.command = message.CMD.GET
        msg.author_id = author_id
        self.network.send(msg.SerializeToString())

    def getColumnsByProjectId(self, author_id, session_id, project_id):
        msg = message()
        msg.target = message.TARGET.COLUMNS
        msg.command = message.CMD.GET
        msg.author_id = author_id
        msg.columns.project_id = project_id
        self.network.send(msg.SerializeToString())

    def getAllColumns(self, author_id, session_id):
        msg = message()
        msg.target = message.TARGET.COLUMNS
        msg.command = message.CMD.GET
        msg.author_id = author_id
        self.network.send(msg.SerializeToString())

    def getAllCards(self, author_id, session_id):
        msg = message()
        msg.target = message.TARGET.CARDS
        msg.command = message.CMD.GET
        msg.author_id = author_id
        self.network.send(msg.SerializeToString())

    def getCardsByProjectID(self, author_id, session_id, project_id):
        msg = message()
        msg.target = message.TARGET.CARDS
        msg.command = message.CMD.GET
        msg.author_id = author_id
        msg.cards.project_id = project_id
        self.network.send(msg.SerializeToString())

    def getAllUsers(self, author_id, session_id):
        msg = message()
        msg.target = message.TARGET.USERS
        msg.command = message.CMD.GET
        msg.author_id = author_id
        self.network.send(msg.SerializeToString())

    def getUserById(self, author_id, session_id, user_id):
        msg = message()
        msg.target = message.TARGET.USERS
        msg.command = message.CMD.GET
        msg.author_id = author_id
        msg.users.id = user_id
        self.network.send(msg.SerializeToString())
        
