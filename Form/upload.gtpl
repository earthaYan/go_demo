
<html>
<head>
<title></title>
</head>
<body>
<form action="/upload" method="post" enctype="multipart/form-data">
  <input type="file" name="uploadfile" />
  <br/>
  <input type="hidden" name="token" value="{{.}}"/>
  <br/>
  <input type="submit" value="upload" />
</form>
</body>
</html>