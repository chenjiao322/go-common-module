package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gitlab.geinc.cn/services/go-common-module/nacos_client"
	"gitlab.geinc.cn/services/go-common-module/utils"
	"time"
)

var client naming_client.INamingClient

func Init() {
	var err error
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("172.16.0.127", uint64(8848), constant.WithScheme("http")),
	}
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(""),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithUsername("nacos"),
		constant.WithPassword("nacos"),
		constant.WithLogLevel("error"),
	)
	client, err = clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}
	nacosInstance, err := NewNacosInstance()
	if err != nil {
		panic(err)
	}
	success, err := client.RegisterInstance(nacosInstance)
	if err != nil {
		panic(err)
	}
	if !success {
		panic("nacos fail")
	}
}

func NewNacosInstance() (vo.RegisterInstanceParam, error) {
	ip, err := utils.IP(nil)
	if err != nil {
		return vo.RegisterInstanceParam{}, err
	}
	return vo.RegisterInstanceParam{
		Ip:          ip.String(),
		Port:        48082,
		ServiceName: "ttt",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
	}, nil
}

func main() {
	Init()
	c := nacos_client.NewCluster(nil)
	ss, _ := nacos_client.GetAllServiceName(client, "")
	for _, s := range ss {
		err := c.InitNode(client, s, "", nil)
		fmt.Println(err)
		err = c.SubscribeNacos(client, s, "", nil)
		fmt.Println(err)
	}
	time.Sleep(time.Second * 10000)
}
