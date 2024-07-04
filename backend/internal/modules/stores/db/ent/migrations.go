// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"waterfall-backend/internal/modules/stores/db/ent/migrations"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Migrations is the model entity for the Migrations schema.
type Migrations struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Migrated holds the value of the "migrated" field.
	Migrated     int `json:"migrated,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Migrations) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case migrations.FieldID, migrations.FieldMigrated:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Migrations fields.
func (m *Migrations) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case migrations.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case migrations.FieldMigrated:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field migrated", values[i])
			} else if value.Valid {
				m.Migrated = int(value.Int64)
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Migrations.
// This includes values selected through modifiers, order, etc.
func (m *Migrations) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// Update returns a builder for updating this Migrations.
// Note that you need to call Migrations.Unwrap() before calling this method if this Migrations
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Migrations) Update() *MigrationsUpdateOne {
	return NewMigrationsClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Migrations entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Migrations) Unwrap() *Migrations {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Migrations is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Migrations) String() string {
	var builder strings.Builder
	builder.WriteString("Migrations(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("migrated=")
	builder.WriteString(fmt.Sprintf("%v", m.Migrated))
	builder.WriteByte(')')
	return builder.String()
}

// MigrationsSlice is a parsable slice of Migrations.
type MigrationsSlice []*Migrations