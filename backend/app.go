package backend

import (
	"ahut-tool/backend/config"
	"ahut-tool/backend/jwxt"
	"ahut-tool/backend/pay"
	"context"
)

var JwxtInstance *jwxt.Service
var PayInstance *pay.Service
var ConfigInstance *config.Service

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	// 初始化Service实例
	JwxtInstance = jwxt.NewService()
	PayInstance = pay.NewService()
	ConfigInstance = config.NewService()
}
