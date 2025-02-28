// Code generated by ogen, DO NOT EDIT.

package api

import (
	"math/bits"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/validate"
)

// Encode encodes CreateUserBadRequest as json.
func (s *CreateUserBadRequest) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes CreateUserBadRequest from json.
func (s *CreateUserBadRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CreateUserBadRequest to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = CreateUserBadRequest(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *CreateUserBadRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CreateUserBadRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes CreateUserInternalServerError as json.
func (s *CreateUserInternalServerError) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes CreateUserInternalServerError from json.
func (s *CreateUserInternalServerError) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CreateUserInternalServerError to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = CreateUserInternalServerError(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *CreateUserInternalServerError) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CreateUserInternalServerError) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Error) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Error) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("message")
		e.Str(s.Message)
	}
}

var jsonFieldsNameOfError = [1]string{
	0: "message",
}

// Decode decodes Error from json.
func (s *Error) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Error to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "message":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Message = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Error")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfError) {
					name = jsonFieldsNameOfError[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Error) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Error) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes GetAllUsersBadRequest as json.
func (s *GetAllUsersBadRequest) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes GetAllUsersBadRequest from json.
func (s *GetAllUsersBadRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GetAllUsersBadRequest to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = GetAllUsersBadRequest(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GetAllUsersBadRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GetAllUsersBadRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes GetAllUsersInternalServerError as json.
func (s *GetAllUsersInternalServerError) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes GetAllUsersInternalServerError from json.
func (s *GetAllUsersInternalServerError) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GetAllUsersInternalServerError to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = GetAllUsersInternalServerError(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GetAllUsersInternalServerError) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GetAllUsersInternalServerError) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes GetAllUsersOKApplicationJSON as json.
func (s GetAllUsersOKApplicationJSON) Encode(e *jx.Encoder) {
	unwrapped := []User(s)

	e.ArrStart()
	for _, elem := range unwrapped {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes GetAllUsersOKApplicationJSON from json.
func (s *GetAllUsersOKApplicationJSON) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GetAllUsersOKApplicationJSON to nil")
	}
	var unwrapped []User
	if err := func() error {
		unwrapped = make([]User, 0)
		if err := d.Arr(func(d *jx.Decoder) error {
			var elem User
			if err := elem.Decode(d); err != nil {
				return err
			}
			unwrapped = append(unwrapped, elem)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = GetAllUsersOKApplicationJSON(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s GetAllUsersOKApplicationJSON) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GetAllUsersOKApplicationJSON) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes GetUserByIdBadRequest as json.
func (s *GetUserByIdBadRequest) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes GetUserByIdBadRequest from json.
func (s *GetUserByIdBadRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GetUserByIdBadRequest to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = GetUserByIdBadRequest(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GetUserByIdBadRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GetUserByIdBadRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes GetUserByIdInternalServerError as json.
func (s *GetUserByIdInternalServerError) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes GetUserByIdInternalServerError from json.
func (s *GetUserByIdInternalServerError) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GetUserByIdInternalServerError to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = GetUserByIdInternalServerError(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GetUserByIdInternalServerError) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GetUserByIdInternalServerError) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *User) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *User) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("id")
		e.Int(s.ID)
	}
	{
		e.FieldStart("name")
		e.Str(s.Name)
	}
	{
		e.FieldStart("email")
		e.Str(s.Email)
	}
}

var jsonFieldsNameOfUser = [3]string{
	0: "id",
	1: "name",
	2: "email",
}

// Decode decodes User from json.
func (s *User) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode User to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int()
				s.ID = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "name":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Name = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		case "email":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Str()
				s.Email = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"email\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode User")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfUser) {
					name = jsonFieldsNameOfUser[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *User) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *User) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *UserCategory) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *UserCategory) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("id")
		e.Int(s.ID)
	}
	{
		e.FieldStart("name")
		e.Str(s.Name)
	}
	{
		e.FieldStart("category")
		e.Str(s.Category)
	}
}

var jsonFieldsNameOfUserCategory = [3]string{
	0: "id",
	1: "name",
	2: "category",
}

// Decode decodes UserCategory from json.
func (s *UserCategory) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UserCategory to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int()
				s.ID = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "name":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Name = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		case "category":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Str()
				s.Category = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"category\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UserCategory")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfUserCategory) {
					name = jsonFieldsNameOfUserCategory[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UserCategory) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UserCategory) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
