
import json
import logging
import math
import os
import shutil

from app import db
from app.apiv1 import api
from flask import request, jsonify, current_app, g
from sqlalchemy import func

from app.models import Bug,Device,Bugtype

@api.route('/bug/<int:id>', methods=['GET'])
def get_bug(id):
	bug = Bug.query.get_or_404(id)

	return jsonify({'success':True,
                    'error_code':0,
                    'records':bug.to_json(),
                    })

@api.route('/bug', methods=['POST'])
def create_bug():
	print(request.json)
	reason = request.json.get('reason')
	start_time = request.json.get('start_time')
	end_time = request.json.get('end_time')

	device_id = request.json.get('device_id')
	device = Device.query.filter_by(id=device_id).first()
 
	if device is None:
		return jsonify({'success':False,'error_code':-1,'errmsg':'device_id不存在'})	

	bugtype_id = request.json.get('bugtype_id')
	bugtype = Bugtype.query.filter_by(id=bugtype_id).first()
 
	if bugtype is None:
		return jsonify({'success':False,'error_code':-1,'errmsg':'bugtype_id不存在'})	

	bug = Bug(reason=reason,start_time=start_time,end_time=end_time,device_id=device.id,bugtype_id=bugtype.id,)

	db.session.add(bug)
	try:
		db.session.commit()
	except Exception as e:
		db.session.rollback()
		logging.error(f'添加数据库发生错误,已经回退:{e}')
		return jsonify({'success': False, 'error_code': -123, 'errmsg': '数据库插入错误，请查看日志'})

	return jsonify({'success':True,
                    'error_code':0,
                    })

@api.route('/bug/<int:id>', methods=['PUT'])
def modify_bug(id):
	print('put json:',request.json)
	bug = Bug.query.get_or_404(id)
	reason = request.json.get('reason')
	start_time = request.json.get('start_time')
	end_time = request.json.get('end_time')
	device_id = request.json.get('device_id')
	device = Device.query.filter_by(id=device_id).first()
	if device is None:
		return jsonify({'success':False,'error_code':-1,'errmsg':'device_id不存在'})	
	bugtype_id = request.json.get('bugtype_id')
	bugtype = Bugtype.query.filter_by(id=bugtype_id).first()
	if bugtype is None:
		return jsonify({'success':False,'error_code':-1,'errmsg':'bugtype_id不存在'})	
	bug.reason = reason or bug.reason
	bug.start_time = start_time or bug.start_time
	bug.end_time = end_time or bug.end_time
	bug.device_id = device.id
	bug.bugtype_id = bugtype.id
	db.session.add(bug)

	try:
		db.session.commit()
	except Exception as e:
		db.session.rollback()
		logging.error(f'修改数据库发生错误,已经回退:{e}')
	return jsonify({'success':True,
                    'error_code':0,
                    })

@api.route('/bug', methods=['DELETE'])
def delete_bug():
	print('delete json:',request.json)
	ids = request.json.get('ids')
	for id in ids:
		bug = Bug.query.get(id)
		if bug is None:
			return jsonify({'success': False, 'error_code': -123, 'errmsg': f'删除错误，id： {id} 不存在'})
		db.session.delete(bug)

		try:
			db.session.commit()
		except Exception as e:
			db.session.rollback()
			logging.error(f'删除数据库发生错误,已经回退:{e}')
			return jsonify({'success': False, 'error_code': -123, 'errmsg': f'删除数据发生错误， {e} '})

	return jsonify({'success':True,
                'error_code':0,
                })

@api.route('/bug/list', methods=['GET'])
def list_bug():
	print(request.args)
	sorter = request.args.get('sorter')
	page = int(request.args.get('current', 1))
	pageSize = int(request.args.get('pageSize', current_app.config['PER_PAGE']))
	pageSize = 20 if pageSize < 10 else pageSize
	total_bugs = Bug.query

	device_id = request.args.get('device_id')
	if device_id is not None:
		device = Device.query.filter_by(id=device_id).first()
		if device is None:
			return jsonify({'success':False,'error_code':-1,'errmsg':'device_id不存在'})
		else:
			total_bugs = total_bugs.filter_by(device_id=device.id)

	bugtype_id = request.args.get('bugtype_id')
	if bugtype_id is not None:
		bugtype = Bugtype.query.filter_by(id=bugtype_id).first()
		if bugtype is None:
			return jsonify({'success':False,'error_code':-1,'errmsg':'bugtype_id不存在'})
		else:
			total_bugs = total_bugs.filter_by(bugtype_id=bugtype.id)
	reason = request.args.get('reason')
	if reason is not None:
		total_bugs = total_bugs.filter(Bug.reason.ilike(f'%{reason}%'))

	if sorter:
		sorter = json.loads(sorter)
		pass
	totalcount = total_bugs.with_entities(func.count(Bug.id)).scalar()
	page = math.ceil(totalcount/pageSize) if  math.ceil(totalcount/pageSize) < page else page
	pagination = total_bugs.paginate(page, per_page = pageSize, error_out = False)
	bugs = pagination.items

	return jsonify({
                    'success':True,
                    'error_code':0,
                    'total':totalcount,
                    "pageSize" : pageSize,
                    "current" : page,
                    "pagecount": pagination.pages,
                    'data':[bug.to_json() for bug in bugs]
                    })

