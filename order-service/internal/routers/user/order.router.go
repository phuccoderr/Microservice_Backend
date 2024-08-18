package user

import "github.com/gin-gonic/gin"

type OrderRouter struct {
}

func (pr *OrderRouter) InitOrderRouter(Router *gin.RouterGroup) {
	//public
	routerPublic := Router.Group("/api/v1/order")
	{
		routerPublic.GET("")
	}
	//private
}
