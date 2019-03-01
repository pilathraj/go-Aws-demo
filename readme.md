## Links
- [Download install AWS Command line Interface](https://aws.amazon.com/cli/)
- [Doc AWS Lambda Deployment Package in Go](https://docs.aws.amazon.com/lambda/latest/dg/lambda-go-how-to-create-deployment-package.html)

## Dev note: Install packages.
Download the  zip tool first:
1. go get -u github.com/aws/aws-lambda-go/cmd/build-lambda-zip
## Application
1. go get github.com/aws/aws-lambda-go/lambd

### create build in windows.
- Cmd> set GOOS=linux
- go build -o main main.go
- D:\go\bin\build-lambda-zip.exe -o main.zip main
- **upload main.zip** when create a lambda function in AWS
