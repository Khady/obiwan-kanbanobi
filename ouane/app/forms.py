from flask.ext.wtf import Form, TextField, PasswordField, TextAreaField
from flask.ext.wtf import Required, validators

class LoginForm(Form):
    login = TextField('Login', validators = [Required()])
    password = PasswordField('Password', validators = [Required()])

class AddProjectForm(Form):
    name = TextField('Name', validators = [Required()])
    description = TextAreaField('Description', validators = [Required()])
    
class AddUserForm(Form):
    login = TextField('Login', [validators.Length(min=4, max=25)])
    email = TextField('Email Address', [validators.Length(min=6, max=35)])
    password = PasswordField('New Password', [
        validators.Required(),
        validators.EqualTo('confirm', message='Passwords must match')
    ])
    confirm = PasswordField('Repeat Password')
