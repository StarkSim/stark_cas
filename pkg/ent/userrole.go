// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cas/pkg/ent/role"
	"cas/pkg/ent/user"
	"cas/pkg/ent/userrole"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// UserRole is the model entity for the UserRole schema.
type UserRole struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy int64 `json:"created_by"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy int64 `json:"updated_by"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at"`
	// UserID holds the value of the "user_id" field.
	UserID int64 `json:"user_id,omitempty"`
	// RoleID holds the value of the "role_id" field.
	RoleID int64 `json:"role_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserRoleQuery when eager-loading is set.
	Edges UserRoleEdges `json:"edges"`
}

// UserRoleEdges holds the relations/edges for other nodes in the graph.
type UserRoleEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Role holds the value of the role edge.
	Role *Role `json:"role,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserRoleEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// RoleOrErr returns the Role value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserRoleEdges) RoleOrErr() (*Role, error) {
	if e.loadedTypes[1] {
		if e.Role == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: role.Label}
		}
		return e.Role, nil
	}
	return nil, &NotLoadedError{edge: "role"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserRole) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userrole.FieldID, userrole.FieldCreatedBy, userrole.FieldUpdatedBy, userrole.FieldUserID, userrole.FieldRoleID:
			values[i] = new(sql.NullInt64)
		case userrole.FieldCreatedAt, userrole.FieldUpdatedAt, userrole.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UserRole", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserRole fields.
func (ur *UserRole) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userrole.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ur.ID = int64(value.Int64)
		case userrole.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				ur.CreatedBy = value.Int64
			}
		case userrole.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				ur.UpdatedBy = value.Int64
			}
		case userrole.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ur.CreatedAt = value.Time
			}
		case userrole.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ur.UpdatedAt = value.Time
			}
		case userrole.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ur.DeletedAt = value.Time
			}
		case userrole.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ur.UserID = value.Int64
			}
		case userrole.FieldRoleID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field role_id", values[i])
			} else if value.Valid {
				ur.RoleID = value.Int64
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the UserRole entity.
func (ur *UserRole) QueryUser() *UserQuery {
	return (&UserRoleClient{config: ur.config}).QueryUser(ur)
}

// QueryRole queries the "role" edge of the UserRole entity.
func (ur *UserRole) QueryRole() *RoleQuery {
	return (&UserRoleClient{config: ur.config}).QueryRole(ur)
}

// Update returns a builder for updating this UserRole.
// Note that you need to call UserRole.Unwrap() before calling this method if this UserRole
// was returned from a transaction, and the transaction was committed or rolled back.
func (ur *UserRole) Update() *UserRoleUpdateOne {
	return (&UserRoleClient{config: ur.config}).UpdateOne(ur)
}

// Unwrap unwraps the UserRole entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ur *UserRole) Unwrap() *UserRole {
	_tx, ok := ur.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserRole is not a transactional entity")
	}
	ur.config.driver = _tx.drv
	return ur
}

// String implements the fmt.Stringer.
func (ur *UserRole) String() string {
	var builder strings.Builder
	builder.WriteString("UserRole(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ur.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", ur.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", ur.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ur.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ur.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(ur.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ur.UserID))
	builder.WriteString(", ")
	builder.WriteString("role_id=")
	builder.WriteString(fmt.Sprintf("%v", ur.RoleID))
	builder.WriteByte(')')
	return builder.String()
}

// UserRoles is a parsable slice of UserRole.
type UserRoles []*UserRole

func (ur UserRoles) config(cfg config) {
	for _i := range ur {
		ur[_i].config = cfg
	}
}
