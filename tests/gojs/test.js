function htmlDecode(str){  
    var s = "";
    if(str.length == 0) return "";
    s = str.replace(/&amp;/g,"&");
    s = s.replace(/&lt;/g,"<");
    s = s.replace(/&gt;/g,">");
    s = s.replace(/&nbsp;/g," ");
    s = s.replace(/&#39;/g,"\'");
    s = s.replace(/&quot;/g,"\"");
    return s;  
}

function process(r) {
    console.log(r.Type);
    r.Type = "liuzl";

    var s = '<td class="a-content"><p>发信人: T2457207290 (东方欲晓||万象四猛之老四||无人敌), 信区: Universal <br /> 标&nbsp;&nbsp;题: Re: 中国社会九大阶层最新划分 <br /> 发信站: 水木社区 (Thu May 18 10:41:34 2017), 站内 <br />&nbsp;&nbsp;<br /> 我不会让你上台的 <br /> 【 在 summoner 的大作中提到: 】 <br /> <font class="f006">: 我10，你就是我的目标 </font> <br /> -- <br />&nbsp;&nbsp;<br /> <font class="f000"></font><font class="f006">※ 来源:·水木社区 <a target="_blank" href="http://m.newsmth.net">http://m.newsmth.net</a>·[FROM: 43.250.200.*]</font><font class="f000"> <br /> </font></p></td>';

    var pattern = /<td class="a-content"><p>发信人: (.+?) \((.*?)\), 信区: (.+?) <br \/> 标&nbsp;&nbsp;题: (.+?) <br \/> 发信站: (.+?) \((.+?)\), 站内 (.*?)<font class="f006">※ 来源:·(.+?) <a target="_blank" href="(.+?)">(?:.+?)<\/a>·\[FROM: (.+?)\]/g;
    var names = new Array("user", "nick", "board", "title", "mailsite", "time", "content", "source", "site", "ip");

    var r = "";
    var obj = new Object();
    if (r = pattern.exec(s)) {
        for (var i = 1; i < r.length; i++) {
            var ret = htmlDecode(r[i].replace(/<br \/>/g, "\n").replace(/<(?:.|\s)*?>/g, ""));
            obj[names[i-1]] = ret.trim();
        }
    }
    console.log(JSON.stringify(obj));

    return true;
}

