## DevOps
>- DevOps（Development 和 Operations 的组合词）是一种重视“软件开发人员（Dev）”和“IT 运维技术人员（Ops）”之间沟通合作的文化、运动或惯例。
>- 透过自动化“软件交付”和“架构变更”的流程，来使得构建、测试、发布软件能够更加地快捷、频繁和可靠。

## CI/CD
>- CI 持续集成（Continuous Integration）
>- CD 持续交付（Continuous Delivery）
>- CD 持续部署（Continuous Deployment）

## DockerHub & gitlab 
>- 参考网站 https://docs.docker.com/ci-cd/github-actions/ 

#### 分别创建账号 
>- https://hub.docker.com/ 
>- https://github.com/

#### GitHub 创建仓库及代码 
>- New repository  -- https://github.com/new
>- 本地新建仓库代码 or  git clone git@github.com:jianzhuliu/docker-whale.git
>- 本地github 建立连接  git remote add github git@github.com:yourname/yourrepository.git 
>- 提交代码 git push github main

#### GitHub 与 DockerHub 配置秘钥信息

#### 1、DockerHub 创建 Access Token
>- 点击个人 Account Settings -> Security  New Access Token
>- 描述随便填，如你的  DockerHub id, 创建后会获得一个 Token, 复制保存 

#### 2、GitHub 配置 DockerHub 信息
>- 对应仓库点击 Settings->Secrets   New repository secret 创建2个
>- Name: DOCKER_HUB_USERNAME , Value: 你的 DockerHub id  
>- Name:DOCKER_HUB_ACCESS_TOKEN ,Value: 刚才DockerHub创建的 Token

#### 3、GitHub 设置 workflow
>- 对应仓库点击 Action -> New workflow
>- main.yml

```
# This is a basic workflow to help you get started with Actions

name: CI to Docker Hub

# Controls when the action will run. 
on:
  # Triggers the workflow on push or pull request events but only for the main branch
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

  # Allows you to run this workflow manually from the Actions tab
  workflow_dispatch:

# A workflow run is made up of one or more jobs that can run sequentially or in parallel
jobs:
  # This workflow contains a single job called "build"
  build:
    # The type of runner that the job will run on
    runs-on: ubuntu-latest

    # Steps represent a sequence of tasks that will be executed as part of the job
    steps:
      # Checks-out your repository under $GITHUB_WORKSPACE, so your job can access it
      - name: Check Out Repo
        uses: actions/checkout@v2
        
      - name: Login to Docker Hub
        uses: docker/login-action@v1
        with:
          username: ${{ secrets.DOCKER_HUB_USERNAME }}
          password: ${{ secrets.DOCKER_HUB_ACCESS_TOKEN }}
      
      - name: Set up Docker Buildx
        id: buildx
        uses: docker/setup-buildx-action@v1
        
      - name: Build and push
        id: docker_build
        uses: docker/build-push-action@v2
        with:
          context: ./
          file: ./Dockerfile
          push: true
          tags: ${{ secrets.DOCKER_HUB_USERNAME }}/docker-whale:latest

      - name: Image digest
        run: echo ${{ steps.docker_build.outputs.digest }}

```

>- 如需使用缓存
```
# This is a basic workflow to help you get started with Actions

name: CI to Docker Hub

# Controls when the action will run. 
on:
  # Triggers the workflow on push or pull request events but only for the main branch
  push:
    branches: [ main ]
	tags:
      - "v*.*.*"
  pull_request:
    branches: [ main ]

  # Allows you to run this workflow manually from the Actions tab
  workflow_dispatch:

# A workflow run is made up of one or more jobs that can run sequentially or in parallel
jobs:
  # This workflow contains a single job called "build"
  build:
    # The type of runner that the job will run on
    runs-on: ubuntu-latest

    # Steps represent a sequence of tasks that will be executed as part of the job
    steps:
      # Checks-out your repository under $GITHUB_WORKSPACE, so your job can access it
      - name: Check Out Repo
        uses: actions/checkout@v2
		
	  - name: Login to Docker Hub
        uses: docker/login-action@v1
        with:
          username: ${{ secrets.DOCKER_HUB_USERNAME }}
          password: ${{ secrets.DOCKER_HUB_ACCESS_TOKEN }}
      - name: Build and push
        id: docker_build
        uses: docker/build-push-action@v2
        with:
          context: ./
          file: ./Dockerfile
          builder: ${{ steps.buildx.outputs.name }}
          push: true
          tags:  ${{ secrets.DOCKER_HUB_USERNAME }}/simplewhale:latest
          cache-from: type=local,src=/tmp/.buildx-cache
          cache-to: type=local,dest=/tmp/.buildx-cache
      - name: Image digest
        run: echo ${{ steps.docker_build.outputs.digest }}
		
```