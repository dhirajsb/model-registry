bin/gqlmapper: tools/gqlmapper/gqlmapper.* tools/gqlmapper/* internal/model/library/*.go
	go build -o bin/gqlmapper tools/gqlmapper/gqlmapper.go

LIB_SRC := $(wildcard config/metadata-library/*.yaml)
GRAPHQLS_SRC := $(patsubst %yaml,%graphqls,$(LIB_SRC:config/metadata-library/%=api/graphql/%))
CONVERTER_SRC := $(patsubst %yaml,%converter.go,$(LIB_SRC:config/metadata-libary/%=internal/converter/%))

$(GRAPHQLS_SRC) $(CONVERTER_SRC) : .sentinel ;

.sentinel: $(LIB_SRC)
	bin/gqlmapper
	@touch $@

.PHONY: gen/gqlmapper
gen/gqlmapper: bin/goverter bin/gqlmapper $(GRAPHQLS_SRC) $(CONVERTER_SRC)
