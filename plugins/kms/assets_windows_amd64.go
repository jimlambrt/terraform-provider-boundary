package kms_plugin_assets

import (
	"embed"
)

// content is our static kms plugin content.
//go:embed assets/windows/amd64
var content embed.FS