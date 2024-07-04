// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"waterfall-backend/internal/modules/stores/db/ent/seedmigrations"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// SeedMigrations is the model entity for the SeedMigrations schema.
type SeedMigrations struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Migrated holds the value of the "migrated" field.
	Migrated     int `json:"migrated,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SeedMigrations) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case seedmigrations.FieldID, seedmigrations.FieldMigrated:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SeedMigrations fields.
func (sm *SeedMigrations) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case seedmigrations.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sm.ID = int(value.Int64)
		case seedmigrations.FieldMigrated:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field migrated", values[i])
			} else if value.Valid {
				sm.Migrated = int(value.Int64)
			}
		default:
			sm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SeedMigrations.
// This includes values selected through modifiers, order, etc.
func (sm *SeedMigrations) Value(name string) (ent.Value, error) {
	return sm.selectValues.Get(name)
}

// Update returns a builder for updating this SeedMigrations.
// Note that you need to call SeedMigrations.Unwrap() before calling this method if this SeedMigrations
// was returned from a transaction, and the transaction was committed or rolled back.
func (sm *SeedMigrations) Update() *SeedMigrationsUpdateOne {
	return NewSeedMigrationsClient(sm.config).UpdateOne(sm)
}

// Unwrap unwraps the SeedMigrations entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sm *SeedMigrations) Unwrap() *SeedMigrations {
	_tx, ok := sm.config.driver.(*txDriver)
	if !ok {
		panic("ent: SeedMigrations is not a transactional entity")
	}
	sm.config.driver = _tx.drv
	return sm
}

// String implements the fmt.Stringer.
func (sm *SeedMigrations) String() string {
	var builder strings.Builder
	builder.WriteString("SeedMigrations(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sm.ID))
	builder.WriteString("migrated=")
	builder.WriteString(fmt.Sprintf("%v", sm.Migrated))
	builder.WriteByte(')')
	return builder.String()
}

// SeedMigrationsSlice is a parsable slice of SeedMigrations.
type SeedMigrationsSlice []*SeedMigrations