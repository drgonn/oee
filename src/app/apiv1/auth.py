from flask import g, jsonify, request
from app.tools.auth import untie_token
from app.models import User
from app.apiv1 import api
from app import db
#在所有的访问前做token或密码认证
@api.before_request
def before_request():
	token = request.args.get("token")
	if request.endpoint[:8] == "api.test":  # 跳过认证
		return
	if token:
		token_dir = untie_token(token)
		if token_dir.get("error"):
			return jsonify({"success": False, "error_code": 201001, "errmsg": f"{token_dir.get('error')}"}),401
		uid = token_dir["data"].get("uid")
		role = token_dir["data"].get("role")
		if uid:
			g.current_user =  User.query.filter_by(uid=uid).first()
			g.role = role
			g.token_used = True
		else:
			return jsonify({"success": False, "error_code": 201001, "errmsg": "token 失效"}),401
		if g.current_user is None:
			user = User(uid=uid)
			db.session.add(user)
			db.session.commit()
			g.current_user = user
	else:
		return jsonify({"success": False, "error_code": -4, "errmsg": "tocken 失效"}),401
@api.teardown_request
def teardown_request(exception=None):
	db.session.close()

@api.route("/test", methods=["GET"])
def test():
	return jsonify({
		"success": True,
		"error_code": 0,
		"qrurl" : "order test OK"
	})
