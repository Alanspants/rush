# Docker

## 1. Docker 基础操作

### 1.1 启动容器

启动一个nginx容器：

```
docker run --name webserver-nginx -d -p 80:80 nginx
```

--name: 容器名

-d: 在后台运行

-p 80:80: 端口指定，宿主机端口：容器端口

nginx：镜像名称，可以添加“:tag_num"

### 1.2 如何进入容器

```
docker attach container ID
```

用attach进入容器后，如果退出，容器就会退出，不推荐

```shell
docker exec -it container ID bash
```

exec退出后，容器不会退出

### 1.3 停止容器

```
docker stop container ID
```

可以接多个containerID

### 1.4 删除镜像

1. 停止容器：docker stop container
2. 删除容器：docker rm container
3. 删除镜像：docker image rmi imageName

### 1.5 下载镜像

```
docker pull [option] [docker Registry address [:port]/] repo_name [:tag]
```

### 1.6 用commit提交镜像

```
$ docker exec -it webserver-nginx bash #进入容器
$ echo '<h1>Hello World<\h1>' > /usr/share/nginx/html/index.html
$ docker commit --author "alannchen" --message "modify default homepage" webserver-nginx nginx-custome:v1
```

--author: 作者

--message: 修改comment

webserver-nginx: 修改的容器名

nginx-custome:v1: 生成后的img名和版本

### 1.7 挂载volume

```shell
docker volume create --name test
```

创建一个volume

```shell
docker run --name ubuntu-volume-test -d -it -v test:/data ubuntu /bin/bash
```

已这个镜像创建一个容器，并挂在一个volume

--name: 容器名

-d: 后台保持

-it: 用交互式线程保证ubuntu容器能运行

-v: 挂在对应的数据卷 容器内路径：物理机路径

## 2. Dockerfile

### 2.1 常见的dockerfile指令和语法

#### 2.1.1 RUN

1. RUN 命令一行会加一层镜像
2. RUN命令中的语句可以用&&连接

```shell
RUN buildDeps='gcc libc6-dev make' \
$$ apt-get update
```

#### 2.1.2 CMD

容器需要一个主进程一直在运行，CDM用于启动容器默认的主进程

```shell
shell格式：CMD <command>
exec格式：CMD [command, argv2, arg2]
```

```shell
CMD ["ngxin","-g","daemon off;"]
```

#### 2.1.3 ENTRYPOINT

将命令包含在容器的运行时中

```dockerfile
FROM centos:7.2
ENTRYPOINT ["bin/cat"]
```

对于这个dockerfile，镜像运行的时候则自带有cat功能，如果我们在运行这个镜像的时候写上一路径，就可以进行测试。

```shell
docker run -it testImage /etc/hosts
```

#### 2.1.4 COPY & ADD

 把某个文件/脚本从一个路径复制到另一个路径。

```dockerfile
COPY <source_path> <target_path>
COPY ["source_path","target_path"]

e.g.
COPY install* /opt/shell
COPY install /opt/shell
```

相对于COPY，ADD可以将source中包含的压缩包自动解压，同时将URL进行自动下载，所以如果是单纯的复制，还是建议使用COPY

```dockerfile
ADD https://baidu.com/img/209123123
```

 #### 2.1.5 EXPOSE

端口暴露

```dockerfile
EXPOSE 物理机port 容器port
```

### 2.2 Example

```dockerfile
FROM oraclelinux:7-slim
ENV PACKAGE_URL https://repo.mysql.com/yum/...

RUN rpmkeys --import http://repo.mysql.com \
&& yum install -y $PACKAGE_URL \
$$ yum install -y libpwquality \

VOLUME /var/lib/mysql

COPY docker-entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

EXPOSE 3306 33060
CMD ["mysqld"]
```

# K8S

## 1. 组件

### 1.1 pod

k8s的最小管理单元，存放着一组有关联关系的容器，这组容器一起共用相同的网络和存储资源

### 1.2 node

节点

master node中运行着API server、scheduler、RC等关键业务。相当于k8s的大脑

work node中运行着一个或多个pod、kubelet、kube-proxy

### 1.3 cluster

一个或多个master + 一个或多个work组成集群

### 1.4 ectd

etcd是一个分布式的键值对数据库，存放着k8s的一些状态信息。

同时提供共享配置和服务的注册发现功能。

### 1.5 API Server

k8s的神经系统，提供各类资源对象的crud和http rest接口。

集群内部组件通过它进行交互通信，然后将交互的信息持久化道etcd中。

### 1.6 Scheduler

资源调度。

整体调度流程分为：预选调度 + 最优调度

1. 预选调度：遍历所有的目标节点，筛选出符合要求的候选节点。
2. 最优调度：在预选调度过程的基础上，采用优选策略计算出每个候选节点的积分，选最高分节点。

### 1.7 Replication Controller

RC，控制器的一种。

通过RC来控制pod的数量，确保pod以指定的副本数运行。

### 1.8 ReplicaSet

RC的升级版。唯一的区别是RS支持集合式标签选择器，而RC只支持等式标签选择器。

（集合式：标签可以有多个值；等值：只能选择一个值或不选）

### 1.9 Deployment

通过deployment操作RC和RS。简单理解为通过yml文件的声明，可以在其中定义pod的数量、更新方式、镜像、资源等。

### 1.10 Label

键值对形式的tag，用来区分资源。

### 1.11 Selector

标签选择器。

### 1.12 Kubelet

运行在work节点上的关键服务，负责与master节点沟通。

master中的apiserver会发送pod的一些配置信息到work节点里，由kubelet负责接收。

同时kubelet定期向master汇报work节点的使用情况：资源使用情况，容器状态等。
