package main

import (
	"flag"
	"fmt"
	"net"
)

// Start 转发
func (c *Client) Start(localPort int) {
	buf := make([]byte, 1024)
	defer func() {
		c.Stop()
	}()
	for {
		// 先从远程读数据
		n, err := c.remoteConn.Read(buf)
		if err != nil {
			continue
		}
		data := buf[:n]
		// 建立与本地80服务的连接
		local, err := net.Dial("tcp", fmt.Sprintf(":%d", localPort))
		if err != nil {
			continue
		}
		// 向80服务写数据
		_, err = local.Write(data)
		if err != nil {
			continue
		}
		// 读取80服务返回的数据
		n, err = local.Read(buf)
		// 关闭80服务，因为本地80服务是http服务，不是持久连接
		// 一个请求结束，就会自动断开。所以在for循环里我们要不断Dial，然后关闭。
		local.Close()
		if err != nil {
			continue
		}
		data = buf[:n]
		// 向远程写数据
		_, err = c.remoteConn.Write(data)
		if err != nil {
			continue
		}
	}
}

// Client 客户端
type Client struct {
	running    chan int
	host       string
	remoteConn net.Conn
}

// New 新建连接
func New(host string) (c *Client, err error) {
	c = &Client{
		host:    host,
		running: make(chan int),
	}
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return
	}
	c.remoteConn = conn
	return
}

// Wait 等待关闭
func (c *Client) Wait() {
	fmt.Printf("客户端运行中...:Server:%s", c.host)
	r := <-c.running
	fmt.Printf("客户端关闭:%d。", r)
}

// Stop 关闭
func (c *Client) Stop() {
	close(c.running)
}
func main() {
	// 参数解析
	host := flag.String("host", "asppj.top", "服务器地址")
	remotePort := flag.Int("remotePort", 5009, "服务器端口")
	localPort := flag.Int("localPort", 6006, "本地端口")
	flag.Parse()
	// if flag.NFlag() != 3 {
	flag.PrintDefaults()
	// 	os.Exit(1)
	// }
	// 建立与服务器的连接
	// remote, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *remotePort))
	c, err := New(fmt.Sprintf("%s:%d", *host, *remotePort))
	if err != nil {
		fmt.Println(err)
	}
	c.Start(*localPort)
	c.Wait()
}
