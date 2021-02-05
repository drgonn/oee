import hashlib
from datetime import datetime
from time import time

from dateutil import tz
from dateutil.tz import tzlocal

permission_dir = {"one":1,"two":3,"three":7,"four":15,"five":31}
#生成散列
def md5(arg):#这是加密函数，将传进来的函数加密
    md5_pwd = hashlib.md5(bytes(arg,encoding='utf-8'))
    md5_pwd.update(bytes("chenrong",encoding='utf-8'))
    return md5_pwd.hexdigest()#返回加密的数据
# 生成随机散列
def randomd5():
	return md5(str(time()))
'''
 utc时间转本地时间,并转换为字符
'''
def utc_switch(utc,time_zone="GTM+8"):
	if utc:
		# time_zone = time_zone or "GTM+8"
		# local_zone = time_zone or datetime.now(tzlocal()).tzname()
		# # UTC Zone
		# from_zone = tz.gettz('UTC')
		# # China Zone
		# to_zone = tz.gettz(local_zone)
		# utc = utc.replace(tzinfo=to_zone)
		# local = utc.astimezone(from_zone)
		return datetime.strftime(utc, "%Y-%m-%d %H:%M:%S")

