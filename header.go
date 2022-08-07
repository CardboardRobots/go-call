package call

import (
	"net/http"
)

type HeaderMap = map[string]string

type Header struct {
	*http.Header
}

func NewHeader(header *http.Header) Header {
	return Header{
		Header: header,
	}
}

const HEADER_CONTENT_TYPE = "Content-Type"
const HEADER_AUTHORIZATION = "Authorization"

type ContentType string

const CONTENT_TYPE_FORM_URLENCODED ContentType = "application/x-www-form-urlencoded; param=value"
const CONTENT_TYPE_JSON ContentType = "application/json"

func (h *Header) SetContentType(contentType ContentType) {
	h.Add(HEADER_CONTENT_TYPE, string(contentType))
}

func (h *Header) SetAuthorization(authorization string) {
	h.Add(HEADER_AUTHORIZATION, authorization)
}
