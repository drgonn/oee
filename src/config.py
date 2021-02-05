import os
basedir = os.path.abspath(os.path.dirname(__file__))
from datetime import timedelta
from celery.schedules import crontab

class Config:
	SECRET_KEY = os.environ.get("SECRET_KEY") or "2d21c91217794dc687100033cf2e47c9"
	SSL_REDIRECT = False
	SQLALCHEMY_TRACK_MODIFICATIONS = False
	SQLALCHEMY_RECORD_QUERIES = True
	PER_PAGE = 20
	FLASKY_SLOW_DB_QUERY_TIME = 0.5

	STATIC_FOLDER = os.path.join(basedir,"app","static")

	STATIC_HOST = os.environ.get("STATIC_HOST") or "http://localhost:8002"



	@staticmethod
	def init_app(app):
		pass
		
class DevelopmentConfig(Config):
	DEBUG = True
	MAIN_HOST = os.environ.get("MAIN_HOST") or "http://127.0.0.1:5001"
	# Redis configuration
	REDIS_HOST = "localhost"
	REDIS_PORT = 6379
	# Celery configuration
	BROKER_URL = "redis://localhost:6379/0"
	CELERY_RESULT_BACKEND = "redis://localhost:6379/0"
	# Database configuration
	SQL_NAME = os.environ.get("SQL_NAME") or "root"
	SQL_PASSWORD = os.environ.get("SQL_PASSWORD") or "7811175yy"
	SQL_HOST = os.environ.get("SQL_HOST") or "127.0.0.1:3306"
	SQL_DATABASE = os.environ.get("SQL_DATABASE") or "oee"
	SQLALCHEMY_DATABASE_URI = f"mysql+pymysql://{SQL_NAME}:{SQL_PASSWORD}@{SQL_HOST}/{SQL_DATABASE}"

	# SQLALCHEMY_DATABASE_URI = "mysql+pymysql://root:7811175yy@127.0.0.1:3306/dcv3qiot"



	INFLUX_USERNAME = os.environ.get("INFLUX_USERNAME") or  "admin"
	INFLUX_PASSWORD = os.environ.get("INFLUX_PASSWORD") or  "668899"
	INFLUX_DBNAME   = os.environ.get("INFLUX_DBNAME") or  "db0"
	INFLUX_HOST     = os.environ.get("INFLUX_HOST") or  "influxdb"
	OTA_URL = os.environ.get("OTA_URL") or "127.0.0.1"
	OTA_URL_ALL = f"http://{OTA_URL}:5010"
	QIOT_HOST = os.environ.get("QIOT_HOST") or  "0.0.0.0:5001"
	QIOT_URL = os.environ.get("QIOT_URL") or "https://iot.sealan.tech/mina/"


class TestingConfig(Config):
	TESTING = True
	# Redis configuration
	REDIS_HOST = "localhost"
	REDIS_PORT = 6379
	# Celery configuration
	BROKER_URL = "redis://localhost:6379/0"
	CELERY_RESULT_BACKEND = "redis://localhost:6379/0"
	# Database configuration
	SQLALCHEMY_DATABASE_URI = os.environ.get("TEST_DATABASE_URL") or \
		"mysql+pymysql://root:7811175yy@127.0.0.1:3306/test"
	WTF_CSRF_ENABLED = False
	QIOT_URL = os.environ.get("QIOT_URL") or "https://iot.sealan.tech/mina/"



class ProductionConfig(Config):
	MAIN_HOST = os.environ.get("MAIN_HOST") or "https://qiot.sealan-tech.com"

	# Redis configuration
	REDIS_HOST = os.environ.get("REDIS_HOST") or "redis"
	REDIS_PORT = os.environ.get("REDIS_PORT") or 6379
	REDIS_DEV  = os.environ.get("REDIS_DEV") or 0
	# Celery configuration
	BROKER_URL = f"redis://{REDIS_HOST}:{REDIS_PORT}/{REDIS_DEV}"
	CELERY_RESULT_BACKEND = f"redis://{REDIS_HOST}:{REDIS_PORT}/{REDIS_DEV}"
	# Database configuration
	SQL_NAME = os.environ.get("SQL_NAME") or "root"
	SQL_PASSWORD = os.environ.get("SQL_PASSWORD") or "668899"
	SQL_HOST = os.environ.get("SQL_HOST") or "172.17.0.5"
	SQL_DATABASE = os.environ.get("SQL_DATABASE") or "oee"
	SQLALCHEMY_DATABASE_URI = f"mysql+pymysql://{SQL_NAME}:{SQL_PASSWORD}@{SQL_HOST}/{SQL_DATABASE}"

	ROLE_IP = os.environ.get("ROLE_IP")   or "127.0.0.1"
	ROLE_PORT = os.environ.get("ROLR_PORT") or "50051"

	INFLUX_USERNAME = os.environ.get("INFLUX_USERNAME") or  "admin"
	INFLUX_PASSWORD = os.environ.get("INFLUX_PASSWORD") or  "668899"
	INFLUX_DBNAME   = os.environ.get("INFLUX_DBNAME") or  "db0"
	INFLUX_HOST     = os.environ.get("INFLUX_HOST") or  "influxdb"
	QIOT_HOST = os.environ.get("QIOT_HOST") or  "dc.sealan-tech.com"
	QIOT_URL = os.environ.get("QIOT_URL") or "https://iot.sealan.tech/mina/"

config = {
	"development": DevelopmentConfig,
	"testing": TestingConfig,
	"production": ProductionConfig,
	"default": DevelopmentConfig
}
