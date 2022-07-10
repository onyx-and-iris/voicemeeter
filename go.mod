module github.com/onyx-and-iris/voicemeeter-api-go

go 1.18

retract (
	// package files moved into root of repository
	[v1.0.0, v1.1.0]
)

require (
	github.com/stretchr/testify v1.8.0
	golang.org/x/sys v0.0.0-20220708085239-5a0f0661e09d
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
