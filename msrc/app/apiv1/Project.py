
import json
import logging
import math
import os
import shutil

from app import db
from app.apiv1 import api
from flask import request, jsonify, current_app, g
from sqlalchemy import func

from app.models import Project

@api.route('/project/<int:id>', methods=['GET'])
def get_project(id):
	project = Project.query.get_or_404(id)

	return jsonify({'success':True,
                    'error_code':0,
                    'records':project.to_json(),
                    })

@api.route('/project', methods=['POST'])
def create_project():
	print(request.json)
	name = request.json.get('name')

	project = Project(name=name,)

	db.session.add(project)
	try:
		db.session.commit()
	except Exception as e:
		db.session.rollback()
		logging.error(f'添加数据库发生错误,已经回退:{e}')
		return jsonify({'success': False, 'error_code': -123, 'errmsg': '数据库插入错误，请查看日志'})

	return jsonify({'success':True,
                    'error_code':0,
                    })

@api.route('/project/<int:id>', methods=['PUT'])
def modify_project(id):
	print('put json:',request.json)
	project = Project.query.get_or_404(id)
	name = request.json.get('name')
	project.name = name or project.name
	db.session.add(project)

	try:
		db.session.commit()
	except Exception as e:
		db.session.rollback()
		logging.error(f'修改数据库发生错误,已经回退:{e}')
	return jsonify({'success':True,
                    'error_code':0,
                    })

@api.route('/project', methods=['DELETE'])
def delete_project():
	print('delete json:',request.json)
	ids = request.json.get('ids')
	for id in ids:
		project = Project.query.get(id)
		if project is None:
			return jsonify({'success': False, 'error_code': -123, 'errmsg': f'删除错误，id： {id} 不存在'})
		if project.plans.first() is not None:
			return jsonify({'success':False,'error_code':-1,'errmsg':'project还拥有plan，不能删除'})
		db.session.delete(project)

		try:
			db.session.commit()
		except Exception as e:
			db.session.rollback()
			logging.error(f'删除数据库发生错误,已经回退:{e}')
			return jsonify({'success': False, 'error_code': -123, 'errmsg': f'删除数据发生错误， {e} '})

	return jsonify({'success':True,
                'error_code':0,
                })

@api.route('/project/list', methods=['GET'])
def list_project():
	print(request.args)
	sorter = request.args.get('sorter')
	page = int(request.args.get('current', 1))
	pageSize = int(request.args.get('pageSize', current_app.config['PER_PAGE']))
	pageSize = 20 if pageSize < 10 else pageSize
	total_projects = Project.query
	name = request.args.get('name')
	if name is not None:
		total_projects = total_projects.filter(Project.name.ilike(f'%{name}%'))

	if sorter:
		sorter = json.loads(sorter)
		pass
	totalcount = total_projects.with_entities(func.count(Project.id)).scalar()
	page = math.ceil(totalcount/pageSize) if  math.ceil(totalcount/pageSize) < page else page
	pagination = total_projects.paginate(page, per_page = pageSize, error_out = False)
	projects = pagination.items

	return jsonify({
                    'success':True,
                    'error_code':0,
                    'total':totalcount,
                    "pageSize" : pageSize,
                    "current" : page,
                    "pagecount": pagination.pages,
                    'data':[project.to_json() for project in projects]
                    })

