
# -*- coding: utf-8 -*-
from flask import Flask
from app import create_app, celery
app = create_app("production")
app.app_context().push()
if __name__ == '__main__':
	app.run()
