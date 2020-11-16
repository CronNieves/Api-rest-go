module Api-Rest-Go

go 1.15

require github.com/gorilla/mux v1.8.0

require github.com/CronNieves/Api-rest-go/src/handlers v0.0.0

replace (
	github.com/CronNieves/Api-rest-go/src/handlers v0.0.0 => ./src/handlers
	github.com/CronNieves/Api-rest-go/src/models => ./src/models
)
