# 项目位置
>- https://gitee.com/jianzhuliu
>- https://github.com/jianzhuliu

# 下载 git 
https://git-scm.com/

# win下配置 git
>- 任意位置右键，选择 Git Bash Here
>- cd ~/.ssh/  如果没有，则创建 mkdir ~/.ssh ，然后进入
>- 配置全局的 name 和 email 
git config --global user.name "yourname"
git config --global user.email "youremail"

>- ssh-keygen -t rsa -C "youremail"
直接按三次回车，无需密码

>- 生成私秘钥在
C:\Users\Administrator\.ssh

>- 打开 id_rsa.pub 到gitee 或者github 添加公钥即可

# 相关命令
>- git init

>- git add *

>- git commit -m "description"

>- git remote add gitee git@gitee.com:jianzhuliu/godemo.git

>- git remote add github git@github.com:jianzhuliu/godemo.git

>- git push gitee master

>- git push github master
