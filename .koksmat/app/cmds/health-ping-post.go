// -------------------------------------------------------------------
// Generated by 365admin-publish
// -------------------------------------------------------------------
/*
---
title: Ping
---
*/
package cmds

import (
	"context"

	"github.com/magicbutton/ai-pnp/execution"
	"github.com/magicbutton/ai-pnp/utils"
)

func HealthPingPost(ctx context.Context, args []string) (*string, error) {

	result, pwsherr := execution.ExecutePowerShell("john", "*", "ai-pnp", "00-health", "10-ping.ps1", "", "-pong", args[0])
	if pwsherr != nil {
		return nil, pwsherr
	}
	utils.PrintSkip2FirstAnd2LastLines(string(result))
	return nil, nil

}
