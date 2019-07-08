var vm = new Vue({
    el: '#data_wrapper',
    data: {
        showList: true,
        showEdit: false,
        showView: false,
        title: null,
        model: {
            projectId: null, // 项目ID
            projectName: null, // 项目名称
            version: null, // 版本
            content: null, // 内容
            enable: null, // 是否启用
            id: null
        }
    },
    methods: {
        showDiff: function () {
            diffview.diffUsingJS(0, vm.model.beforeContent, vm.model.content);
        },
        showPublicDiff: function () {
            diffview.diffUsingJS(0, vm.model.config_before, vm.model.config);
        },
        view: function (id) {
            var id = id || null;
            if (id == null) {
                Alert('请选择修改数据');
                return;
            }

            var url = dudu.ctx + "/admin/configpublic/get/" + id;
            dudu.get(url, function (result) {
                vm.showList = false;
                vm.showEdit = false;
                vm.showView = true;

                vm.title = "查看发布内容";
                vm.model = result.data;
            });
        },
        add: function () {
            var url = dudu.ctx + "/admin/configpublic/getProject";
            dudu.get(url, function (result) {
                vm.showList = false;
                vm.showEdit = true;
                vm.showView = false;

                vm.title = "配置发布";
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
                var url = dudu.ctx + "/admin/configpublic/delete/" + id;
                $.ajax({
                    type: "POST",
                    url: url,
                    success: function (result) {
                        if (result.code === 0) {
                            Alert('操作成功', function (index) {
                                vm.reload();
                            });
                        } else {
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

            Confirm('确定要发布数据？', function () {
                var url = dudu.ctx + "/admin/configpublic/save";
                $.ajax({
                    type: "POST",
                    url: url,
                    data: vm.model,
                    success: function (result) {
                        if (result.code === 0) {
                            Alert('操作成功', function (index) {
                                vm.reload();
                            });
                        } else {
                            ErrorInfo(result.msg);
                        }
                    }
                });
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
        url: dudu.ctx + "admin/configpublic/jqgrid",
        mtype: "POST",
        styleUI: 'Bootstrap',
        datatype: "json",
        colModel: [
            {label: "id", name: 'id', width: 75, hidden: true, key: true},
            {label: "项目名称", name: 'projectName', width: 200, sortable: true,},
            {label: "版本", name: 'version', width: 400, sortable: true},
            {label: "更新时间", name: 'updateTime', width: 240},
            {label: "更新人", name: 'updateName', width: 160, sortable: false},
            {label: "创建时间", name: 'createTime', width: 240},
            {label: "创建人", name: 'createName', width: 160, sortable: false},
            {
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
        rowList:[20,50,100,200,500],
        width: 1050,
        height: 630,
        rowNum: 20,
        caption: "配置发布内容表列表",
        pager: "#jqGridPager"
    });

    // 宽高自适应
    $("#jqGrid").setGridHeight($(window).height() - 160);
    $(window).resize(function () {
        $(window).unbind("onresize");
        $("#jqGrid").setGridHeight($(window).height() - 160).jqGrid('setGridWidth', $('#data_wrapper').width() - 5);
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