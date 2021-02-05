import os

import redis
from flask import Flask
from flask_sqlalchemy import SQLAlchemy
from flask_cors import CORS
from flask_mail import Mail
from celery import Celery
from flasgger import Swagger
from flask_admin import Admin, BaseView, expose

from config import config

db = SQLAlchemy()
swagger = Swagger()
celery = Celery()
mail = Mail()
from app.admin import admin
def create_app(config_name = "default"):
	app = Flask(__name__, static_url_path="")
	app.config.from_object(config[config_name])
	config[config_name].init_app(app)
	mail.init_app(app)
	db.init_app(app)
	swagger.init_app(app)
	CORS(app, supports_credentials=True)
	pool = redis.ConnectionPool(host = app.config["REDIS_HOST"], port = app.config["REDIS_PORT"], decode_responses = True)
	app.sredis = redis.StrictRedis(connection_pool = pool)
	app.sredisPipe = app.sredis.pipeline(transaction = True)
	celery.conf.update(app.config)
	admin.init_app(app)
	from app.apiv1 import api as api_blueprint
	app.register_blueprint(api_blueprint, url_prefix="/api/v1/oee")
	return app
