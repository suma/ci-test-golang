#/bin/sh
# http://www.songmu.jp/riji/entry/2015-01-12-ci-golint.html
rm -f .golint.txt
for os in "linux" "darwin" "freebsd" "windows"; do
    GOOS=$os golint ./... | tee -a .golint.txt
done
test ! -s .golint.txt
