from functools import wraps
from flask import abort
from app.tools.auth import get_permission,Permission
#用作权限访问的装饰器
def permission_required(permission):
	def decorator(f):
		@wraps(f)
		def decorated_function(*args, **kwargs):
			if not get_permission(permission):
				abort(403)
			return f(*args, **kwargs)
		return decorated_function
	return decorator
def admin_required(f):
	return permission_required(Permission.ADMIN)(f)
