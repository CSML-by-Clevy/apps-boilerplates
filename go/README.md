# Go lambda function boilerplate

## Install dependencies
The compressed function file will need all its dependencies so all install them in a folder called package.
```
$ cd src
$ go get github.com/aws/aws-lambda-go/lambda
```
#### IMPORTANT NOTE
Make sure you do compiled the executable with `GOOS=linux go build main.go` even if you compile it in a non-Linux environment

## Compress function
Create a deployment package by packaging the executable in a ZIP file
```
$ zip function.zip main
```

#### IMPORTANT NOTE
for Go The Handler is name `main`
