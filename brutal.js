var brutal={
    selectValId:function(id){
        if(document.getElementById(id)){
            var select_list = document.getElementById(id);
            if(select_list.options.length>0){
                return select_list[select_list.selectedIndex].value;
            }
            return null;
        }
        return null;
    },
    selectValHtmlId:function(id){
        if(document.getElementById(id)){
            var select_list = document.getElementById(id);
            if(select_list.options.length>0){
                var ret = {
                    val:select_list[select_list.selectedIndex].value,
                    html:select_list[select_list.selectedIndex].innerHTML
                }
                return ret;
            }
            return null;
        }
        return null;
    },
    selectHtmlId:function(val,id){
        if(document.getElementById(id)){
            var select_list = document.getElementById(id);
            for(var i=0;i<select_list.length;i++){
                if(val == select_list[i].value){
                    return select_list[i].innerHTML;
                }
            }
            return null;
        }
        return null;
    },
    createOptionsInt:function(n,common){
        var html ='';
        for(var i=1;i<n;i++){
            html += '<option value="'+i+'">'+common+'</option>';
        }
        return html;
    },
        
    setOptionId:function(id,val){
        var select_list = document.getElementById(id);
        for(var i=0;i<select_list.options.length;i++){
            if(select_list.options[i].value == val){
                select_list.options[i].selected = true;
                break;
            }
        }
    },
    
    checkedListName:function(name){
        if(document.getElementsByName(name)){
            var checkbox_list = document.getElementsByName(name);
            var check_arr = [];
            for(var i=0;i<checkbox_list.length;i++){
                if(checkbox_list[i].checked){
                    check_arr.push(checkbox_list[i].value);
                }
            }
            return check_arr;
        }
        else{
            return [];
        }
    },
    
    checkValId:function(id){
        if(document.getElementsByName(name)){
            return document.getElementById(id).value;
        }
        else{
            return null;
        }
    },
    
    setCheckedName:function(name){
        if( document.getElementsByName(name) ){
            var checkbox_list = document.getElementsByName(name);
            for(var i=0;i<checkbox_list.length;i++){
                checkbox_list[i].checked = true;
            }
        }
    },
    isChecked:function(id){
        if(document.getElementById(id)){
           return (document.getElementById(id).checked?true:false);
        }
        else{
            return false;
        }
    },
    setUnCheckedName:function(name){
        if( document.getElementsByName(name) ){
            var checkbox_list = document.getElementsByName(name);
            for(var i=0;i<checkbox_list.length;i++){
                checkbox_list[i].checked = false;
            }
        }
    },
    
    setCheckedId:function(id){
        if(document.getElementById(id)){
           document.getElementById(id).checked = true;
        }
    },
    setUnCheckedId:function(id){
        if(document.getElementById(id)){
           document.getElementById(id).checked = false;
        }
    },
    
    radioCheckedName:function(name){
        if( document.getElementsByName(name) ){
            var radio_list = document.getElementsByName(name);
            for(var i=0;i<radio_list.length;i++){
                if(radio_list[i].checked){
                    return radio_list[i].value;
                }
            }
        }
    },
    
    isRadioChecked:function(){
        
    },

    valId:function(id,val){
        if( document.getElementById(id) ){
            if( val === undefined){
                return document.getElementById(id).value;
            }
            else{
                document.getElementById(id).value = val;
            }
        }

    },
    
    htmlId:function(id,val){
        if( document.getElementById(id) ){
            if( val === undefined){
                return document.getElementById(id).innerHTML;
            }
            else{
                document.getElementById(id).innerHTML = val;
            }
        }
    },
    classNameId:function(id,val){
        if( document.getElementById(id) ){
            if( val === undefined){
                return document.getElementById(id).className;
            }
            else{
                document.getElementById(id).className = val;
            }
        }
    },

    timeOut:function(t,func,args){
        if( args === undefined){
            setTimeout(function(){func();},t);
        }
        else{
            setTimeout(function(){func(args);},t);
        }
    },
    setInfoClean:function(xrg){
        brutal.htmlId(xrg.id,xrg.info);
        setTimeout(function(){brutal.htmlId(xrg.id,'');},xrg.t);
    },
    jsontodata:function(jsd){
        var data = '';
        for(var k in jsd){
            if(typeof jsd[k] == 'object'){
                jsd[k] = jsd[k].toString();
            }
            data += k+'='+encodeURIComponent(jsd[k])+'&';
        }
        return data.substring(0,data.length);
    },

    eleAttrId:function(id,a,v){
        if(document.getElementById(id)){
            if(typeof a !== 'string'){
                return;
            }
            if(document.getElementById(id)[a]!==undefined){
                if(v===undefined){
                    return document.getElementById(id)[a];
                }
                else{
                    document.getElementById(id)[a]=v;
                }
            }
        }
    },

    confirmBlock:function(id,html,class_name){
        if(class_name!==undefined){
            brutal.classNameId(id,class_name);
        }
        brutal.htmlId(id,html);
    },

    confirmCancel:function(id,class_name){
        brutal.htmlId(id,'');
        if(class_name===undefined){
            brutal.classNameId(id,'');
        }
        else{
            brutal.classNameId(id,class_name);
        }
    }
    
}