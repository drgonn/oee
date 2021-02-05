from flask import Blueprint
api = Blueprint('api', __name__)
from app.apiv1 import auth, Alarm, Alarmtype, Bug, Bugtype, Device, Plan, Project, public, Valve, Valvetime, Valvetype, Worktime,Devicestatu