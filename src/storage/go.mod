module Api-Rest-Go/storage

go 1.15

require (
	github.com/CronNieves/Api-rest-go/src/models v0.0.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
    github.com/aws/aws-sdk-go v1.35.31
)

replace github.com/CronNieves/Api-rest-go/src/models v0.0.0 => ./../models
