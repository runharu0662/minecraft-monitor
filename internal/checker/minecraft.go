package checker

import (
	"context"
	"github.com/mcstatus-io/mcutil/v4/status"
)

type Minecraft struct{ Address string }

func (m Minecraft) Check(ctx context.Context) error {
	_, err := status.Modern(ctx, m.Address)
	return err
}
