module github.com/n3wscott/releases-demo/samples

go 1.17

replace github.com/n3wscott/releases-demo/subcomponent/v2 => ./../subcomponent/v2

replace github.com/n3wscott/releases-demo/v2 => ./../v2

require (
	github.com/n3wscott/releases-demo/subcomponent/v2 v2.0.0-00010101000000-000000000000
	github.com/n3wscott/releases-demo/v2 v2.0.0-00010101000000-000000000000
)
