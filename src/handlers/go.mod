module Api-Rest-Go/handlers

go 1.15

require (
	github.com/CronNieves/Api-rest-go/src/models v0.0.0
	github.com/gorilla/mux v1.8.0
	github.com/CronNieves/Api-rest-go/src/storage v0.0.0
)

replace (
	github.com/CronNieves/Api-rest-go/src/storage v0.0.0 => ./../storage
	github.com/CronNieves/Api-rest-go/src/models v0.0.0 => ./../models
)


