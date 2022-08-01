package client

import (
	"fmt"
	"math/rand"
	"sync"

	"google.golang.org/grpc"
)

type ClientConn struct {
	name     string
	addr     string
	uniqueId string
	conn     *grpc.ClientConn
}

type Client struct {
	sync.RWMutex
	conns       map[string][]*ClientConn
	currentConn int
}

var (
	client *Client
	once   sync.Once
)

func GetInstance() *Client {
	once.Do(func() {
		client = New()
	})
	return client
}

func New() *Client {
	return &Client{
		conns: make(map[string][]*ClientConn),
	}
}

func (c *Client) AddConn(name, addr, uniqueId string) error {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	// 延迟关闭连接
	// defer conn.Close()
	c.conns[name] = append(c.conns[name], &ClientConn{
		name:     name,
		addr:     addr,
		uniqueId: uniqueId,
		conn:     conn,
	})
	return nil
}

func (c *Client) DelConn(uniqueId string) {
	fmt.Println(uniqueId)
	for _, conn1 := range c.conns {
		for k, conn2 := range conn1 {
			if conn2.uniqueId == uniqueId {
				c.conns[conn2.name] = append(c.conns[conn2.name][:k], c.conns[conn2.name][k+1:]...)
				conn2.conn.Close()
			}
		}
	}
}

func (c *Client) Pick(name string) *grpc.ClientConn {
	return c.RoundRobinBalance(name)
}

func (c *Client) RoundRobinBalance(name string) *grpc.ClientConn {
	c.RLock()
	defer c.RUnlock()
	if conns, exist := c.conns[name]; !exist {
		return nil
	} else {
		connsLen := len(conns)
		switch connsLen {
		case 0:
			return nil
		case 1:
			return conns[0].conn
		default:
			c.currentConn = (c.currentConn + 1) % len(conns)
			return conns[c.currentConn].conn
		}
	}
}

func (c *Client) RandomBalance(name string) *grpc.ClientConn {
	c.RLock()
	defer c.RUnlock()
	if conns, exist := c.conns[name]; !exist {
		return nil
	} else {
		// 纯随机取节点
		idx := rand.Intn(len(conns))
		for _, v := range conns {
			if idx == 0 {
				return v.conn
			}
			idx--
		}
	}
	return nil
}
