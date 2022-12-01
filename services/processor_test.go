package services

import (
	"testing"

	"github.com/EronAlves1996/GoTeste/graph/model"
)

func TestProcessPassword(t *testing.T) {
	t.Run("Test processing password with upperCase rule", func(t *testing.T) {
		rules := []*model.Rule{}
		rules = append(rules, &model.Rule{Rule: "minUpperCase", Value: 2})

		got := ProcessPassword("WithTwoUpperacase", rules)
		expect := &model.Return{
			Verify:  boolPointer(true),
			NoMatch: []*string{},
		}

		assertProcessingPassword(got, expect, t)
	})

	t.Run("Test processing password with upperCase and digit rule", func(t *testing.T) {
		rules := []*model.Rule{}
		rules = append(rules, &model.Rule{Rule: "minUppercase", Value: 5}, &model.Rule{Rule: "minDigit", Value: 3})

		got := ProcessPassword("WithTwoUpperacaseand123", rules)
		expect := &model.Return{
			Verify:  boolPointer(false),
			NoMatch: []*string{},
		}

		assertProcessingPassword(got, expect, t)
	})
}

func assertProcessingPassword(got, expect *model.Return, t *testing.T) {
	t.Helper()
	if *got.Verify != *expect.Verify {
		t.Errorf("Expect: %t got %t on Verify", *expect.Verify, *got.Verify)
	}

	if *got.Verify && len(got.NoMatch) > 0 {
		t.Errorf("Unconsistent return on noMatch. Expect %v and got %v", expect.NoMatch, got.NoMatch)
	}
}

func boolPointer(b bool) *bool {
	return &b
}
