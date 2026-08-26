package converter

import (
	"config-transpiler/config"
	"reflect"
	"testing"
)

func TestUserConverter(t *testing.T) {
	tests := []struct {
		in  config.CloudConfig
		out config.Butane
	}{
		// test case 1
		{
			config.CloudConfig{
				Users: []config.CloudConfigUser{
					{Name: "example", Gecos: "example-Gecos", Shell: "/bin/bash", Uid: 1004},
					{Name: "example2", Gecos: "example-2-Gecos", Shell: "/bin/bash", Uid: 1023},
				},
			},
			config.Butane{
				Passwd: config.Passwd{
					Users: []config.ButaneUser{
						{Name: "example", Gecos: "example-Gecos", Shell: "/bin/bash", Uid: 1004},
						{Name: "example2", Gecos: "example-2-Gecos", Shell: "/bin/bash", Uid: 1023},
					},
				},
			},
		},
		// test case 2
		{
			config.CloudConfig{
				Users: []config.CloudConfigUser{
					{Name: "example3", Gecos: "", Shell: "", Uid: 2300},
				},
			},
			config.Butane{
				Passwd: config.Passwd{
					Users: []config.ButaneUser{
						{Name: "example3", Gecos: "", Shell: "", Uid: 2300},
					},
				},
			},
		},
		// test case 3
		{
			config.CloudConfig{},
			config.Butane{},
		},
	}
	for i, test := range tests {
		var out config.Butane
		userConvert(test.in, &out)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("test %d, Expected: %+v Received: %+v \n \n", i, test.out, out)
		}
	}
}
