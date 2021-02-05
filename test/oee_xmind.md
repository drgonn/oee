# oee

## 机器

- 机器创建

	- 请求方法

		- POST

	- 请求地址

		- http://localhost:8002/api/v1/order/device?token={token}

	- 请求参数示例

		- {"sn": "2SLxD41W", "name": "K7TcvYn3", "ip": "g53cVBLE", "img": "sGcwyg1l", "type": "sdNZPWkw"}

- 机器列表

	- 请求方法

		- POST

	- 请求地址

		- http://localhost:8002/api/v1/order/device/list?token={token}

	- 请求参数示例

		- {"sn": "2SLxD41W", "name": "K7TcvYn3", "ip": "g53cVBLE", "img": "sGcwyg1l", "type": "sdNZPWkw", "current": 1, "pageSize": 15}

- 机器单个获取

	- 请求方法

		- GET

	- 请求地址

		- http://localhost:8002/api/v1/order/device/id/1?token={token}

	- 请求参数示例

		- {"sn": "2SLxD41W", "name": "K7TcvYn3", "ip": "g53cVBLE", "img": "sGcwyg1l", "type": "sdNZPWkw", "current": 1, "pageSize": 15}

- 机器修改

	- 请求方法

		- PUT

	- 请求地址

		- http://localhost:8002/api/v1/order/device/id/1?token={token}

	- 请求参数示例

		- {"sn": "yaL8mGNn", "name": "SEZpk8yB", "ip": "Snj7ZqkL", "img": "s40eGY6y", "type": "iYBckRfq"}

- 机器删除

	- 请求方法

		- DELETE

	- 请求地址

		- http://localhost:8002/api/v1/order/device/id?token={token}

	- 请求参数示例

		- {"sn": "yaL8mGNn", "name": "SEZpk8yB", "ip": "Snj7ZqkL", "img": "s40eGY6y", "type": "iYBckRfq"}

## 机台工作时间

- 机台工作时间创建

	- 请求方法

		- POST

	- 请求地址

		- http://localhost:8002/api/v1/order/worktime?token={token}

	- 请求参数示例

		- {"start_time": null, "end_time": null, "seconds": 3, "type": 3, "amount": 3, "good": 3, "glue": 58.454}

- 机台工作时间列表

	- 请求方法

		- POST

	- 请求地址

		- http://localhost:8002/api/v1/order/worktime/list?token={token}

	- 请求参数示例

		- {"start_time": null, "end_time": null, "seconds": 3, "type": 3, "amount": 3, "good": 3, "glue": 58.454, "current": 1, "pageSize": 15}

- 机台工作时间单个获取

	- 请求方法

		- GET

	- 请求地址

		- http://localhost:8002/api/v1/order/worktime/id/2?token={token}

	- 请求参数示例

		- {"start_time": null, "end_time": null, "seconds": 3, "type": 3, "amount": 3, "good": 3, "glue": 58.454, "current": 1, "pageSize": 15}

- 机台工作时间修改

	- 请求方法

		- PUT

	- 请求地址

		- http://localhost:8002/api/v1/order/worktime/id/5?token={token}

	- 请求参数示例

		- {"start_time": null, "end_time": null, "seconds": 3, "type": 2, "amount": 2, "good": 3, "glue": 54.404}

- 机台工作时间删除

	- 请求方法

		- DELETE

	- 请求地址

		- http://localhost:8002/api/v1/order/worktime/id/5?token={token}

	- 请求参数示例

		- {"start_time": null, "end_time": null, "seconds": 3, "type": 2, "amount": 2, "good": 3, "glue": 54.404}

## 机台开关状态

- 机台开关状态创建

	- 请求方法

		- POST

	- 请求地址

		- http://localhost:8002/api/v1/order/devicestatu?token={token}

	- 请求参数示例

		- {"start_time": null, "status": 3}

- 机台开关状态列表

	- 请求方法

		- POST

	- 请求地址

		- http://localhost:8002/api/v1/order/devicestatu/list?token={token}

	- 请求参数示例

		- {"start_time": null, "status": 3, "current": 1, "pageSize": 15}

- 机台开关状态单个获取

	- 请求方法

		- GET

	- 请求地址

		- http://localhost:8002/api/v1/order/devicestatu/id?token={token}

	- 请求参数示例

		- {"start_time": null, "status": 3, "current": 1, "pageSize": 15}

- 机台开关状态修改

	- 请求方法

		- PUT

	- 请求地址

		- http://localhost:8002/api/v1/order/devicestatu/id/2?token={token}

	- 请求参数示例

		- {"start_time": null, "status": 5}

