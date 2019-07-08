var dudu = {
    code: {
        ok: 0,
        fail: -1,
        error: -99,
        authError: -401
    },
    // select: function (_url, $select, _name, _value, _selected,_exclude) {
    select: function (_url, paramExtend, _async) {
        var paramSrc = {
            select: $("#select"), // select对象
            name: "name", // 展示名称
            value: "value", // value值
            selected: 0, // 选中元素
            exclude: 0 // 不展示元素
        }

        var param = $.extend(paramSrc, paramExtend)
        dudu.get(_url, function (result) {
            if (result.code == 0) {
                var arr = result.data || null;
                if (arr != null && arr.length > 0) {
                    var optionStr = "";
                    for (var p in arr) {
                        if (param.exclude == arr[p][param.value]) {
                            continue;
                        }
                        optionStr += '<option value="';
                        optionStr += arr[p][param.value];
                        optionStr += '" ';
                        optionStr += (arr[p][param.value] == param.selected) ? 'selected' : '';
                        optionStr += '>';
                        optionStr += arr[p][param.name];
                        optionStr += '</option>';
                    }
                    param.select.append(optionStr);
                }
            } else {
                ErrorInfo('select操作失败：' + result.msg);
            }
        }, _async);
    },
    setValue: function (_url, $element, _value) {
        dudu.get(_url, function (result) {
            if (result.code == 0) {
                var value = result.data[_value];
                if ($element.is('input')) {
                    $element.val($element.val() + value);
                } else if ($element.is('span') || $element.is('dict')) {
                    $element.html($element.html() + value);
                } else {
                    $element.text($element.text() + value);
                }
            } else {
                ErrorInfo('select操作失败：' + result.msg);
            }
        });
    },
    // 数据字典列表
    dictSelect: function ($select, _dictType, _selected, _async) {
        var url = "system/dictdetail/data?dictType=" + _dictType;
        var name = "name";
        var value = "id";
        dudu.select(url, {
            $select: $select,
            name: name,
            value: value,
            selected: _selected
        }, _async);
    },
    // 数据字典Code列表
    dictCodeSelect: function ($select, _dictType, _selected, _async) {
        var url = "system/dictdetail/data?dictType=" + _dictType;
        var name = "name";
        var value = "code";
        dudu.select(url, {
            $select: $select,
            name: name,
            value: value,
            selected: _selected
        }, _async);
    },
    dictValue: function ($element, _id) {
        var _url = "system/dictdetail/data/" + _id;
        dudu.setValue(_url, $element, "name");
    },
    dictCode: function ($element, _dictType, _code) {
        var _url = "system/dictdetail/data/0?dictType=" + _dictType + "&code=" + _code;
        dudu.setValue(_url, $element, "name");
    },
    get: function (_url, succ_callback, _async) {
        _async = _async === false ? false : true;
        // _async = false;
        // alert(_async)
        jQuery.ajax({
            type: 'GET',
            beforeSend: dudu.headToken,
            url: _url,
            data: null,
            async: _async,
            success: function (r) {
                if (r.code == dudu.code.authError) {
                    alert(r.msg);
                    window.top.location.href = "login";
                } else {
                    succ_callback(r)
                }
            },
            error: function (html) {
                var flag = (typeof console != 'undefined');
                if (flag) console.log("服务器忙，提交数据失败，代码:" + html.status + "，请联系管理员！");
                alert("服务器忙，提交数据失败，代码:" + html.status + "，请联系管理员！");
            }
        });
    },
    post: function (_url, _param, succ_callback, _async) {
        _async = _async === false ? false : true;
        jQuery.ajax({
            type: 'POST',
            beforeSend: dudu.headToken,
            url: _url,
            data: _param,
            success: function (r) {
                if (r.code == dudu.code.authError) {
                    alert(r.msg);
                    window.top.location.href = "login";
                } else {
                    succ_callback(r)
                }
            },
            async: _async,
            error: function (html) {
                var flag = (typeof console != 'undefined');
                if (flag) console.log("服务器忙，提交数据失败，代码:" + html.status + "，请联系管理员！");
                alert("服务器忙，提交数据失败，代码:" + html.status + "，请联系管理员！");
            }
        });
    },
    /**
     * 重置表单
     */
    resetForm: function () {
        $(".tableSearch input:not(.btn1)").val("");
        $(".tableSearch select").val("-1");
    },
    /**
     * 避免传递空""值过来value时判断为是取值操作而不是赋值操作
     */
    jqName: function (name, value) {
        if (value === undefined) {
            return jQuery(":input[name='" + name + "']").val() || '';
        } else {
            return jQuery(":input[name='" + name + "']").val(value);
        }
    },
    /*********************************************/
    /* Trim the both side blank of the string    */
    /*********************************************/
    trim: function (s) {
        return _trimRight(_trimLeft(s));
    },
    /****************************************/
    /* Trim the left blank of the string    */
    /****************************************/
    _trimLeft: function (s) {
        while (s.charAt(0) == " " || s.charAt(0) == "") {
            s = s.substr(1, s.length - 1);
        }
        return s;
    },
    /*****************************************/
    /* Trim the right blank of the string    */
    /*****************************************/
    _trimRight: function (s) {
        while (s.charAt(s.length - 1) == " " || s.charAt(s.length - 1) == "") {
            s = s.substr(0, s.length - 1);
        }
        return s;
    },
    log: function (str) {
        var flag = (typeof console != 'undefined');
        if (flag) console.log('dudu--> ' + str);
    },
    headToken: function (xhr) {
        if (window.localStorage.getItem("token") != null) {
            xhr.setRequestHeader("Authorization", "Bearer " + window.localStorage.getItem("token"));
        } else {
            alert("登录超时，请重新登录")
            window.top.location.href = "login";
        }
    },
    loadAuth: function (r) {
        if (r.code == dudu.code.authError) {
            Alert(r.msg, function () {
                window.top.location.href = "login";
            })
        }
    },
};
