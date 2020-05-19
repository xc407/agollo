package roundrobin

import (
	"sync"

	"github.com/xc407/agollo/v3/env/config"
	"github.com/xc407/agollo/v3/loadbalance"
)

func init() {
	loadbalance.SetLoadBalance(&RoundRobin{})
}

//InitLoadBalance 初始化负载均衡器
func InitLoadBalance() {

}

//RoundRobin 轮询调度
type RoundRobin struct {
}

//Load 负载均衡
func (r *RoundRobin) Load(servers *sync.Map) *config.ServerInfo {
	var returnServer *config.ServerInfo
	servers.Range(func(k, v interface{}) bool {
		server := v.(*config.ServerInfo)
		// if some node has down then select next node
		if server.IsDown {
			return true
		}
		returnServer = server
		return false
	})
	return returnServer
}
