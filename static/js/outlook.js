function setBgColor(select, color){
    document.getElementById(select).style.backgroundColor = color;
}

function numFormat(num, id){
    res = num + "";
    while(num >= 10){
        num /= 10;
        n++;
    }
    while(n < 4){
        res = '0' + res;
        n++;
    }
    document.getElementById(id).innerHTML = res;
}