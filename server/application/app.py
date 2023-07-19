from flask import Flask
from application.blueprints.routes import web
from application.blueprints.api import api

app = Flask(__name__)
app.config.from_object('application.config.Config')
app.register_blueprint(web, url_prefix='/')
app.register_blueprint(api, url_prefix='/api')

print(f"Secret Token: {app.config['SECRET_TOKEN']}")