package converter

import (
	"config-transpiler/config"
	"reflect"
	"testing"
)

func TestWriteFiles(t *testing.T) {
	tests := []struct {
		in  config.CloudConfig
		out config.Butane
	}{
		{
			config.CloudConfig{
				WriteFiles: []config.CloudConfigWriteFiles{
					{
						Path:        "/etc/demo1.txt",
						Content:     "Write files test 1",
						Permissions: "0644",
						Owner:       "root:root",
					},
					{
						Path:        "/etc/demo2.conf",
						Content:     "Write files test 2",
						Permissions: "0644",
						Owner:       "devs:bosses",
					},
				},
			},
			config.Butane{
				Storage: config.Storage{
					Files: []config.File{
						{
							Path: "/etc/demo1.txt",
							Contents: config.FileContents{
								Inline: "Write files test 1",
							},
							Mode: 420,
							User: config.FileUser{
								Name: "root",
							},
							Group: config.FileGroup{
								Name: "root",
							},
						},
						{
							Path: "/etc/demo2.conf",
							Contents: config.FileContents{
								Inline: "Write files test 2",
							},
							Mode: 420,
							User: config.FileUser{
								Name: "devs",
							},
							Group: config.FileGroup{
								Name: "bosses",
							},
						},
					},
				},
			},
		},
		{
			config.CloudConfig{
				WriteFiles: []config.CloudConfigWriteFiles{
					{
						Path:        "/etc/demo1.txt",
						Content:     "Write files test 1",
						Permissions: "755",
						Owner:       "root:root",
					},
				},
			},
			config.Butane{
				Storage: config.Storage{
					Files: []config.File{
						{
							Path: "/etc/demo1.txt",
							Contents: config.FileContents{
								Inline: "Write files test 1",
							},
							Mode: 493,
							User: config.FileUser{
								Name: "root",
							},
							Group: config.FileGroup{
								Name: "root",
							},
						},
					},
				},
			},
		},
	}
	for i, test := range tests {
		var out config.Butane
		writeFilesConvert(test.in, &out)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("Failed Test %d, Expected %+v, Received %+v", i, test.out, out)
		}
	}

}
