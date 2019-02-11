$.flag=function(){
	/*
	 * $.flag()用于判断是当前页面父级是否有拥setUrlAndTitle方法;
	 * 
	 * return  true|false
	 * */
	 return  window.parent.setUrlAndTitle?true:false;
}

$.reg = {
        names : /^\S{2,7}$/ ,//名字
        email : /^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/ ,//验证邮箱
        pass : /^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,16}$/ ,//密码必须由6-16位数字加字母组成
        date : /^\d{4}-\d{1,2}-\d{1,2}/ ,//日期
        QQ : /^[1-9][0-9]{4,11}/ ,//qq号
        num : /^(0|[1-9][0-9]*)$/ ,//整数
        phone : /^1[3|4|5|8|7][0-9]\d{8}$/ , //验证手机号
        empty : /(^ +| +$)/g , //去除首尾空格
        name : /^[\u4E00-\u9FA5]{2,4}$/ , /* 后台真实姓名 */
}
$.regs = 'match、exec、test、search、replace、split';
/* cookie*/

$.cookie = {
        /*  设置cookie */
        set : function ( name , value , expires , path , domain ) {
                if ( typeof expires === "undefined" ) {
                        expires = new Date ( new Date ().getTime () + (1000 * 60 * 60 * 24) );
                } else {
                        expires = new Date ( new Date ().getTime () + expires );
                }
                document.cookie = name + "=" + escape ( value ) + ((expires) ? ";expires=" + expires.toGMTString () : "") + ((path) ? ";path=" + path : ";path=/") + ((domain) ? ";domain=" + domain : "");
        } ,
        /*  获取cookie */
        get : function ( name ) {
                var arr = document.cookie.match ( new RegExp ( "(^| )" + name + "=([^;]*)(;|$)" ) );
                if ( arr != null ) {
                        return unescape ( arr[ 2 ] );
                }
                return null;
        } ,
        /*  清除cookie */
        clear : function ( name , path , domain ) {
                if ( this.get ( name ) ) {
                        document.cookie = name + "=" + ((path) ? "; path=" + path : "; path=/") + ((domain) ? "; domain=" + domain : "") + ";expires=Fri, 02-Jan-1970 00:00:00 GMT";
                }
        } ,
        /* 返回时间 */
        date : function ( num ) {
                //设置cookie失效时间 默认30天
                num = num ? num : 30;
                //var time=new Date ( new Date ().getTime () + (1000 * 60 * 60 * 24 * num) )
                var time = 1000 * 60 * 60 * 24 * num;
                return time;
        }
};

/* 禁用浏览器回退按钮 */
$.disableBack = function () {
       /* $(function() {
                if (window.history && window.history.pushState) {
                        $(window).on('popstate', function () {
                                window.history.pushState('forward', null, '#');
                                window.history.forward(1);
                        });
                }
                window.history.pushState('forward', null, '#'); //在IE中必须得有这两行
                window.history.forward(1);
        })*/
        $.cookie.set('data-url',window.location.href)//设置点击返回的cookie(上级页面的cookie)
        $.cookie.set('data-title',$('title').html())//设置点击返回的cookie(上级页面的cookie)
       return false

        
}
/* 获取 url 参数 */
$.urlValue = function ( key ) {
    var ret = {};
    var url = window.location.href;
    var pat = /([^?=&|#]+)=([^?=&|#]+)/g;
    url.replace ( pat , function () {
        ret[ arguments[ 1 ] ] = arguments[ 2 ]
    } );
    return ret[ key ] ? ret[ key ] : ret;
}


/*页面跳转*/
$.clickJump = function (obj){
        var dataUrl=$(obj).attr('data-url');
        var dataTitle = $(obj).attr('data-title');
        var html=$(obj).attr('data-html');
        if (html==="返回"){
                window.history.go(-1);
                return;
        }
        if (!dataUrl || !dataTitle){
                return;
        }else {
                $.cookie.set('url',window.location.href)//设置点击返回的cookie(上级页面的cookie)
                $.cookie.set('title',window.location.href)//设置点击返回的cookie(上级页面的cookie)

                $.cookie.set('data-url',dataUrl)//设置下次刷新cookie
                $.cookie.set('data-title',dataTitle)//设置下次刷新cookie
        }
        
        //window.parent.document  获取iframe父页面
        if($.flag()){
       		window.parent.setUrlAndTitle(dataUrl,dataTitle)	
        }else{
        	window.location.href=dataUrl;
        }
}

//分页记录URL
$.pageSetUrl=function(url,title){
        $.cookie.clear('data-url')//清除cookie
        $.cookie.clear('data-title')//清除cookie
        $.cookie.clear('url')//清除cookie
        $.cookie.clear('title')//清除cookie
        
        $.cookie.set('url',url)//设置点击返回的cookie
        $.cookie.set('title',title)//设置点击返回的cookie

        $.cookie.set('data-url',url)//设置下次刷新cookie
        $.cookie.set('data-title',title)//设置下次刷新cookie
        if($.flag()){
       		window.parent.setUrlAndTitle(url,title); 	
        }else{
        	window.location.href=url;
        }
       
}

//ajsx记录URL
$.ajaxSetUrl=function(url,title){
        $.cookie.clear('data-url')//清除cookie
        $.cookie.clear('data-title')//清除cookie
        $.cookie.clear('url')//清除cookie
        $.cookie.clear('title')//清除cookie
        
        $.cookie.set('data-url',url)//设置下次刷新cookie
        $.cookie.set('data-title',title)//设置下次刷新cookie
        
        $.cookie.set('url',url)//设置点击返回的cookie
        $.cookie.set('title',title)//设置点击返回的cookies
        
}
//表单审核返回历史记录
$.formUrl=function(){
        var url=$.cookie.get('url')
        var title=$.cookie.get('data-title')
        
         if($.flag()){
       		 window.parent.setUrlAndTitle(url,title)
        }else{
        	window.location.href=url;
        }
}

//自动记录
$.automaticUrl=function(){
        $.cookie.set('data-url',window.location.href)//设置点击返回的cookie(上级页面的cookie)
        $.cookie.set('data-title',$('title').html())//设置点击返回的cookie(上级页面的cookie)
}


$(function(){
	
	//获取当前页面的url
	 var btnEleList=$('#nvaTopTitle button',parent.document);
	 
	if(btnEleList){
		 	var url=window.location.origin+window.location.pathname;
			 $.each(btnEleList, function(index,ele){
			 		//console.log(ele, index)
			 		var eleUrl=$(ele).attr('data-url');
			 		if(eleUrl==url){
			 			$(ele).parent().css({display:'block'})
			 			$(ele).removeClass('layui-btn-primary').addClass('layui-btn-normal');
			 			
			 			//所有的兄弟元素变为不选中状态
			 			$(ele).siblings().removeClass("layui-btn-normal");
			 			$(ele).siblings().addClass("layui-btn-primary");
			 		}
			 });
	}
	
});


$(function(){
	/*设置 父级的 iframe padding*/
    function setIframePadding(){
       	 var $body=$(('body'));//获取子页面的body
		 var indexPadding=$($body).attr('data-body')//获取自定义属性
		 if(indexPadding==undefined){
        	 $($body).attr('data-body','body')
      		 $($body).css({paddingBottom:"80px"})//获取iframe中body元素，其他的话自己用$c('#aaa')去获取吧
		}	
    }
    setIframePadding();//执行
})
