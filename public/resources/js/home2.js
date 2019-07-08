layui.config({
    base: dudu.ctx + '/resources/js/home/'
}).use(['navtab'], function () {
    window.jQuery = window.$ = layui.jquery;
    window.layer = layui.layer;
    var element = layui.element(),
        navtab = layui.navtab({
            elem: '.larry-tab-box'
        });

    //iframe自适应
    $(window).on('resize', function () {
        var $content = $('#larry-tab .layui-tab-content');
        $content.height($(this).height() - 140);
        $content.find('iframe').each(function () {
            $(this).height($content.height());
        });
    }).resize();

    $(function () {
        $('#larry-nav-side').click(function () {
            if ($(this).attr('lay-filter') !== undefined) {
                $(this).children('ul').find('li').each(function () {
                    var $this = $(this);
                    if ($this.find('dl').length > 0) {
                        var $dd = $this.find('dd').each(function () {
                            $(this).on('click', function () {
                                var $a = $(this).children('a');
                                var href = $a.data('url');
                                var icon = $a.children('i:first').data('icon');
                                var title = $a.children('span').text();
                                var data = {
                                    href: href,
                                    icon: icon,
                                    title: title
                                }
                                navtab.tabAdd(data);
                            });
                        });
                    } else {
                        $this.on('click', function () {
                            var $a = $(this).children('a');
                            var href = $a.data('url');
                            var icon = $a.children('i:first').data('icon');
                            var title = $a.children('span').text();
                            var data = {
                                href: href,
                                icon: icon,
                                title: title
                            }
                            navtab.tabAdd(data);
                        });
                    }
                });
            }
        }).trigger("click");
    });
});


layui.use(['jquery', 'layer', 'element'], function () {
    window.jQuery = window.$ = layui.jquery;
    window.layer = layui.layer;
    var element = layui.element();

    // larry-side-menu向左折叠
    $('.larry-side-menu').click(function () {
        var sideWidth = $('#larry-side').width();
        if (sideWidth === 200) {
            $('#larry-body').animate({
                left: '0'
            });
            $('#larry-footer').animate({
                left: '0'
            });
            $('#larry-side').animate({
                width: '0'
            });

            $('.larry-side-menu i').removeClass('fa-arrow-left').addClass('fa-arrow-right');
        } else {
            $('#larry-body').animate({
                left: '200px'
            });
            $('#larry-footer').animate({
                left: '200px'
            });
            $('#larry-side').animate({
                width: '200px'
            });

            $('.larry-side-menu i').addClass('fa-arrow-left').removeClass('fa-arrow-right');
        }
    });
});

//生成菜单
var menuItem = Vue.extend({
    name: 'menu-item',
    props: {item: {}},
    template: [
        '<li class="layui-nav-item">',
        '<a v-if="item.type === 1" href="javascript:;">',
        '<i v-if="item.icon != null" :class="item.icon"></i>',
        '<span>{{item.name}}</span>',
        '<em class="layui-nav-more"></em>',
        '</a>',
        '<dl v-if="item.type === 1" class="layui-nav-child">',
        '<dd v-for="item in item.childs">',
        '<a v-if="item.type === 2" href="javascript:;" :data-url="dudu.ctx + item.url"><i v-if="item.icon != null" :class="item.icon" :data-icon="item.icon"></i> <span>{{item.name}}</span></a>',
        '</dd>',
        '</dl>',
        '<a v-if="item.type === 2" href="javascript:;" :data-url="dudu.ctx + item.url"><i v-if="item.icon != null" :class="item.icon" :data-icon="item.icon"></i> <span>{{item.name}}</span></a>',
        '</li>'
    ].join('')
});

//注册菜单组件
Vue.component('menuItem', menuItem);

var vm = new Vue({
    el: '#layui_layout',
    data: {
        user: {},
        menuList: {},
        password: '',
        newPassword: '',
        navTitle: "控制台"
    },
    methods: {
        getMenuList: function (event) {
            dudu.get(dudu.ctx + "/system/menu/user?_" + $.now(), function (r) {
                vm.menuList = r.data;
                window.permissions = r.perms;
            });
        },
        getUser: function () {
            dudu.get(dudu.ctx + "/system/user/info?_" + $.now(), function (r) {
                vm.user = r.data;
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
                    if(vm.password=='') {
                        layer.alert('原密码不能为空');
                        return;
                    }
                    var newPassword = vm.newPassword || '';
                    if(newPassword=='') {
                        layer.alert('新密码不能为空');
                        return;
                    }

                    if(newPassword.length < 6) {
                        layer.alert('新密码长度必须大于6');
                        return;
                    }
                    var data = "password=" + hex_md5(vm.password) + "&newPassword=" + hex_md5(newPassword);
                    $.ajax({
                        type: "POST",
                        url: dudu.ctx + "/system/user/password",
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
        ,updateProject: function(){
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
                area: ['350px', '170px'],
                shadeClose: false,
                content: jQuery("#projectLayer"),
                btn: ['修改', '取消'],
                btn1: function (index) {
                    var projectId = $("#projectLayer [name='projectId']").val();

                    if(projectId=='' || projectId==-1) {
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
    },
    created: function () {
        this.getMenuList();
        this.getUser();
    }
});
