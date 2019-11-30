# openTrace 演进
1. 埋点：gin-web,http请求，grpc请求封装并埋点；
2. 收集：zipkin,AppDash,jaeger 选择其一（阿里云日志服务优先）
3. 展示：

## 埋点

### gin-web 埋点

#### 1. 添加中间件的方式

### http 埋点

#### 1. 封装GET、POST、DELET、PUT等http请求方法

### grpc 埋点

#### 1. grpc 使用grpc 提供的钩子函数

## opentrance-go 介绍

#### 1. Trace 调用链

#### 2. Span 调用链节点

#### 3. Span-Tag 调用链节点标签

#### 4. Span-Log 调用链节点日志