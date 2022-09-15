package nacos_client

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/sirupsen/logrus"
)

type Cluster struct {
	nodes  map[string]*Nodes
	logger *logrus.Logger
}

func NewCluster(logger *logrus.Logger) *Cluster {
	return &Cluster{nodes: make(map[string]*Nodes, 0), logger: logger}
}

func (c *Cluster) GetNode(serviceName string, selector SelectorName, hash string) (*Node, error) {
	if v, ok := c.nodes[serviceName]; ok {
		return v.GetNode(selector, hash)
	} else {
		return nil, NoAvailableNodeError
	}
}

func (c *Cluster) GetAllNode(serviceName string) []*Node {
	if v, ok := c.nodes[serviceName]; ok {
		return v.Member()
	} else {
		return make([]*Node, 0)
	}
}

func (c *Cluster) InitNode(client naming_client.INamingClient, ServiceName, groupName string, cluster []string) error {
	c.nodes[ServiceName] = NewNodes(NewRandomSelector(), NewHashSelector())
	services, err := client.SelectAllInstances(vo.SelectAllInstancesParam{
		Clusters:    cluster,
		ServiceName: ServiceName,
		GroupName:   groupName,
	})
	if err != nil {
		return err
	}
	var tmp []*Node
	for _, service := range services {
		if service.Healthy {
			tmp = append(tmp, NewNode(service.Ip, int(service.Port), int(service.Weight)))
		}
	}
	c.nodes[ServiceName].SetNode(tmp...)
	return nil
}

func (c *Cluster) SubscribeNacos(client naming_client.INamingClient, ServiceName, groupName string, cluster []string) error {
	// 这个订阅好像有bug, 非常不准确.
	return client.Subscribe(&vo.SubscribeParam{
		ServiceName: ServiceName,
		Clusters:    cluster,
		GroupName:   groupName,
		SubscribeCallback: func(services []model.SubscribeService, err error) {
			if c.logger != nil {
				c.logger.Infof("subscribe trigger services:%+v err:%+v", services, err)
			}
			var tmp []*Node
			for _, service := range services {
				tmp = append(tmp, NewNode(service.Ip, int(service.Port), int(service.Weight)))
			}
			c.nodes[ServiceName].SetNode(tmp...)
		},
	})
}

func GetAllServiceName(client naming_client.INamingClient, nameSpace string) ([]string, error) {
	var ans []string
	var pageSize = uint32(100)
	for page := uint32(1); page <= 10; page++ {
		serviceInfos, err := client.GetAllServicesInfo(vo.GetAllServiceInfoParam{
			NameSpace: nameSpace,
			PageNo:    page,
			PageSize:  pageSize,
		})
		if err != nil {
			return nil, err
		}
		ans = append(ans, serviceInfos.Doms...)
		if serviceInfos.Count < int64(pageSize) {
			return ans, nil
		}
	}
	fmt.Println("Service count is more than 1000!")
	return ans, nil
}
