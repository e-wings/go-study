<!DOCTYPE html>
<html lang="zh-cn">
<head>
<meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
<title></title>
</head>
<body>
<form action="/register?username=adam" method="post">
    用户名：<input type="text" name="username"><br/>
    密码：<input type="password" name="password"><br/>
    年龄：<input type="text" name="age"><br/>
    中文姓名：<input type="text" name="name-cn"><br/>
    英文姓名：<input type="text" name="name-en"><br/>
    email：<input type="text" name="email"><br/>
    手机：<input type="text" name="mobile"><br/>
	身份证号：<input type="text" name="id-card"><br/>
    性别：	<input type="radio" name="gender" value="1">男
			<input type="radio" name="gender" value="0">女
			<input type="radio" name="gender" value="2">错误的性别<br/>
    最喜欢的水果：
    <select name="fruit">
		<option value=0>苹果</option>
		<option value=1>香蕉</option>
		<option value=2>梨</option>
		<option value=3>超出范围的选项</option>
	</select><br/>
	喜欢的运动：
	<input type="checkbox" name="interest" value="football">足球
	<input type="checkbox" name="interest" value="basketball">篮球
	<input type="checkbox" name="interest" value="tennis">网球
	<input type="checkbox" name="interest" value="wrong">错误选项<br/>
	出生日期：<input type="date" name="birth-date"><br/>

    <input type="submit" value="提交">
</form>
</body>
</html>