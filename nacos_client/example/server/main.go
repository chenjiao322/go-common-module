package main

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gitlab.geinc.cn/services/go-common-module/utils"
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
		constant.WithUpdateCacheWhenEmpty(true),
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
		Port:        48083,
		ServiceName: "ttt",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
	}, nil
}

func main() {
	Init()
	//ip, err := utils.IP(nil)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//time.Sleep(time.Second * 3000)
	//a, e := client.DeregisterInstance(vo.DeregisterInstanceParam{
	//	Ip:          ip.String(),
	//	Port:        48083,
	//	ServiceName: "ttt",
	//	Ephemeral:   true,
	//})
	//fmt.Println(a, e)
}
