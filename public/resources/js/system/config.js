var vm = new Vue({
    el: '#data_wrapper',
    data: {
        showList: true,
        showEdit: false,
        showView: false,
        title: null,
        operType: 1,
        model: {
            name: null, // 名称
            key: null, // 键
            value: null, // 值
            code: null, // 编码
            parentId: null, // 类型
            sort: 10, // 排序号
            enable: 1, // 是否启用
            id: null,
            copyStatus: 1,//拷贝
            changeStatus: 2//是否可以更新
        }
    },
    methods: {
        view: function (id) {
            var id = id || null;
            if (id == null) {
                Alert('请选择修改数据');
                return;
            }

            var url = dudu.ctx + "/system/config/get/" + id;
            dudu.get(url, function (result) {
                vm.showList = false;
                vm.showEdit = false;
                vm.showView = true;

                vm.title = "查看系统配置表";
                vm.model = result.data;
            });
        },
        add: function () {
            vm.showList = false;
            vm.showEdit = true;
            vm.showView = false;

            var url = dudu.ctx + "/system/config/type";
            $("#showEdit [name='parentId']").empty().append('<option value="-1">--请选择类型--</option>');
            dudu.select(url, {
                select: $("#showEdit [name='parentId']"),
                name: "name",
                value: "id",
                selected: 0
            });

            vm.title = "新增系统配置表";
            vm.model = {
                name: null, // 名称
                key: null, // 键
                value: null, // 值
                code: null, // 编码
                parentId: null, // 类型
                sort: 10, // 排序号
                enable: 1, // 是否启用
                id: null,
                copyStatus: 1,//拷贝
                changeStatus: 2
            };
        },
        update: function (id) {
            var id = id || null;
            if (id == null) {
                Alert('请选择修改数据');
                return;
            }

            var url = dudu.ctx + "/system/config/get/" + id;
            dudu.get(url, function (result) {
                vm.showList = false;
                vm.showEdit = true;
                vm.showView = false;

                if (result.data.parentId == 0) {
                    vm.operType = 2;
                } else {
                    vm.operType = 1;
                }

                var url = dudu.ctx + "/system/config/type";
                $("#showEdit [name='parentId']").empty().append('<option value="-1">--请选择类型--</option>');
                dudu.select(url, {
                    select: $("#showEdit [name='parentId']"),
                    name: "name",
                    value: "id",
                    selected: result.data.parentId
                }, false);

                vm.title = "修改系统配置表";
                vm.model = result.data;
            });
        },
        del: function (id) {
            var id = id || null;
            if (id == null) {
                Alert('请选择删除数据');
                return;
            }

            Confirm('确定要删除选中的记录？', function () {
                var url = dudu.ctx + "/system/config/delete/" + id;
                dudu.post(url, null, function (result) {
                    if (result.code === 0) {
                        Alert('操作成功', function (index) {
                            vm.reload();
                        });
                    } else {
                        ErrorInfo(result.msg);
                    }
                });
            });
        },
        save: function () {
            if (!validForm()) {
                return;
            }

            var url = dudu.ctx + "/system/config/save";
            dudu.post(url, vm.model, function (result) {
                if (result.code === 0) {
                    Alert('操作成功', function (index) {
                        vm.reload();
                    });
                } else {
                    ErrorInfo(result.msg);
                }
            });
        },
        init: function () {
            var url = dudu.ctx + "/system/config/type";
            $("#showList [name='parentId']").empty().append('<option value="-1">--请选择类型--</option>');
            dudu.select(url, {
                select: $("#showList [name='parentId']"),
                name: "name",
                value: "id",
                selected: 0
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
    vm.init();
    // 加载jqgrid
    var editStr = $('#jqGridEdit').html();
    $('#jqGrid').jqGrid({
        url: dudu.ctx + "system/config/jqgrid",
        mtype: "POST",
        loadBeforeSend: dudu.headToken,
        loadComplete: dudu.loadAuth,
        styleUI: 'Bootstrap',
        datatype: "json",
        colModel: [
            {label: "id", name: 'id', width: 75, hidden: true, key: true},
            {
                label: "类型", name: 'typeName', width: 120, sortable: true,
                formatter: function (cellvalue, options, rowObject) {
                    if (cellvalue == '' || cellvalue == null) {
                        return '类型'
                    }
                    return cellvalue;
                }
            },
            {label: "名称", name: 'name', width: 200, sortable: true},
            {label: "键", name: 'key', width: 200, sortable: true},
            {label: "值", name: 'value', width: 200, sortable: true},
            {label: "编码", name: 'code', width: 120, sortable: true},
            {label: "排序号", name: 'sort', width: 120, sortable: true},
            {label: "更新时间", name: 'updateTime', width: 240},
            {label: "更新人", name: 'updateName', width: 160, sortable: false},
            {label: "创建时间", name: 'createTime', width: 240},
            {label: "创建人", name: 'createName', width: 160, sortable: false},
            {
                name: '操作', index: '', width: 140, fixed: true, sortable: false, resize: false,
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
        caption: "系统配置表列表",
        pager: "#jqGridPager"
    });

    // 宽高自适应
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