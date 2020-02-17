package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["Teching-assistant-system/controllers:CourseController"] = append(beego.GlobalControllerRouter["Teching-assistant-system/controllers:CourseController"],
        beego.ControllerComments{
            Method: "GetStudentCourse",
            Router: `/student/course`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teching-assistant-system/controllers:CourseController"] = append(beego.GlobalControllerRouter["Teching-assistant-system/controllers:CourseController"],
        beego.ControllerComments{
            Method: "JionCourse",
            Router: `/student/joincourse/:cid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teching-assistant-system/controllers:CourseController"] = append(beego.GlobalControllerRouter["Teching-assistant-system/controllers:CourseController"],
        beego.ControllerComments{
            Method: "GetTeacherCourse",
            Router: `/teacher/course`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teching-assistant-system/controllers:CourseController"] = append(beego.GlobalControllerRouter["Teching-assistant-system/controllers:CourseController"],
        beego.ControllerComments{
            Method: "DeleteCourse",
            Router: `/teacher/dropcourse/:cid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teching-assistant-system/controllers:MainController"] = append(beego.GlobalControllerRouter["Teching-assistant-system/controllers:MainController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teching-assistant-system/controllers:MainController"] = append(beego.GlobalControllerRouter["Teching-assistant-system/controllers:MainController"],
        beego.ControllerComments{
            Method: "GetHelp",
            Router: `/help`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teching-assistant-system/controllers:MainController"] = append(beego.GlobalControllerRouter["Teching-assistant-system/controllers:MainController"],
        beego.ControllerComments{
            Method: "InTheRow",
            Router: `/intherow`,
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

    beego.GlobalControllerRouter["Teching-assistant-system/controllers:UsersController"] = append(beego.GlobalControllerRouter["Teching-assistant-system/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetStudent",
            Router: `/student`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teching-assistant-system/controllers:UsersController"] = append(beego.GlobalControllerRouter["Teching-assistant-system/controllers:UsersController"],
        beego.ControllerComments{
            Method: "PostSearchCourse",
            Router: `/student`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teching-assistant-system/controllers:UsersController"] = append(beego.GlobalControllerRouter["Teching-assistant-system/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetSetStudent",
            Router: `/student/setting`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teching-assistant-system/controllers:UsersController"] = append(beego.GlobalControllerRouter["Teching-assistant-system/controllers:UsersController"],
        beego.ControllerComments{
            Method: "PostSetStudent",
            Router: `/student/setting`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teching-assistant-system/controllers:UsersController"] = append(beego.GlobalControllerRouter["Teching-assistant-system/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetTeacher",
            Router: `/teacher`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teching-assistant-system/controllers:UsersController"] = append(beego.GlobalControllerRouter["Teching-assistant-system/controllers:UsersController"],
        beego.ControllerComments{
            Method: "PostCreateCourse",
            Router: `/teacher`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teching-assistant-system/controllers:UsersController"] = append(beego.GlobalControllerRouter["Teching-assistant-system/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetSetTeacher",
            Router: `/teacher/setting`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teching-assistant-system/controllers:UsersController"] = append(beego.GlobalControllerRouter["Teching-assistant-system/controllers:UsersController"],
        beego.ControllerComments{
            Method: "PostSetTeacher",
            Router: `/teacher/setting`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
