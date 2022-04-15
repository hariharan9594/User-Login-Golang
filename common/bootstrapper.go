package common

func Startup() {
	CreateLog()
	initConfig()
	createDb()
}
