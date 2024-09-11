package files

import "embed"

//go:embed data/*
var DataFS embed.FS

//go:embed ui/*
var UIFS embed.FS
