package converter

import (
	"config-transpiler/config"
	"strconv"
	"strings"
)

func writeFilesConvert(in config.CloudConfig, out *config.Butane) {
	for _, writeFile := range in.WriteFiles {
		// handle owner parameter, separate username and groupname
		// It is of the form -> owner : <username>:<groupname>
		owner_info := strings.Split(writeFile.Owner, ":")
		username := owner_info[0]
		groupname := owner_info[1]
		mode, err := strconv.ParseInt(writeFile.Permissions, 8, 32)
		if err != nil {
			panic(err)
		}

		butaneStorageFile := config.File{
			Path: writeFile.Path,
			Mode: int(mode),
			Contents: config.FileContents{
				Inline: writeFile.Content,
			},
			User: config.FileUser{
				Name: username,
			},
			Group: config.FileGroup{
				Name: groupname,
			},
		}
		out.Storage.Files = append(out.Storage.Files, butaneStorageFile)
	}
}

func init() {
	register(writeFilesConvert)
}
