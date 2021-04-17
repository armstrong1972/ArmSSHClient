# ArmSSHClient
<hr>
SSHClient by Golang
<hr>
xShell收费，FinalShell太慢，自己用Go语言开发一个简单易用的 SSH Client 套餐包，包括：

<br><b>1) shkey.exe : 生成一对RAS2048的密钥对字符串 (如不重新编译代码，可跳过此步)<br></b>

  <li>公钥: 复制到 shpwd.go 代码文件中，用于为 SSH 密码加密</li>
  <li>私钥: 复制到 sh.go 代码文件中，用于解密 配置文件 中的密码</li>
  
<br><b>2) shpwd.exe : 用公钥加密你的SSH站点的密码 ，将结果复制到配置文件的 “cipher” 参数中<br></b>
&nbsp;&nbsp;&nbsp;&nbsp;用法： > shpwd.exe yourpassword
  
<br><b>3) 增加 SSH配置文件（.json）， 配置文件放在 ./config 目录下<br></b>
    配置文件格式如下 ：<br>
    {<br>
      &nbsp;&nbsp;&nbsp;&nbsp;"mod"  : "pem", <br>
      &nbsp;&nbsp;&nbsp;&nbsp;"addr" : "15.15.15.15:22",<br>
      &nbsp;&nbsp;&nbsp;&nbsp;"user" : "ec2-user",<br>
      &nbsp;&nbsp;&nbsp;&nbsp;"cipher" : "aws.pem" <br>
    }<br>
    其中：<br>
      <li>mod : 有2个值可选。"pem" : 公钥模式 或 “pwd” : 密码模式</li>
      <li>cipher ：<br>
      &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;公钥模式 : pem文件名， 放在 ./config/pem 目录下<br>
      &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;密码模式 : 存放 shpwd.exe 程序加密后的 字符串
      </li>
                
<br><b>4) <font color=red>sh.exe</font>> : SSH主程序，带一个参数<br></b>
&nbsp;&nbsp;&nbsp;&nbsp;用法： > sh.exe config_file (不带扩展名)<br>
&nbsp;&nbsp;&nbsp;&nbsp;案例： > sh.exe demo1 <br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;自动加载 config 目录下的 demo1.json 配置文件，并连接SSH<br>

<img scr="demo.jpg">
		
