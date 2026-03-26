module github.com/Yandex-Practicum/tracker

go 1.25.0

require (
	actioninfo v0.0.0-00010101000000-000000000000
	daysteps v0.0.0-00010101000000-000000000000
	personaldata v0.0.0-00010101000000-000000000000
	trainings v0.0.0-00010101000000-000000000000
)

require spentenergy v0.0.0-00010101000000-000000000000 // indirect

replace (
	actioninfo => ./internal/actioninfo
	daysteps => ./internal/daysteps
	personaldata => ./internal/personaldata
	spentenergy => ./internal/spentenergy
	trainings => ./internal/trainings
)
