default: clean prepare test build

clean:
	@rm -rf build

prepare:
	go get -v github.com/modocache/gover
	go get -v golang.org/x/tools/cmd/cover
	go get -v github.com/mattn/goveralls
	go get -v github.com/onsi/ginkgo/ginkgo
	go get -v github.com/onsi/gomega
	go get -v github.com/onsi/gomega/gstruct
	go get github.com/kardianos/govendor
	govendor sync


test:
	@ginkgo --r  --cover --coverpkg=./...  --trace --race

build:
	@go build -a
	
