package server

func InitServer() {
	route := initRoute()
	route.Run(":8081")
}
