git config --global url.ssh://git@github.com/.insteadOf https://github.com/
git checkout master
git pull

go mod vendor
sudo systemctl restart goadmin.service