import json
from flask import render_template, flash, redirect, g, request, url_for, session, Response
from app  import app, a, event_stream, red
from forms import LoginForm, AddProjectForm, AddUserForm
from functools import wraps
from dbUtils import Projects, Columns

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
        a.sendLogin(form.login.data, form.password.data)
        return redirect(url_for('checklogin', name = form.login.data))
    return render_template('login.html', 
        title = 'Sign In',
        form = form)

@app.route('/checklogin', methods = ['GET', 'POST'])
@app.route('/checklogin/<name>', methods = ['GET', 'POST'])
def checklogin(name = None):
    if name != None and request.args.get('login') == None:
        return render_template('checklogin.html', name = name)
    else:
        if a.checkIfConnected(request.args.get('login')) == True and name == request.args.get('login'):
            session['logged_in'] = True
            session['user_login'] = request.args.get('login')
            return "OK"
        else:
            return "KO"

@app.route('/project', methods = ['GET', 'POST'])
@app.route("/project/<int:id>", methods = ['GET', 'POST'])
@login_required
def project(id = 0):
    connection = a.getUserConnectionData(session['user_login'])
    a.getColumnsByProjectId(connection['author_id'], connection['session_id'], id)
    data = Columns.query.filter_by(project_id = id).order_by(Columns.id).all()
    print data
    return render_template('project.html')

@app.route("/", methods = ['GET', 'POST'])
@app.route("/index", methods = ['GET', 'POST'])
@login_required
def index():
    data = Projects.query.order_by(Projects.id).all()
    connection = a.getUserConnectionData(session['user_login'])
    form = AddProjectForm()
    a.getAllProjetList(connection['author_id'], connection['session_id'], connection['author_id'])
    if form.validate_on_submit():
        a.createProject(connection['author_id'], connection['session_id'], form.name.data, form.description.data)
    return render_template('index.html', data = data, form=form)

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

@app.route('/stream')
@login_required
def stream():
    return Response(event_stream(),
                    mimetype="text/event-stream")
