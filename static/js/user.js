function judgeEmpty(f, p, q){
    var x = document.forms[f][p].value;
    if(x == ""){
        alert("请输入" + q);
        return false;
    }
}

function judgeFile(){
    var x = document.getElementById("uploadfile").value;
    if(x == ""){
        alert("请选择文件")
        return false
    }
}

function judgeForm(p, p1 ,p2, p3) {
    var x = document.forms[p][p1].value;
    var y = document.forms[p][p2].value;
    var z = document.forms[p][p3].value;
    if(x == ""){
        alert("请选择身份信息！");
        return false;
    }
    if(y == ""){
        alert("请输入账号！");
        return false;
    }
    if(z == ""){
        alert("请输入密码！");
        return false;
    }    
}

function judgeForm2(p, p1 ,p2, p3, p4) {
    var x = document.forms[p][p1].value;
    var y = document.forms[p][p2].value;
    var z = document.forms[p][p3].value;
    var a = document.forms[p][p4].value;
    if(x == ""){
        alert("请选择身份信息！");
        return false;
    }
    if(y == ""){
        alert("请输入账号！");
        return false;
    }
    if(z == ""){
        alert("请输入姓名！");
        return false;
    }
    if(a == ""){
        alert("请输入密码！");
        return false;
    }    
}

function judgeForm3(p, p1 ,p2) {
    var x = document.forms[p][p1].value;
    var y = document.forms[p][p2].value;
    if(x == ""){
        alert("请输入新姓名！");
        return false;
    }
    if(y == ""){
        alert("请输入新密码！");
        return false;
    } 
}

function judgePassword(f, p1, p2) {
    var x = document.forms[f][p1].value;
    var y = document.forms[f][p2].value;

    if(x != y){
        alert("两次密码输入不同");
        return false;
    }
}

function notice(notice){
    if(notice != ""){
        alert(notice);
    }
}

function returnToUser(identity) {
    if(identity == "teacher"){
        location.href = "/teacher"
    }else if(identity == "student"){
        location.href = "/student"
    }
}

function confirmAddStudent(sid, cid, nid){
    if(confirm('是否允许这名学生加入课堂？')){
        location.href="/teacher/addstudent/"+sid+"&&"+cid+"&&"+nid;
    }else{
        location.href="/teacher/refusestudent/"+sid+"&&"+cid+"&&"+nid;
    }
}

function setSignPassword(){
    var password = prompt("请输入签到密码", "");
    if(password != null && password != ""){
        location.href="/teacher/course/setclockin/" + password;
    }
}

function clockIn(){
    var password = prompt("请输入签到密码", "");
    if(password != null && password != ""){
        location.href="/student/course/clockin/" + password;
    }
}

function quitCourse(cid){
    if(confirm('请确认退出课堂')){
        location.href= "/student/quitcourse/" + cid;
        return true;
    }else{
        return false;
    }
}

function deletequestion(qid){
    if(confirm('请确认删除帖子')){
        location.href="/community/delete/" + qid;
        return true
    }else {
        return false
    }
}

function judge100Percent(){
    var c1 = document.forms["exam_form"][C1].value;
    var c2 = document.forms["exam_form"][C2].value;
    var c3 = document.forms["exam_form"][C3].value;
    var c4 = document.forms["exam_form"][C4].value;
    var c5 = document.forms["exam_form"][C5].value;
    var c6 = document.forms["exam_form"][C6].value;
    var c7 = document.forms["exam_form"][C7].value;
    var c8 = document.forms["exam_form"][C8].value;
    var c9 = document.forms["exam_form"][C9].value;
    var c10 = document.forms["exam_form"][C10].value;
    var c11 = document.forms["exam_form"][C11].value;
    var c12 = document.forms["exam_form"][C12].value;

    if(c1 + c2 + c3 + c4 + c5 + c6 + c7 + c8 + c9 + c10 + c11 + c12 != 100){
        alert("请确认章节比例总和是否为100%")
        return false;
    }

    return true;
}