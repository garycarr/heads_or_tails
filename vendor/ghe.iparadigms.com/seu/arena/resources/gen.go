package resources

//go:generate go get github.com/jteeuwen/go-bindata/...
//go:generate rm resources.go
//go:generate go-bindata -o resources.go -pkg resources ./...
//go:generate gofmt -w resources.go
