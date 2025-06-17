package cmd

import (
	"context"
	"gmanager/internal/admin/logic"
	"gmanager/internal/library/cache"
)

func initData(ctx context.Context) {
	cache.InitCache(ctx)
	_ = logic.Config.Refresh(ctx)
}
