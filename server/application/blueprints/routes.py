from flask import Blueprint, session, current_app, redirect, render_template, request, url_for
from application.database.util import getAllDiscordUsers

web = Blueprint('web', __name__)

@web.route('/')
@web.route('/index')
def index():
    if not session:
        return redirect(url_for('.login'))
    return render_template('index.html', )

@web.route('/login', methods=['GET', 'POST'])
def login():
    error = None
    if session:
        return redirect(url_for(".index"))
    if request.method == "POST":
        token = request.form.get('token')
        if not token:
            error = "Provide token"
        elif token == current_app.config.get('SECRET_TOKEN'):
           session['auth'] = True
           return redirect(url_for(".index"))
        else:
            error = "Invalid token"
    return render_template('login.html', error=error)

@web.route('/logs', methods=['GET'])
def logs():
    if not session:
        return redirect(url_for('.login'))
    discord_users = getAllDiscordUsers()
    return render_template('logs.html', users=discord_users)