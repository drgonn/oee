
import json
import logging
import math
import os
import shutil

from app import db
from app.apiv1 import api
from flask import request, jsonify, current_app, g
from sqlalchemy import func

from app.models import Plan,Project,User

@api.route('/plan/<int:id>', methods=['GET'])
def get_plan(id):
	plan = Plan.query.get_or_404(id)

	return jsonify({'success':True,
                    'error_code':0,
                    'records':plan.to_json(),
                    })

@api.route('/plan', methods=['POST'])
def create_plan():
	print(request.json)
	name = request.json.get('name')
	week = request.json.get('week')
	current = request.json.get('current')
	follow = request.json.get('follow')

	project_id = request.json.get('project_id')
	project = Project.query.filter_by(id=project_id).first()
 
	if project is None:
		return jsonify({'success':False,'error_code':-1,'errmsg':'project_id不存在'})	
	user = g.current_user
 
	plan = Plan(name=name,week=week,current=current,follow=follow,project_id=project.id,user_id=user.id,)

	db.session.add(plan)
	try:
		db.session.commit()
	except Exception as e:
		db.session.rollback()
		logging.error(f'添加数据库发生错误,已经回退:{e}')
		return jsonify({'success': False, 'error_code': -123, 'errmsg': '数据库插入错误，请查看日志'})

	return jsonify({'success':True,
                    'error_code':0,
                    })

@api.route('/plan/<int:id>', methods=['PUT'])
def modify_plan(id):
	print('put json:',request.json)
	plan = Plan.query.get_or_404(id)
	name = request.json.get('name')
	week = request.json.get('week')
	project_id = request.json.get('project_id')
	project = Project.query.filter_by(id=project_id).first()
	if project is None:
		return jsonify({'success':False,'error_code':-1,'errmsg':'project_id不存在'})	
	user_id = request.json.get('user_id')
	user = User.query.filter_by(id=user_id).first()
	if user is None:
		return jsonify({'success':False,'error_code':-1,'errmsg':'user_id不存在'})	
	plan.name = name or plan.name
	plan.week = week or plan.week
	plan.project_id = project.id
	plan.user_id = user.id
	db.session.add(plan)

	try:
		db.session.commit()
	except Exception as e:
		db.session.rollback()
		logging.error(f'修改数据库发生错误,已经回退:{e}')
	return jsonify({'success':True,
                    'error_code':0,
                    })

@api.route('/plan', methods=['DELETE'])
def delete_plan():
	print('delete json:',request.json)
	ids = request.json.get('ids')
	for id in ids:
		plan = Plan.query.get(id)
		if plan is None:
			return jsonify({'success': False, 'error_code': -123, 'errmsg': f'删除错误，id： {id} 不存在'})
		db.session.delete(plan)

		try:
			db.session.commit()
		except Exception as e:
			db.session.rollback()
			logging.error(f'删除数据库发生错误,已经回退:{e}')
			return jsonify({'success': False, 'error_code': -123, 'errmsg': f'删除数据发生错误， {e} '})

	return jsonify({'success':True,
                'error_code':0,
                })

@api.route('/plan/list', methods=['GET'])
def list_plan():
	print(request.args)
	sorter = request.args.get('sorter')
	page = int(request.args.get('current', 1))
	pageSize = int(request.args.get('pageSize', current_app.config['PER_PAGE']))
	pageSize = 20 if pageSize < 10 else pageSize
	total_plans = Plan.query

	project_id = request.args.get('project_id')
	if project_id is not None:
		project = Project.query.filter_by(id=project_id).first()
		if project is None:
			return jsonify({'success':False,'error_code':-1,'errmsg':'project_id不存在'})
		else:
			total_plans = total_plans.filter_by(project_id=project.id)

	user_id = request.args.get('user_id')
	if user_id is not None:
		user = User.query.filter_by(id=user_id).first()
		if user is None:
			return jsonify({'success':False,'error_code':-1,'errmsg':'user_id不存在'})
		else:
			total_plans = total_plans.filter_by(user_id=user.id)
	name = request.args.get('name')
	if name is not None:
		total_plans = total_plans.filter(Plan.name.ilike(f'%{name}%'))

	week = request.args.get('week')
	if week is not None:
		total_plans = total_plans.filter_by(week=week)

	if sorter:
		sorter = json.loads(sorter)
		pass
	totalcount = total_plans.with_entities(func.count(Plan.id)).scalar()
	page = math.ceil(totalcount/pageSize) if  math.ceil(totalcount/pageSize) < page else page
	pagination = total_plans.paginate(page, per_page = pageSize, error_out = False)
	plans = pagination.items

	return jsonify({
                    'success':True,
                    'error_code':0,
                    'total':totalcount,
                    "pageSize" : pageSize,
                    "current" : page,
                    "pagecount": pagination.pages,
                    'data':[plan.to_json() for plan in plans]
                    })

