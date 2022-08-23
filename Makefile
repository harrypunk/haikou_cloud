.PHONY: clean

LDFLAGS = --ldflags "-extldflags -static"

build/aliyun_migrate.zip: build/aliyun_migrate
	zip -j $@ $<

build/aliyun_migrate: func/aliyun_migrate/app.go
	GOOS=linux CGO_ENABLED=0 go build  -o $@ $<

clean:
	rm build/*