# ZH

![](./static/imgs/LMP-logo.png)
# Linux microscope

LMP是一个基于BCC(BPF Compiler Collection)的Linux系统性能数据实时展示的web工具，它使用BPF(Berkeley Packet Filters)，也叫eBPF，目前LMP在ubuntu18.04上测试通过，内核版本4.15.0。

## 项目架构

![](./static/imgs/LMP-arch4.png)

## 界面截图

![homepage](./static/imgs/homepage.png)

![homepage](./static/imgs/grafana1.png)

![grafana4](./static/imgs/grafana4.png)


##  运行lmp

###  Ubuntu-source

#### 从源码构建lmp，需要的基本环境：

- golang：go1.12及以上；
- docker：influxdb、grafana；
- bcc环境

###  安装依赖docker镜像

```
# For grafana
sudo docker pull grafana/grafana
# For Influxdb
sudo docker pull influxdb
```


### 编译并安装

```
 git clone https://github.com/linuxkerneltravel/lmp
 cd lmp
 make
```

## 单机节点，本地运行

```
# 项目的所有配置均位于config.yaml中，grafana的默认端口为3000端口，influxdb的默认端口为8086，修改配置信息的方式如下：
 vim lmp/config.yaml

#run grafana
 sudo docker run -d \
   -p 3000:3000 \
   --name=grafana \
   grafana/grafana

#run influxdb，按照如下命令启动influxdb之后，会自动带有database lmp，influxdb的用户名和密码位于config.yaml中
 sudo docker run -d \
    -p 8083:8083 \
    -p 8086:8086 \
    --name influxdb \
    -v ${YOUR_PROJECT_PATH}/lmp/test/influxdb_config/default.conf:/etc/influxdb/influxdb.conf \
    -v ${YOUR_PROJECT_PATH}/lmp/test/influxdb_config/data:/var/lib/influxdb/data \
    -v ${YOUR_PROJECT_PATH}/lmp/test/influxdb_config/meta:/var/lib/influxdb/meta \
    -v ${YOUR_PROJECT_PATH}/lmp/test/influxdb_config/wal:/var/lib/influxdb/wal influxdb

#run lmp
 cd lmp/
 make
 sudo ./lmp
```

### 观测-步骤

在单机节点上部署完成lmp并启动之后，通过浏览器访问8080端口即可。如果是本地查看，则访问localhost:8080，如果远程访问，则访问remoteip:8080即可。

8080端口返回页面如下，该页面仅用于观测指标下发，输入栏输入的是观测时间，单位是分钟。可一次下发多个指标，但是注意实际环境使用bcc的开销问题，建议单个指标下发对比开销之后，再组合多个指标观测。

![homepage](./static/imgs/homepage.png)

另外开启grafana页面，通过浏览器访问3000端口即可。如果是本地查看，则访问localhost:3000，如果远程访问，则访问remoteip:3000即可。

grafana用于指标数据观测，进入grafana之后，首先需要登录进入grafana，初始用户名和密码均为admin，之后需要配置grafana连接influxdb：

![grafana1](./static/imgs/grafana1.png)

按照自己的ip地址配置完成以后，点击save&test按钮，测试influxdb是否连接成功，出现如下提示说明连接成功：

![grafana2](./static/imgs/grafana2.png)

接下来导入/lmp/test/grafana-JSON下的lmp.json文件，即可自动创建grafana的dashboard：

![grafana3](./static/imgs/grafana3.png)

在Dashboards中Manage中，点击Import，上传lmp.json文件即可观测：

![grafana4](./static/imgs/grafana4.png)

统计时间到时后，lmp会自动关闭后台bcc插件，之后继续在8080端口页面下发指标即可。



![](./static/imgs/LMP-logo.png)

# Linux microscope

LMP is a web tool for real-time display of Linux system performance data based on BCC (BPF Compiler Collection), which uses BPF (Berkeley Packet Filter), also known as eBPF. Currently, LMP is tested on ubuntu18.04 and the kernel version is 4.15.0.

## Project architecture

![](./static/imgs/LMP-arch4.png)

## Interface screenshot

![homepage2](./static/imgs/homepage2.png)

![homepage](./static/imgs/grafana.png)

![homepage](./static/imgs/data.png)


## Project structure overview  

<details>
<summary>Expand to view</summary>
<pre><code>.
.
├── LICENSE
├── README.md
├── bcctest            Hold all test codes, including bcc、influxdb, etc
├── cmd                Store LMP pid number after startup
├── config.yaml        Project profile
├── controllers        Controller layer code stored in CLD layers
├── dao                Dao layer code stored in CLD layers
├── logger             Zap Log Library Initialization Related Code
├── logic              Logic layer code stored in CLD layers
├── main.go
├── makefile
├── middlewares        Holds middleware, such as JWT, used in the project
├── models             Data structure, such as a user、BpfScan, used in a storage project
├── pkg                A third-party library, such as JWT、snowflake, used in a project
├── plugins            Storage bcc plugins
├── routes             Store initialization routing code
├── settings           Hold viper Initialize related code
├── static             Hold static HTML files, pictures, etc
├── test               Store influxdb initial configuration, files, etc
└── vendor             Storage of project dependencies
</code></pre>
</details>


##  install lmp

###  Ubuntu-source

#### Build lmp from source，The basic environment required is as follows：

- golang
- docker
- bcc

###  Install dependent docker image

```
# For prometheus 
sudo docker pull prom/prometheus
# For grafana
sudo docker pull grafana/grafana
# For Influxdb
sudo docker pull influxdb
```

### Compile and install

```
 git clone https://github.com/linuxkerneltravel/lmp
 cd lmp
 make
 sudo make install
```

##  Single machine node, Run locally

```
# Modify configuration file
 vim lmp/config.yaml

#run grafana
 sudo docker run -d \
   -p 3000:3000 \
   --name=grafana \
   -v /opt/grafana-storage:/var/lib/grafana \
   grafana/grafana
   
#run influxdb
 sudo docker run -d \
    -p 8083:8083 \
    -p 8086:8086 \
    --name influxdb \
    -v ${YOUR_PROJECT_PATH}/lmp/test/influxdb_config/default.conf:/etc/influxdb/influxdb.conf \
    -v ${YOUR_PROJECT_PATH}/lmp/test/influxdb_config/data:/var/lib/influxdb/data \
    -v ${YOUR_PROJECT_PATH}/lmp/test/influxdb_config/meta:/var/lib/influxdb/meta \
    -v ${YOUR_PROJECT_PATH}/lmp/test/influxdb_config/wal:/var/lib/influxdb/wal influxdb
    
#run lmp
 cd lmp/
 make
 sudo ./lmp
```

### observation

http://localhost:8080/  After logging in to grafana, view it.

### Uninstall

```
make clean
```

## Thanks for the support of the following open source projects

- [Gin] - [https://gin-gonic.com/](https://gin-gonic.com/)
- [bcc] - [https://github.com/iovisor/bcc](https://github.com/iovisor/bcc)

# 