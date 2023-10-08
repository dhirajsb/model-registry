bin/restmapper: tools/restmapper/restmapper.* tools/restmapper/* internal/model/library/*.go
	go build -o bin/restmapper tools/restmapper/restmapper.go

LIB_SRC := $(wildcard config/metadata-library/*.yaml)
REST_SCHEMA_SRC := $(LIB_SRC:config/metadata-library/%=api/rest/schemas/%)
REST_CONVERTER_SRC := $(patsubst %yaml,%rest.converter.go,$(LIB_SRC:config/metadata-libary/%=internal/converter/%))

$(REST_SCHEMA_SRC) $(REST_CONVERTER_SRC) : .sentinel ;

.rest.sentinel: $(LIB_SRC)
	bin/restmapper
	@touch $@

.PHONY: gen/restmapper
gen/restmapper: bin/goverter bin/restmapper $(REST_SCHEMA_SRC) $(REST_CONVERTER_SRC)
