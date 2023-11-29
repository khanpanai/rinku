package ui

import "embed"

//go:embed dist/*
var EmbedDirStatic embed.FS
