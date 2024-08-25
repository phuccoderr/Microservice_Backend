package routers

import (
	"order-service/internal/routers/manage"
	"order-service/internal/routers/user"
)

type RouterGroup struct {
	Manage manage.ManageRouterGroup
	User   user.UserRouterGroup
}

var RouterGroupApp = new(RouterGroup)
