module github.com/btnguyen2k/docms

go 1.13

require (
	github.com/urfave/cli v1.22.12
	gopkg.in/yaml.v3 v3.0.1
	main v0.0.0-00010101000000-000000000000
)

replace main => ./be-api
