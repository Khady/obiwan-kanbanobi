from flask import render_template, flash, redirect, g, request, url_for, session, Response
from app  import app, a, event_stream
from forms import LoginForm, AddProjectForm, AddUserForm, AddColumnForm, AddCardForm
from functools import wraps
from dbUtils import Projects, Columns, Cards

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
            connection = a.getUserConnectionData(session['user_login'])
            session['author_id'] = connection['author_id']
            session['session_id'] = connection['session_id']
            return "OK"
        else:
            return "KO"

@app.route('/project', methods = ['GET', 'POST'])
@app.route("/project/<int:id>", methods = ['GET', 'POST'])
@login_required
def project(id = 0):
    #a.getColumnsByProjectId(session['author_id'], session['session_id'], id)
    data = Columns.query.filter_by(project_id = id).order_by(Columns.id).all()
    card = {}
    for d in data:
        c = Cards.query.filter_by(column_id = d.id).order_by(Cards.id).all()
        card[d.id] = c
    form = AddColumnForm(prefix="form")
    formCard = AddCardForm(prefix="formCard")
    if form.validate_on_submit() and form.submit.data:
        a.createColumn(session['author_id'], session['session_id'], id, form.name.data, form.description.data)
    if formCard.validate_on_submit() and formCard.submit.data:
        formCard.idColumn.data
        a.createCard(session['author_id'], session['session_id'], id, formCard.name.data, formCard.description.data, int(formCard.idColumn.data))
    return render_template('project.html', columns=data, card=card, form=form, formCard = formCard, id = id)

@app.route("/", methods = ['GET', 'POST'])
@app.route("/index", methods = ['GET', 'POST'])
@login_required
def index():
    d = Projects.query.order_by(Projects.id).all()
    data = []
    for p in d:
        t = p.read.split(" ")
        if len(data) == 0:
            data.append(p)
        elif str(session['author_id']) in t:
            data.append(p)
        elif t.count("0") == len(t):
            data.append(p)
    form = AddProjectForm()
    #a.getAllProjetList(session['author_id'], session['session_id'], session['author_id'])
    if form.validate_on_submit():
        a.createProject(session['author_id'], session['session_id'], form.name.data, form.description.data)
    return render_template('index.html', data = data, form = form)

@app.route("/admin", methods = ['GET', 'POST'])
@login_required
def admin():
    form = AddUserForm()
    if form.validate_on_submit():
        a.createUser(session['author_id'], session['session_id'], form.login.data, form.email.data, form.password.data, False)
    elif request.method == 'POST' and form.validate() == False:
        flash("Error during the user creation!")
    return render_template('admin.html', form=form)

@app.route('/stream')
@login_required
def stream():
    return Response(event_stream(),
                    mimetype="text/event-stream")


@app.route("/modifCard", methods = ['POST'])
@login_required
def modifCard():
    c = Cards.query.filter_by(id = int(request.form['idCard'])).all()
    print request.form['idColumn']
    print request.form['idCard']
    print c
    c = c[0]
    a.modifyCard(session['author_id'], session['session_id'], int(request.form['idCard']), c.project_id, c.name, c.content, int(request.form['idColumn']))
    return "OK"