- 机台开关状态删除

	- 请求方法

		- DELETE

	- 请求地址

		- http://localhost:8002/api/v1/order/devicestatu/id/3?token={token}

	- 请求参数示例

		- {"start_time": null, "status": 5}

## 阀工作时间

- 阀工作时间创建

	- 请求方法

		- POST

	- 请求地址

		- http://localhost:8002/api/v1/order/valvetime?token={token}

	- 请求参数示例

		- {"start_time": null, "end_time": null, "seconds": 11.859, "volt": 22.746, "amount": 2, "good": 0, "glue": 36.678}

- 阀工作时间列表

	- 请求方法

		- POST

	- 请求地址

		- http://localhost:8002/api/v1/order/valvetime/list?token={token}

	- 请求参数示例

		- {"start_time": null, "end_time": null, "seconds": 11.859, "volt": 22.746, "amount": 2, "good": 0, "glue": 36.678, "current": 1, "pageSize": 15}

- 阀工作时间单个获取

	- 请求方法

		- GET

	- 请求地址

		- http://localhost:8002/api/v1/order/valvetime/id/4?token={token}

	- 请求参数示例

		- {"start_time": null, "end_time": null, "seconds": 11.859, "volt": 22.746, "amount": 2, "good": 0, "glue": 36.678, "current": 1, "pageSize": 15}

- 阀工作时间修改

	- 请求方法

		- PUT

	- 请求地址

		- http://localhost:8002/api/v1/order/valvetime/id/1?token={token}

	- 请求参数示例

		- {"start_time": null, "end_time": null, "seconds": 10.94, "volt": 43.103, "amount": 3, "good": 2, "glue": 88.691}

- 阀工作时间删除

	- 请求方法

		- DELETE

	- 请求地址

		- http://localhost:8002/api/v1/order/valvetime/id/1?token={token}

	- 请求参数示例

		- {"start_time": null, "end_time": null, "seconds": 10.94, "volt": 43.103, "amount": 3, "good": 2, "glue": 88.691}

## 阀类型

- 阀类型创建

	- 请求方法

		- POST

	- 请求地址

		- http://localhost:8002/api/v1/order/valvetype?token={token}

	- 请求参数示例

		- {"name": "Nb2TLpUx"}

- 阀类型列表

	- 请求方法

		- POST

	- 请求地址

		- http://localhost:8002/api/v1/order/valvetype/list?token={token}

	- 请求参数示例

		- {"name": "Nb2TLpUx", "current": 1, "pageSize": 15}

- 阀类型单个获取

	- 请求方法

		- GET

	- 请求地址

		- http://localhost:8002/api/v1/order/valvetype/id/2?token={token}

	- 请求参数示例

		- {"name": "Nb2TLpUx", "current": 1, "pageSize": 15}

- 阀类型修改

	- 请求方法

		- PUT

	- 请求地址

		- http://localhost:8002/api/v1/order/valvetype/id?token={token}

	- 请求参数示例

		- {"name": "MAKoNmTV"}

- 阀类型删除

	- 请求方法

		- DELETE

	- 请求地址

		- http://localhost:8002/api/v1/order/valvetype/id/2?token={token}

	- 请求参数示例

		- {"name": "MAKoNmTV"}

## 点胶阀

- 点胶阀创建

	- 请求方法

		- POST

	- 请求地址

		- http://localhost:8002/api/v1/order/valve?token={token}

	- 请求参数示例

		- {"sn": "feR8Fjk4", "name": "FwshpeGM"}

- 点胶阀列表

	- 请求方法

		- POST

	- 请求地址

		- http://localhost:8002/api/v1/order/valve/list?token={token}

	- 请求参数示例

		- {"sn": "feR8Fjk4", "name": "FwshpeGM", "current": 1, "pageSize": 15}

- 点胶阀单个获取

	- 请求方法

		- GET

	- 请求地址

		- http://localhost:8002/api/v1/order/valve/id/1?token={token}

	- 请求参数示例

		- {"sn": "feR8Fjk4", "name": "FwshpeGM", "current": 1, "pageSize": 15}

- 点胶阀修改

	- 请求方法

		- PUT

	- 请求地址

		- http://localhost:8002/api/v1/order/valve/id/2?token={token}

	- 请求参数示例

		- {"sn": "cxyMkQU2", "name": "nzq409At"}

