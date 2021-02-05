from flask import Blueprint
api = Blueprint('api', __name__)
from app.apiv1 import auth, Alarm, Alarmtype, Bug, Bugtype, Device, Devicestatu, machine, Message, Plan, Project, public, Valve, Valvetime, Valvetype, Worktime