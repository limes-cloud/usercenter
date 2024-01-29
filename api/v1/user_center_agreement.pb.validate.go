// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user_center_agreement.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Agreement with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Agreement) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Agreement with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AgreementMultiError, or nil
// if none found.
func (m *Agreement) ValidateAll() error {
	return m.validate(true)
}

func (m *Agreement) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Status

	// no validation rules for Content

	// no validation rules for Description

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if len(errors) > 0 {
		return AgreementMultiError(errors)
	}

	return nil
}

// AgreementMultiError is an error wrapping multiple validation errors returned
// by Agreement.ValidateAll() if the designated constraints aren't met.
type AgreementMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AgreementMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AgreementMultiError) AllErrors() []error { return m }

// AgreementValidationError is the validation error returned by
// Agreement.Validate if the designated constraints aren't met.
type AgreementValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AgreementValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AgreementValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AgreementValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AgreementValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AgreementValidationError) ErrorName() string { return "AgreementValidationError" }

// Error satisfies the builtin error interface
func (e AgreementValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAgreement.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AgreementValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AgreementValidationError{}

// Validate checks the field values on PageAgreementRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PageAgreementRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PageAgreementRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PageAgreementRequestMultiError, or nil if none found.
func (m *PageAgreementRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PageAgreementRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPage() <= 0 {
		err := PageAgreementRequestValidationError{
			field:  "Page",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetPageSize(); val <= 0 || val > 50 {
		err := PageAgreementRequestValidationError{
			field:  "PageSize",
			reason: "value must be inside range (0, 50]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.Name != nil {
		// no validation rules for Name
	}

	if len(errors) > 0 {
		return PageAgreementRequestMultiError(errors)
	}

	return nil
}

// PageAgreementRequestMultiError is an error wrapping multiple validation
// errors returned by PageAgreementRequest.ValidateAll() if the designated
// constraints aren't met.
type PageAgreementRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageAgreementRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageAgreementRequestMultiError) AllErrors() []error { return m }

// PageAgreementRequestValidationError is the validation error returned by
// PageAgreementRequest.Validate if the designated constraints aren't met.
type PageAgreementRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageAgreementRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageAgreementRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageAgreementRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageAgreementRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageAgreementRequestValidationError) ErrorName() string {
	return "PageAgreementRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PageAgreementRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPageAgreementRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageAgreementRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageAgreementRequestValidationError{}

// Validate checks the field values on GetAgreementRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetAgreementRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAgreementRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetAgreementRequestMultiError, or nil if none found.
func (m *GetAgreementRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAgreementRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := GetAgreementRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetAgreementRequestMultiError(errors)
	}

	return nil
}

// GetAgreementRequestMultiError is an error wrapping multiple validation
// errors returned by GetAgreementRequest.ValidateAll() if the designated
// constraints aren't met.
type GetAgreementRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAgreementRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAgreementRequestMultiError) AllErrors() []error { return m }

// GetAgreementRequestValidationError is the validation error returned by
// GetAgreementRequest.Validate if the designated constraints aren't met.
type GetAgreementRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAgreementRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAgreementRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAgreementRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAgreementRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAgreementRequestValidationError) ErrorName() string {
	return "GetAgreementRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetAgreementRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAgreementRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAgreementRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAgreementRequestValidationError{}

// Validate checks the field values on PageAgreementReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PageAgreementReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PageAgreementReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PageAgreementReplyMultiError, or nil if none found.
func (m *PageAgreementReply) ValidateAll() error {
	return m.validate(true)
}

func (m *PageAgreementReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PageAgreementReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PageAgreementReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PageAgreementReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return PageAgreementReplyMultiError(errors)
	}

	return nil
}

// PageAgreementReplyMultiError is an error wrapping multiple validation errors
// returned by PageAgreementReply.ValidateAll() if the designated constraints
// aren't met.
type PageAgreementReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageAgreementReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageAgreementReplyMultiError) AllErrors() []error { return m }

// PageAgreementReplyValidationError is the validation error returned by
// PageAgreementReply.Validate if the designated constraints aren't met.
type PageAgreementReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageAgreementReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageAgreementReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageAgreementReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageAgreementReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageAgreementReplyValidationError) ErrorName() string {
	return "PageAgreementReplyValidationError"
}

