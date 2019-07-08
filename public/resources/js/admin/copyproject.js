var vm = new Vue({
    el:'#data_wrapper',
    data:{
        showCopy:true,
        title: null,
        model:{
            srcProjectId:null,
            srcProjectName:null,
            destProjectId:null,
            destProjectName:null,
            srcProjectConfig:null,
            destProjectConfig:null,
            projectList:null,
        }
},
    created:function(){
        //为原项目名赋值
        dudu.get(dudu.ctx+"/admin/project/srcProject",function(result){
            if(result.code == 0){
                vm.model.srcProjectId = result.data.id;
                vm.model.srcProjectName=result.data.name;
            }
        });
        //为项目下拉列表赋值
        dudu.get(dudu.ctx + "/admin/project/data",function (result) {
            if(result.code == 0){
                vm.model.projectList = result.data;
            }
        });
        dudu.get(dudu.ctx+"/admin/project/srcDestProjectInfo",function (result) {
            if(result.code == 0){
                vm.model.srcProjectConfig = result.data.srcProjectConfig;
            }
        })
    },
methods: {
    //项目拷贝
    copy:function () {
        $("#copyButton").attr("disabled",true);
        vm.showCopy=true;
        vm.title = "拷贝项目";

        var deskProjectId = $("#deskProject").val();
        if(deskProjectId == ""){
            Alert("请选择要对比复制的项目",function () {
                $("#copyButton").attr("disabled",false);
            });
            return;
        }else{
            Confirm("您确定要复制此项目吗?",function(){
                dudu.get(dudu.ctx+"/admin/project/copy?destProjectId="+deskProjectId+"&type="+vm.model.type,function (result) {
                    if(result.code == 0){
                        Alert("拷贝成功",function(){
                            //重新获取原项目和新项目的配置文件
                            vm.selectProject();
                            $("#copyButton").attr("disabled",false);
                        });
                    }
                });
            },function(){
                $("#copyButton").attr("disabled",false);
            });
        }
    },
    selectProject:function () {
        vm.model.srcProjectConfig = null;
        vm.model.destProjectConfig = null;
        var val = $("#deskProject").val();
        if(val != ""){
            dudu.get(dudu.ctx+"/admin/project/srcDestProjectInfo?destProjectId="+val,function (result) {
                if(result.code == 0){
                    vm.model.srcProjectConfig = result.data.srcProjectConfig;
                    vm.model.destProjectConfig = result.data.destProjectConfig;
                }
            })
            vm.showType();
        }else{
            //弹框

        }
    },
    showPublicDiff: function () {
        var deskProjectId = $("#deskProject").val();
        if(deskProjectId == ""){
            Alert("请选择要对比复制的项目");
            return;
        }
        diffview.diffUsingJS(0, vm.model.srcProjectConfig, vm.model.destProjectConfig);
    }
}
});