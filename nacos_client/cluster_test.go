package nacos_client

import (
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/sirupsen/logrus"
	"reflect"
	"testing"
)

const (
	ServiceA = "SA"
	ServiceB = "SB"
)

var ClusterA = []string{"CA"}

type MockNamingClient struct {
}

func (m MockNamingClient) RegisterInstance(_ vo.RegisterInstanceParam) (bool, error) {
	panic("implement me")
}

func (m MockNamingClient) DeregisterInstance(_ vo.DeregisterInstanceParam) (bool, error) {
	panic("implement me")
}

func (m MockNamingClient) UpdateInstance(_ vo.UpdateInstanceParam) (bool, error) {
	panic("implement me")
}

func (m MockNamingClient) GetService(_ vo.GetServiceParam) (model.Service, error) {
	panic("implement me")
}

func (m MockNamingClient) SelectAllInstances(_ vo.SelectAllInstancesParam) ([]model.Instance, error) {
	return []model.Instance{
		{Ip: "127.0.0.1", Port: 80, Weight: 100},
		{Ip: "127.0.0.1", Port: 81, Weight: 100},
		{Ip: "127.0.0.1", Port: 82, Weight: 100},
	}, nil
}

func (m MockNamingClient) SelectInstances(_ vo.SelectInstancesParam) ([]model.Instance, error) {
	panic("implement me")
}

func (m MockNamingClient) SelectOneHealthyInstance(_ vo.SelectOneHealthInstanceParam) (*model.Instance, error) {
	return &model.Instance{Ip: "127.0.0.1", Port: 80, Weight: 100}, nil
}

func (m MockNamingClient) Subscribe(_ *vo.SubscribeParam) error {
	return nil
}

func (m MockNamingClient) Unsubscribe(_ *vo.SubscribeParam) error {
	panic("implement me")
}

func (m MockNamingClient) GetAllServicesInfo(_ vo.GetAllServiceInfoParam) (model.ServiceList, error) {
	return model.ServiceList{Count: 2, Doms: []string{ServiceA, ServiceB}}, nil
}

func TestGetNode(t *testing.T) {
	cluster := NewCluster(logrus.StandardLogger())
	_ = cluster.SubscribeNacos(&MockNamingClient{}, "", "", ClusterA)
	_ = cluster.InitNode(&MockNamingClient{}, "", "", ClusterA)
	_, _ = cluster.GetNode(ServiceA, RandomSelector, "")
	_, _ = GetAllServiceName(&MockNamingClient{}, "")
}

func TestGetAllServiceName(t *testing.T) {
	type args struct {
		client    naming_client.INamingClient
		nameSpace string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"ok", args{client: &MockNamingClient{}, nameSpace: ""}, []string{ServiceA, ServiceB}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAllServiceName(tt.args.client, tt.args.nameSpace)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllServiceName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllServiceName() got = %v, want %v", got, tt.want)
			}
		})
	}
}
