module github.com/n3wscott/releases-demo/samples

go 1.17

replace github.com/n3wscott/releases-demo/v2 => ./../v2

require (
	github.com/n3wscott/releases-demo/subcomponent/v2 v2.0.19
	github.com/n3wscott/releases-demo/v2 v2.0.19
)
