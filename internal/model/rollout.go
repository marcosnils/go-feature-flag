package model

import (
	"fmt"
	"strings"
	"time"
)

type Rollout struct {
	// Experimentation is your struct to configure an experimentation, it will allow you to configure a start date and
	// an end date for your flag.
	// When the experimentation is not running, the flag will serve the default value.
	Experimentation *Experimentation `json:"experimentation,omitempty" yaml:"experimentation,omitempty" toml:"experimentation,omitempty" slack_short:"false"` // nolint: lll

	// Progressive is your struct to configure a progressive rollout deployment of your flag.
	// It will allow you to ramp up the percentage of your flag over time.
	// You can decide at which percentage you starts and at what percentage you ends in your release ramp.
	// Before the start date we will serve the initial percentage and after we will serve the end percentage.
	Progressive *Progressive `json:"progressive,omitempty" yaml:"progressive,omitempty" toml:"progressive,omitempty" slack_short:"false"` // nolint: lll

	Scheduled *ScheduledRollout `json:"scheduled,omitempty" yaml:"scheduled,omitempty" toml:"scheduled,omitempty" slack_short:"false"` // nolint: lll
}

func (e Rollout) String() string {
	if e.Experimentation == nil {
		return ""
	}
	return "experimentation: " + e.Experimentation.String()
}

type Experimentation struct {
	// Start is the starting time of the experimentation
	Start *time.Time `json:"start,omitempty" yaml:"start,omitempty" toml:"start,omitempty"`

	// End is the ending time of the experimentation
	End *time.Time `json:"end,omitempty" yaml:"end,omitempty" toml:"end,omitempty"`
}

func (e Experimentation) String() string {
	buf := make([]string, 0)
	lo, _ := time.LoadLocation("UTC")

	if e.Start != nil {
		buf = append(buf, fmt.Sprintf("start:[%v]", e.Start.In(lo).Format(time.RFC3339)))
	}
	if e.End != nil {
		buf = append(buf, fmt.Sprintf("end:[%v]", e.End.In(lo).Format(time.RFC3339)))
	}
	return strings.Join(buf, " ")
}

// Progressive is the configuration struct to define a progressive rollout.
type Progressive struct {
	// Percentage is where you can configure at what percentage your progressive rollout start
	// and at what percentage it ends.
	// This field is optional
	Percentage ProgressivePercentage `json:"percentage,omitempty" yaml:"percentage,omitempty" toml:"percentage,omitempty"`

	// ReleaseRamp is the defining when the progressive rollout starts and ends.
	// This field is mandatory if you want to use a progressive rollout.
	// If any field missing we ignore the progressive rollout.
	ReleaseRamp ProgressiveReleaseRamp `json:"releaseRamp,omitempty" yaml:"releaseRamp,omitempty" toml:"releaseRamp,omitempty"` // nolint: lll
}

type ProgressivePercentage struct {
	// Initial is the initial percentage before the rollout start date.
	// This field is optional
	// Default: 0.0
	Initial float64 `json:"initial,omitempty" yaml:"initial,omitempty" toml:"initial,omitempty"`

	// End is the target percentage we want to reach at the end of the rollout phase.
	// This field is optional
	// Default: 100.0
	End float64 `json:"end,omitempty" yaml:"end,omitempty" toml:"end,omitempty"`
}

type ProgressiveReleaseRamp struct {
	// Start is the starting time of the ramp
	Start *time.Time `json:"start,omitempty" yaml:"start,omitempty" toml:"start,omitempty"`

	// End is the ending time of the ramp
	End *time.Time `json:"end,omitempty" yaml:"end,omitempty" toml:"end,omitempty"`
}

type ScheduledRollout struct {
	Steps []ScheduledStep `json:"steps,omitempty" yaml:"steps,omitempty" toml:"steps,omitempty"`
	// scheduled:
	// steps:
	// - date: 2021-04-10T00:00:00.10-05:00
	// rule: internal eq true
	// percentage: 100
	//
	// - date: 2021-04-14T00:00:00.10-05:00
	// rule: internal eq true OR beta eq true
	// percentage: 100
	//
	// - date: 2021-05-21T00:00:00.10-05:00
	// rule: "" // no restriction on who is affected by the flag
	// percentage: 100
}

type ScheduledStep struct {
	FlagData `yaml:",inline"`
	Date     *time.Time `json:"date,omitempty" yaml:"date,omitempty" toml:"date,omitempty"`
}
