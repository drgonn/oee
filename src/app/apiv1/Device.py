
import json
import logging
import math
import os
import shutil

from app import db
from app.apiv1 import api
from flask import request, jsonify, current_app, g
from sqlalchemy import func

from app.models import Device

@api.route('/device/<int:id>', methods=['GET'])
def get_device(id):
	device = Device.query.get_or_404(id)

	return jsonify({'success':True,
                    'error_code':0,
                    'records':device.to_json(),
                    })

@api.route('/device', methods=['POST'])
def create_device():
	print(request.json)
	sn = request.json.get('sn')
	name = request.json.get('name')
	ip = request.json.get('ip')
	img = request.json.get('img')
	type = request.json.get('type')

	device = Device(sn=sn,name=name,ip=ip,img=img,type=type,)
	static_folder = current_app.config['STATIC_FOLDER']
	user_dir = os.path.join(static_folder, 'user_folder', f"{g.current_user.uid}")
	tmp_file_path = os.path.join(user_dir, img)
	if not os.path.exists(tmp_file_path):
		return jsonify({'success':False,'error_code':-1,'errmsg':f'文件{tmp_file_path}不存在'})

	db.session.add(device)
	try:
		db.session.commit()
		dst_dir = os.path.join(static_folder, 'deviceimg', f"{device.id}")
		dst_file_path = os.path.join(dst_dir, img)
		os.makedirs(dst_dir,exist_ok=True)
		shutil.move(tmp_file_path,dst_file_path)
		shutil.rmtree(user_dir)
	except Exception as e:
		db.session.rollback()
		logging.error(f'添加数据库发生错误,已经回退:{e}')
		return jsonify({'success': False, 'error_code': -123, 'errmsg': '数据库插入错误，请查看日志'})

	return jsonify({'success':True,
                    'error_code':0,
                    })

@api.route('/device/<int:id>', methods=['PUT'])
def modify_device(id):
	print('put json:',request.json)
	device = Device.query.get_or_404(id)
	sn = request.json.get('sn')
	name = request.json.get('name')
	ip = request.json.get('ip')
	img = request.json.get('img')
	type = request.json.get('type')
	device.sn = sn or device.sn
	device.name = name or device.name
	device.ip = ip or device.ip
	device.img = img or device.img
	device.type = type or device.type
	db.session.add(device)

	try:
		db.session.commit()
	except Exception as e:
		db.session.rollback()
		logging.error(f'修改数据库发生错误,已经回退:{e}')
	return jsonify({'success':True,
                    'error_code':0,
                    })

@api.route('/device', methods=['DELETE'])
def delete_device():
	print('delete json:',request.json)
	ids = request.json.get('ids')
	for id in ids:
		device = Device.query.get(id)
		if device is None:
			return jsonify({'success': False, 'error_code': -123, 'errmsg': f'删除错误，id： {id} 不存在'})
		if device.worktimes.first() is not None:
			return jsonify({'success':False,'error_code':-1,'errmsg':'device还拥有worktime，不能删除'})
		if device.devicestatus.first() is not None:
			return jsonify({'success':False,'error_code':-1,'errmsg':'device还拥有devicestatu，不能删除'})
		if device.valvetimes.first() is not None:
			return jsonify({'success':False,'error_code':-1,'errmsg':'device还拥有valvetime，不能删除'})
		if device.valves.first() is not None:
			return jsonify({'success':False,'error_code':-1,'errmsg':'device还拥有valve，不能删除'})
		if device.bugs.first() is not None:
			return jsonify({'success':False,'error_code':-1,'errmsg':'device还拥有bug，不能删除'})
		if device.alarms.first() is not None:
			return jsonify({'success':False,'error_code':-1,'errmsg':'device还拥有alarm，不能删除'})
		db.session.delete(device)

		try:
			db.session.commit()
			static_folder = current_app.config['STATIC_FOLDER']
			dst_dir = os.path.join(static_folder, 'device', f"{device.id}")
			shutil.rmtree(dst_dir)
		except Exception as e:
			db.session.rollback()
			logging.error(f'删除数据库发生错误,已经回退:{e}')
			return jsonify({'success': False, 'error_code': -123, 'errmsg': f'删除数据发生错误， {e} '})

	return jsonify({'success':True,
                'error_code':0,
                })

@api.route('/device/list', methods=['GET'])
def list_device():
	print(request.args)
	sorter = request.args.get('sorter')
	page = int(request.args.get('current', 1))
	pageSize = int(request.args.get('pageSize', current_app.config['PER_PAGE']))
	pageSize = 20 if pageSize < 10 else pageSize
	total_devices = Device.query
	sn = request.args.get('sn')
	if sn is not None:
		total_devices = total_devices.filter(Device.sn.ilike(f'%{sn}%'))

	name = request.args.get('name')
	if name is not None:
		total_devices = total_devices.filter(Device.name.ilike(f'%{name}%'))

	type = request.args.get('type')
	if type is not None:
		total_devices = total_devices.filter_by(type=type)

	if sorter:
		sorter = json.loads(sorter)
		pass
	totalcount = total_devices.with_entities(func.count(Device.id)).scalar()
	page = math.ceil(totalcount/pageSize) if  math.ceil(totalcount/pageSize) < page else page
	pagination = total_devices.paginate(page, per_page = pageSize, error_out = False)
	devices = pagination.items

	return jsonify({
                    'success':True,
                    'error_code':0,
                    'total':totalcount,
                    "pageSize" : pageSize,
                    "current" : page,
                    "pagecount": pagination.pages,
                    'data':[device.to_json() for device in devices]
                    })

