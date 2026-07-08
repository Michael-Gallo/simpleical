module github.com/michael-gallo/simpleical/benchmarks

go 1.26.0

require (
	github.com/apognu/gocal v0.9.1
	github.com/arran4/golang-ical v0.3.5
	github.com/michael-gallo/simpleical v0.0.0
	github.com/teambition/rrule-go v1.8.2
)

require github.com/ChannelMeter/iso8601duration v0.0.0-20150204201828-8da3af7a2a61 // indirect

replace github.com/michael-gallo/simpleical => ../
