import redis

class Redis:

	def __init__(self, connStr):
		conn = connStr.split(":")
		self.conn = redis.Redis(host=conn[0], port=conn[1], db=0)

	def client(self):
		return self.conn
	
	
