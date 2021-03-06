// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: jmartin127/user/v1/user.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _user_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on AddUserRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *AddUserRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetUser() == nil {
		return AddUserRequestValidationError{
			field:  "User",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddUserRequestValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// AddUserRequestValidationError is the validation error returned by
// AddUserRequest.Validate if the designated constraints aren't met.
type AddUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddUserRequestValidationError) ErrorName() string { return "AddUserRequestValidationError" }

// Error satisfies the builtin error interface
func (e AddUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddUserRequestValidationError{}

// Validate checks the field values on ModifyUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ModifyUserRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetUser() == nil {
		return ModifyUserRequestValidationError{
			field:  "User",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ModifyUserRequestValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ModifyUserRequestValidationError is the validation error returned by
// ModifyUserRequest.Validate if the designated constraints aren't met.
type ModifyUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ModifyUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ModifyUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ModifyUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ModifyUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ModifyUserRequestValidationError) ErrorName() string {
	return "ModifyUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ModifyUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sModifyUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ModifyUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ModifyUserRequestValidationError{}

// Validate checks the field values on GetUserRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetUserRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetUsername()) < 1 {
		return GetUserRequestValidationError{
			field:  "Username",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// GetUserRequestValidationError is the validation error returned by
// GetUserRequest.Validate if the designated constraints aren't met.
type GetUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserRequestValidationError) ErrorName() string { return "GetUserRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserRequestValidationError{}

// Validate checks the field values on GetUserReply with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetUserReply) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetUserReplyValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetUserReplyValidationError is the validation error returned by
// GetUserReply.Validate if the designated constraints aren't met.
type GetUserReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserReplyValidationError) ErrorName() string { return "GetUserReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetUserReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserReplyValidationError{}

// Validate checks the field values on RemoveUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *RemoveUserRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetUsername()) < 1 {
		return RemoveUserRequestValidationError{
			field:  "Username",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// RemoveUserRequestValidationError is the validation error returned by
// RemoveUserRequest.Validate if the designated constraints aren't met.
type RemoveUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveUserRequestValidationError) ErrorName() string {
	return "RemoveUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveUserRequestValidationError{}

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *User) Validate() error {
	if m == nil {
		return nil
	}

	if !_User_Username_Pattern.MatchString(m.GetUsername()) {
		return UserValidationError{
			field:  "Username",
			reason: "value does not match regex pattern \"^[a-z0-9_-]{3,16}$\"",
		}
	}

	if err := m._validateEmail(m.GetEmail()); err != nil {
		return UserValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
	}

	if utf8.RuneCountInString(m.GetFirstName()) < 1 {
		return UserValidationError{
			field:  "FirstName",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetLastName()) < 1 {
		return UserValidationError{
			field:  "LastName",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

func (m *User) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *User) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}

var _User_Username_Pattern = regexp.MustCompile("^[a-z0-9_-]{3,16}$")
