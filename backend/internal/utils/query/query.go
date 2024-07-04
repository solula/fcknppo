package query

import (
	"github.com/google/go-querystring/query"
)

func Encode(opts interface{}) string {
	v, _ := query.Values(opts)
	return v.Encode()
}
