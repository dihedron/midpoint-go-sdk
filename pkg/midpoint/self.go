package midpoint

import (
	"context"
	"log/slog"
)

type SelfService struct {
	Service
}

func (s *SelfService) Read(ctx context.Context) (*UserData, error) {
	self, err := s.client.Get[UserData](ctx, "self?options=raw")
	if err != nil {
		slog.Error("error reading self", "error", err)
		return nil, err
	}
	return self, nil
}
