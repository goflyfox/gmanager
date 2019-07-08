function EmptyFunc(){
	return;
}
//普通alert
function Alert(message,handler){
	handler = handler || null;
	BootstrapDialog.alert(message,function(){
		if(handler != null) f(handler);
	});
}

//成功信息SucceedInfo
function SucceedInfo(message,handler){
	handler = handler || null;
	BootstrapDialog.show({
         type: BootstrapDialog.TYPE_SUCCESS,
         title: '信息',
         message: message,
         buttons: [{
             label: '确认',
             action: function(dialog) {
            	 if(handler != null) f(handler);
            	 dialog.close();
             }
         }]
     });
}

//失败信息ErrorInfo
function ErrorInfo(message,handler){
	handler = handler || null;
	BootstrapDialog.show({
        type: BootstrapDialog.TYPE_DANGER,
        title: '信息',
        message: message,
        buttons: [{
            label: '确认',
            action: function(dialog) {
           	 if(handler != null) f(handler);
           	 dialog.close();
            }
        }]
    });
}

//询问信息 
function Confirm (message,ok_fun,cancel_fun){ 
	BootstrapDialog.confirm(message, function(result){
        if(result) {
        	f(ok_fun);
        }else {
        	if(cancel_fun!=null){
				f(cancel_fun);
			} 
        }
    });
}

function f(fn) {
   fn(); 
}
