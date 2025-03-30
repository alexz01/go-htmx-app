module example.com/go-htmx-app

go 1.24.0

require (
	example.com/go-htmx-app/greetings v0.0.0-00010101000000-000000000000
	example.com/go-htmx-app/server v0.0.0-00010101000000-000000000000
	rsc.io/quote v1.5.2
)

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)

replace example.com/go-htmx-app/greetings => ./greetings

replace example.com/go-htmx-app/server => ./server
