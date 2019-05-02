// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"oas_sudamericana/liga_sudamericana_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/partido",
			beego.NSInclude(
				&controllers.PartidoController{},
			),
		),

		beego.NSNamespace("/posicion",
			beego.NSInclude(
				&controllers.PosicionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
