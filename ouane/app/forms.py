from flask.ext.wtf import Form, TextField, PasswordField, TextAreaField, HiddenField, SubmitField
from flask.ext.wtf import Required, validators

class LoginForm(Form):
    login = TextField('Login', validators = [Required()])
    password = PasswordField('Password', validators = [Required()])

class AddProjectForm(Form):
    name = TextField('Name', validators = [Required()])
    description = TextAreaField('Description', validators = [Required()])
    
class AddUserForm(Form):
    login = TextField('Login', [validators.Length(min=4, max=25)])
    email = TextField('Email Address', [validators.Required()])
    password = PasswordField('New Password', [
            validators.Required(),
            validators.EqualTo('confirm', message='Passwords must match')
            ])
    confirm = PasswordField('Repeat Password')
    
class AddColumnForm(Form):
    name = TextField('Name', [validators.Length(min=1)])
    description = TextAreaField('Description', [validators.Length(min=1)])
    submit = SubmitField('Submit',  validators = [Required()])
class AddCardForm(Form):
    name = TextField('Name', [validators.Length(min=1)])
    description = TextAreaField('Description', [validators.Length(min=1)])
    idColumn = HiddenField('IdColumn', validators = [Required()])
    submit = SubmitField('Submit',  validators = [Required()])
