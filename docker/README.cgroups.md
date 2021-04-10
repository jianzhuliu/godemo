## Cgroups -- control groups
Cgroups 是Linux系统的一个特性，它可以实现限制进程或者进程组的资源，如CPU、内存、磁盘IO等

## cgroups 主要提供的功能
>- 资源限制: 限制资源的使用量，不如内存上限
>- 优先级控制:不同的组可以有不同的资源使用优先级
>- 审计:计算控制组的资源使用情况
>- 控制: 控制进程的挂起或者恢复

## cgroups 功能实现依赖的三个核心概念
>- 子系统（subsystem）：是一个内核的组件，一个子系统代表一类资源调度控制器。例如内存子系统可以限制内存的使用量，CPU 子系统可以限制 CPU 的使用时间。
>- 控制组（cgroup）：表示一组进程和一组带有参数的子系统的关联关系。例如，一个进程使用了 CPU 子系统来限制 CPU 的使用时间，则这个进程和 CPU 子系统的关联关系称为控制组。
>- 层级树（hierarchy）：是由一系列的控制组按照树状结构排列组成的。这种排列方式可以使得控制组拥有父子关系，子控制组默认拥有父控制组的属性，也就是子控制组会继承于父控制组。比如，系统中定义了一个控制组 c1，限制了 CPU 可以使用 1 核，然后另外一个控制组 c2 想实现既限制 CPU 使用 1 核，同时限制内存使用 2G，那么 c2 就可以直接继承 c1，无须重复定义 CPU 限制。

## 查看当前系统已经挂载的 cgroups 信息
>- mount -t cgroup  

>- ls -l /sys/fs/cgroup

## cpu 子系统

##### 在 cpu 子系统下创建 cgroup
>- mkdir /sys/fs/cgroup/cpu/mydocker
>- ls -l /sys/fs/cgroup/cpu/mydocker
>- cd /sys/fs/cgroup/cpu/mydocker
>- echo 50000 > cpu.cfs_quota_us    --cpu 限制时间为 0.5 核

#### 创建进程，加入 cgroup
>- cd /sys/fs/cgroup/cpu/mydocker
>- echo $$ > tasks  -- 把当前运行的 shell 进程加入 cgroup 
>- cat tasks

#### 执行 CPU 耗时任务
>- while true; do echo ; done;
>- 重新打开一个 shell 窗口，使用 top -p 查看当前 cpu 使用率, 比如 top -p 5661


## memory 子系统

#### 在 memory 子系统下创建 cgroup
>- mkdir /sys/fs/cgroup/memory/mydocker
>- ls -l /sys/fs/cgroup/memory/mydocker
>- echo $((1<<30)) > memory.limit_in_bytes  --限制内存使用 1G
>- cat memory.limit_in_bytes    -- 1073741824, 2^10=1024=1KB, 2^20=1024*1024=1MB,2^30=1024*1024*1024=1GB

#### 创建进程，加入 cgroup
>- cd /sys/fs/cgroup/memory/mydocker
>- echo $$ > tasks  -- 把当前运行的 shell 进程加入 cgroup 
>- cat tasks


## 删除 cgroup
>- 退出当前 shell 窗口，新开一个 shell 窗口，root执行
>- rmdir /sys/fs/cgroup/cpu/mydocker
>- rmdir /sys/fs/cgroup/memory/mydocker

