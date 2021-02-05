
import json
import logging
import math
import os
import shutil

from app import db
from app.apiv1 import api
from flask import request, jsonify, current_app, g
from sqlalchemy import func

from app.models import Alarmtype

@api.route('/alarmtype/<int:id>', methods=['GET'])
def get_alarmtype(id):
	alarmtype = Alarmtype.query.get_or_404(id)

	return jsonify({'success':True,
                    'error_code':0,
                    'records':alarmtype.to_json(),
                    })

@api.route('/alarmtype', methods=['POST'])
def create_alarmtype():
	print(request.json)
	code = request.json.get('code')
	mean = request.json.get('mean')
	cause = request.json.get('cause')
	solution = request.json.get('solution')

	alarmtype = Alarmtype(code=code,mean=mean,cause=cause,solution=solution,)

	db.session.add(alarmtype)
	try:
		db.session.commit()
	except Exception as e:
		db.session.rollback()
		logging.error(f'添加数据库发生错误,已经回退:{e}')
		return jsonify({'success': False, 'error_code': -123, 'errmsg': '数据库插入错误，请查看日志'})

	return jsonify({'success':True,
                    'error_code':0,
                    })

@api.route('/alarmtype/<int:id>', methods=['PUT'])
def modify_alarmtype(id):
	print('put json:',request.json)
	alarmtype = Alarmtype.query.get_or_404(id)
	code = request.json.get('code')
	mean = request.json.get('mean')
	cause = request.json.get('cause')
	solution = request.json.get('solution')
	alarmtype.code = code or alarmtype.code
	alarmtype.mean = mean or alarmtype.mean
	alarmtype.cause = cause or alarmtype.cause
	alarmtype.solution = solution or alarmtype.solution
	db.session.add(alarmtype)

	try:
		db.session.commit()
	except Exception as e:
		db.session.rollback()
		logging.error(f'修改数据库发生错误,已经回退:{e}')
	return jsonify({'success':True,
                    'error_code':0,
                    })

@api.route('/alarmtype', methods=['DELETE'])
def delete_alarmtype():
	print('delete json:',request.json)
	ids = request.json.get('ids')
	for id in ids:
		alarmtype = Alarmtype.query.get(id)
		if alarmtype is None:
			return jsonify({'success': False, 'error_code': -123, 'errmsg': f'删除错误，id： {id} 不存在'})
		if alarmtype.alarms.first() is not None:
			return jsonify({'success':False,'error_code':-1,'errmsg':'alarmtype还拥有alarm，不能删除'})
		db.session.delete(alarmtype)

		try:
			db.session.commit()
		except Exception as e:
			db.session.rollback()
			logging.error(f'删除数据库发生错误,已经回退:{e}')
			return jsonify({'success': False, 'error_code': -123, 'errmsg': f'删除数据发生错误， {e} '})

	return jsonify({'success':True,
                'error_code':0,
                })

@api.route('/alarmtype/list', methods=['GET'])
def list_alarmtype():
	print(request.args)
	sorter = request.args.get('sorter')
	page = int(request.args.get('current', 1))
	pageSize = int(request.args.get('pageSize', current_app.config['PER_PAGE']))
	pageSize = 20 if pageSize < 10 else pageSize
	total_alarmtypes = Alarmtype.query
	if sorter:
		sorter = json.loads(sorter)
		pass
	totalcount = total_alarmtypes.with_entities(func.count(Alarmtype.id)).scalar()
	page = math.ceil(totalcount/pageSize) if  math.ceil(totalcount/pageSize) < page else page
	pagination = total_alarmtypes.paginate(page, per_page = pageSize, error_out = False)
	alarmtypes = pagination.items

	return jsonify({
                    'success':True,
                    'error_code':0,
                    'total':totalcount,
                    "pageSize" : pageSize,
                    "current" : page,
                    "pagecount": pagination.pages,
                    'data':[alarmtype.to_json() for alarmtype in alarmtypes]
                    })

