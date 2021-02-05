from flask import request, jsonify, current_app, g
from app import db
from datetime import datetime,date
from app.tools import utc_switch





class Statu(db.Model):
	__tablename__='status'
	id = db.Column(db.Integer, primary_key=True)
	start_time = db.Column(db.DateTime,  default=datetime.now)
	status = db.Column(db.Integer)



class Process (db.Model):
	__tablename__='process'
	id = db.Column(db.Integer, primary_key=True)
	ct_time = db.Column(db.DateTime,  default=datetime.now)
	good = db.Column(db.Integer, primary_key=True)
	count = db.Column(db.Integer)

