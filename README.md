# tinytest

This project holds minimal examples to explore TinyGo wasm. The examples
are served via GitHub pages and the [docs/](docs) directory on
[tinytest.juli.au](https://tinytest.juli.au).

It currently demonstrates a potential tinygo memory corruption issue:

- Go to https://tinytest.juli.au/
- Open browser developer tools
- Click `Main`

However, executing `go run main.go` does not show that issue.

## Development

The tools used by this repository in the [bin/](bin) directory (e.g. go,
make, golangci-lint) are automatically downloaded on demand by
[Hermit]. Hermit ensures that developers on Mac or Linus as well CI use
the same version of the same tools. No further installation is required
beyond cloning this repo. 

To build and serve the wasm samples locally run `./bin/make serve`. Run
`./bin/make` to view help for other make targets.

[Hermit]: https://cashapp.github.io/hermit/
