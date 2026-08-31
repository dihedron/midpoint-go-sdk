package midpoint

import "errors"

type Service struct {
	client *Client
}

var ErrNotImplemented = errors.New("not implemented")
