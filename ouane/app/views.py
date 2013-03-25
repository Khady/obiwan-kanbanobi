from flask import render_template, flash, redirect, g, request, url_for, session
from app  import app, a#, jug
from forms import LoginForm, AddProjectForm, AddUserForm
from functools import wraps
from dbUtils import Projects

# index view function suppressed for brevity

def login_required(f):
    @wraps(f)
    def decorated_function(*args, **kwargs):
        if 'logged_in' not in  session or a.checkIfConnected(session['user_login']) == False:
            session.pop('logged_in', None)
            return redirect(url_for('login'))
        return f(*args, **kwargs)
    return decorated_function

@app.route('/login', methods = ['GET', 'POST'])
def login():
    if 'logged_in' in session:
        return redirect(url_for('index'))
    form = LoginForm()
    if form.validate_on_submit():
        # print form.login.data
        # print form.password.data
        a.sendLogin(form.login.data, form.password.data)
        return redirect(url_for('checklogin', name = form.login.data))
    return render_template('login.html', 
        title = 'Sign In',
        form = form)

@app.route('/checklogin', methods = ['GET', 'POST'])
@app.route('/checklogin/<name>', methods = ['GET', 'POST'])
def checklogin(name = None):
#    print request.args.get('login')
    if name != None and request.args.get('login') == None:
#        print "test"
        return render_template('checklogin.html', name = name)
    else:
#        print "mdr"
        if a.checkIfConnected(request.args.get('login')) == True and name == request.args.get('login'):
            session['logged_in'] = True
            session['user_login'] = request.args.get('login')
            return "OK"
        else:
            return "KO"

@app.route('/project', methods = ['GET', 'POST'])
@app.route("/project/<name>", methods = ['GET', 'POST'])
@login_required
def project(name = ""):
    return render_template('project.html')

@app.route("/", methods = ['GET', 'POST'])
@app.route("/index", methods = ['GET', 'POST'])
@login_required
def index():
    try:
        data = Projects.query.all()
        print data
    except:
        data = []
    form = AddProjectForm()
    if form.validate_on_submit():
        print form.name.data
        print form.description.data
#        jug.publish('test', form.name.data)
        connection = a.getUserConnectionData(session['user_login'])
        a.createProject(connection['author_id'], connection['session_id'], form.name.data, form.description.data)
    return render_template('index.html', data=data, form=form)

@app.route("/admin", methods = ['GET', 'POST'])
@login_required
def admin():
    form = AddUserForm()
    if form.validate_on_submit():
        connection = a.getUserConnectionData(session['user_login'])
        a.createUser(connection['author_id'], connection['session_id'], form.login.data, form.email.data, form.password.data, False)
    elif request.method == 'POST' and form.validate() == False:
        flash("Error during the user creation!")
    return render_template('admin.html', form=form)
