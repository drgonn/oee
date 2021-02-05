import os

import redis
from flask import Flask
from flask_sqlalchemy import SQLAlchemy
from flask_cors import CORS
from flask_mail import Mail
from celery import Celery
from flasgger import Swagger

from config import config

db = SQLAlchemy()
swagger = Swagger()
celery = Celery()
mail = Mail()
def create_app(config_name = "default"):
	app = Flask(__name__, static_url_path="")
	app.config.from_object(config[config_name])
	config[config_name].init_app(app)
	mail.init_app(app)
	db.init_app(app)
	swagger.init_app(app)
	CORS(app, supports_credentials=True)


	return app
