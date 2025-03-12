package services

func CloseServices() {
	GenaiClient.Close()
	RedisClient.Close()
}
