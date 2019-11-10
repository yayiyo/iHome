package routers

import (
	"iHome/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/api/v1.0/areas", &controllers.AreaController{},"get:GetAreas")
    beego.Router("/api/v1.0/houses/index", &controllers.HouseIndexController{},"get:GetHouseIndex")
    beego.Router("/api/v1.0/session", &controllers.SessionController{},"get:GetSessionData;delete:DeleteSessionData")
    beego.Router("/api/v1.0/users", &controllers.UserController{},"post:Register")
	beego.Router("/api/v1.0/sessions", &controllers.SessionController{},"post:Login")
}
