package converter

import (
	"config-transpiler/config"
	"reflect"
	"testing"
)


func TestGroupsConvert(t *testing.T){
	tests := []struct{
		in config.CloudConfig
		out config.Butane
	}{
		{
			config.CloudConfig{
				Groups: []string{"devs", "bosses"},
			},
			config.Butane{
				Passwd: config.Passwd{
					Groups: []config.ButanePasswdGroups{
						{Name: "devs"},
						{Name: "bosses"},
					},
				},
			},
		},

		{
			config.CloudConfig{
				Groups: []string{"employees"},
			},
			config.Butane{
				Passwd: config.Passwd{
					Groups: []config.ButanePasswdGroups{
						{Name: "employees"},
					},
				},
			},
		},
	}

	for _, test := range tests{
		var but config.Butane
		groupsConvert(test.in,&but)
		if !reflect.DeepEqual(test.out,but){
			t.Errorf("Expected : %+v , Received : %+v for groupsConvert()", test.out, but)
		}
	}
}
