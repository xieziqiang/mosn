// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/type/matcher/struct.proto

package envoy_type_matcher

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
var _struct_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on StructMatcher with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *StructMatcher) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetPath()) < 1 {
		return StructMatcherValidationError{
			field:  "Path",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetPath() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StructMatcherValidationError{
					field:  fmt.Sprintf("Path[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.GetValue() == nil {
		return StructMatcherValidationError{
			field:  "Value",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StructMatcherValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// StructMatcherValidationError is the validation error returned by
// StructMatcher.Validate if the designated constraints aren't met.
type StructMatcherValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StructMatcherValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StructMatcherValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StructMatcherValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StructMatcherValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StructMatcherValidationError) ErrorName() string { return "StructMatcherValidationError" }

// Error satisfies the builtin error interface
func (e StructMatcherValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStructMatcher.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StructMatcherValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StructMatcherValidationError{}

// Validate checks the field values on StructMatcher_PathSegment with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *StructMatcher_PathSegment) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Segment.(type) {

	case *StructMatcher_PathSegment_Key:

		if utf8.RuneCountInString(m.GetKey()) < 1 {
			return StructMatcher_PathSegmentValidationError{
				field:  "Key",
				reason: "value length must be at least 1 runes",
			}
		}

	default:
		return StructMatcher_PathSegmentValidationError{
			field:  "Segment",
			reason: "value is required",
		}

	}

	return nil
}

// StructMatcher_PathSegmentValidationError is the validation error returned by
// StructMatcher_PathSegment.Validate if the designated constraints aren't met.
type StructMatcher_PathSegmentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StructMatcher_PathSegmentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StructMatcher_PathSegmentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StructMatcher_PathSegmentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StructMatcher_PathSegmentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StructMatcher_PathSegmentValidationError) ErrorName() string {
	return "StructMatcher_PathSegmentValidationError"
}

// Error satisfies the builtin error interface
func (e StructMatcher_PathSegmentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStructMatcher_PathSegment.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StructMatcher_PathSegmentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StructMatcher_PathSegmentValidationError{}
