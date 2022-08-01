package etcd

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// 创建租约注册服务
type ServiceRegister struct {
	client        *clientv3.Client
	lease         clientv3.Lease
	leaseResp     *clientv3.LeaseGrantResponse
	canclefunc    func()
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
	key           string
}

func NewServiceRegister(addr []string, timeNum int64) (*ServiceRegister, error) {
	conf := clientv3.Config{
		Endpoints:   addr,
		DialTimeout: 5 * time.Second,
	}

	var client *clientv3.Client

	if clientTem, err := clientv3.New(conf); err == nil {
		client = clientTem
	} else {
		return nil, err
	}

	ser := &ServiceRegister{
		client: client,
	}

	if err := ser.setLease(timeNum); err != nil {
		return nil, err
	}
	go ser.ListenLeaseRespChan()
	return ser, nil
}

// 设置租约
func (sr *ServiceRegister) setLease(timeNum int64) error {
	lease := clientv3.NewLease(sr.client)

	// 设置租约时间
	leaseResp, err := lease.Grant(context.TODO(), timeNum)
	if err != nil {
		return err
	}

	// 设置续租
	ctx, cancelFunc := context.WithCancel(context.TODO())
	sr.canclefunc = cancelFunc

	leaseRespChan, err := lease.KeepAlive(ctx, leaseResp.ID)
	if err != nil {
		return err
	}

	sr.lease = lease
	sr.leaseResp = leaseResp
	sr.keepAliveChan = leaseRespChan
	return nil
}

// 监听 续租情况
func (sr *ServiceRegister) ListenLeaseRespChan() {
	for {
		select {
		case leaseKeepResp := <-sr.keepAliveChan:
			if leaseKeepResp == nil {
				fmt.Printf("已经关闭续租功能\n")
				return
			} else {
				fmt.Printf("续租成功\n")
			}
		}
	}
}
