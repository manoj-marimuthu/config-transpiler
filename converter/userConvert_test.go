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
				[]config.User{
					{Name: "example", Gecos: "example-Gecos"},
					{Name: "example2", Gecos: "example-2-Gecos"},
				},
			},
			config.Butane{
				Variant:"flatcar",
				version:"1.0.0",
				config.Passwd{
					[]config.User{
						{Name: "example", Gecos: "example-Gecos"},
						{Name: "example2", Gecos: "example-2-Gecos"},
					},
				},
			},
		},
		// test case 2
		{
			config.CloudConfig{
				[]config.User{
					{Name: "example3", Gecos: ""},
				},
			},
			config.Butane{
				Variant:"flatcar",
				Version:"1.0.0",
				config.Passwd{
					[]config.User{
						{Name: "example3", Gecos: ""},
					},
				},
			},
		},
		// test case 3
		{
			config.CloudConfig{},
			config.Butane{
				Variant:"flatcar",
				Version:"1.0.0",
			},
		},
	}
	for i, test := range tests {
		var out config.Butane
		userConvert(test.in, &out)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("test %d, Expected: %+v Received: %+v \n", i, test.out, out)
		}
	}
}
