package server

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
)

// MidServer server
type MidServer struct {
	// 客户端监听
	clientLis *net.TCPListener
	// 后端服务连接
	transferLis *net.TCPListener
	// 所有通道
	channels map[int]*Channel
	// 当前通道ID
	curChanelID int
	running     chan int
}

// Channel 通道
type Channel struct {
	// 通道ID
	id int
	// 客户端连接
	client net.Conn
	// 后端服务连接
	transfer net.Conn
	// 客户端接收消息
	clientRecvMsg chan []byte
	// 后端服务发送消息
	transferSendMsg chan []byte
}

// Transfer 转发
func (c *Channel) Transfer() {
	n, err := io.Copy(c.client, c.transfer)
	if err != nil {
		return
	}
	log.Println(n)
}

// New 创建一个服务器
func New() *MidServer {
	return &MidServer{
		channels:    make(map[int]*Channel),
		curChanelID: 0,
		running:     make(chan int),
	}
}

// Start 启动服务
func (s *MidServer) Start(clientPort int, transferPort int) error {
	// 监听客户端
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", clientPort))
	if err != nil {
		return err
	}
	s.clientLis, err = net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}
	// 转发
	addr, err = net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", transferPort))
	if err != nil {
		return err
	}
	s.transferLis, err = net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}
	go s.AcceptLoop()
	return nil
}

// Stop 关闭服务
func (s *MidServer) Stop() {
	_ = s.clientLis.Close()
	_ = s.transferLis.Close()
	// 循环关闭通道连接
	for _, v := range s.channels {
		_ = v.client.Close()
		_ = v.transfer.Close()
	}
	close(s.running)
}

// Wait 等待关闭
func (s *MidServer) Wait() {
	log.Printf("服务运行中...:Channelds Count:%d", len(s.channels))
	r := <-s.running
	log.Printf("服务关闭:%d。", r)
}

// DelChannel 删除通道
func (s *MidServer) DelChannel(id int) {
	chs := s.channels
	delete(chs, id)
	s.channels = chs
}

// AcceptLoop 处理连接
func (s *MidServer) AcceptLoop() {
	transfer, err := s.transferLis.Accept()
	if err != nil {
		return
	}
	defer func() {
		s.Stop()
	}()
	for {
		// 获取连接
		client, err := s.clientLis.Accept()
		if err != nil {
			continue
		}

		// 创建一个通道
		ch := &Channel{
			id:              s.curChanelID,
			client:          client,
			transfer:        transfer,
			clientRecvMsg:   make(chan []byte),
			transferSendMsg: make(chan []byte),
		}
		s.curChanelID++

		// 把通道加入channels中
		chs := s.channels
		chs[ch.id] = ch
		s.channels = chs
		// ch.Transfer()
		// 启一个goroutine处理客户端消息
		go s.ClientMsgLoop(ch)
		// 启一个goroutine处理后端服务消息
		go s.TransferMsgLoop(ch)
		go s.MsgLoop(ch)
	}
}

// ClientMsgLoop 处理客户端消息
func (s *MidServer) ClientMsgLoop(ch *Channel) {
	defer func() {
		log.Println("ClientMsgLoop exit")
	}()
	for {
		data, isClose := <-ch.transferSendMsg
		{
			// 判断channel是否关闭，如果是则返回
			if !isClose {
				return
			}
			_, err := ch.client.Write(data)
			if err != nil {
				return
			}
		}
	}
}

// TransferMsgLoop 处理后端服务消息
func (s *MidServer) TransferMsgLoop(ch *Channel) {
	defer func() {
		log.Println("TransferMsgLoop exit")
	}()
	for {
		data, isClose := <-ch.clientRecvMsg
		{
			// 判断channel是否关闭，如果是则返回
			if !isClose {
				return
			}
			_, err := ch.transfer.Write(data)
			if err != nil {
				return
			}
		}
	}
}

// MsgLoop 客户端与后端服务消息处理
func (s *MidServer) MsgLoop(ch *Channel) {
	defer func() {
		// 关闭channel，好让ClientMsgLoop与TransferMsgLoop退出
		close(ch.clientRecvMsg)
		close(ch.transferSendMsg)
		s.DelChannel(ch.id)
		log.Println("MsgLoop exit")
	}()
	buf := make([]byte, 1024)
	for {
		// TODO io.Copy
		n, err := ch.client.Read(buf)
		if err != nil {
			return
		}
		ch.clientRecvMsg <- buf[:n]
		n, err = ch.transfer.Read(buf)
		if err != nil {
			return
		}
		ch.transferSendMsg <- buf[:n]
	}
}

var s *MidServer

// NewServer 创建服务
func NewServer() {
	// 参数解析
	localPort := flag.Int("localPort", 5008, "客户端访问端口")
	remotePort := flag.Int("remotePort", 5009, "服务访问端口")
	flag.Parse()
	// if flag.NFlag() != 2 {
	flag.PrintDefaults()
	// 	os.Exit(1)
	// }
	if s != nil {
		return
	}
	s = New()
	// 启动服务
	_ = s.Start(*localPort, *remotePort)
	// 循环
	s.Wait()
}

// StopServer 关闭服务
func StopServer() {
	if s != nil {
		s.Stop()
	}
}
