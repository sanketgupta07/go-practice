binary=go-practice
tidy:
	go mod tidy

vet:
	go vet ./...

test:
	go test -v ./...

build: tidy vet test	
	go build -o ${binary} -v

run:build
	./${binary}

#make git-update -e m=<comment>
m="update"
git-update:
	go fmt ./...
	git status
	git add -A
	git commit -m "${m}"
	git push

#layering
#all:vet test run


