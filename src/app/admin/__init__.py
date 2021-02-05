
import flask_admin as admin
from .views import *
from app.models import *
admin = admin.Admin(name='Admin', template_mode='bootstrap3')
admin.add_view(ModelView(User, db.session))
admin.add_view(ModelView(Device, db.session))
admin.add_view(ModelView(Worktime, db.session))
admin.add_view(ModelView(Devicestatu, db.session))
admin.add_view(ModelView(Valvetime, db.session))
admin.add_view(ModelView(Valvetype, db.session))
admin.add_view(ModelView(Valve, db.session))
admin.add_view(ModelView(Bug, db.session))
admin.add_view(ModelView(Bugtype, db.session))
admin.add_view(ModelView(Alarmtype, db.session))
admin.add_view(ModelView(Alarm, db.session))
admin.add_view(ModelView(Project, db.session))
admin.add_view(ModelView(Plan, db.session))
admin.add_view(ModelView(Message, db.session))
