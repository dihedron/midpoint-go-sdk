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
	//entity, result, err := s.client.Get[userWrapper](ctx, "self?options=raw")
	response, err := s.client.client.R().Get(ctx, "/self?options=raw")
	if err != nil {
		slog.Error("error reading self", "error", err, "result", result)
		return nil, err
	}
	return entity.User, nil
}
