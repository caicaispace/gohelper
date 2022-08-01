package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"github.com/pkg/errors"
	clientV3 "go.etcd.io/etcd/client/v3"
)

const (
	_ttl = 10
)

type Register struct {
	client    *clientV3.Client
	leaseId   clientV3.LeaseID
	lease     clientV3.Lease
	nodeInfo  *NodeInfo
	closeChan chan error
}

func NewRegister(nodeInfo *NodeInfo, conf clientV3.Config) (reg *Register, err error) {
	r := &Register{}
	r.closeChan = make(chan error)
	r.nodeInfo = nodeInfo
	r.client, err = clientV3.New(conf)
	return r, err
}

func (r *Register) Run() {
	dur := time.Second * 2
	timer := time.NewTicker(dur)
	r.register()
	for {
		select {
		case <-timer.C:
			r.keepAlive()
		case <-r.closeChan:
			goto EXIT
		}
	}
EXIT:
	log.Printf("[Register] Run exit...")
}

func (r *Register) Stop() {
	r.revoke()
	close(r.closeChan)
}

func (r *Register) register() (err error) {
	r.leaseId = 0
	kv := clientV3.NewKV(r.client)
	r.lease = clientV3.NewLease(r.client)
	leaseResp, err := r.lease.Grant(context.TODO(), _ttl)
	if err != nil {
		err = errors.Wrapf(err, "[Register] register Grant err")
		return
	}
	data, err := json.Marshal(r.nodeInfo)
	if err != nil {
		err = errors.Wrapf(err, "[Register] register json.Marshal err %s-%+v", r.nodeInfo.Name, string(data))
		return
	}
	_, err = kv.Put(context.TODO(), r.nodeInfo.UniqueId, string(data), clientV3.WithLease(leaseResp.ID))
	if err != nil {
		err = errors.Wrapf(err, "[Register] register kv.Put err %s-%+v", r.nodeInfo.Name, string(data))
		return
	}
	r.leaseId = leaseResp.ID
	return
}

func (r *Register) keepAlive() (err error) {
	_, err = r.lease.KeepAliveOnce(context.TODO(), r.leaseId)
	if err != nil {
		// 租约丢失，重新注册
		if err == rpctypes.ErrLeaseNotFound {
			r.register()
			err = nil
		}
		err = errors.Wrapf(err, "[Register] keepAlive err")
	}
	log.Printf(fmt.Sprintf("[Register] keepalive... leaseId:%+v", r.leaseId))
	return err
}

func (r *Register) revoke() (err error) {
	_, err = r.client.Revoke(context.TODO(), r.leaseId)
	if err != nil {
		err = errors.Wrapf(err, "[Register] revoke err")
		return
	}
	log.Printf(fmt.Sprintf("[Register] revoke node:%+v", r.leaseId))
	return
}
