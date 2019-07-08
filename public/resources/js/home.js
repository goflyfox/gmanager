//生成菜单
var menuItem = Vue.extend({
    name: 'menu-item',
    props: {item: {}},
    template: [
        '<li>',
        '<a v-if="item.type === 1" href="javascript:;">',
        '<i v-if="item.icon != null " :class="item.icon"></i><i v-else class="fa fa-circle-o"></i>',
        '<span>{{item.name}}</span>',
        '<i class="fa fa-angle-left pull-right"></i>',
        '</a>',
        '<ul v-if="item.type === 1" class="treeview-menu">',
        '<menu-item :item="item" v-for="item in item.childs"></menu-item>',
        '</ul>',
        '<a v-if="item.type === 2" :href="\'#\'+item.url"><i v-if="item.icon != null " :class="item.icon"></i><i v-else class="fa fa-circle-o"></i> {{item.name}}</a>',
        '</li>'

    ].join('')
});

//iframe自适应
$(window).on('resize', function () {
    var $content = $('.content');
    $content.height($(this).height() - 125);
    $content.find('iframe').each(function () {
        $(this).height($content.height());
    });
}).resize();

//注册菜单组件
Vue.component('menuItem', menuItem);

var vm = new Vue({
    el: '#rrapp',
    data: {
        user: {},
        menuList: {},
        main: dudu.ctx + "/admin/welcome.html",
        password: '',
        newPassword: '',
        navTitle: "控制台"
    },
    methods: {
        getMenuList: function () {
            dudu.get(dudu.ctx + "/system/user/menu?_" + $.now(), function (r) {
                if (r.code != 0) {
                    console.info("getMenuList fail:" + r)
                    return
                }

                var tmpMenus = [];
                for (var i = 0; i < r.data.length; i++) {
                    if (r.data[i].level == 1) {
                        r.data[i].childs = [];
                        tmpMenus.push(r.data[i]);
                    }
                }

                for (var j = 0; j < tmpMenus.length; j++) {
                    for (var i = 0; i < r.data.length; i++) {
                        if (r.data[i].level == 2 && r.data[i].parentid == tmpMenus[j].id) {
                            tmpMenus[j].childs.push(r.data[i]);
                        }
                    }
                }

                vm.menuList = tmpMenus;
                // 暂时不现实按钮权限
                // window.permissions = r.perms;
            });
        },
        getUser: function () {
            dudu.get(dudu.ctx + "/system/user/info?_" + $.now(), function (r) {
                if (r.code != 0) {
                    console.info("getUser fail:" + r)
                    return
                }

                vm.user = r.data;
                vm.getMenuList()

            });
        },
        updatePassword: function () {
            layer.open({
                type: 1,
                skin: 'layui-layer-molv',
                title: "修改密码",
                area: ['550px', '270px'],
                shadeClose: false,
                content: jQuery("#passwordLayer"),
                btn: ['修改', '取消'],
                btn1: function (index) {
                    if (vm.password == '') {
                        layer.alert('原密码不能为空');
                        return;
                    }
                    var newPassword = vm.newPassword || '';
                    if (newPassword == '') {
                        layer.alert('新密码不能为空');
                        return;
                    }

                    if (newPassword.length < 6) {
                        layer.alert('新密码长度必须大于6');
                        return;
                    }
                    var data = "password=" + hex_md5(vm.password) + "&newPassword=" + hex_md5(newPassword);
                    dudu.post(dudu.ctx + "/system/user/password", data, function (result) {
                        if (result.code == 0) {
                            layer.close(index);
                            layer.alert('修改成功', function (index) {
                                location.reload();
                            });
                        } else {
                            layer.alert(result.msg);
                        }
                    });
                }
            });
        }
        , updateProject: function () {
            var url = dudu.ctx + "/admin/project/data";
            $("#projectLayer [name='projectId']").empty().append('<option value="-1">--请选择项目--</option>');
            dudu.select(url, {
                select: $("#projectLayer [name='projectId']"),
                name: "name",
                value: "id",
                selected: 0
            });

            layer.open({
                type: 1,
                skin: 'layui-layer-molv',
                title: "修改项目",
                area: ['550px', '200px'],
                shadeClose: false,
                content: jQuery("#projectLayer"),
                btn: ['修改', '取消'],
                btn1: function (index) {
                    var projectId = $("#projectLayer [name='projectId']").val();

                    if (projectId == '' || projectId == -1) {
                        layer.alert('请选择项目');
                        return;
                    }
                    var data = "projectId=" + projectId;
                    $.ajax({
                        type: "POST",
                        url: dudu.ctx + "/system/user/project",
                        data: data,
                        dataType: "json",
                        success: function (result) {
                            if (result.code == 0) {
                                layer.close(index);
                                layer.alert('修改成功', function (index) {
                                    location.reload();
                                });
                            } else {
                                layer.alert(result.msg);
                            }
                        }
                    });
                }
            });
        }
        , logout: function () {
            dudu.get(dudu.ctx + "/user/logout?_" + $.now(), function (r) {
                if (r.code == 0) {
                    window.top.location.href = "login";
                } else {
                    alert(r.msg);
                    window.top.location.href = "login";
                }
            });
        }
    },
    created: function () {
        // this.getMenuList();
        this.getUser();
    },
    updated: function () {
        //路由
        var router = new Router();
        routerList(router, vm.menuList);
        router.start();
    }
});

function routerList(router, menuList) {
    for (var key in menuList) {
        var menu = menuList[key];
        if (menu.type == 1) {
            routerList(router, menu.childs);
        } else if (menu.type == 2) {
            router.add('#' + menu.url, function () {
                var url = window.location.hash;

                //替换iframe的url
                vm.main = url.replace('#', '');

                //导航菜单展开
                $(".treeview-menu li").removeClass("active");
                $("a[href='" + url + "']").parents("li").addClass("active");

                vm.navTitle = $("a[href='" + url + "']").text();
            });
        }
    }
}
