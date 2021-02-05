
import json
import logging
import math
import os
import shutil

from app import db
from app.apiv1 import api
from flask import request, jsonify, current_app, g
from sqlalchemy import func

from app.models import Message,User

@api.route('/message/<int:id>', methods=['GET'])
def get_message(id):
	message = Message.query.get_or_404(id)

	return jsonify({'success':True,
                    'error_code':0,
                    'records':message.to_json(),
                    })

@api.route('/message', methods=['POST'])
def create_message():
	print(request.json)
	title = request.json.get('title')
	type = request.json.get('type')
	description = request.json.get('description')
	user = g.current_user
 
	message = Message(title=title,type=type,description=description,user_id=user.id,)

	db.session.add(message)
	try:
		db.session.commit()
	except Exception as e:
		db.session.rollback()
		logging.error(f'添加数据库发生错误,已经回退:{e}')
		return jsonify({'success': False, 'error_code': -123, 'errmsg': '数据库插入错误，请查看日志'})

	return jsonify({'success':True,
                    'error_code':0,
                    })

@api.route('/message/<int:id>', methods=['PUT'])
def modify_message(id):
	print('put json:',request.json)
	message = Message.query.get_or_404(id)
	title = request.json.get('title')
	type = request.json.get('type')
	read = request.json.get('read')
	user_id = request.json.get('user_id')

	message.title = title or message.title
	message.type = type or message.type
	message.read = read or message.read
	db.session.add(message)

	try:
		db.session.commit()
	except Exception as e:
		db.session.rollback()
		logging.error(f'修改数据库发生错误,已经回退:{e}')
	return jsonify({'success':True,
                    'error_code':0,
                    })

@api.route('/message', methods=['DELETE'])
def delete_message():
	print('delete json:',request.json)
	ids = request.json.get('ids')
	for id in ids:
		message = Message.query.get(id)
		if message is None:
			return jsonify({'success': False, 'error_code': -123, 'errmsg': f'删除错误，id： {id} 不存在'})
		db.session.delete(message)

		try:
			db.session.commit()
		except Exception as e:
			db.session.rollback()
			logging.error(f'删除数据库发生错误,已经回退:{e}')
			return jsonify({'success': False, 'error_code': -123, 'errmsg': f'删除数据发生错误， {e} '})

	return jsonify({'success':True,
                'error_code':0,
                })

@api.route('/message/list', methods=['GET'])
def list_message():
	print(request.args)
	sorter = request.args.get('sorter')
	page = int(request.args.get('current', 1))
	pageSize = int(request.args.get('pageSize', current_app.config['PER_PAGE']))
	pageSize = 20 if pageSize < 10 else pageSize
	total_messages = Message.query

	user_id = request.args.get('user_id')
	if user_id is not None:
		user = User.query.filter_by(id=user_id).first()
		if user is None:
			return jsonify({'success':False,'error_code':-1,'errmsg':'user_id不存在'})
		else:
			total_messages = total_messages.filter_by(user_id=user.id)
	title = request.args.get('title')
	if title is not None:
		total_messages = total_messages.filter(Message.title.ilike(f'%{title}%'))

	type = request.args.get('type')
	if type is not None:
		total_messages = total_messages.filter_by(type=type)

	read = request.args.get('read')
	if read is not None:
		total_messages = total_messages.filter_by(read=read)
		print(read,total_messages.all())

	if sorter:
		sorter = json.loads(sorter)
		pass
	totalcount = total_messages.with_entities(func.count(Message.id)).scalar()
	page = math.ceil(totalcount/pageSize) if  math.ceil(totalcount/pageSize) < page else page
	pagination = total_messages.paginate(page, per_page = pageSize, error_out = False)
	messages = pagination.items

	return jsonify({
                    'success':True,
                    'error_code':0,
                    'total':totalcount,
                    "pageSize" : pageSize,
                    "current" : page,
                    "pagecount": pagination.pages,
                    'data':[message.to_json() for message in messages]
                    })

