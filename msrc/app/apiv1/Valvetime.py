
import json
import logging
import math
import os
import shutil

from app import db
from app.apiv1 import api
from flask import request, jsonify, current_app, g
from sqlalchemy import func

from app.models import Valvetime,Device,Valve

@api.route('/valvetime/<int:id>', methods=['GET'])
def get_valvetime(id):
	valvetime = Valvetime.query.get_or_404(id)

	return jsonify({'success':True,
                    'error_code':0,
                    'records':valvetime.to_json(),
                    })

@api.route('/valvetime', methods=['POST'])
def create_valvetime():
	print(request.json)
	start_time = request.json.get('start_time')
	end_time = request.json.get('end_time')
	seconds = request.json.get('seconds')
	volt = request.json.get('volt')
	amount = request.json.get('amount')
	good = request.json.get('good')
	glue = request.json.get('glue')

	device_id = request.json.get('device_id')
	device = Device.query.filter_by(id=device_id).first()
 
	if device is None:
		return jsonify({'success':False,'error_code':-1,'errmsg':'device_id不存在'})	

	valve_id = request.json.get('valve_id')
	valve = Valve.query.filter_by(id=valve_id).first()
 
	if valve is None:
		return jsonify({'success':False,'error_code':-1,'errmsg':'valve_id不存在'})	

	valvetime = Valvetime(start_time=start_time,end_time=end_time,seconds=seconds,volt=volt,amount=amount,good=good,glue=glue,device_id=device.id,valve_id=valve.id,)

	db.session.add(valvetime)
	try:
		db.session.commit()
	except Exception as e:
		db.session.rollback()
		logging.error(f'添加数据库发生错误,已经回退:{e}')
		return jsonify({'success': False, 'error_code': -123, 'errmsg': '数据库插入错误，请查看日志'})

	return jsonify({'success':True,
                    'error_code':0,
                    })

@api.route('/valvetime/<int:id>', methods=['PUT'])
def modify_valvetime(id):
	print('put json:',request.json)
	valvetime = Valvetime.query.get_or_404(id)
	start_time = request.json.get('start_time')
	end_time = request.json.get('end_time')
	seconds = request.json.get('seconds')
	volt = request.json.get('volt')
	amount = request.json.get('amount')
	good = request.json.get('good')
	glue = request.json.get('glue')
	device_id = request.json.get('device_id')
	device = Device.query.filter_by(id=device_id).first()
	if device is None:
		return jsonify({'success':False,'error_code':-1,'errmsg':'device_id不存在'})	
	valve_id = request.json.get('valve_id')
	valve = Valve.query.filter_by(id=valve_id).first()
	if valve is None:
		return jsonify({'success':False,'error_code':-1,'errmsg':'valve_id不存在'})	
	valvetime.start_time = start_time or valvetime.start_time
	valvetime.end_time = end_time or valvetime.end_time
	valvetime.seconds = seconds or valvetime.seconds
	valvetime.volt = volt or valvetime.volt
	valvetime.amount = amount or valvetime.amount
	valvetime.good = good or valvetime.good
	valvetime.glue = glue or valvetime.glue
	valvetime.device_id = device.id
	valvetime.valve_id = valve.id
	db.session.add(valvetime)

	try:
		db.session.commit()
	except Exception as e:
		db.session.rollback()
		logging.error(f'修改数据库发生错误,已经回退:{e}')
	return jsonify({'success':True,
                    'error_code':0,
                    })

@api.route('/valvetime', methods=['DELETE'])
def delete_valvetime():
	print('delete json:',request.json)
	ids = request.json.get('ids')
	for id in ids:
		valvetime = Valvetime.query.get(id)
		if valvetime is None:
			return jsonify({'success': False, 'error_code': -123, 'errmsg': f'删除错误，id： {id} 不存在'})
		db.session.delete(valvetime)

		try:
			db.session.commit()
		except Exception as e:
			db.session.rollback()
			logging.error(f'删除数据库发生错误,已经回退:{e}')
			return jsonify({'success': False, 'error_code': -123, 'errmsg': f'删除数据发生错误， {e} '})

	return jsonify({'success':True,
                'error_code':0,
                })

@api.route('/valvetime/list', methods=['GET'])
def list_valvetime():
	print(request.args)
	sorter = request.args.get('sorter')
	page = int(request.args.get('current', 1))
	pageSize = int(request.args.get('pageSize', current_app.config['PER_PAGE']))
	pageSize = 20 if pageSize < 10 else pageSize
	total_valvetimes = Valvetime.query

	device_id = request.args.get('device_id')
	if device_id is not None:
		device = Device.query.filter_by(id=device_id).first()
		if device is None:
			return jsonify({'success':False,'error_code':-1,'errmsg':'device_id不存在'})
		else:
			total_valvetimes = total_valvetimes.filter_by(device_id=device.id)

	valve_id = request.args.get('valve_id')
	if valve_id is not None:
		valve = Valve.query.filter_by(id=valve_id).first()
		if valve is None:
			return jsonify({'success':False,'error_code':-1,'errmsg':'valve_id不存在'})
		else:
			total_valvetimes = total_valvetimes.filter_by(valve_id=valve.id)
	if sorter:
		sorter = json.loads(sorter)
		pass
	totalcount = total_valvetimes.with_entities(func.count(Valvetime.id)).scalar()
	page = math.ceil(totalcount/pageSize) if  math.ceil(totalcount/pageSize) < page else page
	pagination = total_valvetimes.paginate(page, per_page = pageSize, error_out = False)
	valvetimes = pagination.items

	return jsonify({
                    'success':True,
                    'error_code':0,
                    'total':totalcount,
                    "pageSize" : pageSize,
                    "current" : page,
                    "pagecount": pagination.pages,
                    'data':[valvetime.to_json() for valvetime in valvetimes]
                    })