- 点胶阀删除

	- 请求方法

		- DELETE

	- 请求地址

		- http://localhost:8002/api/v1/order/valve/id/1?token={token}

	- 请求参数示例

		- {"sn": "cxyMkQU2", "name": "nzq409At"}

## 故障

- 故障创建

	- 请求方法

		- POST

	- 请求地址

		- http://localhost:8002/api/v1/order/bug?token={token}

	- 请求参数示例

		- {"reason": null, "start_time": null, "end_time": null}

- 故障列表

	- 请求方法

		- POST

	- 请求地址

		- http://localhost:8002/api/v1/order/bug/list?token={token}

	- 请求参数示例

		- {"reason": null, "start_time": null, "end_time": null, "current": 1, "pageSize": 15}

- 故障单个获取

	- 请求方法

		- GET

	- 请求地址

		- http://localhost:8002/api/v1/order/bug/id?token={token}

	- 请求参数示例

		- {"reason": null, "start_time": null, "end_time": null, "current": 1, "pageSize": 15}

- 故障修改

	- 请求方法

		- PUT

	- 请求地址

		- http://localhost:8002/api/v1/order/bug/id/4?token={token}

	- 请求参数示例

		- {"reason": null, "start_time": null, "end_time": null}

- 故障删除

	- 请求方法

		- DELETE

	- 请求地址

		- http://localhost:8002/api/v1/order/bug/id/2?token={token}

	- 请求参数示例

		- {"reason": null, "start_time": null, "end_time": null}

## 故障类型

- 故障类型创建

	- 请求方法

		- POST

	- 请求地址

		- http://localhost:8002/api/v1/order/bugtype?token={token}

	- 请求参数示例

		- {"name": "plHq19RZ", "sn": "HW4S8D3p", "about": null}

- 故障类型列表

	- 请求方法

		- POST

	- 请求地址

		- http://localhost:8002/api/v1/order/bugtype/list?token={token}

	- 请求参数示例

		- {"name": "plHq19RZ", "sn": "HW4S8D3p", "about": null, "current": 1, "pageSize": 15}

- 故障类型单个获取

	- 请求方法

		- GET

	- 请求地址

		- http://localhost:8002/api/v1/order/bugtype/id/2?token={token}

	- 请求参数示例

		- {"name": "plHq19RZ", "sn": "HW4S8D3p", "about": null, "current": 1, "pageSize": 15}

- 故障类型修改

	- 请求方法

		- PUT

	- 请求地址

		- http://localhost:8002/api/v1/order/bugtype/id/2?token={token}

	- 请求参数示例

		- {"name": "UySOwMIi", "sn": "1a0L2hnO", "about": null}

- 故障类型删除

	- 请求方法

		- DELETE

	- 请求地址

		- http://localhost:8002/api/v1/order/bugtype/id/1?token={token}

	- 请求参数示例

		- {"name": "UySOwMIi", "sn": "1a0L2hnO", "about": null}

## 报警

- 报警创建

	- 请求方法

		- POST

	- 请求地址

		- http://localhost:8002/api/v1/order/alarmtype?token={token}

	- 请求参数示例

		- {"code": "rlsEKLNz", "mean": null, "cause": null, "solution": null}

- 报警列表

	- 请求方法

		- POST

	- 请求地址

		- http://localhost:8002/api/v1/order/alarmtype/list?token={token}

	- 请求参数示例

		- {"code": "rlsEKLNz", "mean": null, "cause": null, "solution": null, "current": 1, "pageSize": 15}

- 报警单个获取

	- 请求方法

		- GET

	- 请求地址

		- http://localhost:8002/api/v1/order/alarmtype/id?token={token}

	- 请求参数示例

		- {"code": "rlsEKLNz", "mean": null, "cause": null, "solution": null, "current": 1, "pageSize": 15}

- 报警修改

	- 请求方法

		- PUT

	- 请求地址

		- http://localhost:8002/api/v1/order/alarmtype/id/3?token={token}

	- 请求参数示例

		- {"code": "8zrcEpJx", "mean": null, "cause": null, "solution": null}

- 报警删除

	- 请求方法

		- DELETE

	- 请求地址

		- http://localhost:8002/api/v1/order/alarmtype/id/5?token={token}

	- 请求参数示例

		- {"code": "8zrcEpJx", "mean": null, "cause": null, "solution": null}

## 报警

- 报警创建

	- 请求方法

		- POST

	- 请求地址

		- http://localhost:8002/api/v1/order/alarm?token={token}

	- 请求参数示例

		- {}

