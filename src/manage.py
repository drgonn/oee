import os
from flask_script import Manager, Shell
from flask_migrate import Migrate, MigrateCommand
from app import create_app, db, celery
from app.models import  *
app = create_app(os.getenv("FLASK_CONFIG") or "default")
manager = Manager(app)
migrate = Migrate(app, db)
app.app_context().push()
def make_shell_context():
	return dict(app=app,db=db,User = User,Device = Device,Worktime = Worktime,Devicestatu = Devicestatu,Valvetime = Valvetime,Valvetype = Valvetype,Valve = Valve,Bug = Bug,Bugtype = Bugtype,Alarmtype = Alarmtype,Alarm = Alarm,Project = Project,Plan = Plan,Message = Message)
manager.add_command("shell", Shell(make_context=make_shell_context))
manager.add_command("db", MigrateCommand)
@manager.command
def db_init():
	db.create_all()
@manager.command 
def test():
	import unittest
	tests = unittest.TestLoader().discover("tests")
	unittest.TextTestRunner(verbosity=2).run(tests)
if __name__ == "__main__":
	manager.run()
