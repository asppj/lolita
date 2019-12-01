# openTrace 演进
1. 埋点：gin-web,http请求，grpc请求封装并埋点；
2. 收集：zipkin,AppDash,jaeger 选择其一（阿里云日志服务优先）
3. 展示：

## 埋点
    1. Inject 请求方注入
        - http header
        - grpc with接口
    2. Extract 服务方提取
        - http header
        - grpc with接口

### gin-web服务方 埋点

#### 1. 添加中间件的方式

### http请求方 埋点

#### 1. 封装GET、POST、DELET、PUT等http请求方法

### grpc 埋点

#### 1. grpc 使用grpc 提供的钩子函数

## opentrance-go 介绍

#### 1. Trace 调用链

#### 2. Span 调用链节点

#### 3. Span-Tag 调用链节点标签

#### 4. Span-Log 调用链节点日志

#####
```text
1. opentracing.GlobalTracer().Extract 方法提取HTTP头中的spanContexts
2. opentracing.ChildOf 方法基于提取出来的spanContexts生成新的child spanContexts
3. opentracing.GlobalTracer().StartSpan 方法生成一个新的span
4. github.com/opentracing/opentracing-go/ext 通过ext可以为追踪添加一些tag来展示更多信息，比如URL，请求类型(GET，POST...), 返回码
5. sp.Finish() 结束这一个span
```


