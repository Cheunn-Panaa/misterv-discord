# Fetch dependencies.
go get -d -v
go build -ldflags='-w -s -extldflags "-static"' -a -o /usr/src/app .