// Error satisfies the builtin error interface
func (e PageAgreementReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPageAgreementReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageAgreementReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageAgreementReplyValidationError{}

// Validate checks the field values on AddAgreementRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AddAgreementRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddAgreementRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddAgreementRequestMultiError, or nil if none found.
func (m *AddAgreementRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddAgreementRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := AddAgreementRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Status

	if utf8.RuneCountInString(m.GetContent()) < 1 {
		err := AddAgreementRequestValidationError{
			field:  "Content",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDescription()) < 1 {
		err := AddAgreementRequestValidationError{
			field:  "Description",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return AddAgreementRequestMultiError(errors)
	}

	return nil
}

// AddAgreementRequestMultiError is an error wrapping multiple validation
// errors returned by AddAgreementRequest.ValidateAll() if the designated
// constraints aren't met.
type AddAgreementRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddAgreementRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddAgreementRequestMultiError) AllErrors() []error { return m }

// AddAgreementRequestValidationError is the validation error returned by
// AddAgreementRequest.Validate if the designated constraints aren't met.
type AddAgreementRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddAgreementRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddAgreementRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddAgreementRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddAgreementRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddAgreementRequestValidationError) ErrorName() string {
	return "AddAgreementRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AddAgreementRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddAgreementRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddAgreementRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddAgreementRequestValidationError{}

// Validate checks the field values on AddAgreementReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AddAgreementReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddAgreementReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddAgreementReplyMultiError, or nil if none found.
func (m *AddAgreementReply) ValidateAll() error {
	return m.validate(true)
}

func (m *AddAgreementReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return AddAgreementReplyMultiError(errors)
	}

	return nil
}

// AddAgreementReplyMultiError is an error wrapping multiple validation errors
// returned by AddAgreementReply.ValidateAll() if the designated constraints
// aren't met.
type AddAgreementReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddAgreementReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddAgreementReplyMultiError) AllErrors() []error { return m }

// AddAgreementReplyValidationError is the validation error returned by
// AddAgreementReply.Validate if the designated constraints aren't met.
type AddAgreementReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddAgreementReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddAgreementReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddAgreementReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddAgreementReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddAgreementReplyValidationError) ErrorName() string {
	return "AddAgreementReplyValidationError"
}

// Error satisfies the builtin error interface
func (e AddAgreementReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddAgreementReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddAgreementReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddAgreementReplyValidationError{}

// Validate checks the field values on UpdateAgreementRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateAgreementRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateAgreementRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateAgreementRequestMultiError, or nil if none found.
func (m *UpdateAgreementRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateAgreementRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := UpdateAgreementRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := UpdateAgreementRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetContent()) < 1 {
		err := UpdateAgreementRequestValidationError{
			field:  "Content",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDescription()) < 1 {
		err := UpdateAgreementRequestValidationError{
			field:  "Description",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.Status != nil {
		// no validation rules for Status
	}

	if len(errors) > 0 {
		return UpdateAgreementRequestMultiError(errors)
	}

	return nil
}

// UpdateAgreementRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateAgreementRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateAgreementRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateAgreementRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateAgreementRequestMultiError) AllErrors() []error { return m }

// UpdateAgreementRequestValidationError is the validation error returned by
// UpdateAgreementRequest.Validate if the designated constraints aren't met.
type UpdateAgreementRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateAgreementRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateAgreementRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateAgreementRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateAgreementRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateAgreementRequestValidationError) ErrorName() string {
	return "UpdateAgreementRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateAgreementRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateAgreementRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateAgreementRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateAgreementRequestValidationError{}

// Validate checks the field values on DeleteAgreementRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteAgreementRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteAgreementRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteAgreementRequestMultiError, or nil if none found.
func (m *DeleteAgreementRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteAgreementRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := DeleteAgreementRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteAgreementRequestMultiError(errors)
	}

	return nil
}

// DeleteAgreementRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteAgreementRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteAgreementRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteAgreementRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteAgreementRequestMultiError) AllErrors() []error { return m }

// DeleteAgreementRequestValidationError is the validation error returned by
// DeleteAgreementRequest.Validate if the designated constraints aren't met.
type DeleteAgreementRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteAgreementRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteAgreementRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteAgreementRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteAgreementRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteAgreementRequestValidationError) ErrorName() string {
	return "DeleteAgreementRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteAgreementRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteAgreementRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteAgreementRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteAgreementRequestValidationError{}
