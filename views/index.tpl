<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>首页|前端开发者</title>
</head>
<body>
<div style="width: 500px;margin: 0 auto">
    {{range .users}}
    用户名：{{.Name}}<br>
    密 码：{{.Password}}<br>
    手机号码:{{.Mobile}}<br>
    <hr style="border: 1px double #444">
    {{end}}<br>
    {{ .userinfo.Name}}
    <a href="/logout">注销</a>
</div>
<script src="/static/js/yifan.toast.js"></script>
</body>
</html>