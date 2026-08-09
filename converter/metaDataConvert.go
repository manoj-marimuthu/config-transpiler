package converter

import (
	"config-transpiler/config"
)

func metaDataConvert(in config.CloudConfig, out *config.Butane) {
	out.Variant = "flatcar"
	out.Version = "1.0.0"
}

func init() {
	register(metaDataConvert)
}
