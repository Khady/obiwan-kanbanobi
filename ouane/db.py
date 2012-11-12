class Db:
    def __init__(self):
        self.users = {}
        self.projectsById = {}
        self.cardsByID = {}
        self.columnsById = {}
        self.commentById = {}

    def addUser(self, id, name):
        self.users[id] = name

    def addProjectsById(self, id, name, admins_id=[], read=[]):
        self.projectsById[id] = {'name' : name, 'admins_id' : admins_id, 'read' : read}
        
    def addColumnsById(self, id, name, desc="", tags="", scripts_ids="", write=[]):
        self.columnsById[id] = {'name' : name, 'desc' : desc, "tags" : tags, 'scripts_ids' : scripts_ids, 'write': write}

    def getAllUser(self):
        return self.users

    def getAllProjectsById(self):
        return self.projectsById

    def getAllColumnsById(self):
        return self.columnsById
