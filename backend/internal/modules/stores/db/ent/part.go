// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"waterfall-backend/internal/modules/stores/db/ent/part"
	"waterfall-backend/internal/modules/stores/db/ent/release"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Part is the model entity for the Part schema.
type Part struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Number holds the value of the "number" field.
	Number int `json:"number,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Annotation holds the value of the "annotation" field.
	Annotation *string `json:"annotation,omitempty"`
	// ReleaseUUID holds the value of the "release_uuid" field.
	ReleaseUUID *string `json:"release_uuid,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PartQuery when eager-loading is set.
	Edges        PartEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PartEdges holds the relations/edges for other nodes in the graph.
type PartEdges struct {
	// Chapters holds the value of the chapters edge.
	Chapters []*Chapter `json:"chapters,omitempty"`
	// Release holds the value of the release edge.
	Release *Release `json:"release,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ChaptersOrErr returns the Chapters value or an error if the edge
// was not loaded in eager-loading.
func (e PartEdges) ChaptersOrErr() ([]*Chapter, error) {
	if e.loadedTypes[0] {
		return e.Chapters, nil
	}
	return nil, &NotLoadedError{edge: "chapters"}
}

// ReleaseOrErr returns the Release value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PartEdges) ReleaseOrErr() (*Release, error) {
	if e.loadedTypes[1] {
		if e.Release == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: release.Label}
		}
		return e.Release, nil
	}
	return nil, &NotLoadedError{edge: "release"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Part) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case part.FieldNumber:
			values[i] = new(sql.NullInt64)
		case part.FieldID, part.FieldTitle, part.FieldAnnotation, part.FieldReleaseUUID:
			values[i] = new(sql.NullString)
		case part.FieldCreatedAt, part.FieldUpdatedAt, part.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Part fields.
func (pa *Part) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case part.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				pa.ID = value.String
			}
		case part.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pa.CreatedAt = value.Time
			}
		case part.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pa.UpdatedAt = value.Time
			}
		case part.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				pa.DeletedAt = new(time.Time)
				*pa.DeletedAt = value.Time
			}
		case part.FieldNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field number", values[i])
			} else if value.Valid {
				pa.Number = int(value.Int64)
			}
		case part.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				pa.Title = value.String
			}
		case part.FieldAnnotation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field annotation", values[i])
			} else if value.Valid {
				pa.Annotation = new(string)
				*pa.Annotation = value.String
			}
		case part.FieldReleaseUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field release_uuid", values[i])
			} else if value.Valid {
				pa.ReleaseUUID = new(string)
				*pa.ReleaseUUID = value.String
			}
		default:
			pa.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Part.
// This includes values selected through modifiers, order, etc.
func (pa *Part) Value(name string) (ent.Value, error) {
	return pa.selectValues.Get(name)
}

// QueryChapters queries the "chapters" edge of the Part entity.
func (pa *Part) QueryChapters() *ChapterQuery {
	return NewPartClient(pa.config).QueryChapters(pa)
}

// QueryRelease queries the "release" edge of the Part entity.
func (pa *Part) QueryRelease() *ReleaseQuery {
	return NewPartClient(pa.config).QueryRelease(pa)
}

// Update returns a builder for updating this Part.
// Note that you need to call Part.Unwrap() before calling this method if this Part
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Part) Update() *PartUpdateOne {
	return NewPartClient(pa.config).UpdateOne(pa)
}

// Unwrap unwraps the Part entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pa *Part) Unwrap() *Part {
	_tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Part is not a transactional entity")
	}
	pa.config.driver = _tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Part) String() string {
	var builder strings.Builder
	builder.WriteString("Part(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pa.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pa.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pa.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := pa.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("number=")
	builder.WriteString(fmt.Sprintf("%v", pa.Number))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(pa.Title)
	builder.WriteString(", ")
	if v := pa.Annotation; v != nil {
		builder.WriteString("annotation=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := pa.ReleaseUUID; v != nil {
		builder.WriteString("release_uuid=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Parts is a parsable slice of Part.
type Parts []*Part
