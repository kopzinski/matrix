echo "Setting up the GOPATH to: $PWD"
export GOPATH=$PWD

echo "Running the Matrix program tests."
go test ./...
