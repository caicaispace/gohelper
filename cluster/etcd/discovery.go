package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/caicaispace/gohelper/server/grpc/client"

	clientV3 "go.etcd.io/etcd/client/v3"
)

type Discovery struct {
	client   *clientV3.Client
	nodeInfo *NodeInfo
	nodes    *NodesManager
}

func NewDiscovery(nodeInfo *NodeInfo, conf clientV3.Config, manager *NodesManager) (dis *Discovery, err error) {
	d := &Discovery{}
	d.nodeInfo = nodeInfo
	if manager == nil {
		return nil, fmt.Errorf("[Discovery] manager == nil")
	}
	d.nodes = manager
	d.client, err = clientV3.New(conf)
	return d, err
}

func (d *Discovery) Pull() {
	kv := clientV3.NewKV(d.client)
	resp, err := kv.Get(context.TODO(), "discovery/", clientV3.WithPrefix())
	if err != nil {
		log.Fatalf("[Discovery] kv.Get err:%+v", err)
		return
	}
	for _, v := range resp.Kvs {
		node := &NodeInfo{}
		err = json.Unmarshal(v.Value, node)
		if err != nil {
			log.Printf("[Discovery] json.Unmarshal err:%+v", err)
			continue
		}
		d.nodes.AddNode(node)
		log.Printf("[Discovery] pull node:%+v", node)
	}
}

func (d *Discovery) Watch() {
	watcher := clientV3.NewWatcher(d.client)
	watchChan := watcher.Watch(context.TODO(), "discovery", clientV3.WithPrefix())
	for {
		select {
		case resp := <-watchChan:
			d.watchEvent(resp.Events)
		}
	}
}

var c = client.GetInstance()

func (d *Discovery) watchEvent(evs []*clientV3.Event) {
	for _, ev := range evs {
		switch ev.Type {
		case clientV3.EventTypePut:
			node := &NodeInfo{}
			err := json.Unmarshal(ev.Kv.Value, node)
			if err != nil {
				log.Printf("[Discovery] json.Unmarshal err:%+v", err)
				continue
			}
			d.nodes.AddNode(node)
			log.Printf(fmt.Sprintf("[Discovery] new node:%s", string(ev.Kv.Value)))
			c.AddConn(node.Name, node.Addr, node.UniqueId)
		case clientV3.EventTypeDelete:
			d.nodes.DelNode(string(ev.Kv.Key))
			log.Printf(fmt.Sprintf("[Discovery] del node:%s data:%s", string(ev.Kv.Key), string(ev.Kv.Value)))
			c.DelConn(string(ev.Kv.Key))
		}
	}
}
