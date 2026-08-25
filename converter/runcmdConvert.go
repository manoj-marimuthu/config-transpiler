package converter

import (
	"config-transpiler/config"
	"fmt"
)

func runCmdConvert(in config.CloudConfig, out *config.Butane){
	for i, command := range in.RunCmd{
		butaneUnit := config.ButaneSystemdUnit{
			Name: fmt.Sprintf("cloud-init-runcmd-%d.service",i),
			Enabled: true,
			Contents: fmt.Sprintf(`[Unit]
Description=cloud-init runcmd unit
[Service]
Type=oneshot
ExecStart=/bin/sh -c '%s'
[Install]
WantedBy=multi-user.target`, command),
		}
		out.Systemd.Units = append(out.Systemd.Units, butaneUnit)
	}
}
func init(){
	register(runCmdConvert)
}
