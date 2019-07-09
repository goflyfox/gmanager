var vm = new Vue({
    el: '#data_wrapper',
    data: {
        showList: true,
        showEdit: false,
        showView: false,
        title: null,
        model: {
            parentId: null, // 上级机构
            name: null, // 部门
            code: null, // 机构编码
            sort: null, // 序号
            linkman: null, // 联系人
            linkmanNo: null, // 联系人电话
            remark: null, // 机构描述
            enable: 1, // 是否启用
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

            var url = dudu.ctx + "/system/department/get/" + id;
            dudu.get(url, function (result) {
                vm.showList = false;
                vm.showEdit = false;
                vm.showView = true;

                vm.title = "查看部门";
                vm.model = result.data;
            });
        },
        add: function (selectedId) {
            selectedId = selectedId || 0;
            vm.showList = false;
            vm.showEdit = true;
            vm.showView = false;

            var url = dudu.ctx + "/system/department/list";
            $("#parentId").empty().append('<option value="-1">--请选择类型--</option>');
            dudu.select(url, {
                select: $("#parentId"),
                name: "name",
                value: "id",
                selected: 0
            });

            vm.title = "新增部门";
            vm.model = {
                parentId: selectedId, // 上级机构
                name: null, // 部门
                code: null, // 机构编码
                sort: null, // 序号
                linkman: null, // 联系人
                linkmanNo: null, // 联系人电话
                remark: null, // 机构描述
                enable: 1, // 是否启用
                id: null
            };
        },
        update: function (id) {
            var id = id || null;
            if (id == null) {
                Alert('请选择修改数据');
                return;
            }

            var url = dudu.ctx + "/system/department/get/" + id;
            dudu.get(url, function (result) {
                vm.showList = false;
                vm.showEdit = true;
                vm.showView = false;

                var url = dudu.ctx + "/system/department/list";
                $("#parentId").empty().append('<option value="-1">--请选择类型--</option>');
                dudu.select(url, {
                    select: $("#parentId"),
                    name: "name",
                    value: "id",
                    selected: result.data.parentId,
                    exclude: result.data.id
                }, false);

                vm.title = "修改部门";
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
                var url = dudu.ctx + "/system/department/delete/" + id;
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

            var url = dudu.ctx + "/system/department/save";
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

            ztree.refresh();

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

var ztree = {
    zTreeObj: null,
    selected: null,
    setting: {
        check: {
            enable: false
        },
        view: {
            addHoverDom: function (treeId, treeNode) {
                ztree.addHoverDom(treeId, treeNode);
            },
            removeHoverDom: function (treeId, treeNode) {
                ztree.removeHoverDom(treeId, treeNode);
            },
            dblClickExpand: false,
            showLine: true,
            selectedMulti: false
        },
        data: {
            simpleData: {
                enable: true,
                idKey: "id",
                pIdKey: "pId",
                rootPId: ""
            }
        },
        callback: {
            beforeClick: function (treeId, treeNode) {
                // 设置选中ID
                // ztree.selected = treeNode.id;
                // oper.jqgrid(treeNode.id);
                var zTree = $.fn.zTree.getZTreeObj("tree");
                // $("#parentid").val(treeNode.id);
                // vm.reload();
                if (treeNode.isParent) {
                    zTree.expandNode(treeNode);
                    return false;
                } else {
                    return true;
                }
            }
        }
    }
    , addHoverDom: function (treeId, treeNode) {
        var sObj = $("#" + treeNode.tId + "_span");
        if (treeNode.editNameFlag || $("#addBtn_" + treeNode.tId).length > 0) return;
        var addStr = "<span class='button add' id='addBtn_" + treeNode.tId + "'></span>";
        addStr += "<span class='button edit' id='editBtn_" + treeNode.tId + "'></span>";
        addStr += "<span class='button remove' id='removeBtn_" + treeNode.tId
            + "' title='add node' onfocus='this.blur();'></span>";
        sObj.after(addStr);

        var addBtn = $("#addBtn_" + treeNode.tId);
        if (addBtn) addBtn.bind("click", function () {
            vm.add(treeNode.id);
            return false;
        });

        var editBtn = $("#editBtn_" + treeNode.tId);
        if (editBtn) editBtn.bind("click", function () {
            vm.update(treeNode.id);
            return false;
        });

        var delBtn = $("#removeBtn_" + treeNode.tId);
        if (delBtn) delBtn.bind("click", function () {
            vm.del(treeNode.id);
            return false;
        });

    }
    , removeHoverDom: function (treeId, treeNode) {
        $("#addBtn_" + treeNode.tId).unbind().remove();
        $("#editBtn_" + treeNode.tId).unbind().remove();
        $("#removeBtn_" + treeNode.tId).unbind().remove();
    }
    , refresh: function () {
        var url = "system/department/list";
        dudu.post(url, "", function (treeData) {
            var zNodes = new Array();
            // zNodes.push({
            //     id: 0,
            //     pId: -1,
            //     name: "根节点",
            //     open: true // (tmp.parentId == 0)
            // });
            for (i in treeData.data) {
                var tmp = treeData.data[i];
                var obj = {
                    id: tmp.id,
                    pId: tmp.parentId,
                    name: tmp.name,
                    open: true // (tmp.parentId == 0)
                };
                zNodes.push(obj);
            }

            var $tree = $("#tree");
            var zTreeObj = $.fn.zTree.init($tree, ztree.setting, zNodes);
            //var zTree = $.fn.zTree.getZTreeObj("tree");
            // 选中之前选中的节点
            if (ztree.selected != null) {
                zTreeObj.selectNode(zTreeObj.getNodeByParam("id", ztree.selected));
            }
            // 设置
            ztree.zTreeObj = zTreeObj;
        });
    }
};

// 初始化
jQuery(function ($) {
    // 刷新ztree
    ztree.refresh();

    // 加载jqgrid
    var editStr = $('#jqGridEdit').html();
    $('#jqGrid').jqGrid({
        url: dudu.ctx + "system/department/jqgrid",
        mtype: "POST",
        loadBeforeSend: dudu.headToken,
        loadComplete: dudu.loadAuth,
        styleUI: 'Bootstrap',
        datatype: "json",
        colModel: [
            {label: "id", name: 'id', width: 75, hidden: true, key: true},
            {
                label: "上级机构", name: 'parentName', width: 300, sortable: false,
                formatter: function (cellvalue, options, rowObject) {
                    cellvalue = cellvalue || '';
                    return cellvalue == '' ? '根节点' : cellvalue;
                }
            },
            {label: "机构名称", name: 'name', width: 300},
            {label: "排序号", name: 'sort', width: 120},
            {label: "机构编码", name: 'code', width: 120},
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
        rowList: [20, 50, 100, 200, 500],
        width: 1050,
        height: 630,
        rowNum: 20,
        caption: "部门列表",
        pager: "#jqGridPager"
    });

    // 宽高自适应
    $(window).resize(function () {
        $(window).unbind("onresize");
        $("#jqGrid").setGridHeight($(window).height() - 165).jqGrid('setGridWidth', $('#data_content').width() - 5);
        $("#tree").height($(window).height() - 165);
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