echo '>> running the unit test...'

cd ../ # project base path
go test ./...
if [ $? = 1 ]; then
    exit 1
fi