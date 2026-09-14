package models

import (
	"sync"
	"testing"

	"gorm.io/gorm/schema"
)

func TestRequiredCategoryRestrictsDeletion(t *testing.T) {
	for _, model := range []interface{}{&Category{}, &Post{}} {
		s, err := schema.Parse(model, &sync.Map{}, schema.NamingStrategy{})
		if err != nil {
			t.Fatal(err)
		}
		found := false
		for _, relation := range s.Relationships.Relations {
			constraint := relation.ParseConstraint()
			if constraint == nil || constraint.Name != "fk_categories_posts" {
				continue
			}
			if constraint.OnDelete != "RESTRICT" {
				t.Fatalf("required category must restrict deletion, got %q", constraint.OnDelete)
			}
			found = true
		}
		if !found {
			t.Fatalf("category constraint missing for %T", model)
		}
	}
}
