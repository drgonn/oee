
import json
import logging
import math
import os
import shutil

from app import db
from app.apiv1 import api
from flask import request, jsonify, current_app, g
from sqlalchemy import func

from app.models import Valve,Valvetype,Device

@api.route('/valve/<int:id>', methods=['GET'])
def get_valve(id):
	valve = Valve.query.get_or_404(id)

	return jsonify({'success':True,
                    'error_code':0,
                    'records':valve.to_json(),
                    })

@api.route('/valve', methods=['POST'])
def create_valve():
	print(request.json)
	sn = request.json.get('sn')
	name = request.json.get('name')

	valvetype_id = request.json.get('valvetype_id')
	valvetype = Valvetype.query.filter_by(id=valvetype_id).first()
 
	if valvetype is None:
		return jsonify({'success':False,'error_code':-1,'errmsg':'valvetype_id不存在'})	

	device_id = request.json.get('device_id')
	device = Device.query.filter_by(id=device_id).first()
 
	if device is None:
		return jsonify({'success':False,'error_code':-1,'errmsg':'device_id不存在'})	

	valve = Valve(sn=sn,name=name,valvetype_id=valvetype.id,device_id=device.id,)

	db.session.add(valve)
	try:
		db.session.commit()
	except Exception as e:
		db.session.rollback()
		logging.error(f'添加数据库发生错误,已经回退:{e}')
		return jsonify({'success': False, 'error_code': -123, 'errmsg': '数据库插入错误，请查看日志'})

	return jsonify({'success':True,
                    'error_code':0,
                    })

@api.route('/valve/<int:id>', methods=['PUT'])
def modify_valve(id):
	print('put json:',request.json)
	valve = Valve.query.get_or_404(id)
	sn = request.json.get('sn')
	name = request.json.get('name')
	valvetype_id = request.json.get('valvetype_id')
	valvetype = Valvetype.query.filter_by(id=valvetype_id).first()
	if valvetype is None:
		return jsonify({'success':False,'error_code':-1,'errmsg':'valvetype_id不存在'})	
	device_id = request.json.get('device_id')
	device = Device.query.filter_by(id=device_id).first()
	if device is None:
		return jsonify({'success':False,'error_code':-1,'errmsg':'device_id不存在'})	
	valve.sn = sn or valve.sn
	valve.name = name or valve.name
	valve.valvetype_id = valvetype.id
	valve.device_id = device.id
	db.session.add(valve)

	try:
		db.session.commit()
	except Exception as e:
		db.session.rollback()
		logging.error(f'修改数据库发生错误,已经回退:{e}')
	return jsonify({'success':True,
                    'error_code':0,
                    })

@api.route('/valve', methods=['DELETE'])
def delete_valve():
	print('delete json:',request.json)
	ids = request.json.get('ids')
	for id in ids:
		valve = Valve.query.get(id)
		if valve is None:
			return jsonify({'success': False, 'error_code': -123, 'errmsg': f'删除错误，id： {id} 不存在'})
		if valve.valvetimes.first() is not None:
			return jsonify({'success':False,'error_code':-1,'errmsg':'valve还拥有valvetime，不能删除'})
		if valve.alarms.first() is not None:
			return jsonify({'success':False,'error_code':-1,'errmsg':'valve还拥有alarm，不能删除'})
		db.session.delete(valve)

		try:
			db.session.commit()
		except Exception as e:
			db.session.rollback()
			logging.error(f'删除数据库发生错误,已经回退:{e}')
			return jsonify({'success': False, 'error_code': -123, 'errmsg': f'删除数据发生错误， {e} '})

	return jsonify({'success':True,
                'error_code':0,
                })

@api.route('/valve/list', methods=['GET'])
def list_valve():
	print(request.args)
	sorter = request.args.get('sorter')
	page = int(request.args.get('current', 1))
	pageSize = int(request.args.get('pageSize', current_app.config['PER_PAGE']))
	pageSize = 20 if pageSize < 10 else pageSize
	total_valves = Valve.query

	valvetype_id = request.args.get('valvetype_id')
	if valvetype_id is not None:
		valvetype = Valvetype.query.filter_by(id=valvetype_id).first()
		if valvetype is None:
			return jsonify({'success':False,'error_code':-1,'errmsg':'valvetype_id不存在'})
		else:
			total_valves = total_valves.filter_by(valvetype_id=valvetype.id)

	device_id = request.args.get('device_id')
	if device_id is not None:
		device = Device.query.filter_by(id=device_id).first()
		if device is None:
			return jsonify({'success':False,'error_code':-1,'errmsg':'device_id不存在'})
		else:
			total_valves = total_valves.filter_by(device_id=device.id)
	sn = request.args.get('sn')
	if sn is not None:
		total_valves = total_valves.filter(Valve.sn.ilike(f'%{sn}%'))

	name = request.args.get('name')
	if name is not None:
		total_valves = total_valves.filter(Valve.name.ilike(f'%{name}%'))

	if sorter:
		sorter = json.loads(sorter)
		pass
	totalcount = total_valves.with_entities(func.count(Valve.id)).scalar()
	page = math.ceil(totalcount/pageSize) if  math.ceil(totalcount/pageSize) < page else page
	pagination = total_valves.paginate(page, per_page = pageSize, error_out = False)
	valves = pagination.items

	return jsonify({
                    'success':True,
                    'error_code':0,
                    'total':totalcount,
                    "pageSize" : pageSize,
                    "current" : page,
                    "pagecount": pagination.pages,
                    'data':[valve.to_json() for valve in valves]
                    })

