// -------------------------------------------------------------------
// Generated by 365admin-publish
// -------------------------------------------------------------------
/*
---
title: Core Version
---
*/
package cmds

import (
	"context"

	"github.com/magicbutton/ai-pnp/execution"
	"github.com/magicbutton/ai-pnp/utils"
)

func HealthCoreversionPost(ctx context.Context, args []string) (*string, error) {

	result, pwsherr := execution.ExecutePowerShell("john", "*", "ai-pnp", "00-health", "20-coreversion.ps1", "")
	if pwsherr != nil {
		return nil, pwsherr
	}
	utils.PrintSkip2FirstAnd2LastLines(string(result))
	return nil, nil

}