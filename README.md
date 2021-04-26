# go-mod-details
---
An action that returns details about a `go.mod` file.

## Inputs
### `modfile`
Path to the `go.mod` file that should be parsed (*optional*)

## Outputs
### `modfile`
Path to the `go.mod` file that was parsed

### `go_version`
The go version that is defined in the `go.mod` file

### `module`
The module name defined in the `go.mod` file



## Example
```yaml
steps:
  -
    name: Checkout code
    uses: actions/checkout@v2
  -
    name: Get go.mod details
    uses: Eun/go-mod-details@v1
    id: go-mod-details
  -
    name: Install Go
    uses: actions/setup-go@v1
    with:
      go-version: ${{ steps.go-mod-details.outputs.go_version }}
  -
    name: Test
    run: go test -v ./...
```
