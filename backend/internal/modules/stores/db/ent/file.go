// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"waterfall-backend/internal/models"
	"waterfall-backend/internal/models/files"
	"waterfall-backend/internal/modules/stores/db/ent/file"
	"waterfall-backend/internal/modules/stores/db/ent/user"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// File is the model entity for the File schema.
type File struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Filename holds the value of the "filename" field.
	Filename string `json:"filename,omitempty"`
	// MimeType holds the value of the "mime_type" field.
	MimeType string `json:"mime_type,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// CreatorUUID holds the value of the "creator_uuid" field.
	CreatorUUID *string `json:"creator_uuid,omitempty"`
	// ObjectType holds the value of the "object_type" field.
	ObjectType models.ObjectType `json:"object_type,omitempty"`
	// ObjectRef holds the value of the "object_ref" field.
	ObjectRef string `json:"object_ref,omitempty"`
	// Type holds the value of the "type" field.
	Type files.Type `json:"type,omitempty"`
	// Temp holds the value of the "temp" field.
	Temp bool `json:"temp,omitempty"`
	// SequenceNumber holds the value of the "sequence_number" field.
	SequenceNumber *uint `json:"sequence_number,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FileQuery when eager-loading is set.
	Edges        FileEdges `json:"edges"`
	selectValues sql.SelectValues
}

// FileEdges holds the relations/edges for other nodes in the graph.
type FileEdges struct {
	// Creator holds the value of the creator edge.
	Creator *User `json:"creator,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CreatorOrErr returns the Creator value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FileEdges) CreatorOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Creator == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Creator, nil
	}
	return nil, &NotLoadedError{edge: "creator"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*File) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case file.FieldTemp:
			values[i] = new(sql.NullBool)
		case file.FieldSequenceNumber:
			values[i] = new(sql.NullInt64)
		case file.FieldID, file.FieldFilename, file.FieldMimeType, file.FieldDescription, file.FieldCreatorUUID, file.FieldObjectType, file.FieldObjectRef, file.FieldType:
			values[i] = new(sql.NullString)
		case file.FieldCreatedAt, file.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the File fields.
func (f *File) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case file.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				f.ID = value.String
			}
		case file.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				f.CreatedAt = value.Time
			}
		case file.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				f.UpdatedAt = value.Time
			}
		case file.FieldFilename:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field filename", values[i])
			} else if value.Valid {
				f.Filename = value.String
			}
		case file.FieldMimeType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mime_type", values[i])
			} else if value.Valid {
				f.MimeType = value.String
			}
		case file.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				f.Description = value.String
			}
		case file.FieldCreatorUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field creator_uuid", values[i])
			} else if value.Valid {
				f.CreatorUUID = new(string)
				*f.CreatorUUID = value.String
			}
		case file.FieldObjectType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field object_type", values[i])
			} else if value.Valid {
				f.ObjectType = models.ObjectType(value.String)
			}
		case file.FieldObjectRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field object_ref", values[i])
			} else if value.Valid {
				f.ObjectRef = value.String
			}
		case file.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				f.Type = files.Type(value.String)
			}
		case file.FieldTemp:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field temp", values[i])
			} else if value.Valid {
				f.Temp = value.Bool
			}
		case file.FieldSequenceNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sequence_number", values[i])
			} else if value.Valid {
				f.SequenceNumber = new(uint)
				*f.SequenceNumber = uint(value.Int64)
			}
		default:
			f.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the File.
// This includes values selected through modifiers, order, etc.
func (f *File) Value(name string) (ent.Value, error) {
	return f.selectValues.Get(name)
}

// QueryCreator queries the "creator" edge of the File entity.
func (f *File) QueryCreator() *UserQuery {
	return NewFileClient(f.config).QueryCreator(f)
}

// Update returns a builder for updating this File.
// Note that you need to call File.Unwrap() before calling this method if this File
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *File) Update() *FileUpdateOne {
	return NewFileClient(f.config).UpdateOne(f)
}

// Unwrap unwraps the File entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *File) Unwrap() *File {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: File is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *File) String() string {
	var builder strings.Builder
	builder.WriteString("File(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("created_at=")
	builder.WriteString(f.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(f.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("filename=")
	builder.WriteString(f.Filename)
	builder.WriteString(", ")
	builder.WriteString("mime_type=")
	builder.WriteString(f.MimeType)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(f.Description)
	builder.WriteString(", ")
	if v := f.CreatorUUID; v != nil {
		builder.WriteString("creator_uuid=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("object_type=")
	builder.WriteString(fmt.Sprintf("%v", f.ObjectType))
	builder.WriteString(", ")
	builder.WriteString("object_ref=")
	builder.WriteString(f.ObjectRef)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", f.Type))
	builder.WriteString(", ")
	builder.WriteString("temp=")
	builder.WriteString(fmt.Sprintf("%v", f.Temp))
	builder.WriteString(", ")
	if v := f.SequenceNumber; v != nil {
		builder.WriteString("sequence_number=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Files is a parsable slice of File.
type Files []*File
