package xerr

import (
	"mime"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func Parse(r *http.Request, v any) error {

	if err := httpx.ParsePath(r, v); err != nil {
		return err
	}

	if err := httpx.ParseHeaders(r, v); err != nil {
		return err
	}

	ct := r.Header.Get("Content-Type")
	d, _, _ := mime.ParseMediaType(ct)
	if d == "application/json" {
		if err := httpx.ParseJsonBody(r, v); err != nil {
			return err
		}
	} else {
		if err := httpx.ParseForm(r, v); err != nil {
			return err
		}
	}

	return nil
}
