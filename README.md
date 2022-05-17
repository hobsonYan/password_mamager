### 密码管理工具

##### 创建带图标的可执行文件
1. 创建rc文件（main.rc），输入 
    
    IDI_ICON1 ICON "password_manage.ico"

2. 创建syso文件

    windres -o main.syso main.rc

3. 执行 build 命令

    go build -o ../bin/password_manager.exe

    创建 linux 下运行的文件

    set GOOS=linux
    set GOARCH=amd64
    go build -o ../bin/password_manager_linux

##### 创建 docker 镜像
    docker build -f ./Dockerfile -t password_manager .