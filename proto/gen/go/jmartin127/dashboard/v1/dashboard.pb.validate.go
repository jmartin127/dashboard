// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: jmartin127/dashboard/v1/dashboard.proto

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
var _dashboard_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on GetDashboardRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetDashboardRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetOriginAddress()) < 1 {
		return GetDashboardRequestValidationError{
			field:  "OriginAddress",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetDestinationAddress()) < 1 {
		return GetDashboardRequestValidationError{
			field:  "DestinationAddress",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// GetDashboardRequestValidationError is the validation error returned by
// GetDashboardRequest.Validate if the designated constraints aren't met.
type GetDashboardRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDashboardRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDashboardRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDashboardRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDashboardRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDashboardRequestValidationError) ErrorName() string {
	return "GetDashboardRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetDashboardRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDashboardRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDashboardRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDashboardRequestValidationError{}

// Validate checks the field values on GetDashboardReply with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetDashboardReply) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTraffic()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetDashboardReplyValidationError{
				field:  "Traffic",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetWeather()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetDashboardReplyValidationError{
				field:  "Weather",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetDashboardReplyValidationError is the validation error returned by
// GetDashboardReply.Validate if the designated constraints aren't met.
type GetDashboardReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDashboardReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDashboardReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDashboardReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDashboardReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDashboardReplyValidationError) ErrorName() string {
	return "GetDashboardReplyValidationError"
}

// Error satisfies the builtin error interface
func (e GetDashboardReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDashboardReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDashboardReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDashboardReplyValidationError{}