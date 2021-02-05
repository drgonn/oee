import pymysql
import json
import logging
import math

import os
import shutil
from datetime import datetime
from app import db
from app.apiv1 import api
from flask import request, jsonify, current_app, g
from sqlalchemy import func,extract

from app.models import Worktime,Device,Devicestatu


def s2hour(seconds):
	hour = seconds//3600
	hs = seconds - hour*3600
	minute = hs // 60
	ms = hs - minute*60
	if hour:
		r = f"{hour}小时{minute}分{ms}秒"
	elif minute:
		r = f"{minute}分{ms}秒"
	else:
		r = f"{ms}秒"
	return r



@api.route('/machine/detail', methods=['GET'])
def machine_detail():
	print(request.args)
	sorter = request.args.get('sorter')
	page = int(request.args.get('current', 1))
	pageSize = int(request.args.get('pageSize', current_app.config['PER_PAGE']))
	pageSize = 20 if pageSize < 10 else pageSize
	total_worktimes = Worktime.query
	total_status = Devicestatu.query.order_by(Devicestatu.start_time.desc())
	day = request.args.get('day')   # 查看某天结果
	month = request.args.get('month') # 查看某月结果





	device_id = request.args.get('device_id')
	if device_id is not None:
		device = Device.query.filter_by(id=device_id).first()
		if device is None:
			return jsonify({'success':False,'error_code':-1,'errmsg':'device_id不存在'})

		total_worktimes = total_worktimes.filter_by(device_id=device.id)
		total_status = total_status.filter_by(device_id=device.id)

	wtype = request.args.get('type')
	if wtype is not None:
		total_worktimes = total_worktimes.filter_by(type=wtype)
	if day:
		year = int(day[:4])
		month = int(day[5:7])
		tday = int(day[8:10])
		print(year,month,tday)
		total_worktimes = total_worktimes.filter(
			extract('year',Worktime.start_time) == year,
			extract('month', Worktime.start_time) == month,
			extract('day', Worktime.start_time) == tday,
		)
		total_status = total_status.filter(
			extract('year', Devicestatu.start_time) == year,
			extract('month', Devicestatu.start_time) == month,
			extract('day', Devicestatu.start_time) == tday,
			)
	if sorter:
		sorter = json.loads(sorter)
		pass
	totalcount = total_worktimes.with_entities(func.count(Worktime.id)).scalar()
	page = math.ceil(totalcount/pageSize) if  math.ceil(totalcount/pageSize) < page else page
	pagination = total_worktimes.paginate(page, per_page = pageSize, error_out = False)
	worktimes = pagination.items
	total_seconds = sum([w.seconds for w in worktimes])
	total_time = s2hour(total_seconds)



	return jsonify({
		'success':True,
		'error_code':0,
		'total':totalcount,
		"pageSize" : pageSize,
		"current" : page,
		"pagecount": pagination.pages,
		'data':[worktime.to_json() for worktime in worktimes],
		'statudata':[worktime.to_json() for worktime in total_status],
		'sumup':{
			"total_time":total_time,
		}
	})

@api.route('/calcuate/meachine', methods=['GET'])
def calcuate_worktime():
	print(request.args)
	total_worktimes = Worktime.query
	total_status = Devicestatu.query.order_by(Devicestatu.start_time.desc())


	device_id = request.args.get('device_id')
	if device_id is not None:
		device = Device.query.filter_by(id=device_id).first()
		if device is None:
			return jsonify({'success':False,'error_code':-1,'errmsg':'device_id不存在'})
		else:
			"""查询更新数据"""
			HOST =  device.ip
			PORT = 3306
			USER = 'root'
			PASSWD = '7811175yy'
			DB = 'oee'
			CHARSET = 'utf8'
			try:
				conn = pymysql.connect(
					host = HOST,
					port = PORT,
					user = USER,
					passwd = PASSWD,
					db = DB,
					charset = CHARSET)
				cursor = conn.cursor()
			except Exception:
				pass
			near = device.worktimes.order_by(Worktime.start_time.desc()).first()
			if near:
				neartime = near.start_time
				nearstr = datetime.strftime(neartime, "%Y%m%d%H%M%S")
				sql = f"select * from status where start_time>={nearstr}"
				print(sql)
			else:
				sql = "select * from status"

			try:
				cursor.execute(sql)
				data = cursor.fetchall()
				cursor.close()
				for i,d in enumerate(data[:-1]):
					stime = d[1]
					statu = d[2]
					print(statu)
					if Devicestatu.query.filter_by(start_time =stime).filter_by(device_id = device_id).first() is None:
						ds = Devicestatu(start_time = stime, status = statu,device_id=device_id)
						db.session.add(ds)
					if statu == 1:
						bd = data[i+1]
						if bd[2] != 2:
							print('数据存在连续启动而缺少启动状态')
							continue
						etime = bd[1]
						if Worktime.query.filter_by(start_time =stime).filter_by(type=1).filter_by(device_id = device_id).first() is None:
							print(stime)
							print(etime)
							print(etime - stime)
							seconds = (etime - stime).total_seconds()
							w = Worktime(start_time = stime, end_time = etime,type =1,device_id=device_id,seconds=seconds)
							db.session.add(w)
					elif statu == 2:
						bd = data[i + 1]
						if bd[2] != 1:
							print('数据存在连续启动而缺少启动状态')
							continue
						etime = bd[1]
						if Worktime.query.filter_by(start_time=stime).filter_by(type=2).filter_by(
								device_id=device_id).first() is None:
							print(stime)
							print(etime)
							print(etime - stime)
							seconds = (etime - stime).total_seconds()
							w = Worktime(start_time=stime, end_time=etime, type=2, device_id=device_id, seconds=seconds)
							db.session.add(w)
				db.session.commit()
			except Exception as e:
				print("请求数据没有连上机台数据库",e)


		total_worktimes = total_worktimes.filter_by(device_id=device.id)
		total_status = total_status.filter_by(device_id=device.id)




	return jsonify({
		'success':True,
		'error_code':0,
	})

