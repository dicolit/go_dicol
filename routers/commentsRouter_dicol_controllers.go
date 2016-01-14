package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["dicol/controllers:DcArticleController"] = append(beego.GlobalControllerRouter["dicol/controllers:DcArticleController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["dicol/controllers:DcArticleController"] = append(beego.GlobalControllerRouter["dicol/controllers:DcArticleController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dicol/controllers:DcArticleController"] = append(beego.GlobalControllerRouter["dicol/controllers:DcArticleController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dicol/controllers:DcArticleController"] = append(beego.GlobalControllerRouter["dicol/controllers:DcArticleController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["dicol/controllers:DcArticleController"] = append(beego.GlobalControllerRouter["dicol/controllers:DcArticleController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["dicol/controllers:DcProductController"] = append(beego.GlobalControllerRouter["dicol/controllers:DcProductController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["dicol/controllers:DcProductController"] = append(beego.GlobalControllerRouter["dicol/controllers:DcProductController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dicol/controllers:DcProductController"] = append(beego.GlobalControllerRouter["dicol/controllers:DcProductController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dicol/controllers:DcProductController"] = append(beego.GlobalControllerRouter["dicol/controllers:DcProductController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["dicol/controllers:DcProductController"] = append(beego.GlobalControllerRouter["dicol/controllers:DcProductController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
