package routers

import (
	"order-service/internal/routers/manage"
)

type RouterGroup struct {
	Manage manage.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
