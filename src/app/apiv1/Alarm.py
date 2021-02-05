
import json
import logging
import math
import os
import shutil

from app import db
from app.apiv1 import api
from flask import request, jsonify, current_app, g
from sqlalchemy import func

from app.models import Alarm,Alarmtype,Valve,Device

@api.route('/alarm/<int:id>', methods=['GET'])
def get_alarm(id):
	alarm = Alarm.query.get_or_404(id)

	return jsonify({'success':True,
                    'error_code':0,
                    'records':alarm.to_json(),
                    })

@api.route('/alarm', methods=['POST'])
def create_alarm():
	print(request.json)

	alarmtype_id = request.json.get('alarmtype_id')
	alarmtype = Alarmtype.query.filter_by(id=alarmtype_id).first()
 
	if alarmtype is None:
		return jsonify({'success':False,'error_code':-1,'errmsg':'alarmtype_id不存在'})	

	valve_id = request.json.get('valve_id')
	valve = Valve.query.filter_by(id=valve_id).first()
 
	if valve is None:
		return jsonify({'success':False,'error_code':-1,'errmsg':'valve_id不存在'})	

	device_id = request.json.get('device_id')
	device = Device.query.filter_by(id=device_id).first()
 
	if device is None:
		return jsonify({'success':False,'error_code':-1,'errmsg':'device_id不存在'})	

	alarm = Alarm(alarmtype_id=alarmtype.id,valve_id=valve.id,device_id=device.id,)

	db.session.add(alarm)
	try:
		db.session.commit()
	except Exception as e:
		db.session.rollback()
		logging.error(f'添加数据库发生错误,已经回退:{e}')
		return jsonify({'success': False, 'error_code': -123, 'errmsg': '数据库插入错误，请查看日志'})

	return jsonify({'success':True,
                    'error_code':0,
                    })

@api.route('/alarm/<int:id>', methods=['PUT'])
def modify_alarm(id):
	print('put json:',request.json)
	alarm = Alarm.query.get_or_404(id)
	alarmtype_id = request.json.get('alarmtype_id')
	alarmtype = Alarmtype.query.filter_by(id=alarmtype_id).first()
	if alarmtype is None:
		return jsonify({'success':False,'error_code':-1,'errmsg':'alarmtype_id不存在'})	
	valve_id = request.json.get('valve_id')
	valve = Valve.query.filter_by(id=valve_id).first()
	if valve is None:
		return jsonify({'success':False,'error_code':-1,'errmsg':'valve_id不存在'})	
	device_id = request.json.get('device_id')
	device = Device.query.filter_by(id=device_id).first()
	if device is None:
		return jsonify({'success':False,'error_code':-1,'errmsg':'device_id不存在'})	
	alarm.alarmtype_id = alarmtype.id
	alarm.valve_id = valve.id
	alarm.device_id = device.id
	db.session.add(alarm)

	try:
		db.session.commit()
	except Exception as e:
		db.session.rollback()
		logging.error(f'修改数据库发生错误,已经回退:{e}')
	return jsonify({'success':True,
                    'error_code':0,
                    })

@api.route('/alarm', methods=['DELETE'])
def delete_alarm():
	print('delete json:',request.json)
	ids = request.json.get('ids')
	for id in ids:
		alarm = Alarm.query.get(id)
		if alarm is None:
			return jsonify({'success': False, 'error_code': -123, 'errmsg': f'删除错误，id： {id} 不存在'})
		db.session.delete(alarm)

		try:
			db.session.commit()
		except Exception as e:
			db.session.rollback()
			logging.error(f'删除数据库发生错误,已经回退:{e}')
			return jsonify({'success': False, 'error_code': -123, 'errmsg': f'删除数据发生错误， {e} '})

	return jsonify({'success':True,
                'error_code':0,
                })

@api.route('/alarm/list', methods=['GET'])
def list_alarm():
	print(request.args)
	sorter = request.args.get('sorter')
	page = int(request.args.get('current', 1))
	pageSize = int(request.args.get('pageSize', current_app.config['PER_PAGE']))
	pageSize = 20 if pageSize < 10 else pageSize
	total_alarms = Alarm.query

	alarmtype_id = request.args.get('alarmtype_id')
	if alarmtype_id is not None:
		alarmtype = Alarmtype.query.filter_by(id=alarmtype_id).first()
		if alarmtype is None:
			return jsonify({'success':False,'error_code':-1,'errmsg':'alarmtype_id不存在'})
		else:
			total_alarms = total_alarms.filter_by(alarmtype_id=alarmtype.id)

	valve_id = request.args.get('valve_id')
	if valve_id is not None:
		valve = Valve.query.filter_by(id=valve_id).first()
		if valve is None:
			return jsonify({'success':False,'error_code':-1,'errmsg':'valve_id不存在'})
		else:
			total_alarms = total_alarms.filter_by(valve_id=valve.id)

	device_id = request.args.get('device_id')
	if device_id is not None:
		device = Device.query.filter_by(id=device_id).first()
		if device is None:
			return jsonify({'success':False,'error_code':-1,'errmsg':'device_id不存在'})
		else:
			total_alarms = total_alarms.filter_by(device_id=device.id)
	if sorter:
		sorter = json.loads(sorter)
		pass
	totalcount = total_alarms.with_entities(func.count(Alarm.id)).scalar()
	page = math.ceil(totalcount/pageSize) if  math.ceil(totalcount/pageSize) < page else page
	pagination = total_alarms.paginate(page, per_page = pageSize, error_out = False)
	alarms = pagination.items

	return jsonify({
                    'success':True,
                    'error_code':0,
                    'total':totalcount,
                    "pageSize" : pageSize,
                    "current" : page,
                    "pagecount": pagination.pages,
                    'data':[alarm.to_json() for alarm in alarms]
                    })

