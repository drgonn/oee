from flask import request, jsonify, current_app, g
from app import db
from datetime import datetime,date
from app.tools import utc_switch

class User(db.Model):
	__tablename__='users'
	id = db.Column(db.Integer, primary_key=True)
	uid = db.Column(db.String(64), unique=True, index=True, nullable=False)
	name = db.Column(db.String(64))
	createDate = db.Column(db.DateTime, default=datetime.utcnow)
	
	def to_json(self):
		return{
			'id':self.id,
			'uid': self.uid,
			'name': self.name,
			'createDate': utc_switch(self.createDate),
		}

	def __repr__(self):
		return '<User %r>' % self.name

class Device(db.Model):
	__tablename__='devices'
	id = db.Column(db.Integer, primary_key=True)
	sn = db.Column(db.String(16), index=True)
	name = db.Column(db.String(64))
	ip = db.Column(db.String(64))
	img = db.Column(db.String(256))
	type = db.Column(db.String(64))
	
	def to_json(self):
		static_host = current_app.config['STATIC_HOST']
		return{
			'id':self.id,
			'sn': self.sn,
			'name': self.name,
			'ip': self.ip,
			'img_url': f"{static_host}/deviceimg/{self.id}/"+self.img if self.img else None,
			'img': self.img,
			'type': self.type,
		}

	def __repr__(self):
		return '<Device %r>' % self.name

class Worktime(db.Model):
	__tablename__='worktimes'
	id = db.Column(db.Integer, primary_key=True)
	start_time = db.Column(db.DateTime)
	end_time = db.Column(db.DateTime)
	seconds = db.Column(db.Integer)
	type = db.Column(db.Integer)
	amount = db.Column(db.Integer)
	good = db.Column(db.Integer)
	glue = db.Column(db.Float)
	device_id = db.Column(db.Integer, db.ForeignKey('devices.id'))
	device = db.relationship('Device', backref=db.backref('worktimes', lazy='dynamic'))
	
	def to_json(self):
		return{
			'id':self.id,
			'start_time': utc_switch(self.start_time)[11:],
			'end_time': utc_switch(self.end_time),
			'seconds': self.seconds,
			'type': self.type,
			'amount': self.amount,
			'good': self.good,
			'glue': self.glue,
			'device_name' : self.device.name,
		}

	def __repr__(self):
		return '<Worktime %r>' % self.id

class Devicestatu(db.Model):
	__tablename__='devicestatus'
	id = db.Column(db.Integer, primary_key=True)
	start_time = db.Column(db.DateTime)
	status = db.Column(db.Integer)
	device_id = db.Column(db.Integer, db.ForeignKey('devices.id'))
	device = db.relationship('Device', backref=db.backref('devicestatus', lazy='dynamic'))
	
	def to_json(self):
		return{
			'id':self.id,
			'start_time': utc_switch(self.start_time),
			'status': self.status,
			'device_name' : self.device.name,
		}

	def __repr__(self):
		return '<Devicestatu %r>' % self.id

class Valvetime(db.Model):
	__tablename__='valvetimes'
	id = db.Column(db.Integer, primary_key=True)
	start_time = db.Column(db.DateTime)
	end_time = db.Column(db.DateTime)
	seconds = db.Column(db.Float)
	volt = db.Column(db.Float)
	amount = db.Column(db.Integer)
	good = db.Column(db.Integer)
	glue = db.Column(db.Float)
	device_id = db.Column(db.Integer, db.ForeignKey('devices.id'))
	device = db.relationship('Device', backref=db.backref('valvetimes', lazy='dynamic'))
	valve_id = db.Column(db.Integer, db.ForeignKey('valves.id'))
	valve = db.relationship('Valve', backref=db.backref('valvetimes', lazy='dynamic'))
	
	def to_json(self):
		return{
			'id':self.id,
			'start_time': utc_switch(self.start_time),
			'end_time': utc_switch(self.end_time),
			'seconds': self.seconds,
			'volt': self.volt,
			'amount': self.amount,
			'good': self.good,
			'glue': self.glue,
			'device_name' : self.device.name,
			'valve_name' : self.valve.name,
			'valve_sn' : self.valve.sn,
		}

	def __repr__(self):
		return '<Valvetime %r>' % self.id

class Valvetype(db.Model):
	__tablename__='valvetypes'
	id = db.Column(db.Integer, primary_key=True)
	name = db.Column(db.String(64))
	
	def to_json(self):
		return{
			'id':self.id,
			'name': self.name,
		}

	def __repr__(self):
		return '<Valvetype %r>' % self.id

