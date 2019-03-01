//幻灯片
var isSlide = true
function slide() {
    if(isSlide == false){
        return
    }
    $(".carousel-indicators > li").each(function (key,val) {
        var className = $(val).attr("class")
        var currentKey = 0;
        var nextKey = 0
        var liObj = $(".carousel-indicators > li");
        var divObj =  $(".carousel-inner > div");
        if(className == "active"){
            currentKey = key
            if(currentKey == liObj.length-1){
                nextKey = 0
            }else{
                nextKey = currentKey+1
            }
            divObj.eq(currentKey).removeClass("active")
            liObj.eq(currentKey).removeClass("active")
            liObj.eq(nextKey).addClass("active")
            divObj.eq(nextKey).addClass("active")
            return false
        }
    })
}

ref = setInterval(function(){
    slide();
},5000);


$(".carousel-indicators > li").click(function () {
    var currIndex = $(this).attr("data-index")
    var liObj = $(".carousel-indicators > li");
    var divObj =  $(".carousel-inner > div");
    divObj.removeClass("active")
    liObj.removeClass("active")
    liObj.eq(currIndex).addClass("active")
    divObj.eq(currIndex).addClass("active")
    isSlide = false
    window.setTimeout(function () {
        isSlide = true
    },3000);
})


$(".carousel-right").click(function () {
    var liObj = $(".carousel-indicators > li");
    var divObj =  $(".carousel-inner > div");
    var nextKey = 0
     liObj.each(function (key,val) {
        var className = $(val).attr("class")
        if(className == "active"){
            if(key == liObj.length-1){
                nextKey = 0
            }else{
                nextKey = key+1
            }
            divObj.eq(key).removeClass("active")
            liObj.eq(key).removeClass("active")
            liObj.eq(nextKey).addClass("active")
            divObj.eq(nextKey).addClass("active")
            return false
        }
    })
    isSlide = false
    window.setTimeout(function () {
        isSlide = true
    },3000);
})


$(".carousel-left").click(function () {
    var liObj = $(".carousel-indicators > li");
    var divObj =  $(".carousel-inner > div");
    var nextKey = 0
    liObj.each(function (key,val) {
        var className = $(val).attr("class")
        if(className == "active"){
            if(key < 0){
                nextKey = liObj.length-1
            }else{
                nextKey = key-1
            }
            divObj.eq(key).removeClass("active")
            liObj.eq(key).removeClass("active")
            liObj.eq(nextKey).addClass("active")
            divObj.eq(nextKey).addClass("active")
            return false
        }
    })
    isSlide = false
    window.setTimeout(function () {
        isSlide = true
    },3000);
})


