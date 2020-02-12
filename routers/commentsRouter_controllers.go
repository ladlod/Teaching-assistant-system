package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["Teching-assistant-system/controllers:MainController"] = append(beego.GlobalControllerRouter["Teching-assistant-system/controllers:MainController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teching-assistant-system/controllers:UsersController"] = append(beego.GlobalControllerRouter["Teching-assistant-system/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetSignin",
            Router: `/signin`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teching-assistant-system/controllers:UsersController"] = append(beego.GlobalControllerRouter["Teching-assistant-system/controllers:UsersController"],
        beego.ControllerComments{
            Method: "PostSignin",
            Router: `/signin`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teching-assistant-system/controllers:UsersController"] = append(beego.GlobalControllerRouter["Teching-assistant-system/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetSignup",
            Router: `/signup`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teching-assistant-system/controllers:UsersController"] = append(beego.GlobalControllerRouter["Teching-assistant-system/controllers:UsersController"],
        beego.ControllerComments{
            Method: "PostSignup",
            Router: `/signup`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
