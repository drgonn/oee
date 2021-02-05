
import json
import logging
import math
import os
import shutil

from app import db
from app.apiv1 import api
from flask import request, jsonify, current_app, g
from sqlalchemy import func

from app.models import Worktime,Device

@api.route('/worktime/<int:id>', methods=['GET'])
def get_worktime(id):
	worktime = Worktime.query.get_or_404(id)

	return jsonify({'success':True,
                    'error_code':0,
                    'records':worktime.to_json(),
                    })

@api.route('/worktime', methods=['POST'])
def create_worktime():
	print(request.json)
	start_time = request.json.get('start_time')
	end_time = request.json.get('end_time')
	seconds = request.json.get('seconds')
	type = request.json.get('type')
	amount = request.json.get('amount')
	good = request.json.get('good')
	glue = request.json.get('glue')

	device_id = request.json.get('device_id')
	device = Device.query.filter_by(id=device_id).first()
 
	if device is None:
		return jsonify({'success':False,'error_code':-1,'errmsg':'device_id不存在'})	

	worktime = Worktime(start_time=start_time,end_time=end_time,seconds=seconds,type=type,amount=amount,good=good,glue=glue,device_id=device.id,)

	db.session.add(worktime)
	try:
		db.session.commit()
	except Exception as e:
		db.session.rollback()
		logging.error(f'添加数据库发生错误,已经回退:{e}')
		return jsonify({'success': False, 'error_code': -123, 'errmsg': '数据库插入错误，请查看日志'})

	return jsonify({'success':True,
                    'error_code':0,
                    })

@api.route('/worktime/<int:id>', methods=['PUT'])
def modify_worktime(id):
	print('put json:',request.json)
	worktime = Worktime.query.get_or_404(id)
	start_time = request.json.get('start_time')
	end_time = request.json.get('end_time')
	seconds = request.json.get('seconds')
	type = request.json.get('type')
	amount = request.json.get('amount')
	good = request.json.get('good')
	glue = request.json.get('glue')
	device_id = request.json.get('device_id')
	device = Device.query.filter_by(id=device_id).first()
	if device is None:
		return jsonify({'success':False,'error_code':-1,'errmsg':'device_id不存在'})	
	worktime.start_time = start_time or worktime.start_time
	worktime.end_time = end_time or worktime.end_time
	worktime.seconds = seconds or worktime.seconds
	worktime.type = type or worktime.type
	worktime.amount = amount or worktime.amount
	worktime.good = good or worktime.good
	worktime.glue = glue or worktime.glue
	worktime.device_id = device.id
	db.session.add(worktime)

	try:
		db.session.commit()
	except Exception as e:
		db.session.rollback()
		logging.error(f'修改数据库发生错误,已经回退:{e}')
	return jsonify({'success':True,
                    'error_code':0,
                    })

@api.route('/worktime', methods=['DELETE'])
def delete_worktime():
	print('delete json:',request.json)
	ids = request.json.get('ids')
	for id in ids:
		worktime = Worktime.query.get(id)
		if worktime is None:
			return jsonify({'success': False, 'error_code': -123, 'errmsg': f'删除错误，id： {id} 不存在'})
		db.session.delete(worktime)

		try:
			db.session.commit()
		except Exception as e:
			db.session.rollback()
			logging.error(f'删除数据库发生错误,已经回退:{e}')
			return jsonify({'success': False, 'error_code': -123, 'errmsg': f'删除数据发生错误， {e} '})

	return jsonify({'success':True,
                'error_code':0,
                })

@api.route('/worktime/list', methods=['GET'])
def list_worktime():
	print(request.args)
	sorter = request.args.get('sorter')
	page = int(request.args.get('current', 1))
	pageSize = int(request.args.get('pageSize', current_app.config['PER_PAGE']))
	pageSize = 20 if pageSize < 10 else pageSize
	total_worktimes = Worktime.query

	device_id = request.args.get('device_id')
	if device_id is not None:
		device = Device.query.filter_by(id=device_id).first()
		if device is None:
			return jsonify({'success':False,'error_code':-1,'errmsg':'device_id不存在'})
		else:
			total_worktimes = total_worktimes.filter_by(device_id=device.id)
	type = request.args.get('type')
	if type is not None:
		total_worktimes = total_worktimes.filter_by(type=type)

	if sorter:
		sorter = json.loads(sorter)
		pass
	totalcount = total_worktimes.with_entities(func.count(Worktime.id)).scalar()
	page = math.ceil(totalcount/pageSize) if  math.ceil(totalcount/pageSize) < page else page
	pagination = total_worktimes.paginate(page, per_page = pageSize, error_out = False)
	worktimes = pagination.items

	return jsonify({
                    'success':True,
                    'error_code':0,
                    'total':totalcount,
                    "pageSize" : pageSize,
                    "current" : page,
                    "pagecount": pagination.pages,
                    'data':[worktime.to_json() for worktime in worktimes]
                    })

