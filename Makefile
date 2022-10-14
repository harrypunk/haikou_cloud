funcs = aliyun_migrate init_common init_schools init_mock_families init_mock_teachers init_mock_sessions associate_mock_student_teacher
sls_bucket = sls075

.PHONY: clean $(funcs) sync all

GO_FLAGS = GOOS=linux CGO_ENABLED=0

$(funcs): %: build/%.zip build/%

build/%.zip: build/%
	zip -j $@ $<

build/%: func/%/app.go
	${GO_FLAGS} go build  -o $@ $<

clean:
	rm build/*

all: $(funcs)

sync:
	aliyun oss sync -f -u --include "*.zip" ./build oss://${sls_bucket}
