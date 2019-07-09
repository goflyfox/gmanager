var vm = new Vue({
    el: '#data_wrapper',
    data: {
        showList: true,
        showEdit: false,
        showView: false,
        title: null,
        model: {
            logType: null, // 类型
            operObject: null, // 操作对象
            operTable: null, // 操作表
            operId: null, // 操作主键
            operType: null, // 操作类型
            operRemark: null, // 操作备注
            enable: null, // 是否启用
            updateTime: null, // 更新时间
            updateId: null, // 更新人
            createTime: null, // 创建时间
            createId: null, // 创建者
            id: null
        }
    },
    methods: {
        view: function (id) {
            var id = id || null;
            if (id == null) {
                Alert('请选择修改数据');
                return;
            }

            var url = dudu.ctx + "/system/log/get/" + id;
            dudu.get(url, function (result) {
                vm.showList = false;
                vm.showEdit = false;
                vm.showView = true;

                vm.title = "查看日志";
                vm.model = result.data;
            });
        },
        reload: function () {
            vm.showList = true;
            vm.showEdit = false;
            vm.showView = false;

            var fields = $("#showList form").serializeArray();
            var jsonData = {};
            jQuery.each(fields, function (i, field) {
                jsonData[field.name] = field.value;
            });

            $('#jqGrid').jqGrid('setGridParam', {
                postData: jsonData
            }).trigger('reloadGrid');
        }
    }
});

// 初始化
jQuery(function ($) {
    // 加载jqgrid
    var editStr = $('#jqGridEdit').html();
    $('#jqGrid').jqGrid({
        url: dudu.ctx + "system/log/jqgrid",
        mtype: "POST",
        loadBeforeSend: dudu.headToken,
        loadComplete: dudu.loadAuth,
        styleUI: 'Bootstrap',
        datatype: "json",
        colModel: [
            {label: "id", name: 'id', width: 75, hidden: true, key: true},
            {
                label: "类型",
                name: 'logType',
                width: 120,
                sortable: true,
                formatter: function (cellValue, options, rowObject) {
                    var str = "";
                    if (cellValue == 1) {
                        str = "系统类型";
                    } else {
                        str = "操作类型";
                    }
                    return str;
                }
            },
            {label: "操作对象", name: 'operObject', width: 120, sortable: true},
            {label: "操作表", name: 'operTable', width: 120, sortable: true},
            {label: "操作主键", name: 'operId', width: 120, sortable: true},
            {label: "操作类型", name: 'operType', width: 120, sortable: true},
            {label: "操作备注", name: 'operRemark', width: 120, sortable: true, hidden: true},
            {label: "是否启用", name: 'enable', width: 120, sortable: true, hidden: true},
            {label: "更新时间", name: 'updateTime', width: 240, hidden: true},
            {label: "更新人", name: 'updateName', width: 160, sortable: false, hidden: true},
            {label: "创建时间", name: 'createTime', width: 240},
            {label: "创建人", name: 'createName', width: 160, sortable: false}
            ,{
                name: '操作', index: '', width: 200, fixed: true, sortable: false, resize: false,
                formatter: function (cellvalue, options, rowObject) {
                    var replaceStr = "\\[id\\]";
                    var buttonStr = editStr;
                    try {
                        buttonStr = buttonStr.replace(/\r\n/g, "");
                        buttonStr = buttonStr.replace(/\n/g, "");
                        buttonStr = buttonStr.replace(new RegExp(replaceStr, 'gm'), rowObject.id);
                    } catch (e) {
                        alert(e.message);
                    }
                    return buttonStr;
                }
            }
        ],
        rownumbers: true,
        sortname: 'id',
        sortorder: 'desc',
        viewrecords: true,
        autowidth: true,
        rowList: [20, 50, 100, 200, 500],
        width: 1050,
        height: 630,
        rowNum: 20,
        caption: "日志列表",
        pager: "#jqGridPager"
    });

    // 宽高自适应
    $("#jqGrid").setGridHeight($(window).height() - 165);
    $(window).resize(function () {
        $(window).unbind("onresize");
        $("#jqGrid").setGridHeight($(window).height() - 165).jqGrid('setGridWidth', $('#data_wrapper').width() - 5);
        $(window).bind("onresize", this);
    }).resize();

    $('#jqGrid').jqGrid('navGrid', "#jqGridPager", {
        search: false, // show search button on the toolbar
        add: false,
        edit: false,
        del: false,
        refresh: true,
        view: false
    });

});