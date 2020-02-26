function judgeEmpty(f, p, q){
    var x = document.forms[f][p].value;
    if(x == ""){
        alert("请输入" + q);
        return false;
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