- 报警列表

	- 请求方法

		- POST

	- 请求地址

		- http://localhost:8002/api/v1/order/alarm/list?token={token}

	- 请求参数示例

		- {"current": 1, "pageSize": 15}

- 报警单个获取

	- 请求方法

		- GET

	- 请求地址

		- http://localhost:8002/api/v1/order/alarm/id?token={token}

	- 请求参数示例

		- {"current": 1, "pageSize": 15}

- 报警修改

	- 请求方法

		- PUT

	- 请求地址

		- http://localhost:8002/api/v1/order/alarm/id?token={token}

	- 请求参数示例

		- {}

- 报警删除

	- 请求方法

		- DELETE

	- 请求地址

		- http://localhost:8002/api/v1/order/alarm/id?token={token}

	- 请求参数示例

		- {}

## 项目分类

- 项目分类创建

	- 请求方法

		- POST

	- 请求地址

		- http://localhost:8002/api/v1/order/project?token={token}

	- 请求参数示例

		- {"name": "1iLntbz8"}

- 项目分类列表

	- 请求方法

		- POST

	- 请求地址

		- http://localhost:8002/api/v1/order/project/list?token={token}

	- 请求参数示例

		- {"name": "1iLntbz8", "current": 1, "pageSize": 15}

- 项目分类单个获取

	- 请求方法

		- GET

	- 请求地址

		- http://localhost:8002/api/v1/order/project/id/5?token={token}

	- 请求参数示例

		- {"name": "1iLntbz8", "current": 1, "pageSize": 15}

- 项目分类修改

	- 请求方法

		- PUT

	- 请求地址

		- http://localhost:8002/api/v1/order/project/id/5?token={token}

	- 请求参数示例

		- {"name": "bP6Ifxcp"}

- 项目分类删除

	- 请求方法

		- DELETE

	- 请求地址

		- http://localhost:8002/api/v1/order/project/id/5?token={token}

	- 请求参数示例

		- {"name": "bP6Ifxcp"}

## 项目进度

- 项目进度创建

	- 请求方法

		- POST

	- 请求地址

		- http://localhost:8002/api/v1/order/plan?token={token}

	- 请求参数示例

		- {"name": "RiUOkD9d", "week": 3, "current": null, "follow": null}

- 项目进度列表

	- 请求方法

		- POST

	- 请求地址

		- http://localhost:8002/api/v1/order/plan/list?token={token}

	- 请求参数示例

		- {"name": "RiUOkD9d", "week": 3, "current": 1, "follow": null, "pageSize": 15}

- 项目进度单个获取

	- 请求方法

		- GET

	- 请求地址

		- http://localhost:8002/api/v1/order/plan/id?token={token}

	- 请求参数示例

		- {"name": "RiUOkD9d", "week": 3, "current": 1, "follow": null, "pageSize": 15}

- 项目进度修改

	- 请求方法

		- PUT

	- 请求地址

		- http://localhost:8002/api/v1/order/plan/id/5?token={token}

	- 请求参数示例

		- {"name": "oEXugfFP", "week": 4}

- 项目进度删除

	- 请求方法

		- DELETE

	- 请求地址

		- http://localhost:8002/api/v1/order/plan/id/1?token={token}

	- 请求参数示例

		- {"name": "oEXugfFP", "week": 4}

## 通知

- 通知创建

	- 请求方法

		- POST

	- 请求地址

		- http://localhost:8002/api/v1/order/message?token={token}

	- 请求参数示例

		- {"title": "16WhY0u3", "type": 1, "description": null}

- 通知列表

	- 请求方法

		- POST

	- 请求地址

		- http://localhost:8002/api/v1/order/message/list?token={token}

	- 请求参数示例

		- {"title": "16WhY0u3", "type": 1, "description": null, "current": 1, "pageSize": 15}

- 通知单个获取

	- 请求方法

		- GET

	- 请求地址

		- http://localhost:8002/api/v1/order/message/id/1?token={token}

	- 请求参数示例

		- {"title": "16WhY0u3", "type": 1, "description": null, "current": 1, "pageSize": 15}

- 通知修改

	- 请求方法

		- PUT

	- 请求地址

		- http://localhost:8002/api/v1/order/message/id/1?token={token}

	- 请求参数示例

		- {"title": "VcdCKBIb", "type": 0, "read": 0}

- 通知删除

	- 请求方法

		- DELETE

	- 请求地址

		- http://localhost:8002/api/v1/order/message/id/5?token={token}

	- 请求参数示例

		- {"title": "VcdCKBIb", "type": 0, "read": 0}

