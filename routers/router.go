package routers

import (
	"dicol/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/article", &controllers.DcArticleController{})
    beego.Router("/product", &controllers.DcProductController{})
    /*
    ns := beego.NewNamespace("/vi",

		beego.NSNamespace("/dc_article",
			beego.NSInclude(
				&controllers.DcArticleController{},
			),
		),

		beego.NSNamespace("/dc_product",
			beego.NSInclude(
				&controllers.DcProductController{},
			),
		),
	)
	beego.AddNamespace(ns)
	*/
}
