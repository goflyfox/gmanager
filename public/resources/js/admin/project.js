var vm = new Vue({
    el:'#data_wrapper',
    data:{
        showList: true,
        showEdit: false,
        showView: false,
        title: null,
        model:{
        name:null, // 项目名称
        type:null, // 类型
        sort:null, // 排序
        remark:null, // 备注
        enable:null, // 是否启用
        id:null
        }
},
methods: {
    view: function(id){
        var id = id || null;
        if(id == null){
            Alert('请选择修改数据');
            return ;
        }

        var url = dudu.ctx + "/admin/project/get/" + id;
        dudu.get(url, function(result){
            vm.showList = false;
            vm.showEdit = false;
            vm.showView = true;

            vm.title = "查看项目表";
            vm.model = result.data;
        });
    },
    add: function(){
        vm.showList = false;
        vm.showEdit = true;
        vm.showView = false;

        vm.title = "新增项目表";
        vm.model = {
            name:null, // 项目名称
            type:1, // 类型
            sort:10, // 排序
            remark:null, // 备注
            enable:1, // 是否启用
            id:null
        };
    },
    update: function (id) {
        var id = id || null;
        if(id == null){
            Alert('请选择修改数据');
            return ;
        }

        var url = dudu.ctx + "/admin/project/get/" + id;
        dudu.get(url, function(result){
            vm.showList = false;
            vm.showEdit = true;
            vm.showView = false;

            vm.title = "修改项目表";
            vm.model = result.data;
        });
    },
    del: function (id) {
        var id = id || null;
        if(id == null){
            Alert('请选择删除数据');
            return ;
        }

        Confirm('确定要删除选中的记录？', function(){
            var url = dudu.ctx + "/admin/project/delete/" + id;
            $.ajax({
                type: "POST",
                url: url,
                success: function(result){
                    if(result.code === 0){
                        Alert('操作成功', function(index){
                            vm.reload();
                        });
                    }else{
                        ErrorInfo(result.msg);
                    }
                }
            });
        });
    },
    save: function () {
        if (!validForm()) {
            return;
        }

        var url = dudu.ctx + "/admin/project/save";
        $.ajax({
            type: "POST",
            url: url,
            data: vm.model,
            success: function(result){
                if(result.code === 0){
                    Alert('操作成功', function(index){
                        vm.reload();
                    });
                }else{
                    ErrorInfo(result.msg);
                }
            }
        });
    },
    reload: function () {
        vm.showList = true;
        vm.showEdit = false;
        vm.showView = false;

        var fields = $("#showList form").serializeArray();
        var jsonData = {};
        jQuery.each( fields, function(i, field){
            jsonData[field.name]=field.value;
        });

        $('#jqGrid').jqGrid('setGridParam', {
            postData : jsonData
        }).trigger('reloadGrid');
    }
}
});

// 初始化
jQuery(function($) {
    // 加载jqgrid
    var editStr = $('#jqGridEdit').html();
    $('#jqGrid').jqGrid({
        url:dudu.ctx + "admin/project/jqgrid",
        mtype: "POST",
        styleUI : 'Bootstrap',
        datatype: "json",
        colModel: [
            {label: "id",name: 'id',width: 75,hidden:true,key:true},
            {label: "项目名称",name: 'name',width: 120,sortable:true},
            {label: "排序",name: 'sort',width: 120,sortable:true},
            {label: "备注",name: 'remark',width: 120,sortable:true},
            {
                label: "是否启用", name: 'enable', width: 180, sortable: true,
                formatter: function (cellValue, options, rowObject) {
                    var str = "";
                    if (cellValue == 1) {
                        str = '<span class="label label-success">启用</span>';
                    } else if (cellValue == 2) {
                        str = '<span class="label label-warning">禁用</span>';
                    } else {
                        str = "";
                    }
                    return str;
                }
            },
            {label: "更新时间", name: 'updateTime', width: 240},
            {label: "更新人", name: 'updateName', width: 160, sortable: false},
            {label: "创建时间", name: 'createTime', width: 240},
            {label: "创建人", name: 'createName', width: 160, sortable: false},
            {
                name: '操作', index: '', width: 200, fixed: true, sortable: false, resize: false,
                formatter: function(cellvalue, options, rowObject) {
                    var replaceStr = "\\[id\\]";
                    var buttonStr = editStr;
                    try{
                        buttonStr = buttonStr.replace(/\r\n/g,"");
                        buttonStr = buttonStr.replace(/\n/g,"");
                        buttonStr = buttonStr.replace(new RegExp(replaceStr,'gm'),rowObject.id );
                    }catch(e) {
                        alert(e.message);
                    }
                    return buttonStr ;
                }
            }
        ],
        rownumbers: true,
        sortname: 'id',
        sortorder: 'desc',
        viewrecords: true,
        autowidth: true,
        rowList:[20,50,100,200,500],
        width: 1050,
        height: 630,
        rowNum: 20,
        caption: "项目表列表",
        pager: "#jqGridPager"
    });

    // 宽高自适应
    $("#jqGrid").setGridHeight($(window).height() - 160);
    $(window).resize(function(){
        $(window).unbind("onresize");
        $("#jqGrid").setGridHeight($(window).height() - 160).jqGrid('setGridWidth', $('#data_wrapper').width() - 5);
        $(window).bind("onresize", this);
    }).resize();

    $('#jqGrid').jqGrid('navGrid',"#jqGridPager", {
        search: false, // show search button on the toolbar
        add: false,
        edit: false,
        del: false,
        refresh: true,
        view: false
    });

});