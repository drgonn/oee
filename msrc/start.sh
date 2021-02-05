#!/bin/bash
python manage.py db upgrade
python manage.py redis
python manage.py init_base
gunicorn  wsgi:app  -c  ./gunicorn.conf.py
