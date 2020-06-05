package main

var mockData []Container = []Container{}

func main() {
	app := App{}

	app.Initialize()
	app.initializeDB()
	app.Run()
}
