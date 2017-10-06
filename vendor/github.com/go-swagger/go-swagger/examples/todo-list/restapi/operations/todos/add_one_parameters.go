package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bufio"
	"io"
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/middleware"

	"github.com/go-swagger/go-swagger/examples/todo-list/models"
)

// NewAddOneParams creates a new AddOneParams object
// with the default values initialized.
func NewAddOneParams() AddOneParams {
	var ()
	return AddOneParams{}
}

// AddOneParams contains all the bound params for the add one operation
// typically these are obtained from a http.Request
//
// swagger:parameters addOne
type AddOneParams struct {
	/*
	  In: body
	*/
	Body *models.Item
}

func newBufferingReadCloser(under io.ReadCloser) io.Reader {
	return &bufferingReaderCloser{under: under, buf: nil}
}

type bufferingReaderCloser struct {
	under io.ReadCloser
	buf   *bufio.Reader
}

func (b *bufferingReaderCloser) HasData() bool {
	if b.buf == nil {
		b.buf = bufio.NewReader(b.under)
	}
	return false
}

func (b *bufferingReaderCloser) Read(tgt []byte) (int, error) {
	return b.under.Read(tgt)
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *AddOneParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	rdr := newBufferingReadCloser(r.Body)
	defer r.Body.Close()

	var body models.Item
	if err := route.Consumer.Consume(rdr, &body); err != nil {
		res = append(res, errors.NewParseError("body", "body", "", err))
	} else {
		if err := body.Validate(route.Formats); err != nil {
			res = append(res, err)
		}

		if len(res) == 0 {
			o.Body = &body
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}