package plugin

import (
	"github.com/metafates/testo/internal/directive"
)

// Plan for running the tests.
type Plan struct {
	// Modify may filter or re-order planned tests in-place.
	// Nil values are ignored.
	//
	// It will not receive subtests.
	Modify func(tests *[]PlannedTest)
}

// PlannedTest is a test to be scheduled for execution.
type PlannedTest interface {
	directive.DoNotImplement

	// Name of the test.
	Name() string

	// Info about this test.
	Info() TestInfo
}

func mergePlans(plugins ...Spec) Plan {
	return Plan{
		Modify: func(tests *[]PlannedTest) {
			for _, p := range plugins {
				// TODO: stop iterating if len(tests) == 0?

				if p.Plan.Modify != nil {
					p.Plan.Modify(tests)
				}
			}
		},
	}
}
