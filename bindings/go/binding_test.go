package tree_sitter_modelica_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-modelica"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_modelica.Language())
	if language == nil {
		t.Errorf("Error loading Modelica grammar")
	}
}
