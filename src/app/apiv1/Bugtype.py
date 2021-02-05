
import json
import logging
import math
import os
import shutil

from app import db
from app.apiv1 import api
from flask import request, jsonify, current_app, g
from sqlalchemy import func

from app.models import Bugtype

@api.route('/bugtype/<int:id>', methods=['GET'])
def get_bugtype(id):
	bugtype = Bugtype.query.get_or_404(id)

	return jsonify({'success':True,
                    'error_code':0,
                    'records':bugtype.to_json(),
                    })

@api.route('/bugtype', methods=['POST'])
def create_bugtype():
	print(request.json)
	name = request.json.get('name')
	sn = request.json.get('sn')
	about = request.json.get('about')

	bugtype = Bugtype(name=name,sn=sn,about=about,)

	db.session.add(bugtype)
	try:
		db.session.commit()
	except Exception as e:
		db.session.rollback()
		logging.error(f'添加数据库发生错误,已经回退:{e}')
		return jsonify({'success': False, 'error_code': -123, 'errmsg': '数据库插入错误，请查看日志'})

	return jsonify({'success':True,
                    'error_code':0,
                    })

@api.route('/bugtype/<int:id>', methods=['PUT'])
def modify_bugtype(id):
	print('put json:',request.json)
	bugtype = Bugtype.query.get_or_404(id)
	name = request.json.get('name')
	sn = request.json.get('sn')
	about = request.json.get('about')
	bugtype.name = name or bugtype.name
	bugtype.sn = sn or bugtype.sn
	bugtype.about = about or bugtype.about
	db.session.add(bugtype)

	try:
		db.session.commit()
	except Exception as e:
		db.session.rollback()
		logging.error(f'修改数据库发生错误,已经回退:{e}')
	return jsonify({'success':True,
                    'error_code':0,
                    })

@api.route('/bugtype', methods=['DELETE'])
def delete_bugtype():
	print('delete json:',request.json)
	ids = request.json.get('ids')
	for id in ids:
		bugtype = Bugtype.query.get(id)
		if bugtype is None:
			return jsonify({'success': False, 'error_code': -123, 'errmsg': f'删除错误，id： {id} 不存在'})
		if bugtype.bugs.first() is not None:
			return jsonify({'success':False,'error_code':-1,'errmsg':'bugtype还拥有bug，不能删除'})
		db.session.delete(bugtype)

		try:
			db.session.commit()
		except Exception as e:
			db.session.rollback()
			logging.error(f'删除数据库发生错误,已经回退:{e}')
			return jsonify({'success': False, 'error_code': -123, 'errmsg': f'删除数据发生错误， {e} '})

	return jsonify({'success':True,
                'error_code':0,
                })

@api.route('/bugtype/list', methods=['GET'])
def list_bugtype():
	print(request.args)
	sorter = request.args.get('sorter')
	page = int(request.args.get('current', 1))
	pageSize = int(request.args.get('pageSize', current_app.config['PER_PAGE']))
	pageSize = 20 if pageSize < 10 else pageSize
	total_bugtypes = Bugtype.query
	name = request.args.get('name')
	if name is not None:
		total_bugtypes = total_bugtypes.filter(Bugtype.name.ilike(f'%{name}%'))

	sn = request.args.get('sn')
	if sn is not None:
		total_bugtypes = total_bugtypes.filter(Bugtype.sn.ilike(f'%{sn}%'))

	if sorter:
		sorter = json.loads(sorter)
		pass
	totalcount = total_bugtypes.with_entities(func.count(Bugtype.id)).scalar()
	page = math.ceil(totalcount/pageSize) if  math.ceil(totalcount/pageSize) < page else page
	pagination = total_bugtypes.paginate(page, per_page = pageSize, error_out = False)
	bugtypes = pagination.items

	return jsonify({
                    'success':True,
                    'error_code':0,
                    'total':totalcount,
                    "pageSize" : pageSize,
                    "current" : page,
                    "pagecount": pagination.pages,
                    'data':[bugtype.to_json() for bugtype in bugtypes]
                    })

