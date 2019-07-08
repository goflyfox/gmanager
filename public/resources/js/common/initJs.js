jQuery(function($) {
    // 查询
    mySearch();

});

function mySearch(){
    // 回车绑定查询按钮
    $('.tableSearch').on('keydown', function (e) {
        var key = e.which;
        if (key == 13 && $(":input[name='search']").length > 0) {
            e.preventDefault();
            $(":input[name='search']").click();
        }
    });
}

