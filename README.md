# my-blog<br>
this resp has a static html , run on localhost:8080/<br>
use beego/orm to create database,and insert,query,delete,update database<br>

http://beego.me/<br>
###安装<br>
进入cmd<br>
####执行go get github.com/astaxie/beego<br>
生成文件夹<br>
D:\GOPATH\src\github.com\astaxie<br>
####执行go get github.com/beego/bee<br>
生成可执行文件 D:\GOPATH\bin\bee.exe<br>
配置D:\GOPATH\bin\到path环境变量中<br>
####创建新项目<br>
![img](https://github.com/willkernel/my-blog/blob/master/pngs/newproj.png)<br>

####运行test<br>
![img](https://github.com/willkernel/my-blog/blob/master/pngs/runtestapp.png)<br>

访问http://127.0.0.1:8080/<br>
![img](https://github.com/willkernel/my-blog/blob/master/pngs/profilehtml.png)<br>

####数据库配置<br>
http://golanghome.com/post/289<br>
####创建数据库<br>
![img](https://github.com/willkernel/my-blog/blob/master/pngs/createdatabase.png)<br>
#### main.go<br>
![img](https://github.com/willkernel/my-blog/blob/master/pngs/maingo.png)<br>
#### user.go<br>
![img](https://github.com/willkernel/my-blog/blob/master/pngs/usergo.png)<br>
#### db.go<br>
![img](https://github.com/willkernel/my-blog/blob/master/pngs/db.png)<br>

####运行bee run<br>
![img](https://github.com/willkernel/my-blog/blob/master/pngs/beerun.png)<br>


