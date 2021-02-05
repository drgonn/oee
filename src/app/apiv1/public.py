import hashlib
import os

from app.apiv1 import api
from flask import request, jsonify, current_app, g
from werkzeug import secure_filename


def gen_file_name_with_md5sum(filename, md5sum):
	"""
	If file was exist already, rename it and return a new name
	"""
	name, extension = os.path.splitext(filename)
	filename = '%s_%s%s' % (name, md5sum, extension)
	return filename
@api.route('/upload', methods=['POST'])
def upload():
	upload_file = request.files.get('file')
	static_folder = current_app.config['STATIC_FOLDER']
	if upload_file:
		md5sum = hashlib.md5(upload_file.read()).hexdigest()
		upload_file.seek(0)
		filename = secure_filename(upload_file.filename)
		upload_file = request.files['file']
		filename = gen_file_name_with_md5sum(filename, md5sum)
		tmp_dir = os.path.join(static_folder, 'user_folder',f"{g.current_user.uid}")
		os.makedirs(tmp_dir, exist_ok=True)
		uploaded_file_path = os.path.join(tmp_dir,filename)
		upload_file.save(uploaded_file_path)
		# get file size after saving
		size = os.path.getsize(uploaded_file_path)
		if size <= 0:
			return jsonify({"success": False,
							"error_code": 0,
							}), 406,

		return jsonify({"success": True,
						"error_code": 0 ,
						"upload_file": filename,
						}), 201,

	return jsonify({"success": False,
					"error_code": 0,
					}), 406,
