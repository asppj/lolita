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

## jaeger 搭建
> all_in_one 镜像 tools->jaeger->docker-compose.yml


#参考资料

1. [几种分布式调用链监控组件的比较](https://juejin.im/post/5a0579e6f265da4326524f0f#heading-0)
2. [opentracing翻译版](https://wu-sheng.gitbooks.io/opentracing-io/content/pages/spec.html)
3. [grpc-opentracing中间件](https://godoc.org/github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc)
4. [openTracing 讲解](https://github.com/yurishkuro/opentracing-tutorial)