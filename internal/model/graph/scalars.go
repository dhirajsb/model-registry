package graph

import (
	"encoding/base64"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"strconv"
)

// Int64 is a scalar type used in gql schema
func MarshalInt64(i int64) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = fmt.Fprintf(w, "%d", i)
	})
}

func UnmarshalInt64(v interface{}) (int64, error) {
	switch v := v.(type) {
	case string:
		return strconv.ParseInt(v, 10, 64)
	case *string:
		return strconv.ParseInt(*v, 10, 64)
	case int64:
		return v, nil
	default:
		return -1, fmt.Errorf("%T is not int64", v)
	}
}

// ByteArray is a scalar type used in gql schema
func MarshalByteArray(b []byte) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = fmt.Fprintf(w, "%q", base64.StdEncoding.EncodeToString(b))
	})
}

func UnmarshalByteArray(v interface{}) ([]byte, error) {
	switch v := v.(type) {
	case string:
		return base64.StdEncoding.DecodeString(v)
	case *string:
		return base64.StdEncoding.DecodeString(*v)
	case []byte:
		return v, nil
	default:
		return nil, fmt.Errorf("%T is not []byte", v)
	}
}
