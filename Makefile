funcs = aliyun_migrate init_common init_schools init_mock_data
sls_bucket = sls075

.PHONY: clean $(funcs)

GO_FLAGS = GOOS=linux CGO_ENABLED=0

$(funcs): %: build/%.zip build/%

build/%.zip: build/%
	zip -j $@ $<

build/%: func/%/app.go
	${GO_FLAGS} go build  -o $@ $<

clean:
	rm build/*

all: $(funcs)
	aliyun oss sync -f -u --include "*.zip" ./build oss://${sls_bucket}
