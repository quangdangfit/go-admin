git reset --hard
git checkout develop
git pull

git config --global url.ssh://git@github.com/.insteadOf https://github.com/
go get github.com/tokoinofficial/common_lib@develop
go get github.com/tokoinofficial/error_lib@develop

go mod vendor

go test -v ./...

echo "Finished"