package jq

import "github.com/cloudfoundry/packit"

func Detect() packit.DetectFunc {
	return func(context packit.DetectContext) (packit.DetectResult, error) {
		return packit.DetectResult{
			Plan: packit.BuildPlan{
				Provides: []packit.BuildPlanProvision{
					{Name: "jq"},
				},
				Requires: []packit.BuildPlanRequirement{
					{Name: "jq"},
				},
			},
		}, nil
	}
}