class Valve(db.Model):
	__tablename__='valves'
	id = db.Column(db.Integer, primary_key=True)
	sn = db.Column(db.String(16), index=True)
	name = db.Column(db.String(64))
	valvetype_id = db.Column(db.Integer, db.ForeignKey('valvetypes.id'))
	valvetype = db.relationship('Valvetype', backref=db.backref('valves', lazy='dynamic'))
	device_id = db.Column(db.Integer, db.ForeignKey('devices.id'))
	device = db.relationship('Device', backref=db.backref('valves', lazy='dynamic'))
	
	def to_json(self):
		return{
			'id':self.id,
			'sn': self.sn,
			'name': self.name,
			'valvetype_name' : self.valvetype.name,
			'device_name' : self.device.name,
		}

	def __repr__(self):
		return '<Valve %r>' % self.id

class Bug(db.Model):
	__tablename__='bugs'
	id = db.Column(db.Integer, primary_key=True)
	reason = db.Column(db.Text)
	start_time = db.Column(db.DateTime)
	end_time = db.Column(db.DateTime)
	device_id = db.Column(db.Integer, db.ForeignKey('devices.id'))
	device = db.relationship('Device', backref=db.backref('bugs', lazy='dynamic'))
	bugtype_id = db.Column(db.Integer, db.ForeignKey('bugtypes.id'))
	bugtype = db.relationship('Bugtype', backref=db.backref('bugs', lazy='dynamic'))
	
	def to_json(self):
		return{
			'id':self.id,
			'reason': self.reason,
			'start_time': utc_switch(self.start_time),
			'end_time': utc_switch(self.end_time),
			'device_name' : self.device.name,
			'bugtype_name' : self.bugtype.name,
		}

	def __repr__(self):
		return '<Bug %r>' % self.id

class Bugtype(db.Model):
	__tablename__='bugtypes'
	id = db.Column(db.Integer, primary_key=True)
	name = db.Column(db.String(64))
	sn = db.Column(db.String(64))
	about = db.Column(db.Text)
	
	def to_json(self):
		return{
			'id':self.id,
			'name': self.name,
			'sn': self.sn,
			'about': self.about,
		}

	def __repr__(self):
		return '<Bugtype %r>' % self.id

class Alarmtype(db.Model):
	__tablename__='alarmtypes'
	id = db.Column(db.Integer, primary_key=True)
	code = db.Column(db.String(64))
	mean = db.Column(db.Text)
	cause = db.Column(db.Text)
	solution = db.Column(db.Text)
	
	def to_json(self):
		return{
			'id':self.id,
			'code': self.code,
			'mean': self.mean,
			'cause': self.cause,
			'solution': self.solution,
		}

	def __repr__(self):
		return '<Alarmtype %r>' % self.code

class Alarm(db.Model):
	__tablename__='alarms'
	id = db.Column(db.Integer, primary_key=True)
	create_time = db.Column(db.DateTime, default=datetime.now)
	alarmtype_id = db.Column(db.Integer, db.ForeignKey('alarmtypes.id'))
	alarmtype = db.relationship('Alarmtype', backref=db.backref('alarms', lazy='dynamic'))
	valve_id = db.Column(db.Integer, db.ForeignKey('valves.id'))
	valve = db.relationship('Valve', backref=db.backref('alarms', lazy='dynamic'))
	device_id = db.Column(db.Integer, db.ForeignKey('devices.id'))
	device = db.relationship('Device', backref=db.backref('alarms', lazy='dynamic'))
	
	def to_json(self):
		return{
			'id':self.id,
			'create_time': utc_switch(self.create_time),
			'alarmtype_code' : self.alarmtype.code,
			'alarmtype_mean' : self.alarmtype.mean,
			'alarmtype_cause' : self.alarmtype.cause,
			'alarmtype_solution' : self.alarmtype.solution,
			'valve_name' : self.valve.name,
			'valve_id' : self.valve.id,
		}

	def __repr__(self):
		return '<Alarm %r>' % self.id

class Project(db.Model):
	__tablename__='projects'
	id = db.Column(db.Integer, primary_key=True)
	name = db.Column(db.String(256))
	
	def to_json(self):
		return{
			'id':self.id,
			'name': self.name,
		}

	def __repr__(self):
		return '<Project %r>' % self.name

class Plan(db.Model):
	__tablename__='plans'
	id = db.Column(db.Integer, primary_key=True)
	name = db.Column(db.String(256))
	week = db.Column(db.Integer)
	current = db.Column(db.Text)
	follow = db.Column(db.Text)
	time = db.Column(db.DateTime, default=datetime.utcnow)
	project_id = db.Column(db.Integer, db.ForeignKey('projects.id'))
	project = db.relationship('Project', backref=db.backref('plans', lazy='dynamic'))
	user_id = db.Column(db.Integer, db.ForeignKey('users.id'))
	user = db.relationship('User', backref=db.backref('plans', lazy='dynamic'))
	
	def to_json(self):
		return{
			'id':self.id,
			'name': self.name,
			'week': self.week,
			'current': self.current,
			'follow': self.follow,
			'time': utc_switch(self.time),
			'project_name' : self.project.name,
			'user_name' : self.user.name,
		}

	def __repr__(self):
		return '<Plan %r>' % self.id
