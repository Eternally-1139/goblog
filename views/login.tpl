<!DOCTYPE html>

<html>
<head>
    <title>登录|前端开发者</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <link rel="stylesheet" href="/static/css/bootstrap.min.css">
    <style>
        section.registerSection{
            text-align: center;
            margin:30px
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
    <h3>前端开发者--Login</h3>
    <input id="username" type="text" class="form-control" value="" placeholder="username">
    <div class="margin-top-15"></div>
    <input id="password" type="password" class="form-control" value="" placeholder="password">
    <button id="registerButton" class="btn btn-info" onclick="register()"> 登录 </button><br>
    <a href="/register">I want to Register.</a>
</section>
</body>

<script src="/static/js/jquery.min.js"></script>
<script src="/static/js/bootstrap.min.js"></script>
<script>
    function register() {
        var username = $("#username").val();
        var password = $("#password").val();
        if(username==""||password==""){
            alert("请填写完整信息！");
        }else{
            $.ajax({
                url: '/user/login',
                type: 'post',
                dataType: 'json',
                data:{
                    'username':username,
                    'password':password
                },
                contentType: 'application/x-www-form-urlencoded;charset=UTF-8',
                cache: false,
                success: function(data) {
                    if(data.status==10000){
                        window.location.href="/";
                    }else if(data.status==10002){
                        alert(data.message);
                        // window.location.reload();
                    }else{
                        alert("您的网络异常!");
                    }

                },
                error : function() {
                    alert("Connect SERVER Error!")
                }
            });
        }


    }


</script>
</html>
