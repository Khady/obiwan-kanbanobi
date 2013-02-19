from flask import render_template, flash, redirect, g
from app  import app, a
from forms import LoginForm

# index view function suppressed for brevity

@app.route('/login', methods = ['GET', 'POST'])
def login():
    form = LoginForm()
    if form.validate_on_submit():
        print form.login.data
        print form.password.data
        a.sendLogin(form.login.data, form.password.data)
        return redirect('index')
    return render_template('login.html', 
        title = 'Sign In',
        form = form)

@app.route("/")
@app.route("/index")
def index():
    try:
        data = Users.query.all()
    except:
        data = []
    return render_template('index.html', data=data)
