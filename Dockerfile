
FROM python:3.6.2
RUN mkdir /code
RUN mkdir /code/log
WORKDIR /code
COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt
COPY src/gunicorn.conf.py ./
COPY src/start.sh ./
COPY . .
WORKDIR /code/src
RUN chmod +x start.sh
ENTRYPOINT ["./start.sh"]
