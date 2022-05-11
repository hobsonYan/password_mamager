### 密码管理工具

创建带图标的可执行文件
1、创建rc文件（main.rc），输入 IDI_ICON1 ICON "password_manage.ico"
2、创建syso文件，windres -o main.syso main.rc
3、执行 go build -o ../bin/password_manager.exe