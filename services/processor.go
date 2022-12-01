package service

import (
	"github.com/EronAlves1996/GoTeste/graph/model"
	"github.com/EronAlves1996/GoTeste/services/validators"
)

func processPassword(password string, rules []*model.Rule) *model.Return {
	var noMatch = []*string{}
	var verify bool = true

	for _, r := range rules {
		if !validators.Validators[r.Rule](password, r.Value) {
			noMatch = append(noMatch, &r.Rule)
			verify = false
		}
	}

	return &model.Return{
		Verify:  &verify,
		NoMatch: noMatch,
	}
}