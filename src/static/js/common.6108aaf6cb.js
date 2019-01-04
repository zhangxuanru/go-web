(function (window) {
  //创建空console对象，避免JS报错
  if (!window.console) {
    window.console = {};
  }
  var console = window.console;
  var funcs = ['assert', 'clear', 'count', 'debug', 'dir', 'dirxml',
    'error', 'exception', 'group', 'groupCollapsed', 'groupEnd',
    'info', 'log', 'markTimeline', 'profile', 'profileEnd',
    'table', 'time', 'timeEnd', 'timeStamp', 'trace', 'warn'];
  for (var i = 0, l = funcs.length; i < l; i++) {
    var func = funcs[i];
    if (!console[func]) {
      console[func] = function () {
      };
    }
  }
  if (!console.memory) {
    console.memory = {};
  }
})(window);


function addScript(src) {
  var test = "";
  var s = document.createElement('script');
  s.type = 'text/javascript';
  s.async = true;
  s.src = src;
  document.body.appendChild(s);
}

//addScript('http://c.csdnimg.cn/rabbit/search-service/main.js');

function formatUrl(url) {
  let reg = /(?:[?&]+)([^&]+)=([^&]+)/g;
  let data = {};

  function fn(str, pro, value) {
    data[decodeURIComponent(pro)] = decodeURIComponent(value);
  }

  url.replace(reg, fn);
  return data;
}




