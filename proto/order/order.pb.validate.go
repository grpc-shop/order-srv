// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: order.proto

package order

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
)

// Validate checks the field values on OrderInfo with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *OrderInfo) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for OrderId

	// no validation rules for UserName

	// no validation rules for UserPhone

	// no validation rules for UserAddress

	// no validation rules for TotalAmount

	// no validation rules for PaidAt

	// no validation rules for PaymentNo

	// no validation rules for RefundStatus

	// no validation rules for RefundNo

	// no validation rules for OrderStatus

	// no validation rules for Extra

	// no validation rules for Remark

	return nil
}

// OrderInfoValidationError is the validation error returned by
// OrderInfo.Validate if the designated constraints aren't met.
type OrderInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderInfoValidationError) ErrorName() string { return "OrderInfoValidationError" }

// Error satisfies the builtin error interface
func (e OrderInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderInfoValidationError{}

// Validate checks the field values on OrderItemInfo with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *OrderItemInfo) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for OrderId

	// no validation rules for ProductId

	// no validation rules for ProductSkuId

	// no validation rules for Amount

	// no validation rules for Price

	// no validation rules for Title

	// no validation rules for Image

	return nil
}

// OrderItemInfoValidationError is the validation error returned by
// OrderItemInfo.Validate if the designated constraints aren't met.
type OrderItemInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderItemInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderItemInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderItemInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderItemInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderItemInfoValidationError) ErrorName() string { return "OrderItemInfoValidationError" }

// Error satisfies the builtin error interface
func (e OrderItemInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderItemInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderItemInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderItemInfoValidationError{}

// Validate checks the field values on CreateOrderReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CreateOrderReq) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for UserName

	// no validation rules for UserPhone

	// no validation rules for UserAddress

	// no validation rules for Extra

	// no validation rules for Remark

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateOrderReqValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CreateOrderReqValidationError is the validation error returned by
// CreateOrderReq.Validate if the designated constraints aren't met.
type CreateOrderReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrderReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrderReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrderReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrderReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrderReqValidationError) ErrorName() string { return "CreateOrderReqValidationError" }

// Error satisfies the builtin error interface
func (e CreateOrderReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrderReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrderReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrderReqValidationError{}

// Validate checks the field values on CreateOrderReply with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CreateOrderReply) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	// no validation rules for Msg

	if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateOrderReplyValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateOrderReplyValidationError is the validation error returned by
// CreateOrderReply.Validate if the designated constraints aren't met.
type CreateOrderReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrderReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrderReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrderReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrderReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrderReplyValidationError) ErrorName() string { return "CreateOrderReplyValidationError" }

// Error satisfies the builtin error interface
func (e CreateOrderReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrderReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrderReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrderReplyValidationError{}

// Validate checks the field values on OrderPaidReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *OrderPaidReq) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for OrderId

	// no validation rules for PaymentNo

	// no validation rules for PaidAt

	// no validation rules for PaidAmount

	return nil
}

// OrderPaidReqValidationError is the validation error returned by
// OrderPaidReq.Validate if the designated constraints aren't met.
type OrderPaidReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderPaidReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderPaidReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderPaidReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderPaidReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderPaidReqValidationError) ErrorName() string { return "OrderPaidReqValidationError" }

// Error satisfies the builtin error interface
func (e OrderPaidReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderPaidReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderPaidReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderPaidReqValidationError{}

// Validate checks the field values on OrderPaidReply with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *OrderPaidReply) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	// no validation rules for Msg

	if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrderPaidReplyValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// OrderPaidReplyValidationError is the validation error returned by
// OrderPaidReply.Validate if the designated constraints aren't met.
type OrderPaidReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderPaidReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderPaidReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderPaidReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderPaidReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderPaidReplyValidationError) ErrorName() string { return "OrderPaidReplyValidationError" }

// Error satisfies the builtin error interface
func (e OrderPaidReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderPaidReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderPaidReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderPaidReplyValidationError{}

// Validate checks the field values on CreateOrderReqItem with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateOrderReqItem) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ProductId

	// no validation rules for ProductSkuId

	// no validation rules for Amount

	// no validation rules for Price

	// no validation rules for Title

	// no validation rules for Image

	return nil
}

// CreateOrderReqItemValidationError is the validation error returned by
// CreateOrderReqItem.Validate if the designated constraints aren't met.
type CreateOrderReqItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrderReqItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrderReqItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrderReqItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrderReqItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrderReqItemValidationError) ErrorName() string {
	return "CreateOrderReqItemValidationError"
}

// Error satisfies the builtin error interface
func (e CreateOrderReqItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrderReqItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrderReqItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrderReqItemValidationError{}

// Validate checks the field values on CreateOrderReplyOrder with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateOrderReplyOrder) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// CreateOrderReplyOrderValidationError is the validation error returned by
// CreateOrderReplyOrder.Validate if the designated constraints aren't met.
type CreateOrderReplyOrderValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrderReplyOrderValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrderReplyOrderValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrderReplyOrderValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrderReplyOrderValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrderReplyOrderValidationError) ErrorName() string {
	return "CreateOrderReplyOrderValidationError"
}

// Error satisfies the builtin error interface
func (e CreateOrderReplyOrderValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrderReplyOrder.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrderReplyOrderValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrderReplyOrderValidationError{}

// Validate checks the field values on OrderPaidReplyOrder with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *OrderPaidReplyOrder) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// OrderPaidReplyOrderValidationError is the validation error returned by
// OrderPaidReplyOrder.Validate if the designated constraints aren't met.
type OrderPaidReplyOrderValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderPaidReplyOrderValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderPaidReplyOrderValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderPaidReplyOrderValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderPaidReplyOrderValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderPaidReplyOrderValidationError) ErrorName() string {
	return "OrderPaidReplyOrderValidationError"
}

// Error satisfies the builtin error interface
func (e OrderPaidReplyOrderValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderPaidReplyOrder.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderPaidReplyOrderValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderPaidReplyOrderValidationError{}
