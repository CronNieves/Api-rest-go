module Api-Rest-Go

go 1.15

require github.com/gorilla/mux v1.8.0
require github.com/CronNieves/Api-rest-go/src/handlers v0.0.0-00010101000000-000000000000
require github.com/CronNieves/Api-rest-go/src/models v0.0.0


replace (
	github.com/CronNieves/Api-rest-go/src/handlers => ./src/handlers
	github.com/CronNieves/Api-rest-go/src/models => ./src/models
)
