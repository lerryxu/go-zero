package utils

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

func HttpGet(ctx context.Context, uri string, header http.Header, timeout int, proxy *url.URL) ([]byte, error) {
	return HttpDo(ctx, http.MethodGet, uri, nil, header, timeout, proxy)
}

func HttpPost(ctx context.Context, uri string, data []byte, header http.Header, timeout int, proxy *url.URL) ([]byte, error) {
	return HttpDo(ctx, http.MethodPost, uri, data, header, timeout, proxy)
}

func HttpDo(ctx context.Context, method, uri string, data []byte, header http.Header, timeout int, proxy *url.URL) ([]byte, error) {
	client := &http.Client{}
	if timeout > 0 {
		client.Timeout = time.Duration(timeout) * time.Second
	}
	if proxy != nil {
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxy),
		}
	}
	req, err := http.NewRequest(method, uri, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	if header != nil {
		req.Header = header
	}

	resp, err := client.Do(req)
	fields := []logx.LogField{
		logx.Field("method", method),
		logx.Field("uri", uri),
	}
	if data != nil {
		fields = append(fields, logx.Field("body", string(data)))
	}

	if err != nil {
		logx.WithContext(ctx).WithFields(fields...).Error(err)
		return nil, err
	}
	defer resp.Body.Close()
	ret, err := io.ReadAll(resp.Body)
	logx.WithContext(ctx).WithFields(fields...).Debug(string(ret))
	return ret, err
}
