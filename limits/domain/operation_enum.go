// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package domain

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

const (
	// OperationStatusNew is a OperationStatus of type New.
	OperationStatusNew OperationStatus = iota
	// OperationStatusPending is a OperationStatus of type Pending.
	OperationStatusPending
	// OperationStatusCommitted is a OperationStatus of type Committed.
	OperationStatusCommitted
	// OperationStatusRollback is a OperationStatus of type Rollback.
	OperationStatusRollback
)

var ErrInvalidOperationStatus = errors.New("not a valid OperationStatus")

const _OperationStatusName = "newpendingcommittedrollback"

var _OperationStatusMap = map[OperationStatus]string{
	OperationStatusNew:       _OperationStatusName[0:3],
	OperationStatusPending:   _OperationStatusName[3:10],
	OperationStatusCommitted: _OperationStatusName[10:19],
	OperationStatusRollback:  _OperationStatusName[19:27],
}

// String implements the Stringer interface.
func (x OperationStatus) String() string {
	if str, ok := _OperationStatusMap[x]; ok {
		return str
	}
	return fmt.Sprintf("OperationStatus(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x OperationStatus) IsValid() bool {
	_, ok := _OperationStatusMap[x]
	return ok
}

var _OperationStatusValue = map[string]OperationStatus{
	_OperationStatusName[0:3]:   OperationStatusNew,
	_OperationStatusName[3:10]:  OperationStatusPending,
	_OperationStatusName[10:19]: OperationStatusCommitted,
	_OperationStatusName[19:27]: OperationStatusRollback,
}

// ParseOperationStatus attempts to convert a string to a OperationStatus.
func ParseOperationStatus(name string) (OperationStatus, error) {
	if x, ok := _OperationStatusValue[name]; ok {
		return x, nil
	}
	return OperationStatus(0), fmt.Errorf("%s is %w", name, ErrInvalidOperationStatus)
}

// MarshalText implements the text marshaller method.
func (x OperationStatus) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *OperationStatus) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseOperationStatus(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

var errOperationStatusNilPtr = errors.New("value pointer is nil") // one per type for package clashes

// Scan implements the Scanner interface.
func (x *OperationStatus) Scan(value interface{}) (err error) {
	if value == nil {
		*x = OperationStatus(0)
		return
	}

	// A wider range of scannable types.
	// driver.Value values at the top of the list for expediency
	switch v := value.(type) {
	case int64:
		*x = OperationStatus(v)
	case string:
		*x, err = ParseOperationStatus(v)
	case []byte:
		*x, err = ParseOperationStatus(string(v))
	case OperationStatus:
		*x = v
	case int:
		*x = OperationStatus(v)
	case *OperationStatus:
		if v == nil {
			return errOperationStatusNilPtr
		}
		*x = *v
	case uint:
		*x = OperationStatus(v)
	case uint64:
		*x = OperationStatus(v)
	case *int:
		if v == nil {
			return errOperationStatusNilPtr
		}
		*x = OperationStatus(*v)
	case *int64:
		if v == nil {
			return errOperationStatusNilPtr
		}
		*x = OperationStatus(*v)
	case float64: // json marshals everything as a float64 if it's a number
		*x = OperationStatus(v)
	case *float64: // json marshals everything as a float64 if it's a number
		if v == nil {
			return errOperationStatusNilPtr
		}
		*x = OperationStatus(*v)
	case *uint:
		if v == nil {
			return errOperationStatusNilPtr
		}
		*x = OperationStatus(*v)
	case *uint64:
		if v == nil {
			return errOperationStatusNilPtr
		}
		*x = OperationStatus(*v)
	case *string:
		if v == nil {
			return errOperationStatusNilPtr
		}
		*x, err = ParseOperationStatus(*v)
	}

	return
}

// Value implements the driver Valuer interface.
func (x OperationStatus) Value() (driver.Value, error) {
	return x.String(), nil
}
