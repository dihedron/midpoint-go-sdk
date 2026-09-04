package midpoint

import (
	"context"
	"log/slog"
)

// SelfService is the Service tha retrieves information about the current user's profile.
type SelfService struct {
	Service
}

// Read retrieves information about the current user's profile.
func (s *SelfService) Read(ctx context.Context) (*User, error) {
	response, err := s.client.
		R().
		SetContext(ctx).
		SetQueryParam("options", "raw").
		SetResult(&userWrapper{}).
		Get("/self")
	if err != nil {
		slog.Error("error reading self", "error", err, "result", response)
		return nil, err
	}
	return response.Result().(*userWrapper).User, nil
}
