package models

import (
	"sync"
	"testing"

	"gorm.io/gorm/schema"
)

// Cek wiring many2many User.HasRoles: pastikan GORM nulis ke tabel
// user_has_roles dengan kolom user_id + role_id (dipakai juga di query filter role).
func TestUserHasRolesJoinTable(t *testing.T) {
	s, err := schema.Parse(&User{}, &sync.Map{}, schema.NamingStrategy{})
	if err != nil {
		t.Fatalf("parse User schema: %v", err)
	}

	rel, ok := s.Relationships.Relations["HasRoles"]
	if !ok {
		t.Fatal("relasi HasRoles tidak ada")
	}
	if rel.JoinTable == nil {
		t.Fatal("HasRoles bukan many2many (JoinTable nil)")
	}
	if got := rel.JoinTable.Table; got != "user_has_roles" {
		t.Fatalf("join table = %q, mau user_has_roles", got)
	}
	for _, col := range []string{"user_id", "role_id"} {
		if rel.JoinTable.LookUpField(col) == nil {
			t.Fatalf("kolom %q tidak ada di %s", col, rel.JoinTable.Table)
		}
	}
}
