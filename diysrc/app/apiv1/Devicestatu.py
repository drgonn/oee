
import json
import logging
import math
import os
import shutil

from app import db
from app.apiv1 import api
from flask import request, jsonify, current_app, g
from sqlalchemy import func

from app.models import Devicestatu,Device

@api.route('/devicestatu/<int:id>', methods=['GET'])
def get_devicestatu(id):
	devicestatu = Devicestatu.query.get_or_404(id)

	return jsonify({'success':True,
                    'error_code':0,
                    'records':devicestatu.to_json(),
                    })

@api.route('/devicestatu', methods=['POST'])
def create_devicestatu():
	print(request.json)
	start_time = request.json.get('start_time')
	status = request.json.get('status')

	device_id = request.json.get('device_id')
	device = Device.query.filter_by(id=device_id).first()
 
	if device is None:
		return jsonify({'success':False,'error_code':-1,'errmsg':'device_id不存在'})	

	devicestatu = Devicestatu(start_time=start_time,status=status,device_id=device.id,)

	db.session.add(devicestatu)
	try:
		db.session.commit()
	except Exception as e:
		db.session.rollback()
		logging.error(f'添加数据库发生错误,已经回退:{e}')
		return jsonify({'success': False, 'error_code': -123, 'errmsg': '数据库插入错误，请查看日志'})

	return jsonify({'success':True,
                    'error_code':0,
                    })

@api.route('/devicestatu/<int:id>', methods=['PUT'])
def modify_devicestatu(id):
	print('put json:',request.json)
	devicestatu = Devicestatu.query.get_or_404(id)
	start_time = request.json.get('start_time')
	status = request.json.get('status')
	device_id = request.json.get('device_id')
	device = Device.query.filter_by(id=device_id).first()
	if device is None:
		return jsonify({'success':False,'error_code':-1,'errmsg':'device_id不存在'})	
	devicestatu.start_time = start_time or devicestatu.start_time
	devicestatu.status = status or devicestatu.status
	devicestatu.device_id = device.id
	db.session.add(devicestatu)

	try:
		db.session.commit()
	except Exception as e:
		db.session.rollback()
		logging.error(f'修改数据库发生错误,已经回退:{e}')
	return jsonify({'success':True,
                    'error_code':0,
                    })

@api.route('/devicestatu', methods=['DELETE'])
def delete_devicestatu():
	print('delete json:',request.json)
	ids = request.json.get('ids')
	for id in ids:
		devicestatu = Devicestatu.query.get(id)
		if devicestatu is None:
			return jsonify({'success': False, 'error_code': -123, 'errmsg': f'删除错误，id： {id} 不存在'})
		db.session.delete(devicestatu)

		try:
			db.session.commit()
		except Exception as e:
			db.session.rollback()
			logging.error(f'删除数据库发生错误,已经回退:{e}')
			return jsonify({'success': False, 'error_code': -123, 'errmsg': f'删除数据发生错误， {e} '})

	return jsonify({'success':True,
                'error_code':0,
                })

@api.route('/devicestatu/list', methods=['GET'])
def list_devicestatu():
	print(request.args)
	sorter = request.args.get('sorter')
	page = int(request.args.get('current', 1))
	pageSize = int(request.args.get('pageSize', current_app.config['PER_PAGE']))
	pageSize = 20 if pageSize < 10 else pageSize
	total_devicestatus = Devicestatu.query

	device_id = request.args.get('device_id')
	if device_id is not None:
		device = Device.query.filter_by(id=device_id).first()
		if device is None:
			return jsonify({'success':False,'error_code':-1,'errmsg':'device_id不存在'})
		else:
			total_devicestatus = total_devicestatus.filter_by(device_id=device.id)
	status = request.args.get('status')
	if status is not None:
		total_devicestatus = total_devicestatus.filter_by(status=status)

	if sorter:
		sorter = json.loads(sorter)
		pass
	totalcount = total_devicestatus.with_entities(func.count(Devicestatu.id)).scalar()
	page = math.ceil(totalcount/pageSize) if  math.ceil(totalcount/pageSize) < page else page
	pagination = total_devicestatus.paginate(page, per_page = pageSize, error_out = False)
	devicestatus = pagination.items

	return jsonify({
                    'success':True,
                    'error_code':0,
                    'total':totalcount,
                    "pageSize" : pageSize,
                    "current" : page,
                    "pagecount": pagination.pages,
                    'data':[devicestatu.to_json() for devicestatu in devicestatus]
                    })

