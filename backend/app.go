// backend/app.go
// Wails 应用主结构，初始化和管理各个服务实例

package backend

import (
	"ahut-tool/backend/jwxt"
	"ahut-tool/backend/news"
	"ahut-tool/backend/pay"
	"context"
)

var (
	JwxtInstance *jwxt.Service
	PayInstance  *pay.Service
	NewsInstance *news.Service
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	JwxtInstance = jwxt.NewService()
	PayInstance = pay.NewService()
	NewsInstance = news.NewService()
}
