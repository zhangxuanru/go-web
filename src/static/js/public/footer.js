
$(".top_page_go").click(function () {
    var topPage = $("#top_pagination").val()
    gotoPage(topPage)
})

$(".top_page_go_next").click(function () {
    var topPage = $("#top_pagination").val()
    topPage   = parseInt(topPage)
    var  nextPage = topPage+1
    gotoPage(nextPage)
})

$(".top_page_go_prev").click(function () {
    var topPage = $("#top_pagination").val()
    topPage   = parseInt(topPage)
    var prevPage = topPage-1
    gotoPage(prevPage)
})

$(".bottom-btn-go").click(function () {
    var page = $("#bottom_pagination").val()
    gotoPage(page)
})

$(".bottom_page_go_next").click(function () {
    var page = $("#bottom_pagination").val()
    page = parseInt(page)
    var nextPage = page+1
    gotoPage(nextPage)
})

$(".bottom_page_prev").click(function () {
    var page = $("#bottom_pagination").val()
    page = parseInt(page)
    var prevPage = page-1
    gotoPage(prevPage)
})


function gotoPage(page) {
    page  = parseInt(page)
    if(page > 0 && page <= totalPage){
        var url = window.location.href
        if (url.indexOf("page") > -1){
            url =  url.replace("page="+currPage,"page="+page)
        }else{
            if (url.indexOf("?") > -1){
                url = url+"&page="+page
            }else{
                url = url+"?page="+page
            }
        }
        window.location.href=url
    }
}
