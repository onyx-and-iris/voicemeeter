module github.com/onyx-and-iris/voicemeeter

go 1.18

// package files moved into root of repository
retract [v1.0.0, v1.1.0]

require (
	github.com/sirupsen/logrus v1.9.0
	github.com/stretchr/testify v1.8.0
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
