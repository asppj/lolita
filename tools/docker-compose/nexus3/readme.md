## docker 私有库配置
1. 15003->5003 
    docker proxy 转发阿里云docker镜像服务 https://registry.cn-shanghai.aliyuncs.com
2. 15002->5002
    docker hosted 本地存储docker镜像
3. 15001->5001
    docker group 添加docker proxy 和 docker hosted，但是不能push dockerimage[可以通过nginx-proxy达到效果](./nginx-proxy.conf)