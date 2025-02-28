// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// GetCategoryByIdParams is parameters of getCategoryById operation.
type GetCategoryByIdParams struct {
	ID int
}

func unpackGetCategoryByIdParams(packed middleware.Parameters) (params GetCategoryByIdParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int)
	}
	return params
}

func decodeGetCategoryByIdParams(args [1]string, argsEscaped bool, r *http.Request) (params GetCategoryByIdParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}
