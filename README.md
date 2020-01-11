[![Build Status](https://travis-ci.com/asppj/t-go-opentrace.svg?branch=master)](https://travis-ci.com/asppj/t-go-opentrace)
1. Logging - 用于记录离散的事件。例如，应用程序的调试信息或错误信息。它是我们诊断问题的依据。
2. Metrics - 用于记录可聚合的数据。例如，队列的当前深度可被定义为一个度量值，在元素入队或出队时被更新；HTTP 请求个数可被定义为一个计数器，新请求到来时进行累加。
3. Tracing - 用于记录请求范围内的信息。例如，一次远程方法调用的执行过程和耗时。它是我们排查系统性能问题的利器。

# [opencensus(beta)](https://opencensus.io/language-support/)
#  [OpenTelemetry(alpha)](https://opentelemetry.io/docs/golang/tracing/)

# opentrace 
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

### 1. gin-web服务方 埋点(完成，待测)

>  添加中间件的方式

### 2. requests请求 埋点（完成，待测）

>封装GET、POST、DELET、PUT等http请求方法

### 3. grpc 埋点（完成，待测）

>grpc 使用grpc 提供的钩子函数

### 4. DB驱动埋点

>1. mongodb 埋点 （未开始）
>2. redis 埋点 （未开始）
>3. es 埋点 （未开始）

## opentrance-go 介绍

#### 1. Trace 调用链

#### 2. Span 调用链节点

>#### 1. Span-Tag 调用链节点标签
```text
     ext.PeerHostname.Set(sp, c.Request.Host)
 	ext.PeerAddress.Set(sp, c.Request.RemoteAddr)
 	ext.PeerService.Set(sp, c.ClientIP())
 	ext.HTTPStatusCode.Set(sp, uint16(statusCode))
 	ext.HTTPMethod.Set(sp, c.Request.Method)
 	ext.HTTPUrl.Set(sp, c.Request.URL.Path)
 ```
>#### 2. Span-Log 调用链节点日志
```golang
	span.LogKV("Request:", ctx.Request)
    span.LogFields("Request:", ctx.Request)
```

### 3. Baggage item 全局范围不属于 span
```golang
// set
span.SetBaggageItem("greeting", greeting)
// get
greeting := span.BaggageItem("greeting")
```
### 4. Sampling,采样[0,1]
>sampling.priority - integer
```text
1. const，全量采集，采样率设置0,1 分别对应打开和关闭
2. probabilistic ，概率采集，默认万份之一，0~1之间取值，
3. rateLimiting ，限速采集，每秒只能采集一定量的数据
4. remote ，一种动态采集策略，根据当前系统的访问量调节采集策略
```
#### inject/extract 
>childOf、followOf
```text
1. opentracing.GlobalTracer().Extract 方法提取HTTP头中的spanContexts
2. opentracing.ChildOf 方法基于提取出来的spanContexts生成新的child spanContexts
3. opentracing.GlobalTracer().StartSpan 方法生成一个新的span
4. github.com/opentracing/opentracing-go/ext 通过ext可以为追踪添加一些tag来展示更多信息，比如URL，请求类型(GET，POST...), 返回码
5. sp.Finish() 结束这一个span
```
#### context 传递span
```golang
// 存储到 context 中
ctx := context.Background()
ctx = opentracing.ContextWithSpan(ctx, span)
//....

// 其他过程获取并开始子 span
span, ctx := opentracing.StartSpanFromContext(ctx, "newspan")
defer span.Finish()
// StartSpanFromContext 会将新span保存到ctx中更新

```

### context 传递

#### 网络方式传递
    Extrect 从Header中提取
    Extrect 从grpc Ctx中提取
    inject 注入Header
    Inject 注入grpc Ctx
    进程中：opentracing.SpanFromContext(ctx)或span, ctx := opentracing.StartSpanFromContext(ctx, "newspan")
    gin.Context:ctx.Request.Context()
    - gin.Context!=context.Context!=opentracing.SpanContext
    mongodb/redis/es:
        -context 不再使用DefaultContext()
        统一使用gin.request.Context()[因为携带了span]

## jaeger 搭建
> all_in_one 镜像 tools->jaeger->docker-compose.yml


#参考资料

1. [几种分布式调用链监控组件的比较](https://juejin.im/post/5a0579e6f265da4326524f0f#heading-0)
2. [opentracing翻译版](https://wu-sheng.gitbooks.io/opentracing-io/content/pages/spec.html)
3. [grpc-opentracing中间件](https://godoc.org/github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc)
4. [openTracing 讲解](https://github.com/yurishkuro/opentracing-tutorial)
5. [阿里云日志服务配置aliyun-log-jaeger](https://github.com/aliyun/aliyun-log-jaeger/blob/master/README_CN.md?)

# prometheus metrics 统计

## prometheus 收集数据

## Grafana 展示数据

### prometheus 指标类型
 
 1. Counter 单调增
 2. Guage  可增减
 3. Historygram 直方图
 4. Summary 聚合图（1-10占比，10-20占比。。。）
 
 #### gin-ext 中间件
 1. 启动时长-uptime
 2. 请求总数-http_request_count_total
 3. 请求延时-http_request_duration_seconds
 4. 请求字节-http_request_size_bytes
 5. 响应字节-http_response_size_bytes
 
 #### kafka 驱动封装
 
 ##### 消费者
 
 1. 消费延时-kafka_consumer_duration_seconds
 2. 消费成功数-kafka_consumer_success_total
 3. 消费失败数-kafka_consumer_failed_total
 4. Payload 字节数-kafka_consumer_payload_byte_size

 ##### 生产者
 1. 生产延时-kafka_product_duration_seconds
 2. 生产成功数-kafka_product_success_total
 3. 生产失败数-kafka_product_failed_total
 4. Payload 字节数-kafka_product_payload_byte_size

 
### 官方sdk
* github.com/prometheus/client_golang 

# 参考资料

1. [prometheus介绍](https://www.imhanjm.com/2019/10/06/%E6%B7%B1%E5%85%A5%E7%90%86%E8%A7%A3prometheus(go%20sdk)/)