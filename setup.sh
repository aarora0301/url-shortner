
cd $GOPATH/src/github.com/poc/url-shortner
go build
if [ $? -eq 0 ]; then
    echo  Build OK
else
    echo  Build FAIL
fi

go run app.go