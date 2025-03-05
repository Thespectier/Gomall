# Gomall
TikTok E-commerce Project of ByteDance Youth Training Camp Based on Golang

### 项目结构介绍 ###

本项目的开发参考学习了[字节青训营后端项目](https://github.com/cloudwego/biz-demo/tree/main/gomall),主要采用由字节跳动开源的，用于构建微服务架构的中间件集合[CloudWeGo](https://www.cloudwego.io/zh/),各个微服务架构的生成指令都在Makefile中给出了实例。

本次微服务之间通信，主要采用的是RPC框架。其中各个微服务的服务端放在app文件夹中，各个微服务的接口protobuf文件放在idl文件夹中，客户端程序放在rpc_gen中。在微服务之间进行通信时，只需要在获取rpc模块，就能如同调用本地函数一样调用所有微服务的接口，省去了复杂的网络通信编程。

客户端程序主要是作为调用的接口，所以用户在编程开发时不用考虑客户端的具体代码实现。

对于Kitex框架来说，项目开发是更加需要关注服务端的业务逻辑实现，以及些许环境文件配置。

##### 服务端结构 #####

1. biz/dal层中存储着数据库相关配置文件，在项目开发过程中，若微服务涉及到数据库的操作时，则会在其中创建model模块，专门用来编写并存放基于gorm的数据库中间件。

2. biz/service层中存放着在protobuf中定义的服务框架，所有业务逻辑代码都是在这一层实现的。

3. conf层中存放着yaml文件，也就是微服务的运行环境，如本地开发时rpc对应的端口号，数据库的本地端口号等等，根据部署环境的不同可以在不同文件夹下进行修改，并启动对应yaml，默认是测试环境test。conf.go则是便于开发者在需要调用相关环境时，提供了方法。

4. infra/rpc(rpc)层，当需要调用其余微服务接口时，需要进行一些rpc框架相关的初始化，就会需要在这个文件夹中编写相应微服务调用的初始化代码。

5. 微服务相关的容器部署(本项目基本上只涉及mysql和consul)都在docker-compose文件中，由于一直在本地测试，所以直接使用根目录下的docker-compose即可。

### 测试步骤 ###

1. 在根目录下执行：

``` shell
# 首次运行时执行
docker-compose up -d
# 后续再次运行时
docker-compose restart -d
#or docker compose start -d
```

根据docker-compose文件启动容器。

2. 进入相应的微服务，如cart，然后配置一下Golang的依赖包

``` shell
cd app/cart
go mod tidy
```

3. 由于在项目根目录下已经加载了workspace，所以直接运行即可

``` shell
go run .
```

4. 服务端正常运行之后就可以运行service层中的单元测试代码，对该服务进行测试。

5. 测试结束后关闭容器：
``` shell
# 停止服务
docker compose stop
# 停止并删除服务
docker compose down 
```

### 参考资料 ###

https://github.com/cloudwego/biz-demo/tree/main/gomall

https://gorm.io/zh_CN/docs/index.html

https://www.cloudwego.io/zh/docs/kitex/

https://www.cloudwego.io/zh/docs/hertz/

https://github.com/cloudwego/biz-demo/tree/main/gomall