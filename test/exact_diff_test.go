package test

import (
	"cli-text-compare/pkg"
	"reflect"
	"testing"
)

func TestExactDiff(t *testing.T) {
	textOne := []byte("selam\nnaber\nnasilsin")
	textTwo := []byte("sekam\nna\nnasilsin")

	actions := pkg.Diff(textOne, textTwo)

	knownActionsArr := []pkg.Action{
		"match", "match", "substitute", "match", "match", "match", "match", "match", "delete", "delete", "delete", "match", "match", "match", "match", "match", "match", "match", "match", "match",
	}

	if !reflect.DeepEqual(knownActionsArr, actions) {
		t.Logf("got:  %v", actions)
		t.Logf("want: %v", knownActionsArr)
		t.Fatal("you broke the algorithm")
	}
}
