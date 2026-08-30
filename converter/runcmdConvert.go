package converter

import (
	"config-transpiler/config"
	"fmt"
	"strings"
)

func runCmdConvert(in config.CloudConfig, out *config.Butane) {
	for i, command := range in.RunCmd {
		var execStart string
		if command.Command != ""{
			execStart = fmt.Sprintf("/bin/sh -c '%s'",command.Command)
		}else{
			execStart = strings.Join(command.Args," ")
		}
		butaneUnit := config.ButaneSystemdUnit{
			Name:    fmt.Sprintf("cloud-init-runcmd-%d.service", i),
			Enabled: true,
			Contents: fmt.Sprintf(`[Unit]
Description=cloud-init runcmd unit
[Service]
Type=oneshot
ExecStart=%s
[Install]
WantedBy=multi-user.target`, execStart),
		}
		out.Systemd.Units = append(out.Systemd.Units, butaneUnit)
	}
}

func init() {
	register(runCmdConvert)
}
