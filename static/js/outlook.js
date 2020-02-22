function setBgColor(select, color){
    document.getElementById(select).style.backgroundColor = color;
}

function numFormat(num){
    var res = num + "";
    var n = res.length;
    while(n < 4){
    res = '0' + res;
    n++;
    }
    return res;
}