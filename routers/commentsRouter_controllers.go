package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:ClockinController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:ClockinController"],
        beego.ControllerComments{
            Method: "Clockin",
            Router: `/student/course/clockin/:password`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:ClockinController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:ClockinController"],
        beego.ControllerComments{
            Method: "SetClock",
            Router: `/teacher/course/setclockin/:password`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CommunityController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CommunityController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/community`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CommunityController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CommunityController"],
        beego.ControllerComments{
            Method: "GetQuestion",
            Router: `/community/:qid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CommunityController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CommunityController"],
        beego.ControllerComments{
            Method: "PostNewAnswer",
            Router: `/community/:qid`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CommunityController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CommunityController"],
        beego.ControllerComments{
            Method: "SupportAnswer",
            Router: `/community/answer/support/:aid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CommunityController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CommunityController"],
        beego.ControllerComments{
            Method: "DeleteQuestion",
            Router: `/community/delete/:qid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CommunityController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CommunityController"],
        beego.ControllerComments{
            Method: "GetMyQuestion",
            Router: `/community/myquestions`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CommunityController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CommunityController"],
        beego.ControllerComments{
            Method: "DealNotice",
            Router: `/community/notice/:nqid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CommunityController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CommunityController"],
        beego.ControllerComments{
            Method: "GetNewQuestion",
            Router: `/community/postquestion`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CommunityController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CommunityController"],
        beego.ControllerComments{
            Method: "PostNewQuestion",
            Router: `/community/postquestion`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CommunityController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CommunityController"],
        beego.ControllerComments{
            Method: "SupportQuestion",
            Router: `/community/support/:qid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CourseController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CourseController"],
        beego.ControllerComments{
            Method: "StudentSelectCourse",
            Router: `/student/:cid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CourseController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CourseController"],
        beego.ControllerComments{
            Method: "GetStudentCourse",
            Router: `/student/course`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CourseController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CourseController"],
        beego.ControllerComments{
            Method: "JionCourse",
            Router: `/student/joincourse/:cid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CourseController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CourseController"],
        beego.ControllerComments{
            Method: "QuitCourse",
            Router: `/student/quitcourse/:cid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CourseController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CourseController"],
        beego.ControllerComments{
            Method: "TeacherSelectCourse",
            Router: `/teacher/:cid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CourseController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CourseController"],
        beego.ControllerComments{
            Method: "AddStudent",
            Router: `/teacher/addstudent/:scid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CourseController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CourseController"],
        beego.ControllerComments{
            Method: "GetTeacherCourse",
            Router: `/teacher/course`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CourseController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CourseController"],
        beego.ControllerComments{
            Method: "DeleteCourse",
            Router: `/teacher/dropcourse/:cid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CourseController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:CourseController"],
        beego.ControllerComments{
            Method: "RefuseStudent",
            Router: `/teacher/refusestudent/:scid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:FileController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:FileController"],
        beego.ControllerComments{
            Method: "GetDownloadFile",
            Router: `/download/courseware/:filename`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:FileController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:FileController"],
        beego.ControllerComments{
            Method: "PostUploadFile",
            Router: `/teacher/course`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:FileController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:FileController"],
        beego.ControllerComments{
            Method: "GetDeleteFile",
            Router: `/teacher/delete/:filename`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:HomeworkController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:HomeworkController"],
        beego.ControllerComments{
            Method: "GetStudentHomework",
            Router: `/student/course/homework/:hid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:HomeworkController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:HomeworkController"],
        beego.ControllerComments{
            Method: "PostStudentHomework",
            Router: `/student/course/homework/:hid`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:HomeworkController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:HomeworkController"],
        beego.ControllerComments{
            Method: "GetAddHomework",
            Router: `/teacher/course/homework`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:HomeworkController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:HomeworkController"],
        beego.ControllerComments{
            Method: "PostAddHomework",
            Router: `/teacher/course/homework`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:HomeworkController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:HomeworkController"],
        beego.ControllerComments{
            Method: "GetTeacherHomework",
            Router: `/teacher/course/homework/:hid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:HomeworkController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:HomeworkController"],
        beego.ControllerComments{
            Method: "DownloadHomework",
            Router: `/teacher/download/:filename`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:MainController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:MainController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:MainController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:MainController"],
        beego.ControllerComments{
            Method: "InTheRow",
            Router: `/intherow`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:UsersController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetSignin",
            Router: `/signin`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:UsersController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:UsersController"],
        beego.ControllerComments{
            Method: "PostSignin",
            Router: `/signin`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:UsersController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetSignup",
            Router: `/signup`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:UsersController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:UsersController"],
        beego.ControllerComments{
            Method: "PostSignup",
            Router: `/signup`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:UsersController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetStudent",
            Router: `/student`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:UsersController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:UsersController"],
        beego.ControllerComments{
            Method: "PostSearchCourse",
            Router: `/student`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:UsersController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:UsersController"],
        beego.ControllerComments{
            Method: "DealNotice",
            Router: `/student/dealnotice/:nctid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:UsersController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetSetStudent",
            Router: `/student/setting`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:UsersController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:UsersController"],
        beego.ControllerComments{
            Method: "PostSetStudent",
            Router: `/student/setting`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:UsersController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:UsersController"],
        beego.ControllerComments{
            Method: "PostCreateCourse",
            Router: `/teacher`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:UsersController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetTeacher",
            Router: `/teacher`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:UsersController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetSetTeacher",
            Router: `/teacher/setting`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Teaching-assistant-system/controllers:UsersController"] = append(beego.GlobalControllerRouter["Teaching-assistant-system/controllers:UsersController"],
        beego.ControllerComments{
            Method: "PostSetTeacher",
            Router: `/teacher/setting`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
