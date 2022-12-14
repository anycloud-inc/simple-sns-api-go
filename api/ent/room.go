// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"simple_sns_api/ent/room"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Room is the model entity for the Room schema.
type Room struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// UsersId holds the value of the "usersId" field.
	UsersId string `json:"usersId,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoomQuery when eager-loading is set.
	Edges RoomEdges `json:"edges"`
}

// RoomEdges holds the relations/edges for other nodes in the graph.
type RoomEdges struct {
	// RoomUsers holds the value of the roomUsers edge.
	RoomUsers []*RoomUser `json:"roomUsers,omitempty"`
	// Messages holds the value of the messages edge.
	Messages []*Message `json:"messages,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// RoomUsersOrErr returns the RoomUsers value or an error if the edge
// was not loaded in eager-loading.
func (e RoomEdges) RoomUsersOrErr() ([]*RoomUser, error) {
	if e.loadedTypes[0] {
		return e.RoomUsers, nil
	}
	return nil, &NotLoadedError{edge: "roomUsers"}
}

// MessagesOrErr returns the Messages value or an error if the edge
// was not loaded in eager-loading.
func (e RoomEdges) MessagesOrErr() ([]*Message, error) {
	if e.loadedTypes[1] {
		return e.Messages, nil
	}
	return nil, &NotLoadedError{edge: "messages"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Room) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case room.FieldUsersId:
			values[i] = new(sql.NullString)
		case room.FieldCreatedAt, room.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case room.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Room", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Room fields.
func (r *Room) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case room.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				r.ID = *value
			}
		case room.FieldUsersId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field usersId", values[i])
			} else if value.Valid {
				r.UsersId = value.String
			}
		case room.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case room.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryRoomUsers queries the "roomUsers" edge of the Room entity.
func (r *Room) QueryRoomUsers() *RoomUserQuery {
	return (&RoomClient{config: r.config}).QueryRoomUsers(r)
}

// QueryMessages queries the "messages" edge of the Room entity.
func (r *Room) QueryMessages() *MessageQuery {
	return (&RoomClient{config: r.config}).QueryMessages(r)
}

// Update returns a builder for updating this Room.
// Note that you need to call Room.Unwrap() before calling this method if this Room
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Room) Update() *RoomUpdateOne {
	return (&RoomClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Room entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Room) Unwrap() *Room {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Room is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Room) String() string {
	var builder strings.Builder
	builder.WriteString("Room(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("usersId=")
	builder.WriteString(r.UsersId)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(r.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Rooms is a parsable slice of Room.
type Rooms []*Room

func (r Rooms) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
