from flask.ext.wtf import Form, TextField, PasswordField, TextAreaField
from flask.ext.wtf import Required

class LoginForm(Form):
    login = TextField('Login', validators = [Required()])
    password = PasswordField('Password', validators = [Required()])

class AddProjectForm(Form):
    name = TextField('Name', validators = [Required()])
    description = TextAreaField('Description', validators = [Required()])
    
