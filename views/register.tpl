<!DOCTYPE html>

<html>
<head>
    <title>前端论坛</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <link rel="stylesheet" href="/static/css/bootstrap.min.css">
    <style>
        section.registerSection{
            text-align: center;
            margin: 30px;
        }
        .registerSection input{
            max-width: 400px;
            height: 40px;
            margin: 0 auto;
        }
        .margin-top-15{
            margin-top: 15px;
        }
        section.registerSection h3{
            color:#888;
            margin-bottom: 20px;
        }
        #registerButton{
            margin-top: 15px;
            margin-bottom: 5px;
            width: 150px;
        }
        #loginButton{
            margin-top: 15px;
        }
        .registerSection a{
            margin-top: 15px;
        }
    </style>
</head>
<body>
    <header><div style="height: 180px"></div></header>
    <section class="registerSection">
        <h3>前端开发者-Register</h3>
        <input id="username" type="text" class="form-control" value="" placeholder="username">
        <div class="margin-top-15"></div>
        <input id="password" type="password" class="form-control" value="" placeholder="password">
        <div class="margin-top-15"></div>
        <input id="mobile" type="text" class="form-control" value="" placeholder="mobile">
        <button id="registerButton" class="btn btn-info" onclick="register()"> 注册 </button><br>
        <a href="/login">I want to login.</a>
    </section>
</body>

<script src="/static/js/jquery.min.js"></script>
<script src="/static/js/bootstrap.min.js"></script>
<script src="/static/js/yifan.toast.js"></script>
<script>
    function register() {
        var username = $("#username").val();
        var password = $("#password").val();
        var mobile = $("#mobile").val();
        if(username==""||password==""||mobile==""){
            alert("请填写完整信息！");
        }else{
            $.ajax({
                url: '/user/register',
                type: 'post',
                dataType: 'json',
                data:{
                    'username':username,
                    'password':password,
                    'mobile':mobile
                },
                contentType: 'application/x-www-form-urlencoded;charset=UTF-8',
                cache: false,
                success: function(data) {
                    if(data.status==10000){
                        alert("Register Success,Please Login!");
                        window.location.href = "/login";
                    }else {
                        alert("Register Error");
                        window.location.reload();
                    }
                },
                error : function() {
                    alert("Connect SERVER Error!");
                    window.location.reload();
                }
            });
        }


    }


</script>
</html>
