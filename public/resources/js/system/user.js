var vm = new Vue({
    el: '#data_wrapper',
    data: {
        showList: true,
        showEdit: false,
        showView: false,
        showAuth: false,
        title: null,
        model: {
            uuid: null, // UUID
            username: null, // 登录名
            password: null, // 密码
            realName: null, // 真实姓名
            departId: null, // 部门
            userType: null, // 类型
            status: null, // 状态
            email: null, // email
            tel: null, // 手机号
            address: null, // 地址
            titleUrl: null, // 头像地址
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

            var url = dudu.ctx + "/system/user/get/" + id;
            dudu.get(url, function (result) {
                vm.showList = false;
                vm.showEdit = false;
                vm.showView = true;

                vm.title = "查看用户";
                vm.model = result.data;
            });
        },
        add: function () {
            vm.showList = false;
            vm.showEdit = true;
            vm.showView = false;

            var url = dudu.ctx + "/system/department/list";
            $("#showEdit #departId").empty().append('<option value="-1">--请选择部门--</option>');
            dudu.select(url, {
                select: $("#showEdit #departId"),
                name: "name",
                value: "id",
                selected: 0
            });

            vm.title = "新增用户";
            vm.model = {
                uuid: null, // UUID
                username: null, // 登录名
                password: null, // 密码
                realName: null, // 真实姓名
                departId: null, // 部门
                userType: 2, // 类型
                email: null, // email
                tel: null, // 手机号
                address: null, // 地址
                titleUrl: null, // 头像地址
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

            var url = dudu.ctx + "/system/user/get/" + id;
            dudu.get(url, function (result) {
                vm.showList = false;
                vm.showEdit = true;
                vm.showView = false;

                var url = dudu.ctx + "/system/department/list";
                $("#showEdit #departId").empty().append('<option value="-1">--请选择部门--</option>');
                dudu.select(url, {
                    select: $("#showEdit #departId"),
                    name: "name",
                    value: "id",
                    selected: vm.model.departId
                }, false);

                vm.title = "修改用户";
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
                var url = dudu.ctx + "/system/user/delete/" + id;
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

            var url = dudu.ctx + "/system/user/save";
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
        auth: function (userId) {
            var url = dudu.ctx + "system/user/roleinfo";
            dudu.post(url, {"userId": userId}, function (result) {
                if (result.code == 0) {
                    vm.showList = false;
                    vm.showAuth = true;
                    vm.title = "授权用户";

                    $('#showAuth [name="userid"]').val(userId);
                    $('#showAuth #roleids').val(result.data.roleIds);

                    var roleList = result.data.list;
                    var roleLi = '';
                    for (var i = 0; i < roleList.length; i++) {
                        roleLi += '<li class="list-group-item"><div class="checkbox"><label>';
                        roleLi += '<input type="checkbox" name="roleid" id="roleid_' + roleList[i]["id"] + '" value="' + roleList[i]["id"] + '">&nbsp;&nbsp;&nbsp;&nbsp;' + roleList[i]["name"] + '';
                        roleLi += '</label></div></li>';
                    }
                    $("#showAuth ul").html(roleLi);

                    // 初始化，设置被选中
                    var roleids = $("#roleids").val().split(",");
                    for (var i = 0; i < roleids.length; i++) {
                        if (roleids[i] != '') $("[id='roleid_" + roleids[i] + "']").attr("checked", true);
                    }
                } else {
                    ErrorInfo('操作失败：' + result.msg);
                }
            })
        },
        saveAuth: function () {
            var ids = "";
            $('#showAuth input[name="roleid"]:checked').each(function () {
                ids += $(this).val() + ',';
            });
            if (ids != "") {
                ids = ids.substring(0, ids.length - 1);
            }

            $('#showAuth [name="roleids"]').val(ids);

            var title = '确认要保存该数据么？';
            var url = dudu.ctx + 'system/user/rolesave';
            var params = $("#showAuth form").serialize();
            Confirm(title, function () {
                dudu.post(url, params, function (result) {
                    if (result.code == 0) {
                        vm.reload();
                    } else {
                        ErrorInfo('操作失败：' + result.msg);
                    }
                });
            });
        },
        init: function () {
            var url = dudu.ctx + "/system/department/list";
            $("#showList [name='departId']").empty().append('<option value="-1">--请选择部门--</option>');
            dudu.select(url, {
                select: $("#showList [name='departId']"),
                name: "name",
                value: "id",
                selected: 0
            });
        },
        reload: function () {
            vm.showList = true;
            vm.showEdit = false;
            vm.showView = false;
            vm.showAuth = false;

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
        url: dudu.ctx + "system/user/jqgrid",
        mtype: "POST",
        loadBeforeSend: dudu.headToken,
        loadComplete: dudu.loadAuth,
        styleUI: 'Bootstrap',
        datatype: "json",
        colModel: [
            {label: "ID", name: 'id', width: 75, key: true},
            {label: "部门", name: 'departName', width: 120, sortable: true, sortable: false},
            {label: "用户名", name: 'username', width: 120, sortable: true},
            {label: "真实姓名", name: 'realName', width: 120, sortable: true},
            {
                label: "用户类型",
                name: 'userType',
                width: 120,
                sortable: true,
                formatter: function (cellValue, options, rowObject) {
                    var str = "";
                    if (cellValue == 1) {
                        str = "管理员";
                    } else if (cellValue == 2) {
                        str = "普通用户";
                    } else if (cellValue == 3) {
                        str = "前台用户";
                    } else if (cellValue == 4) {
                        str = "第三方用户";
                    } else if (cellValue == 5) {
                        str = "API用户";
                    }
                    return str;
                }
            },
            {label: "手机号", name: 'tel', width: 120, sortable: true},
            {label: "更新时间", name: 'updateTime', width: 240},
            {label: "更新人", name: 'updateName', width: 160, sortable: false},
            {label: "创建时间", name: 'createTime', width: 240},
            {label: "创建人", name: 'createName', width: 160, sortable: false},
            {
                name: '操作', index: '', width: 280, fixed: true, sortable: false, resize: false,
                formatter: function (cellvalue, options, rowObject) {
                    var replaceStr = "\\[id\\]";
                    var replaceShowStr = "\\[show\\]";
                    var buttonStr = editStr;
                    try {
                        buttonStr = buttonStr.replace(/\r\n/g, "");
                        buttonStr = buttonStr.replace(/\n/g, "");
                        buttonStr = buttonStr.replace(new RegExp(replaceStr, 'gm'), rowObject.id);
                        if (rowObject.userType == 1) {
                            buttonStr = buttonStr.replace(new RegExp(replaceShowStr, 'gm'), "auth_hide");
                        } else {
                            buttonStr = buttonStr.replace(new RegExp(replaceShowStr, 'gm'), "");
                        }
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
        caption: "用户列表",
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