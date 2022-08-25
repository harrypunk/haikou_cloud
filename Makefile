.PHONY: clean

LDFLAGS = --ldflags "-extldflags -static"
GO_FLAGS = GOOS=linux CGO_ENABLED=0

build/aliyun_migrate.zip: build/aliyun_migrate
	zip -j $@ $<

build/aliyun_migrate: func/aliyun_migrate/app.go
	${GO_FLAGS} go build  -o $@ $<

clean:
	rm build/*