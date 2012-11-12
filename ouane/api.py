from network import Network
from message_pb2 import message

class Api:
    def __init__(self, host, port):
        self.network = Network(host, port)

    def getProjetList(self, author_id, session_id, name):
        msg = message()
        msg.target = message.TARGET.PROJECTS
        msg.command = message.CMD.GET
        msg.author_id = author_id
        if name != "":
            msg.projects.name = name
        self.network.send(msg.SerializeToString())

    def getColumnsByProjectId(self, author_id, session_id, project_id):
        msg = message()
        msg.target = message.TARGET.COLUMNS
        msg.command = message.CMD.GET
        msg.author_id = author_id
        if name != "":
            msg.columns.project_id = project_id
        self.network.send(msg.SerializeToString())

    def getCardsByProjectID(self, author_id, session_id, project_id):
        msg = message()
        msg.target = message.TARGET.CARDS
        msg.command = message.CMD.GET
        msg.author_id = author_id
        if name != "":
            msg.cards.project_id = project_id
        self.network.send(msg.SerializeToString())

    def getUserById(self, author_id, session_id, user_id):
        msg = message()
        msg.target = message.TARGET.USERS
        msg.command = message.CMD.GET
        msg.author_id = author_id
        if name != "":
            msg.users.id = user_id
        self.network.send(msg.SerializeToString())
        
