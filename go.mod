module Api-Rest-Go

go 1.15

require github.com/gorilla/mux v1.8.0

require github.com/CronNieves/Api-rest-go/src/handlers v0.0.0

require github.com/CronNieves/Api-rest-go/src/models v0.0.0

require (
	github.com/CronNieves/Api-rest-go/src/storage v0.0.0
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/spec v0.19.13 // indirect
	github.com/go-openapi/swag v0.19.12 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/swaggo/http-swagger v0.0.0-20200308142732-58ac5e232fba
	github.com/swaggo/swag v1.6.9
	github.com/urfave/cli/v2 v2.3.0 // indirect
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b // indirect
	golang.org/x/text v0.3.4 // indirect
	golang.org/x/tools v0.0.0-20201117152513-9036a0f9af11 // indirect
)

replace (
	github.com/CronNieves/Api-rest-go/src/handlers v0.0.0 => ./src/handlers
	github.com/CronNieves/Api-rest-go/src/models v0.0.0 => ./src/models
	github.com/CronNieves/Api-rest-go/src/storage v0.0.0 => ./src/storage
)