(function ($, window, document) {
  $(function () {

    //添加渠道来源
    function addRegSouce() {
      let url = formatUrl(location.href);
      if (url.utm_source) {
        setCookie("source", url.utm_source)
      }
    }

    addRegSouce();

    function setCookie(name, value) {
      const days = 30;
      let expires = new Date();
      expires.setTime(expires.getTime() + days * 24 * 60 * 60 * 1000);
      document.cookie = name + "=" + escape(value) + ";expires=" + expires;
    }

    function getCookie(name) {
      var arr, reg = new RegExp("(^| )" + name + "=([^;]*)(;|$)");
      if (arr = document.cookie.match(reg)) {
        return unescape(arr[2]);
      } else {
        return null;
      }
    }

    var userAgent = window.navigator.userAgent.toUpperCase();
    var browser = '';
    var version = 0;
    var support = false;
    if (/MSIE/.test(userAgent)) {
      browser = 'MSIE';
      // IE11 // Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; rv:11.0) like Gecko
      // IE10 // Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 10.0; WOW64; Trident/8.0; .NET4.0C; .NET4.0E)
      // IE9  // Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 10.0; WOW64; Trident/8.0; .NET4.0C; .NET4.0E)
      // IE8  // Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 10.0; WOW64; Trident/8.0; .NET4.0C; .NET4.0E)
      if (/rv:11/.test(userAgent)) {
        version = 11;
      } else {
        version = parseInt(userAgent.match(/MSIE.*?(?=\;)/)[0].substr(5));
      }
      support = (version >= 10) ? true : false;
    } else if (/CHROME/.test(userAgent) && /SAFARI/.test(userAgent)) {
      browser = 'Chrome';
      // Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.71 Safari/537.36
      version = parseInt(userAgent.match(/CHROME\/.*?(?=\.)/)[0].substr(7));
      support = (version >= 42) ? true : false;
    } else if (/SAFARI/.test(userAgent)) {
      browser = 'Safari';
      // Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12) AppleWebKit/602.1.50 (KHTML, like Gecko) Version/10.0 Safari/602.1.50
      version = parseInt(userAgent.match(/VERSION\/.*?(?=\.)/)[0].substr(8));
      support = (version >= 9) ? true : false;
    } else if (/FIREFOX/.test(userAgent)) {
      browser = 'Firefox';
      // Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:43.0) Gecko/20100101 Firefox/43.0
      version = parseInt(userAgent.match(/FIREFOX\/.*?(?=\.)/)[0].substr(8));
      support = (version >= 42) ? true : false;

    } else {
      browser = window.navigator.userAgent;
      version = 0;
      support = true;
    }


    // console.log('\n browser_unsupport  [',  getCookie('browser_unsupport'),  '] ');
    // console.log('\n browser:', browser , '  version:', version, '  support:',support,  'cookie', document.cookie );

    var topOffset = 120;//($('body').height() - 300)/2;
    var leftOffset = ($('body').width() - 900) / 2;

    var supportDialogStyle = 'top:' + topOffset + 'px; left:' + leftOffset + 'px;';

    var html = '<!-- BrowserSupport DialogHTMLString -->' +
      '<div class="supportModel"  style="background-color: #000;width: 100%;height: 100%; position: fixed;top: 0;left: 0;opacity: 0.4;"> </div>' +
      '<div class="supportDialog" style="position:fixed; padding:0; height:240px; width:900px; z-index:100; ' + supportDialogStyle + ' ">' +
      '       <div class="" >' +
      '           <div class="">' +
      '               <div class="modal-content">' +
      '                   <div class="modal-header">' +
      '                       <button type="button" class="close" id="supportModelClose">×</button>' +
      '                       <div class="modal-title fs-12" id="myModalLabel">提示</div>' +
      '                   </div>' +
      '                   <div class="modal-body pt-15 pb-30 pl-30 pr-30">' +
      '                       <div class="fs-14">我们注意到你正在使用一个与我们的网站不兼容的浏览器</div>  ' +
      '                       <p calss="fs-14 pt-10">你可以继续我们的网站，但它的部分可能无法正常工作。对于最好的体验，我们建议您升级或尝试这些流行的浏览器之一</p>    ' +
      '                       <div class="clearfix mt-15">' +
      '                           <div class="col-sm-3 col-md-3 col-lg-3">' +
      '                               <img src="//goss.vcg.com/html/images/logo/ie.png" />' +
      '                               <a href="https://support.microsoft.com/zh-cn/help/17621" class="fs-14 mt-15" target="_blank">立即下载</a>' +
      '                           </div>' +
      '                           <div class="col-sm-3 col-md-3 col-lg-3">' +
      '                               <img src="//goss.vcg.com/html/images/logo/chrome.png" />' +
      '                               <a href="http://www.google.cn/chrome/browser/desktop/index.html" class="fs-14 mt-15" target="_blank">立即下载</a>' +
      '                           </div>' +
      '                           <div class="col-sm-3 col-md-3 col-lg-3">' +
      '                               <img src="//goss.vcg.com/html/images/logo/safari.png" />' +
      '                               <a href="https://support.apple.com/zh-cn/HT204416" class="fs-14 mt-15" target="_blank">立即下载</a>' +
      '                           </div>' +
      '                           <div class="col-sm-3 col-md-3 col-lg-3">' +
      '                               <img src="//goss.vcg.com/html/images/logo/firefox.png" />' +
      '                               <a href="http://www.firefox.com.cn/" class="fs-14 mt-15" target="_blank">立即下载</a>' +
      '                           </div>' +
      '                       </div>' +
      '                   </div>' +
      '               </div>' +
      '           </div>' +
      '       </div>' +
      '</div>';


    if (!support) {
      if (getCookie('browser_unsupport') == undefined) {
        $('body').append(html);
      }
    }

    $('#supportModelClose').click(function () {
      $('.supportModel').remove();
      $('.supportDialog').remove();
      setCookie('browser_unsupport', true);
    });

  });
})(jQuery, window, document);


