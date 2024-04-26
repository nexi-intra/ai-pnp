package main

import (
	"runtime/debug"
	"strings"

	"github.com/magicbutton/ai-pnp/magicapp"
	"github.com/magicbutton/ai-pnp/utils"
)

func main() {
	info, _ := debug.ReadBuildInfo()

	// split info.Main.Path by / and get the last element
	s1 := strings.Split(info.Main.Path, "/")
	name := s1[len(s1)-1]
	description := `---
title: ai-pnp
description: Describe the main purpose of this kitchen
---

# ai-pnp
`
	magicapp.Setup(".env")
	magicapp.RegisterServeCmd("ai-pnp", description, "0.0.1", 8080)
	magicapp.RegisterCmds()
	utils.RootCmd.PersistentFlags().BoolVarP(&utils.Verbose, "verbose", "v", false, "verbose output")

	magicapp.Execute(name, "ai-pnp", "")
}
