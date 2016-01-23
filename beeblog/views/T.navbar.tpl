{{define "navbar"}}
    <div class="navbar navbar-default navbar-fixed-top">    
      <div class="container">
        <a class="navbar-brand" href="/">我的博客</a>
        <div>
          <ul class="nav navbar-nav">
            <li class="{{if .IsHome}}active{{end}}"><a href="/">首页</a></li>
            <li class="{{if .IsCategory}}active{{end}}"><a href="/category">分类</a></li>
            <li class="{{if .IsTopic}}active{{end}}"><a href="/topic">文章</a></li>
          </ul>
        </div>
        <div class="pull-right">
          <ul class="nav navbar-nav">
          {{if .IsLogin}}
          <li><a href="/login?exit=true">退出</a></li>
          {{else}}
          <li><a href="/login">管理员登录</a></li>
          {{end}}
          </ul>
        </div>
      </div>
    </div>
{{end}}