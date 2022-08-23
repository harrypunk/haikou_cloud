LDFLAGS = --ldflags "-extldflags -static"

build/aliyun_migrate.zip: build/aliyun_migrate
	zip -j $@ $<

build/aliyun_migrate: func/aliyun_migrate/app.go
	GOOS=linux GOARCH=amd64 go build  -o $@ ${LDFLAGS} $<
