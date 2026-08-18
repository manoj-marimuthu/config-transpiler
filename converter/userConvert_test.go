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
					{Name: "example", Gecos: "example-Gecos",Shell:"/bin/bash"},
					{Name: "example2", Gecos: "example-2-Gecos",Shell:"/bin/bash"},
				},
			},
			config.Butane{
				Passwd: config.Passwd{
					Users: []config.ButaneUser{
						{Name: "example", Gecos: "example-Gecos", Shell:"/bin/bash"},
						{Name: "example2", Gecos: "example-2-Gecos", Shell:"/bin/bash"},
					},
				},
			},
		},
		// test case 2
		{
			config.CloudConfig{
				Users: []config.CloudConfigUser{
					{Name: "example3", Gecos: "",Shell:""},
				},
			},
			config.Butane{
				Passwd: config.Passwd{
					Users: []config.ButaneUser{
						{Name: "example3", Gecos: "",Shell:""},
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