; (function ($, h, c) {
  var a = $([]), e = $.resize = $.extend($.resize, {}), i, k = "setTimeout", j = "resize", d = j + "-special-event",
    b = "delay", f = "throttleWindow";
  e[b] = 250;
  e[f] = true;
  $.event.special[j] = {
    setup: function () {
      if (!e[f] && this[k]) {
        return false
      }
      var l = $(this);
      a = a.add(l);
      $.data(this, d, {
        w: l.width(),
        h: l.height()
      });
      if (a.length === 1) {
        g()
      }
    },
    teardown: function () {
      if (!e[f] && this[k]) {
        return false
      }
      var l = $(this);
      a = a.not(l);
      l.removeData(d);
      if (!a.length) {
        clearTimeout(i)
      }
    },
    add: function (l) {
      if (!e[f] && this[k]) {
        return false
      }
      var n;

      function m(s, o, p) {
        var q = $(this)
          , r = $.data(this, d);
        r.w = o !== c ? o : q.width();
        r.h = p !== c ? p : q.height();
        n.apply(this, arguments)
      }

      if ($.isFunction(l)) {
        n = l;
        return m
      } else {
        n = l.handler;
        l.handler = m
      }
    }
  };

  function g() {
    i = h[k](function () {
      a.each(function () {
        var n = $(this)
          , m = n.width()
          , l = n.height()
          , o = $.data(this, d);
        if (m !== o.w || l !== o.h) {
          n.trigger(j, [o.w = m, o.h = l])
        }
      });
      g()
    }, e[b])
  }
})(jQuery, window);
$(function () {
  var isSameWidth = false;
  if ($("#searchfilterdiv")) {
    isSameWidth = $("#searchboxdiv").hasClass("same-width");
  }

  window.onscroll = function () {
    var t = document.documentElement.scrollTop || document.body.scrollTop;

    if ($("#searchfilterdiv")) {
      //var top = $("#searchfilterdiv").offset().top;
      let topicTop = 140;

      if ($("#img-banner").attr('src')) {
        topicTop += 200;
      }
      if (t >= topicTop) {
        $("#bs-example-navbar-collapse-fixed").show()
      } else {
        $("#bs-example-navbar-collapse-fixed").hide()
      }
      if (t >= 87) {

        $("#searchboxdiv").addClass("navbar-fixed-top");
        $("#searchboxdiv").addClass("s-search-ceiling");
        // 搜索框
        // $("#search-form").addClass("search-ceiling");

        // TODO 搜索后搜索框下部显示搜索相似等结果集的DOM操作
        $("#SearchFilterComBox").addClass("navbar-fixed-top");
        // $("#gallery-preview-strip").addClass("visibility");

        $("#search-content").addClass("container");
        $("#search-content").width("");
        $('.as-menu.list-group').css({ top: "51px" });
        $('.as-menu.list-group').width($(".searchFrame").width());
        if (isSameWidth) {
          $("#searchboxdiv").removeClass("same-width");
        }
      } else {
        $("#SearchFilterComBox").removeClass("navbar-fixed-top");
        $("#gallery-preview-strip").removeClass("gallery-preview-strip-fixed");
        $("#gallery-preview-strip").removeClass("visibility");

        $("#searchboxdiv").removeClass("navbar-fixed-top");
        $("#searchboxdiv").removeClass("s-search-ceiling");
        // 搜索框
        // $("#search-form").removeClass("search-ceiling");

        $("#search-content").removeClass("container");
        $("#search-content").width("100%");
        $('.as-menu.list-group').css({ top: "60px" });
        $('.as-menu.list-group').width($(".searchFrame").width());
        if (isSameWidth) {
          $("#searchboxdiv").addClass("same-width");
        }
      }
    }

    var toTopDiv = $("#ih-totop");
    if (toTopDiv) {
      if (t >= 10) {
        toTopDiv.show();
      } else {
        toTopDiv.hide();
      }
    }

    var scrollFunc = function (e) {
      e = e || window.event;
      let $gallerypreviewstrip = $('#gallery-preview-strip');

      if (e.wheelDelta) {  //判断浏览器IE，谷歌滑轮事件
        if (e.wheelDelta > 0 && t >= 87) { //当滑轮向上滚动时
          $gallerypreviewstrip.addClass('gallery-preview-strip-fixed');
          $gallerypreviewstrip.addClass('visibility');
        } else {
          $gallerypreviewstrip.removeClass('gallery-preview-strip-fixed');
          $gallerypreviewstrip.removeClass('visibility');
        }
      } else if (e.detail) {  //Firefox滑轮事件
        if (e.detail < 0 && t >= 87) { //当滑轮向上滚动时
          $gallerypreviewstrip.addClass('visibility');
          $gallerypreviewstrip.addClass('gallery-preview-strip-fixed');
        } else {
          $gallerypreviewstrip.removeClass('visibility');
          $gallerypreviewstrip.removeClass('gallery-preview-strip-fixed');
        }
      }
    };
    //给页面绑定滑轮滚动事件
    if (document.addEventListener) {//firefox
      document.addEventListener('DOMMouseScroll', scrollFunc, false);
    }
    //滚动滑轮触发scrollFunc方法  //ie 谷歌
    window.onmousewheel = document.onmousewheel = scrollFunc;

  };

  // window.onpopstate = function () {
  //   let url = location.href;
  //   location.href = url;
  //   return false;
  // }

})
