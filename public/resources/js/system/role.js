// 权限
var menuZtree = {
    tree: null,
    setting: {
        check: {
            enable: true,
            nocheckInherit: true
        },
        data: {
            simpleData: {
                enable: true,
                idKey: "id",
                pIdKey: "pId",
                rootPId: ""
            }
        },
    },
    selected: function (roleId) {
        dudu.post(dudu.ctx + "/system/role/info", {"roleId": roleId}, function (r) {
            //勾选角色所拥有的菜单
            var menuIds = r.data.menus;
            for (var i = 0; i < menuIds.length; i++) {
                var node = menuZtree.tree.getNodeByParam("id", menuIds[i]);
                menuZtree.tree.checkNode(node, true, false);
            }
        });
    },
    refresh: function () {
        var url = dudu.ctx + "system/menu/list";
        dudu.get(url, function (treeData) {
            var zNodes = new Array();
            for (i in treeData.data) {
                var tmp = treeData.data[i];
                var obj = {
                    id: tmp.id,
                    pId: tmp.parentid,
                    name: tmp.name,
                    open: (tmp.type == 1) ? true : false // (tmp.parentId == 0)
                };
                zNodes.push(obj);
            }

            menuZtree.tree = $.fn.zTree.init($("#menuTree"), menuZtree.setting, zNodes);
        });
    }
};

var vm = new Vue({
    el: '#data_wrapper',
    data: {
        showList: true,
        showEdit: false,
        showView: false,
        title: null,
        model: {
            name: null, // 名称
            status: null, // 状态
            sort: null, // 排序
            remark: null, // 说明
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

            var url = dudu.ctx + "/system/role/get/" + id;
            dudu.get(url, function (result) {
                vm.showList = false;
                vm.showEdit = false;
                vm.showView = true;

                vm.title = "查看角色";
                vm.model = result.data;
            });
        },
        add: function () {
            vm.showList = false;
            vm.showEdit = true;
            vm.showView = false;

            // 展示菜单树
            menuZtree.refresh();

            vm.title = "新增角色";
            vm.model = {
                name: null, // 名称
                status: 1, // 状态
                sort: 10, // 排序
                remark: null, // 说明
                id: null
            };
        },
        update: function (id) {
            var id = id || null;
            if (id == null) {
                Alert('请选择修改数据');
                return;
            }

            // 展示菜单树
            menuZtree.refresh();

            var url = dudu.ctx + "/system/role/get/" + id;
            dudu.get(url, function (result) {
                vm.showList = false;
                vm.showEdit = true;

                vm.showView = false;
                // 选中节点
                menuZtree.selected(id);

                vm.title = "修改角色";
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
                var url = dudu.ctx + "/system/role/delete/" + id;
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

            //获取选择的菜单
            var nodes = menuZtree.tree.getCheckedNodes(true);
            var menuIdArray = new Array();
            for (var i = 0; i < nodes.length; i++) {
                menuIdArray.push(nodes[i].id);
            }
            var menus = menuIdArray.join(",");
            vm.model.menus = menus;

            var url = dudu.ctx + "/system/role/save";
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
        url: dudu.ctx + "system/role/jqgrid",
        mtype: "POST",
        loadBeforeSend: dudu.headToken,
        loadComplete: dudu.loadAuth,
        styleUI: 'Bootstrap',
        datatype: "json",
        colModel: [
            {label: "id", name: 'id', width: 75, hidden: true, key: true},
            {label: "名称", name: 'name', width: 120, sortable: true},
            {
                label: "状态",
                name: 'status',
                width: 120,
                sortable: true,
                formatter: function (cellValue, options, rowObject) {
                    var str = "";
                    if (cellValue == 1) {
                        str = "显示";
                    } else {
                        str = "隐藏";
                    }
                    return str;
                }
            },
            {label: "排序", name: 'sort', width: 120, sortable: true},
            {label: "说明", name: 'remark', width: 120, sortable: true},
            {label: "更新时间", name: 'updateTime', width: 240},
            {label: "更新人", name: 'updateName', width: 160, sortable: false},
            {label: "创建时间", name: 'createTime', width: 240},
            {label: "创建人", name: 'createName', width: 160, sortable: false},
            {
                name: '操作', index: '', width: 280, fixed: true, sortable: false, resize: false,
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
        caption: "角色列表",
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