
import json
import logging
import math
import os
import shutil

from app import db
from app.apiv1 import api
from flask import request, jsonify, current_app, g
from sqlalchemy import func

from app.models import Valvetype

@api.route('/valvetype/<int:id>', methods=['GET'])
def get_valvetype(id):
	valvetype = Valvetype.query.get_or_404(id)

	return jsonify({'success':True,
                    'error_code':0,
                    'records':valvetype.to_json(),
                    })

@api.route('/valvetype', methods=['POST'])
def create_valvetype():
	print(request.json)
	name = request.json.get('name')

	valvetype = Valvetype(name=name,)

	db.session.add(valvetype)
	try:
		db.session.commit()
	except Exception as e:
		db.session.rollback()
		logging.error(f'添加数据库发生错误,已经回退:{e}')
		return jsonify({'success': False, 'error_code': -123, 'errmsg': '数据库插入错误，请查看日志'})

	return jsonify({'success':True,
                    'error_code':0,
                    })

@api.route('/valvetype/<int:id>', methods=['PUT'])
def modify_valvetype(id):
	print('put json:',request.json)
	valvetype = Valvetype.query.get_or_404(id)
	name = request.json.get('name')
	valvetype.name = name or valvetype.name
	db.session.add(valvetype)

	try:
		db.session.commit()
	except Exception as e:
		db.session.rollback()
		logging.error(f'修改数据库发生错误,已经回退:{e}')
	return jsonify({'success':True,
                    'error_code':0,
                    })

@api.route('/valvetype', methods=['DELETE'])
def delete_valvetype():
	print('delete json:',request.json)
	ids = request.json.get('ids')
	for id in ids:
		valvetype = Valvetype.query.get(id)
		if valvetype is None:
			return jsonify({'success': False, 'error_code': -123, 'errmsg': f'删除错误，id： {id} 不存在'})
		if valvetype.valves.first() is not None:
			return jsonify({'success':False,'error_code':-1,'errmsg':'valvetype还拥有valve，不能删除'})
		db.session.delete(valvetype)

		try:
			db.session.commit()
		except Exception as e:
			db.session.rollback()
			logging.error(f'删除数据库发生错误,已经回退:{e}')
			return jsonify({'success': False, 'error_code': -123, 'errmsg': f'删除数据发生错误， {e} '})

	return jsonify({'success':True,
                'error_code':0,
                })

@api.route('/valvetype/list', methods=['GET'])
def list_valvetype():
	print(request.args)
	sorter = request.args.get('sorter')
	page = int(request.args.get('current', 1))
	pageSize = int(request.args.get('pageSize', current_app.config['PER_PAGE']))
	pageSize = 20 if pageSize < 10 else pageSize
	total_valvetypes = Valvetype.query
	name = request.args.get('name')
	if name is not None:
		total_valvetypes = total_valvetypes.filter(Valvetype.name.ilike(f'%{name}%'))

	if sorter:
		sorter = json.loads(sorter)
		pass
	totalcount = total_valvetypes.with_entities(func.count(Valvetype.id)).scalar()
	page = math.ceil(totalcount/pageSize) if  math.ceil(totalcount/pageSize) < page else page
	pagination = total_valvetypes.paginate(page, per_page = pageSize, error_out = False)
	valvetypes = pagination.items

	return jsonify({
                    'success':True,
                    'error_code':0,
                    'total':totalcount,
                    "pageSize" : pageSize,
                    "current" : page,
                    "pagecount": pagination.pages,
                    'data':[valvetype.to_json() for valvetype in valvetypes]
                    })

