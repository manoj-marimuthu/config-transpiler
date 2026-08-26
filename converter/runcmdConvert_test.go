package converter

import (
	"config-transpiler/config"
	"reflect"
	"testing"
)

func TestRunCmd(t *testing.T) {
	tests := []struct {
		in  config.CloudConfig
		out config.Butane
	}{
		{
			config.CloudConfig{
				RunCmd: []string{
					"echo hello",
					"ls -l",
				},
			},
			config.Butane{
				Systemd: config.Systemd{
					Units: []config.ButaneSystemdUnit{
						{
							Name:    "cloud-init-runcmd-0.service",
							Enabled: true,
							Contents: `[Unit]
Description=cloud-init runcmd unit
[Service]
Type=oneshot
ExecStart=/bin/sh -c 'echo hello'
[Install]
WantedBy=multi-user.target`,
						},
						{
							Name:    "cloud-init-runcmd-1.service",
							Enabled: true,
							Contents: `[Unit]
Description=cloud-init runcmd unit
[Service]
Type=oneshot
ExecStart=/bin/sh -c 'ls -l'
[Install]
WantedBy=multi-user.target`,
						},
					},
				},
			},
		},
	}
	for i, test := range tests {
		var out config.Butane
		runCmdConvert(test.in, &out)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("test %d, Expected: %+v, Received: %+v \n \n", i, test.out, out)
		}
	}
}
