module github.com/ttacon/onepassword/cmd/ope

go 1.16

require (
	github.com/jedib0t/go-pretty/v6 v6.2.4 // indirect
	github.com/ttacon/onepassword v0.0.0
	github.com/ttacon/toml2cli v0.1.1 // indirect
	github.com/urfave/cli/v2 v2.3.0 // indirect
)

replace github.com/ttacon/onepassword => ../..
