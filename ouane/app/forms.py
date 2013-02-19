from flask.ext.wtf import Form, TextField, PasswordField
from flask.ext.wtf import Required

class LoginForm(Form):
    login = TextField('Login', validators = [Required()])
    password = PasswordField('Password', validators = [Required()])
