# linux

## 1. 基本命令

### 1.1 登录相关命令

#### 1.1.1 tty

区分当前的窗口是哪个页面，从0开始。

tty -> teletypewriters

#### 1.1.2 whoami

当前登录的用户是谁。

#### 1.1.3 who am i

结合以上两点

#### 1.1.4 who

当前机器所有的登录者 + 终端号

#### 1.1.5 w

看到每个终端的人在干啥。

### 1.2 bash相关命令

#### 1.2.1 hostname

更改主机名

```bash
# hostname 显示主机名
VM-1-2-centos
# hostname alannchen-centos 临时修改主机名
# hostname set-hostname alannchen-centos 永久修改主机名
```

#### 1.2.2 环境变量

```bash
# echo $PATH
/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/root/bin

一共有三个文件夹：
1. /usr/local/sbin
2. /usr/local/bin
3. /usr/sbin
4. /usr/bin
5. /root/bin
对于一个命令，会在以上五个路径中逐步查找，查不到则报错
```

#### 1.2.3 type

显示命令是built-in的，还是外部的

```bash
# type who
who is /usr/bin/who
# type echo
echo is a shell builtin
```

```bash
# type who
who is /usr/bin/who 未使用/首次使用
# who
root     pts/0        2022-11-26 22:39 (120.244.220.179) 首次使用
# type who
who is hashed (/usr/bin/who) hash表示被缓存在系统里
```

#### 1.2.4 hash

查看当前命令缓存，以及清空缓存

```bash
# hash
hits    command
   3    /usr/bin/hostname
   1    /usr/bin/cat
   1    /usr/bin/ps
   3    /usr/bin/who
# hash -r 清空所有缓存
# hash -d hostname 清空某个命令的缓存
```

### 1.3 格式命令

#### 1.3.1 alias

为一个命令设置别名

```bash
# alias host='hostnamectl set-hostname' 设置别名
# host alannchen-centos 使用别名
# alias host 查看别名
# unalias host 删除别名
# unalias -a 取消所有别名
```

优先级：别名 > 内部命令 > 外部命令

alias > hash > $PATH 

注：在命令行中定义的别名，仅对当前shell进程有效

如果想永久有效，要定义在配置文件中

- 仅对当前用户：~/.bashrc
- 对左右用户：/etc/bashrc

#### 1.3.2 快捷参数

使用上一个命令的参数

command !*

#### 1.3.3 重新调用

使用最近的，以某字母开头的命令

!h 调用最近的h开头的命令

### 1.4 常见硬件命令

#### 1.4.1 lscpu

查看cpu信息 

#### 1.4.2 free

查看内存大小

或者使用：free -h 自动换算单位

#### 1.4.3 lsblk

查看硬盘大小

#### 1.4.4 arch

看系统的架构

#### 1.4.5 查看系统版本

cat /etc/os-release

### 1.5 文件相关

#### 1.5.1 文件类型

- . 普通文件
- d 目录文件
- l 链接文件 link
- b 块设备 block
- c 字符设备 char
- p 管道文件 pipe （fifo，类似于消息队列，先进先出）
- s 套接字文件 socket （双方都可读写，比如mysql）

ls -l 的第一位可以查看

#### 1.5.2 cd - 

回到上次的目录

cd -

echo %OLDPWD

#### 1.5.3 ll

ll => ls -alf(ubuntu) => ls -al

ls: -a(显示隐藏文件) -l(显示文件详情) -F(显示文件类型)