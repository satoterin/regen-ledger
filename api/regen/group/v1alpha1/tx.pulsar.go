package groupv1alpha1

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var _ protoreflect.List = (*_MsgCreateGroup_2_list)(nil)

type _MsgCreateGroup_2_list struct {
	list *[]*Member
}

func (x *_MsgCreateGroup_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgCreateGroup_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_MsgCreateGroup_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Member)
	(*x.list)[i] = concreteValue
}

func (x *_MsgCreateGroup_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Member)
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgCreateGroup_2_list) AppendMutable() protoreflect.Value {
	v := new(Member)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgCreateGroup_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_MsgCreateGroup_2_list) NewElement() protoreflect.Value {
	v := new(Member)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgCreateGroup_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgCreateGroup          protoreflect.MessageDescriptor
	fd_MsgCreateGroup_admin    protoreflect.FieldDescriptor
	fd_MsgCreateGroup_members  protoreflect.FieldDescriptor
	fd_MsgCreateGroup_metadata protoreflect.FieldDescriptor
)

func init() {
	file_regen_group_v1alpha1_tx_proto_init()
	md_MsgCreateGroup = File_regen_group_v1alpha1_tx_proto.Messages().ByName("MsgCreateGroup")
	fd_MsgCreateGroup_admin = md_MsgCreateGroup.Fields().ByName("admin")
	fd_MsgCreateGroup_members = md_MsgCreateGroup.Fields().ByName("members")
	fd_MsgCreateGroup_metadata = md_MsgCreateGroup.Fields().ByName("metadata")
}

var _ protoreflect.Message = (*fastReflection_MsgCreateGroup)(nil)

type fastReflection_MsgCreateGroup MsgCreateGroup

func (x *MsgCreateGroup) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCreateGroup)(x)
}

func (x *MsgCreateGroup) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCreateGroup_messageType fastReflection_MsgCreateGroup_messageType
var _ protoreflect.MessageType = fastReflection_MsgCreateGroup_messageType{}

type fastReflection_MsgCreateGroup_messageType struct{}

func (x fastReflection_MsgCreateGroup_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCreateGroup)(nil)
}
func (x fastReflection_MsgCreateGroup_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCreateGroup)
}
func (x fastReflection_MsgCreateGroup_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateGroup
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCreateGroup) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateGroup
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCreateGroup) Type() protoreflect.MessageType {
	return _fastReflection_MsgCreateGroup_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCreateGroup) New() protoreflect.Message {
	return new(fastReflection_MsgCreateGroup)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCreateGroup) Interface() protoreflect.ProtoMessage {
	return (*MsgCreateGroup)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCreateGroup) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Admin != "" {
		value := protoreflect.ValueOfString(x.Admin)
		if !f(fd_MsgCreateGroup_admin, value) {
			return
		}
	}
	if len(x.Members) != 0 {
		value := protoreflect.ValueOfList(&_MsgCreateGroup_2_list{list: &x.Members})
		if !f(fd_MsgCreateGroup_members, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_MsgCreateGroup_metadata, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgCreateGroup) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgCreateGroup.admin":
		return x.Admin != ""
	case "regen.group.v1alpha1.MsgCreateGroup.members":
		return len(x.Members) != 0
	case "regen.group.v1alpha1.MsgCreateGroup.metadata":
		return len(x.Metadata) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateGroup"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateGroup does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroup) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgCreateGroup.admin":
		x.Admin = ""
	case "regen.group.v1alpha1.MsgCreateGroup.members":
		x.Members = nil
	case "regen.group.v1alpha1.MsgCreateGroup.metadata":
		x.Metadata = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateGroup"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateGroup does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCreateGroup) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.group.v1alpha1.MsgCreateGroup.admin":
		value := x.Admin
		return protoreflect.ValueOfString(value)
	case "regen.group.v1alpha1.MsgCreateGroup.members":
		if len(x.Members) == 0 {
			return protoreflect.ValueOfList(&_MsgCreateGroup_2_list{})
		}
		listValue := &_MsgCreateGroup_2_list{list: &x.Members}
		return protoreflect.ValueOfList(listValue)
	case "regen.group.v1alpha1.MsgCreateGroup.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateGroup"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateGroup does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroup) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgCreateGroup.admin":
		x.Admin = value.Interface().(string)
	case "regen.group.v1alpha1.MsgCreateGroup.members":
		lv := value.List()
		clv := lv.(*_MsgCreateGroup_2_list)
		x.Members = *clv.list
	case "regen.group.v1alpha1.MsgCreateGroup.metadata":
		x.Metadata = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateGroup"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateGroup does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroup) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgCreateGroup.members":
		if x.Members == nil {
			x.Members = []*Member{}
		}
		value := &_MsgCreateGroup_2_list{list: &x.Members}
		return protoreflect.ValueOfList(value)
	case "regen.group.v1alpha1.MsgCreateGroup.admin":
		panic(fmt.Errorf("field admin of message regen.group.v1alpha1.MsgCreateGroup is not mutable"))
	case "regen.group.v1alpha1.MsgCreateGroup.metadata":
		panic(fmt.Errorf("field metadata of message regen.group.v1alpha1.MsgCreateGroup is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateGroup"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateGroup does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCreateGroup) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgCreateGroup.admin":
		return protoreflect.ValueOfString("")
	case "regen.group.v1alpha1.MsgCreateGroup.members":
		list := []*Member{}
		return protoreflect.ValueOfList(&_MsgCreateGroup_2_list{list: &list})
	case "regen.group.v1alpha1.MsgCreateGroup.metadata":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateGroup"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateGroup does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCreateGroup) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.MsgCreateGroup", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCreateGroup) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroup) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgCreateGroup) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCreateGroup) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCreateGroup)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Admin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Members) > 0 {
			for _, e := range x.Members {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		l = len(x.Metadata)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgCreateGroup)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Members) > 0 {
			for iNdEx := len(x.Members) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Members[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x12
			}
		}
		if len(x.Admin) > 0 {
			i -= len(x.Admin)
			copy(dAtA[i:], x.Admin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Admin)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgCreateGroup)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateGroup: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateGroup: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Admin = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Members = append(x.Members, &Member{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Members[len(x.Members)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Metadata = append(x.Metadata[:0], dAtA[iNdEx:postIndex]...)
				if x.Metadata == nil {
					x.Metadata = []byte{}
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgCreateGroupResponse          protoreflect.MessageDescriptor
	fd_MsgCreateGroupResponse_group_id protoreflect.FieldDescriptor
)

func init() {
	file_regen_group_v1alpha1_tx_proto_init()
	md_MsgCreateGroupResponse = File_regen_group_v1alpha1_tx_proto.Messages().ByName("MsgCreateGroupResponse")
	fd_MsgCreateGroupResponse_group_id = md_MsgCreateGroupResponse.Fields().ByName("group_id")
}

var _ protoreflect.Message = (*fastReflection_MsgCreateGroupResponse)(nil)

type fastReflection_MsgCreateGroupResponse MsgCreateGroupResponse

func (x *MsgCreateGroupResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCreateGroupResponse)(x)
}

func (x *MsgCreateGroupResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCreateGroupResponse_messageType fastReflection_MsgCreateGroupResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgCreateGroupResponse_messageType{}

type fastReflection_MsgCreateGroupResponse_messageType struct{}

func (x fastReflection_MsgCreateGroupResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCreateGroupResponse)(nil)
}
func (x fastReflection_MsgCreateGroupResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCreateGroupResponse)
}
func (x fastReflection_MsgCreateGroupResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateGroupResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCreateGroupResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateGroupResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCreateGroupResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgCreateGroupResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCreateGroupResponse) New() protoreflect.Message {
	return new(fastReflection_MsgCreateGroupResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCreateGroupResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgCreateGroupResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCreateGroupResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.GroupId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.GroupId)
		if !f(fd_MsgCreateGroupResponse_group_id, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgCreateGroupResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgCreateGroupResponse.group_id":
		return x.GroupId != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateGroupResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateGroupResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroupResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgCreateGroupResponse.group_id":
		x.GroupId = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateGroupResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateGroupResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCreateGroupResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.group.v1alpha1.MsgCreateGroupResponse.group_id":
		value := x.GroupId
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateGroupResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateGroupResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroupResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgCreateGroupResponse.group_id":
		x.GroupId = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateGroupResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateGroupResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroupResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgCreateGroupResponse.group_id":
		panic(fmt.Errorf("field group_id of message regen.group.v1alpha1.MsgCreateGroupResponse is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateGroupResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateGroupResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCreateGroupResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgCreateGroupResponse.group_id":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateGroupResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateGroupResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCreateGroupResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.MsgCreateGroupResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCreateGroupResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroupResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgCreateGroupResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCreateGroupResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCreateGroupResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.GroupId != 0 {
			n += 1 + runtime.Sov(uint64(x.GroupId))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgCreateGroupResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.GroupId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.GroupId))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgCreateGroupResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateGroupResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateGroupResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field GroupId", wireType)
				}
				x.GroupId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.GroupId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_MsgUpdateGroupMembers_3_list)(nil)

type _MsgUpdateGroupMembers_3_list struct {
	list *[]*Member
}

func (x *_MsgUpdateGroupMembers_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgUpdateGroupMembers_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_MsgUpdateGroupMembers_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Member)
	(*x.list)[i] = concreteValue
}

func (x *_MsgUpdateGroupMembers_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Member)
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgUpdateGroupMembers_3_list) AppendMutable() protoreflect.Value {
	v := new(Member)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgUpdateGroupMembers_3_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_MsgUpdateGroupMembers_3_list) NewElement() protoreflect.Value {
	v := new(Member)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgUpdateGroupMembers_3_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgUpdateGroupMembers                protoreflect.MessageDescriptor
	fd_MsgUpdateGroupMembers_admin          protoreflect.FieldDescriptor
	fd_MsgUpdateGroupMembers_group_id       protoreflect.FieldDescriptor
	fd_MsgUpdateGroupMembers_member_updates protoreflect.FieldDescriptor
)

func init() {
	file_regen_group_v1alpha1_tx_proto_init()
	md_MsgUpdateGroupMembers = File_regen_group_v1alpha1_tx_proto.Messages().ByName("MsgUpdateGroupMembers")
	fd_MsgUpdateGroupMembers_admin = md_MsgUpdateGroupMembers.Fields().ByName("admin")
	fd_MsgUpdateGroupMembers_group_id = md_MsgUpdateGroupMembers.Fields().ByName("group_id")
	fd_MsgUpdateGroupMembers_member_updates = md_MsgUpdateGroupMembers.Fields().ByName("member_updates")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateGroupMembers)(nil)

type fastReflection_MsgUpdateGroupMembers MsgUpdateGroupMembers

func (x *MsgUpdateGroupMembers) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupMembers)(x)
}

func (x *MsgUpdateGroupMembers) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateGroupMembers_messageType fastReflection_MsgUpdateGroupMembers_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateGroupMembers_messageType{}

type fastReflection_MsgUpdateGroupMembers_messageType struct{}

func (x fastReflection_MsgUpdateGroupMembers_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupMembers)(nil)
}
func (x fastReflection_MsgUpdateGroupMembers_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupMembers)
}
func (x fastReflection_MsgUpdateGroupMembers_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupMembers
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateGroupMembers) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupMembers
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateGroupMembers) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateGroupMembers_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateGroupMembers) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupMembers)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateGroupMembers) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateGroupMembers)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateGroupMembers) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Admin != "" {
		value := protoreflect.ValueOfString(x.Admin)
		if !f(fd_MsgUpdateGroupMembers_admin, value) {
			return
		}
	}
	if x.GroupId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.GroupId)
		if !f(fd_MsgUpdateGroupMembers_group_id, value) {
			return
		}
	}
	if len(x.MemberUpdates) != 0 {
		value := protoreflect.ValueOfList(&_MsgUpdateGroupMembers_3_list{list: &x.MemberUpdates})
		if !f(fd_MsgUpdateGroupMembers_member_updates, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgUpdateGroupMembers) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupMembers.admin":
		return x.Admin != ""
	case "regen.group.v1alpha1.MsgUpdateGroupMembers.group_id":
		return x.GroupId != uint64(0)
	case "regen.group.v1alpha1.MsgUpdateGroupMembers.member_updates":
		return len(x.MemberUpdates) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupMembers"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupMembers does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMembers) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupMembers.admin":
		x.Admin = ""
	case "regen.group.v1alpha1.MsgUpdateGroupMembers.group_id":
		x.GroupId = uint64(0)
	case "regen.group.v1alpha1.MsgUpdateGroupMembers.member_updates":
		x.MemberUpdates = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupMembers"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupMembers does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateGroupMembers) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupMembers.admin":
		value := x.Admin
		return protoreflect.ValueOfString(value)
	case "regen.group.v1alpha1.MsgUpdateGroupMembers.group_id":
		value := x.GroupId
		return protoreflect.ValueOfUint64(value)
	case "regen.group.v1alpha1.MsgUpdateGroupMembers.member_updates":
		if len(x.MemberUpdates) == 0 {
			return protoreflect.ValueOfList(&_MsgUpdateGroupMembers_3_list{})
		}
		listValue := &_MsgUpdateGroupMembers_3_list{list: &x.MemberUpdates}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupMembers"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupMembers does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMembers) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupMembers.admin":
		x.Admin = value.Interface().(string)
	case "regen.group.v1alpha1.MsgUpdateGroupMembers.group_id":
		x.GroupId = value.Uint()
	case "regen.group.v1alpha1.MsgUpdateGroupMembers.member_updates":
		lv := value.List()
		clv := lv.(*_MsgUpdateGroupMembers_3_list)
		x.MemberUpdates = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupMembers"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupMembers does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMembers) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupMembers.member_updates":
		if x.MemberUpdates == nil {
			x.MemberUpdates = []*Member{}
		}
		value := &_MsgUpdateGroupMembers_3_list{list: &x.MemberUpdates}
		return protoreflect.ValueOfList(value)
	case "regen.group.v1alpha1.MsgUpdateGroupMembers.admin":
		panic(fmt.Errorf("field admin of message regen.group.v1alpha1.MsgUpdateGroupMembers is not mutable"))
	case "regen.group.v1alpha1.MsgUpdateGroupMembers.group_id":
		panic(fmt.Errorf("field group_id of message regen.group.v1alpha1.MsgUpdateGroupMembers is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupMembers"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupMembers does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateGroupMembers) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupMembers.admin":
		return protoreflect.ValueOfString("")
	case "regen.group.v1alpha1.MsgUpdateGroupMembers.group_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.group.v1alpha1.MsgUpdateGroupMembers.member_updates":
		list := []*Member{}
		return protoreflect.ValueOfList(&_MsgUpdateGroupMembers_3_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupMembers"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupMembers does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateGroupMembers) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.MsgUpdateGroupMembers", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateGroupMembers) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMembers) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgUpdateGroupMembers) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateGroupMembers) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateGroupMembers)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Admin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.GroupId != 0 {
			n += 1 + runtime.Sov(uint64(x.GroupId))
		}
		if len(x.MemberUpdates) > 0 {
			for _, e := range x.MemberUpdates {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupMembers)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.MemberUpdates) > 0 {
			for iNdEx := len(x.MemberUpdates) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.MemberUpdates[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x1a
			}
		}
		if x.GroupId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.GroupId))
			i--
			dAtA[i] = 0x10
		}
		if len(x.Admin) > 0 {
			i -= len(x.Admin)
			copy(dAtA[i:], x.Admin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Admin)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupMembers)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupMembers: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupMembers: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Admin = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field GroupId", wireType)
				}
				x.GroupId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.GroupId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MemberUpdates", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.MemberUpdates = append(x.MemberUpdates, &Member{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.MemberUpdates[len(x.MemberUpdates)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgUpdateGroupMembersResponse protoreflect.MessageDescriptor
)

func init() {
	file_regen_group_v1alpha1_tx_proto_init()
	md_MsgUpdateGroupMembersResponse = File_regen_group_v1alpha1_tx_proto.Messages().ByName("MsgUpdateGroupMembersResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateGroupMembersResponse)(nil)

type fastReflection_MsgUpdateGroupMembersResponse MsgUpdateGroupMembersResponse

func (x *MsgUpdateGroupMembersResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupMembersResponse)(x)
}

func (x *MsgUpdateGroupMembersResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateGroupMembersResponse_messageType fastReflection_MsgUpdateGroupMembersResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateGroupMembersResponse_messageType{}

type fastReflection_MsgUpdateGroupMembersResponse_messageType struct{}

func (x fastReflection_MsgUpdateGroupMembersResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupMembersResponse)(nil)
}
func (x fastReflection_MsgUpdateGroupMembersResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupMembersResponse)
}
func (x fastReflection_MsgUpdateGroupMembersResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupMembersResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateGroupMembersResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupMembersResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateGroupMembersResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateGroupMembersResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateGroupMembersResponse) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupMembersResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateGroupMembersResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateGroupMembersResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateGroupMembersResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgUpdateGroupMembersResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupMembersResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupMembersResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMembersResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupMembersResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupMembersResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateGroupMembersResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupMembersResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupMembersResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMembersResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupMembersResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupMembersResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMembersResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupMembersResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupMembersResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateGroupMembersResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupMembersResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupMembersResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateGroupMembersResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.MsgUpdateGroupMembersResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateGroupMembersResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMembersResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgUpdateGroupMembersResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateGroupMembersResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateGroupMembersResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupMembersResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupMembersResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupMembersResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupMembersResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgUpdateGroupAdmin           protoreflect.MessageDescriptor
	fd_MsgUpdateGroupAdmin_admin     protoreflect.FieldDescriptor
	fd_MsgUpdateGroupAdmin_group_id  protoreflect.FieldDescriptor
	fd_MsgUpdateGroupAdmin_new_admin protoreflect.FieldDescriptor
)

func init() {
	file_regen_group_v1alpha1_tx_proto_init()
	md_MsgUpdateGroupAdmin = File_regen_group_v1alpha1_tx_proto.Messages().ByName("MsgUpdateGroupAdmin")
	fd_MsgUpdateGroupAdmin_admin = md_MsgUpdateGroupAdmin.Fields().ByName("admin")
	fd_MsgUpdateGroupAdmin_group_id = md_MsgUpdateGroupAdmin.Fields().ByName("group_id")
	fd_MsgUpdateGroupAdmin_new_admin = md_MsgUpdateGroupAdmin.Fields().ByName("new_admin")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateGroupAdmin)(nil)

type fastReflection_MsgUpdateGroupAdmin MsgUpdateGroupAdmin

func (x *MsgUpdateGroupAdmin) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupAdmin)(x)
}

func (x *MsgUpdateGroupAdmin) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateGroupAdmin_messageType fastReflection_MsgUpdateGroupAdmin_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateGroupAdmin_messageType{}

type fastReflection_MsgUpdateGroupAdmin_messageType struct{}

func (x fastReflection_MsgUpdateGroupAdmin_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupAdmin)(nil)
}
func (x fastReflection_MsgUpdateGroupAdmin_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupAdmin)
}
func (x fastReflection_MsgUpdateGroupAdmin_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupAdmin
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateGroupAdmin) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupAdmin
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateGroupAdmin) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateGroupAdmin_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateGroupAdmin) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupAdmin)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateGroupAdmin) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateGroupAdmin)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateGroupAdmin) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Admin != "" {
		value := protoreflect.ValueOfString(x.Admin)
		if !f(fd_MsgUpdateGroupAdmin_admin, value) {
			return
		}
	}
	if x.GroupId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.GroupId)
		if !f(fd_MsgUpdateGroupAdmin_group_id, value) {
			return
		}
	}
	if x.NewAdmin != "" {
		value := protoreflect.ValueOfString(x.NewAdmin)
		if !f(fd_MsgUpdateGroupAdmin_new_admin, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgUpdateGroupAdmin) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupAdmin.admin":
		return x.Admin != ""
	case "regen.group.v1alpha1.MsgUpdateGroupAdmin.group_id":
		return x.GroupId != uint64(0)
	case "regen.group.v1alpha1.MsgUpdateGroupAdmin.new_admin":
		return x.NewAdmin != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAdmin"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAdmin does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAdmin) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupAdmin.admin":
		x.Admin = ""
	case "regen.group.v1alpha1.MsgUpdateGroupAdmin.group_id":
		x.GroupId = uint64(0)
	case "regen.group.v1alpha1.MsgUpdateGroupAdmin.new_admin":
		x.NewAdmin = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAdmin"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAdmin does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateGroupAdmin) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupAdmin.admin":
		value := x.Admin
		return protoreflect.ValueOfString(value)
	case "regen.group.v1alpha1.MsgUpdateGroupAdmin.group_id":
		value := x.GroupId
		return protoreflect.ValueOfUint64(value)
	case "regen.group.v1alpha1.MsgUpdateGroupAdmin.new_admin":
		value := x.NewAdmin
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAdmin"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAdmin does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAdmin) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupAdmin.admin":
		x.Admin = value.Interface().(string)
	case "regen.group.v1alpha1.MsgUpdateGroupAdmin.group_id":
		x.GroupId = value.Uint()
	case "regen.group.v1alpha1.MsgUpdateGroupAdmin.new_admin":
		x.NewAdmin = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAdmin"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAdmin does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAdmin) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupAdmin.admin":
		panic(fmt.Errorf("field admin of message regen.group.v1alpha1.MsgUpdateGroupAdmin is not mutable"))
	case "regen.group.v1alpha1.MsgUpdateGroupAdmin.group_id":
		panic(fmt.Errorf("field group_id of message regen.group.v1alpha1.MsgUpdateGroupAdmin is not mutable"))
	case "regen.group.v1alpha1.MsgUpdateGroupAdmin.new_admin":
		panic(fmt.Errorf("field new_admin of message regen.group.v1alpha1.MsgUpdateGroupAdmin is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAdmin"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAdmin does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateGroupAdmin) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupAdmin.admin":
		return protoreflect.ValueOfString("")
	case "regen.group.v1alpha1.MsgUpdateGroupAdmin.group_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.group.v1alpha1.MsgUpdateGroupAdmin.new_admin":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAdmin"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAdmin does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateGroupAdmin) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.MsgUpdateGroupAdmin", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateGroupAdmin) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAdmin) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgUpdateGroupAdmin) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateGroupAdmin) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateGroupAdmin)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Admin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.GroupId != 0 {
			n += 1 + runtime.Sov(uint64(x.GroupId))
		}
		l = len(x.NewAdmin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupAdmin)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.NewAdmin) > 0 {
			i -= len(x.NewAdmin)
			copy(dAtA[i:], x.NewAdmin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.NewAdmin)))
			i--
			dAtA[i] = 0x1a
		}
		if x.GroupId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.GroupId))
			i--
			dAtA[i] = 0x10
		}
		if len(x.Admin) > 0 {
			i -= len(x.Admin)
			copy(dAtA[i:], x.Admin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Admin)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupAdmin)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupAdmin: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupAdmin: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Admin = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field GroupId", wireType)
				}
				x.GroupId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.GroupId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NewAdmin", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.NewAdmin = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgUpdateGroupAdminResponse protoreflect.MessageDescriptor
)

func init() {
	file_regen_group_v1alpha1_tx_proto_init()
	md_MsgUpdateGroupAdminResponse = File_regen_group_v1alpha1_tx_proto.Messages().ByName("MsgUpdateGroupAdminResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateGroupAdminResponse)(nil)

type fastReflection_MsgUpdateGroupAdminResponse MsgUpdateGroupAdminResponse

func (x *MsgUpdateGroupAdminResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupAdminResponse)(x)
}

func (x *MsgUpdateGroupAdminResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateGroupAdminResponse_messageType fastReflection_MsgUpdateGroupAdminResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateGroupAdminResponse_messageType{}

type fastReflection_MsgUpdateGroupAdminResponse_messageType struct{}

func (x fastReflection_MsgUpdateGroupAdminResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupAdminResponse)(nil)
}
func (x fastReflection_MsgUpdateGroupAdminResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupAdminResponse)
}
func (x fastReflection_MsgUpdateGroupAdminResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupAdminResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateGroupAdminResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupAdminResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateGroupAdminResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateGroupAdminResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateGroupAdminResponse) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupAdminResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateGroupAdminResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateGroupAdminResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateGroupAdminResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgUpdateGroupAdminResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAdminResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAdminResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAdminResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAdminResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAdminResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateGroupAdminResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAdminResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAdminResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAdminResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAdminResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAdminResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAdminResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAdminResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAdminResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateGroupAdminResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAdminResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAdminResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateGroupAdminResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.MsgUpdateGroupAdminResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateGroupAdminResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAdminResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgUpdateGroupAdminResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateGroupAdminResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateGroupAdminResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupAdminResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupAdminResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupAdminResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupAdminResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgUpdateGroupMetadata          protoreflect.MessageDescriptor
	fd_MsgUpdateGroupMetadata_admin    protoreflect.FieldDescriptor
	fd_MsgUpdateGroupMetadata_group_id protoreflect.FieldDescriptor
	fd_MsgUpdateGroupMetadata_metadata protoreflect.FieldDescriptor
)

func init() {
	file_regen_group_v1alpha1_tx_proto_init()
	md_MsgUpdateGroupMetadata = File_regen_group_v1alpha1_tx_proto.Messages().ByName("MsgUpdateGroupMetadata")
	fd_MsgUpdateGroupMetadata_admin = md_MsgUpdateGroupMetadata.Fields().ByName("admin")
	fd_MsgUpdateGroupMetadata_group_id = md_MsgUpdateGroupMetadata.Fields().ByName("group_id")
	fd_MsgUpdateGroupMetadata_metadata = md_MsgUpdateGroupMetadata.Fields().ByName("metadata")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateGroupMetadata)(nil)

type fastReflection_MsgUpdateGroupMetadata MsgUpdateGroupMetadata

func (x *MsgUpdateGroupMetadata) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupMetadata)(x)
}

func (x *MsgUpdateGroupMetadata) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateGroupMetadata_messageType fastReflection_MsgUpdateGroupMetadata_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateGroupMetadata_messageType{}

type fastReflection_MsgUpdateGroupMetadata_messageType struct{}

func (x fastReflection_MsgUpdateGroupMetadata_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupMetadata)(nil)
}
func (x fastReflection_MsgUpdateGroupMetadata_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupMetadata)
}
func (x fastReflection_MsgUpdateGroupMetadata_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupMetadata
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateGroupMetadata) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupMetadata
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateGroupMetadata) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateGroupMetadata_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateGroupMetadata) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupMetadata)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateGroupMetadata) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateGroupMetadata)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateGroupMetadata) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Admin != "" {
		value := protoreflect.ValueOfString(x.Admin)
		if !f(fd_MsgUpdateGroupMetadata_admin, value) {
			return
		}
	}
	if x.GroupId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.GroupId)
		if !f(fd_MsgUpdateGroupMetadata_group_id, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_MsgUpdateGroupMetadata_metadata, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgUpdateGroupMetadata) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupMetadata.admin":
		return x.Admin != ""
	case "regen.group.v1alpha1.MsgUpdateGroupMetadata.group_id":
		return x.GroupId != uint64(0)
	case "regen.group.v1alpha1.MsgUpdateGroupMetadata.metadata":
		return len(x.Metadata) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupMetadata"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupMetadata does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMetadata) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupMetadata.admin":
		x.Admin = ""
	case "regen.group.v1alpha1.MsgUpdateGroupMetadata.group_id":
		x.GroupId = uint64(0)
	case "regen.group.v1alpha1.MsgUpdateGroupMetadata.metadata":
		x.Metadata = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupMetadata"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupMetadata does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateGroupMetadata) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupMetadata.admin":
		value := x.Admin
		return protoreflect.ValueOfString(value)
	case "regen.group.v1alpha1.MsgUpdateGroupMetadata.group_id":
		value := x.GroupId
		return protoreflect.ValueOfUint64(value)
	case "regen.group.v1alpha1.MsgUpdateGroupMetadata.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupMetadata"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupMetadata does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMetadata) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupMetadata.admin":
		x.Admin = value.Interface().(string)
	case "regen.group.v1alpha1.MsgUpdateGroupMetadata.group_id":
		x.GroupId = value.Uint()
	case "regen.group.v1alpha1.MsgUpdateGroupMetadata.metadata":
		x.Metadata = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupMetadata"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupMetadata does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMetadata) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupMetadata.admin":
		panic(fmt.Errorf("field admin of message regen.group.v1alpha1.MsgUpdateGroupMetadata is not mutable"))
	case "regen.group.v1alpha1.MsgUpdateGroupMetadata.group_id":
		panic(fmt.Errorf("field group_id of message regen.group.v1alpha1.MsgUpdateGroupMetadata is not mutable"))
	case "regen.group.v1alpha1.MsgUpdateGroupMetadata.metadata":
		panic(fmt.Errorf("field metadata of message regen.group.v1alpha1.MsgUpdateGroupMetadata is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupMetadata"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupMetadata does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateGroupMetadata) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupMetadata.admin":
		return protoreflect.ValueOfString("")
	case "regen.group.v1alpha1.MsgUpdateGroupMetadata.group_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.group.v1alpha1.MsgUpdateGroupMetadata.metadata":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupMetadata"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupMetadata does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateGroupMetadata) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.MsgUpdateGroupMetadata", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateGroupMetadata) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMetadata) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgUpdateGroupMetadata) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateGroupMetadata) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateGroupMetadata)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Admin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.GroupId != 0 {
			n += 1 + runtime.Sov(uint64(x.GroupId))
		}
		l = len(x.Metadata)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupMetadata)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
			i--
			dAtA[i] = 0x1a
		}
		if x.GroupId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.GroupId))
			i--
			dAtA[i] = 0x10
		}
		if len(x.Admin) > 0 {
			i -= len(x.Admin)
			copy(dAtA[i:], x.Admin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Admin)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupMetadata)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupMetadata: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Admin = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field GroupId", wireType)
				}
				x.GroupId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.GroupId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Metadata = append(x.Metadata[:0], dAtA[iNdEx:postIndex]...)
				if x.Metadata == nil {
					x.Metadata = []byte{}
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgUpdateGroupMetadataResponse protoreflect.MessageDescriptor
)

func init() {
	file_regen_group_v1alpha1_tx_proto_init()
	md_MsgUpdateGroupMetadataResponse = File_regen_group_v1alpha1_tx_proto.Messages().ByName("MsgUpdateGroupMetadataResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateGroupMetadataResponse)(nil)

type fastReflection_MsgUpdateGroupMetadataResponse MsgUpdateGroupMetadataResponse

func (x *MsgUpdateGroupMetadataResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupMetadataResponse)(x)
}

func (x *MsgUpdateGroupMetadataResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateGroupMetadataResponse_messageType fastReflection_MsgUpdateGroupMetadataResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateGroupMetadataResponse_messageType{}

type fastReflection_MsgUpdateGroupMetadataResponse_messageType struct{}

func (x fastReflection_MsgUpdateGroupMetadataResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupMetadataResponse)(nil)
}
func (x fastReflection_MsgUpdateGroupMetadataResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupMetadataResponse)
}
func (x fastReflection_MsgUpdateGroupMetadataResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupMetadataResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupMetadataResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateGroupMetadataResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupMetadataResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateGroupMetadataResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupMetadataResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupMetadataResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupMetadataResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupMetadataResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupMetadataResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupMetadataResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupMetadataResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupMetadataResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupMetadataResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupMetadataResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupMetadataResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupMetadataResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.MsgUpdateGroupMetadataResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateGroupMetadataResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateGroupMetadataResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupMetadataResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupMetadataResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupMetadataResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupMetadataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgCreateGroupAccount                 protoreflect.MessageDescriptor
	fd_MsgCreateGroupAccount_admin           protoreflect.FieldDescriptor
	fd_MsgCreateGroupAccount_group_id        protoreflect.FieldDescriptor
	fd_MsgCreateGroupAccount_metadata        protoreflect.FieldDescriptor
	fd_MsgCreateGroupAccount_decision_policy protoreflect.FieldDescriptor
)

func init() {
	file_regen_group_v1alpha1_tx_proto_init()
	md_MsgCreateGroupAccount = File_regen_group_v1alpha1_tx_proto.Messages().ByName("MsgCreateGroupAccount")
	fd_MsgCreateGroupAccount_admin = md_MsgCreateGroupAccount.Fields().ByName("admin")
	fd_MsgCreateGroupAccount_group_id = md_MsgCreateGroupAccount.Fields().ByName("group_id")
	fd_MsgCreateGroupAccount_metadata = md_MsgCreateGroupAccount.Fields().ByName("metadata")
	fd_MsgCreateGroupAccount_decision_policy = md_MsgCreateGroupAccount.Fields().ByName("decision_policy")
}

var _ protoreflect.Message = (*fastReflection_MsgCreateGroupAccount)(nil)

type fastReflection_MsgCreateGroupAccount MsgCreateGroupAccount

func (x *MsgCreateGroupAccount) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCreateGroupAccount)(x)
}

func (x *MsgCreateGroupAccount) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCreateGroupAccount_messageType fastReflection_MsgCreateGroupAccount_messageType
var _ protoreflect.MessageType = fastReflection_MsgCreateGroupAccount_messageType{}

type fastReflection_MsgCreateGroupAccount_messageType struct{}

func (x fastReflection_MsgCreateGroupAccount_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCreateGroupAccount)(nil)
}
func (x fastReflection_MsgCreateGroupAccount_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCreateGroupAccount)
}
func (x fastReflection_MsgCreateGroupAccount_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateGroupAccount
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCreateGroupAccount) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateGroupAccount
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCreateGroupAccount) Type() protoreflect.MessageType {
	return _fastReflection_MsgCreateGroupAccount_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCreateGroupAccount) New() protoreflect.Message {
	return new(fastReflection_MsgCreateGroupAccount)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCreateGroupAccount) Interface() protoreflect.ProtoMessage {
	return (*MsgCreateGroupAccount)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCreateGroupAccount) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Admin != "" {
		value := protoreflect.ValueOfString(x.Admin)
		if !f(fd_MsgCreateGroupAccount_admin, value) {
			return
		}
	}
	if x.GroupId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.GroupId)
		if !f(fd_MsgCreateGroupAccount_group_id, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_MsgCreateGroupAccount_metadata, value) {
			return
		}
	}
	if x.DecisionPolicy != nil {
		value := protoreflect.ValueOfMessage(x.DecisionPolicy.ProtoReflect())
		if !f(fd_MsgCreateGroupAccount_decision_policy, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgCreateGroupAccount) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgCreateGroupAccount.admin":
		return x.Admin != ""
	case "regen.group.v1alpha1.MsgCreateGroupAccount.group_id":
		return x.GroupId != uint64(0)
	case "regen.group.v1alpha1.MsgCreateGroupAccount.metadata":
		return len(x.Metadata) != 0
	case "regen.group.v1alpha1.MsgCreateGroupAccount.decision_policy":
		return x.DecisionPolicy != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateGroupAccount"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateGroupAccount does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroupAccount) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgCreateGroupAccount.admin":
		x.Admin = ""
	case "regen.group.v1alpha1.MsgCreateGroupAccount.group_id":
		x.GroupId = uint64(0)
	case "regen.group.v1alpha1.MsgCreateGroupAccount.metadata":
		x.Metadata = nil
	case "regen.group.v1alpha1.MsgCreateGroupAccount.decision_policy":
		x.DecisionPolicy = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateGroupAccount"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateGroupAccount does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCreateGroupAccount) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.group.v1alpha1.MsgCreateGroupAccount.admin":
		value := x.Admin
		return protoreflect.ValueOfString(value)
	case "regen.group.v1alpha1.MsgCreateGroupAccount.group_id":
		value := x.GroupId
		return protoreflect.ValueOfUint64(value)
	case "regen.group.v1alpha1.MsgCreateGroupAccount.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	case "regen.group.v1alpha1.MsgCreateGroupAccount.decision_policy":
		value := x.DecisionPolicy
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateGroupAccount"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateGroupAccount does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroupAccount) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgCreateGroupAccount.admin":
		x.Admin = value.Interface().(string)
	case "regen.group.v1alpha1.MsgCreateGroupAccount.group_id":
		x.GroupId = value.Uint()
	case "regen.group.v1alpha1.MsgCreateGroupAccount.metadata":
		x.Metadata = value.Bytes()
	case "regen.group.v1alpha1.MsgCreateGroupAccount.decision_policy":
		x.DecisionPolicy = value.Message().Interface().(*anypb.Any)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateGroupAccount"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateGroupAccount does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroupAccount) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgCreateGroupAccount.decision_policy":
		if x.DecisionPolicy == nil {
			x.DecisionPolicy = new(anypb.Any)
		}
		return protoreflect.ValueOfMessage(x.DecisionPolicy.ProtoReflect())
	case "regen.group.v1alpha1.MsgCreateGroupAccount.admin":
		panic(fmt.Errorf("field admin of message regen.group.v1alpha1.MsgCreateGroupAccount is not mutable"))
	case "regen.group.v1alpha1.MsgCreateGroupAccount.group_id":
		panic(fmt.Errorf("field group_id of message regen.group.v1alpha1.MsgCreateGroupAccount is not mutable"))
	case "regen.group.v1alpha1.MsgCreateGroupAccount.metadata":
		panic(fmt.Errorf("field metadata of message regen.group.v1alpha1.MsgCreateGroupAccount is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateGroupAccount"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateGroupAccount does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCreateGroupAccount) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgCreateGroupAccount.admin":
		return protoreflect.ValueOfString("")
	case "regen.group.v1alpha1.MsgCreateGroupAccount.group_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.group.v1alpha1.MsgCreateGroupAccount.metadata":
		return protoreflect.ValueOfBytes(nil)
	case "regen.group.v1alpha1.MsgCreateGroupAccount.decision_policy":
		m := new(anypb.Any)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateGroupAccount"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateGroupAccount does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCreateGroupAccount) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.MsgCreateGroupAccount", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCreateGroupAccount) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroupAccount) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgCreateGroupAccount) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCreateGroupAccount) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCreateGroupAccount)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Admin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.GroupId != 0 {
			n += 1 + runtime.Sov(uint64(x.GroupId))
		}
		l = len(x.Metadata)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.DecisionPolicy != nil {
			l = options.Size(x.DecisionPolicy)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgCreateGroupAccount)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.DecisionPolicy != nil {
			encoded, err := options.Marshal(x.DecisionPolicy)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
			i--
			dAtA[i] = 0x1a
		}
		if x.GroupId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.GroupId))
			i--
			dAtA[i] = 0x10
		}
		if len(x.Admin) > 0 {
			i -= len(x.Admin)
			copy(dAtA[i:], x.Admin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Admin)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgCreateGroupAccount)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateGroupAccount: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateGroupAccount: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Admin = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field GroupId", wireType)
				}
				x.GroupId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.GroupId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Metadata = append(x.Metadata[:0], dAtA[iNdEx:postIndex]...)
				if x.Metadata == nil {
					x.Metadata = []byte{}
				}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DecisionPolicy", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.DecisionPolicy == nil {
					x.DecisionPolicy = &anypb.Any{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.DecisionPolicy); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgCreateGroupAccountResponse         protoreflect.MessageDescriptor
	fd_MsgCreateGroupAccountResponse_address protoreflect.FieldDescriptor
)

func init() {
	file_regen_group_v1alpha1_tx_proto_init()
	md_MsgCreateGroupAccountResponse = File_regen_group_v1alpha1_tx_proto.Messages().ByName("MsgCreateGroupAccountResponse")
	fd_MsgCreateGroupAccountResponse_address = md_MsgCreateGroupAccountResponse.Fields().ByName("address")
}

var _ protoreflect.Message = (*fastReflection_MsgCreateGroupAccountResponse)(nil)

type fastReflection_MsgCreateGroupAccountResponse MsgCreateGroupAccountResponse

func (x *MsgCreateGroupAccountResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCreateGroupAccountResponse)(x)
}

func (x *MsgCreateGroupAccountResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCreateGroupAccountResponse_messageType fastReflection_MsgCreateGroupAccountResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgCreateGroupAccountResponse_messageType{}

type fastReflection_MsgCreateGroupAccountResponse_messageType struct{}

func (x fastReflection_MsgCreateGroupAccountResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCreateGroupAccountResponse)(nil)
}
func (x fastReflection_MsgCreateGroupAccountResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCreateGroupAccountResponse)
}
func (x fastReflection_MsgCreateGroupAccountResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateGroupAccountResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCreateGroupAccountResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateGroupAccountResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCreateGroupAccountResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgCreateGroupAccountResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCreateGroupAccountResponse) New() protoreflect.Message {
	return new(fastReflection_MsgCreateGroupAccountResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCreateGroupAccountResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgCreateGroupAccountResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCreateGroupAccountResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Address != "" {
		value := protoreflect.ValueOfString(x.Address)
		if !f(fd_MsgCreateGroupAccountResponse_address, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgCreateGroupAccountResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgCreateGroupAccountResponse.address":
		return x.Address != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateGroupAccountResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateGroupAccountResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroupAccountResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgCreateGroupAccountResponse.address":
		x.Address = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateGroupAccountResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateGroupAccountResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCreateGroupAccountResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.group.v1alpha1.MsgCreateGroupAccountResponse.address":
		value := x.Address
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateGroupAccountResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateGroupAccountResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroupAccountResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgCreateGroupAccountResponse.address":
		x.Address = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateGroupAccountResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateGroupAccountResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroupAccountResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgCreateGroupAccountResponse.address":
		panic(fmt.Errorf("field address of message regen.group.v1alpha1.MsgCreateGroupAccountResponse is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateGroupAccountResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateGroupAccountResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCreateGroupAccountResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgCreateGroupAccountResponse.address":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateGroupAccountResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateGroupAccountResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCreateGroupAccountResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.MsgCreateGroupAccountResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCreateGroupAccountResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateGroupAccountResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgCreateGroupAccountResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCreateGroupAccountResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCreateGroupAccountResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Address)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgCreateGroupAccountResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Address) > 0 {
			i -= len(x.Address)
			copy(dAtA[i:], x.Address)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Address)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgCreateGroupAccountResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateGroupAccountResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateGroupAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Address = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgUpdateGroupAccountAdmin           protoreflect.MessageDescriptor
	fd_MsgUpdateGroupAccountAdmin_admin     protoreflect.FieldDescriptor
	fd_MsgUpdateGroupAccountAdmin_address   protoreflect.FieldDescriptor
	fd_MsgUpdateGroupAccountAdmin_new_admin protoreflect.FieldDescriptor
)

func init() {
	file_regen_group_v1alpha1_tx_proto_init()
	md_MsgUpdateGroupAccountAdmin = File_regen_group_v1alpha1_tx_proto.Messages().ByName("MsgUpdateGroupAccountAdmin")
	fd_MsgUpdateGroupAccountAdmin_admin = md_MsgUpdateGroupAccountAdmin.Fields().ByName("admin")
	fd_MsgUpdateGroupAccountAdmin_address = md_MsgUpdateGroupAccountAdmin.Fields().ByName("address")
	fd_MsgUpdateGroupAccountAdmin_new_admin = md_MsgUpdateGroupAccountAdmin.Fields().ByName("new_admin")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateGroupAccountAdmin)(nil)

type fastReflection_MsgUpdateGroupAccountAdmin MsgUpdateGroupAccountAdmin

func (x *MsgUpdateGroupAccountAdmin) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupAccountAdmin)(x)
}

func (x *MsgUpdateGroupAccountAdmin) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateGroupAccountAdmin_messageType fastReflection_MsgUpdateGroupAccountAdmin_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateGroupAccountAdmin_messageType{}

type fastReflection_MsgUpdateGroupAccountAdmin_messageType struct{}

func (x fastReflection_MsgUpdateGroupAccountAdmin_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupAccountAdmin)(nil)
}
func (x fastReflection_MsgUpdateGroupAccountAdmin_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupAccountAdmin)
}
func (x fastReflection_MsgUpdateGroupAccountAdmin_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupAccountAdmin
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateGroupAccountAdmin) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupAccountAdmin
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateGroupAccountAdmin) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateGroupAccountAdmin_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateGroupAccountAdmin) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupAccountAdmin)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateGroupAccountAdmin) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateGroupAccountAdmin)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateGroupAccountAdmin) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Admin != "" {
		value := protoreflect.ValueOfString(x.Admin)
		if !f(fd_MsgUpdateGroupAccountAdmin_admin, value) {
			return
		}
	}
	if x.Address != "" {
		value := protoreflect.ValueOfString(x.Address)
		if !f(fd_MsgUpdateGroupAccountAdmin_address, value) {
			return
		}
	}
	if x.NewAdmin != "" {
		value := protoreflect.ValueOfString(x.NewAdmin)
		if !f(fd_MsgUpdateGroupAccountAdmin_new_admin, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgUpdateGroupAccountAdmin) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupAccountAdmin.admin":
		return x.Admin != ""
	case "regen.group.v1alpha1.MsgUpdateGroupAccountAdmin.address":
		return x.Address != ""
	case "regen.group.v1alpha1.MsgUpdateGroupAccountAdmin.new_admin":
		return x.NewAdmin != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountAdmin"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountAdmin does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAccountAdmin) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupAccountAdmin.admin":
		x.Admin = ""
	case "regen.group.v1alpha1.MsgUpdateGroupAccountAdmin.address":
		x.Address = ""
	case "regen.group.v1alpha1.MsgUpdateGroupAccountAdmin.new_admin":
		x.NewAdmin = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountAdmin"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountAdmin does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateGroupAccountAdmin) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupAccountAdmin.admin":
		value := x.Admin
		return protoreflect.ValueOfString(value)
	case "regen.group.v1alpha1.MsgUpdateGroupAccountAdmin.address":
		value := x.Address
		return protoreflect.ValueOfString(value)
	case "regen.group.v1alpha1.MsgUpdateGroupAccountAdmin.new_admin":
		value := x.NewAdmin
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountAdmin"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountAdmin does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAccountAdmin) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupAccountAdmin.admin":
		x.Admin = value.Interface().(string)
	case "regen.group.v1alpha1.MsgUpdateGroupAccountAdmin.address":
		x.Address = value.Interface().(string)
	case "regen.group.v1alpha1.MsgUpdateGroupAccountAdmin.new_admin":
		x.NewAdmin = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountAdmin"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountAdmin does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAccountAdmin) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupAccountAdmin.admin":
		panic(fmt.Errorf("field admin of message regen.group.v1alpha1.MsgUpdateGroupAccountAdmin is not mutable"))
	case "regen.group.v1alpha1.MsgUpdateGroupAccountAdmin.address":
		panic(fmt.Errorf("field address of message regen.group.v1alpha1.MsgUpdateGroupAccountAdmin is not mutable"))
	case "regen.group.v1alpha1.MsgUpdateGroupAccountAdmin.new_admin":
		panic(fmt.Errorf("field new_admin of message regen.group.v1alpha1.MsgUpdateGroupAccountAdmin is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountAdmin"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountAdmin does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateGroupAccountAdmin) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupAccountAdmin.admin":
		return protoreflect.ValueOfString("")
	case "regen.group.v1alpha1.MsgUpdateGroupAccountAdmin.address":
		return protoreflect.ValueOfString("")
	case "regen.group.v1alpha1.MsgUpdateGroupAccountAdmin.new_admin":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountAdmin"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountAdmin does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateGroupAccountAdmin) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.MsgUpdateGroupAccountAdmin", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateGroupAccountAdmin) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAccountAdmin) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgUpdateGroupAccountAdmin) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateGroupAccountAdmin) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateGroupAccountAdmin)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Admin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Address)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.NewAdmin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupAccountAdmin)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.NewAdmin) > 0 {
			i -= len(x.NewAdmin)
			copy(dAtA[i:], x.NewAdmin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.NewAdmin)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Address) > 0 {
			i -= len(x.Address)
			copy(dAtA[i:], x.Address)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Address)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Admin) > 0 {
			i -= len(x.Admin)
			copy(dAtA[i:], x.Admin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Admin)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupAccountAdmin)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupAccountAdmin: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupAccountAdmin: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Admin = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Address = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NewAdmin", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.NewAdmin = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgUpdateGroupAccountAdminResponse protoreflect.MessageDescriptor
)

func init() {
	file_regen_group_v1alpha1_tx_proto_init()
	md_MsgUpdateGroupAccountAdminResponse = File_regen_group_v1alpha1_tx_proto.Messages().ByName("MsgUpdateGroupAccountAdminResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateGroupAccountAdminResponse)(nil)

type fastReflection_MsgUpdateGroupAccountAdminResponse MsgUpdateGroupAccountAdminResponse

func (x *MsgUpdateGroupAccountAdminResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupAccountAdminResponse)(x)
}

func (x *MsgUpdateGroupAccountAdminResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateGroupAccountAdminResponse_messageType fastReflection_MsgUpdateGroupAccountAdminResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateGroupAccountAdminResponse_messageType{}

type fastReflection_MsgUpdateGroupAccountAdminResponse_messageType struct{}

func (x fastReflection_MsgUpdateGroupAccountAdminResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupAccountAdminResponse)(nil)
}
func (x fastReflection_MsgUpdateGroupAccountAdminResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupAccountAdminResponse)
}
func (x fastReflection_MsgUpdateGroupAccountAdminResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupAccountAdminResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateGroupAccountAdminResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupAccountAdminResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateGroupAccountAdminResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateGroupAccountAdminResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateGroupAccountAdminResponse) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupAccountAdminResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateGroupAccountAdminResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateGroupAccountAdminResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateGroupAccountAdminResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgUpdateGroupAccountAdminResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountAdminResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountAdminResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAccountAdminResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountAdminResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountAdminResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateGroupAccountAdminResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountAdminResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountAdminResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAccountAdminResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountAdminResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountAdminResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAccountAdminResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountAdminResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountAdminResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateGroupAccountAdminResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountAdminResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountAdminResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateGroupAccountAdminResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.MsgUpdateGroupAccountAdminResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateGroupAccountAdminResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAccountAdminResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgUpdateGroupAccountAdminResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateGroupAccountAdminResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateGroupAccountAdminResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupAccountAdminResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupAccountAdminResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupAccountAdminResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupAccountAdminResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgUpdateGroupAccountDecisionPolicy                 protoreflect.MessageDescriptor
	fd_MsgUpdateGroupAccountDecisionPolicy_admin           protoreflect.FieldDescriptor
	fd_MsgUpdateGroupAccountDecisionPolicy_address         protoreflect.FieldDescriptor
	fd_MsgUpdateGroupAccountDecisionPolicy_decision_policy protoreflect.FieldDescriptor
)

func init() {
	file_regen_group_v1alpha1_tx_proto_init()
	md_MsgUpdateGroupAccountDecisionPolicy = File_regen_group_v1alpha1_tx_proto.Messages().ByName("MsgUpdateGroupAccountDecisionPolicy")
	fd_MsgUpdateGroupAccountDecisionPolicy_admin = md_MsgUpdateGroupAccountDecisionPolicy.Fields().ByName("admin")
	fd_MsgUpdateGroupAccountDecisionPolicy_address = md_MsgUpdateGroupAccountDecisionPolicy.Fields().ByName("address")
	fd_MsgUpdateGroupAccountDecisionPolicy_decision_policy = md_MsgUpdateGroupAccountDecisionPolicy.Fields().ByName("decision_policy")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateGroupAccountDecisionPolicy)(nil)

type fastReflection_MsgUpdateGroupAccountDecisionPolicy MsgUpdateGroupAccountDecisionPolicy

func (x *MsgUpdateGroupAccountDecisionPolicy) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupAccountDecisionPolicy)(x)
}

func (x *MsgUpdateGroupAccountDecisionPolicy) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateGroupAccountDecisionPolicy_messageType fastReflection_MsgUpdateGroupAccountDecisionPolicy_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateGroupAccountDecisionPolicy_messageType{}

type fastReflection_MsgUpdateGroupAccountDecisionPolicy_messageType struct{}

func (x fastReflection_MsgUpdateGroupAccountDecisionPolicy_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupAccountDecisionPolicy)(nil)
}
func (x fastReflection_MsgUpdateGroupAccountDecisionPolicy_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupAccountDecisionPolicy)
}
func (x fastReflection_MsgUpdateGroupAccountDecisionPolicy_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupAccountDecisionPolicy
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicy) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupAccountDecisionPolicy
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicy) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateGroupAccountDecisionPolicy_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicy) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupAccountDecisionPolicy)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicy) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateGroupAccountDecisionPolicy)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicy) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Admin != "" {
		value := protoreflect.ValueOfString(x.Admin)
		if !f(fd_MsgUpdateGroupAccountDecisionPolicy_admin, value) {
			return
		}
	}
	if x.Address != "" {
		value := protoreflect.ValueOfString(x.Address)
		if !f(fd_MsgUpdateGroupAccountDecisionPolicy_address, value) {
			return
		}
	}
	if x.DecisionPolicy != nil {
		value := protoreflect.ValueOfMessage(x.DecisionPolicy.ProtoReflect())
		if !f(fd_MsgUpdateGroupAccountDecisionPolicy_decision_policy, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicy) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy.admin":
		return x.Admin != ""
	case "regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy.address":
		return x.Address != ""
	case "regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy.decision_policy":
		return x.DecisionPolicy != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicy) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy.admin":
		x.Admin = ""
	case "regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy.address":
		x.Address = ""
	case "regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy.decision_policy":
		x.DecisionPolicy = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicy) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy.admin":
		value := x.Admin
		return protoreflect.ValueOfString(value)
	case "regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy.address":
		value := x.Address
		return protoreflect.ValueOfString(value)
	case "regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy.decision_policy":
		value := x.DecisionPolicy
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicy) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy.admin":
		x.Admin = value.Interface().(string)
	case "regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy.address":
		x.Address = value.Interface().(string)
	case "regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy.decision_policy":
		x.DecisionPolicy = value.Message().Interface().(*anypb.Any)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicy) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy.decision_policy":
		if x.DecisionPolicy == nil {
			x.DecisionPolicy = new(anypb.Any)
		}
		return protoreflect.ValueOfMessage(x.DecisionPolicy.ProtoReflect())
	case "regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy.admin":
		panic(fmt.Errorf("field admin of message regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy is not mutable"))
	case "regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy.address":
		panic(fmt.Errorf("field address of message regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicy) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy.admin":
		return protoreflect.ValueOfString("")
	case "regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy.address":
		return protoreflect.ValueOfString("")
	case "regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy.decision_policy":
		m := new(anypb.Any)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicy) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicy) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicy) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicy) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicy) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateGroupAccountDecisionPolicy)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Admin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Address)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.DecisionPolicy != nil {
			l = options.Size(x.DecisionPolicy)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupAccountDecisionPolicy)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.DecisionPolicy != nil {
			encoded, err := options.Marshal(x.DecisionPolicy)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Address) > 0 {
			i -= len(x.Address)
			copy(dAtA[i:], x.Address)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Address)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Admin) > 0 {
			i -= len(x.Admin)
			copy(dAtA[i:], x.Admin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Admin)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupAccountDecisionPolicy)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupAccountDecisionPolicy: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupAccountDecisionPolicy: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Admin = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Address = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DecisionPolicy", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.DecisionPolicy == nil {
					x.DecisionPolicy = &anypb.Any{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.DecisionPolicy); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgUpdateGroupAccountDecisionPolicyResponse protoreflect.MessageDescriptor
)

func init() {
	file_regen_group_v1alpha1_tx_proto_init()
	md_MsgUpdateGroupAccountDecisionPolicyResponse = File_regen_group_v1alpha1_tx_proto.Messages().ByName("MsgUpdateGroupAccountDecisionPolicyResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateGroupAccountDecisionPolicyResponse)(nil)

type fastReflection_MsgUpdateGroupAccountDecisionPolicyResponse MsgUpdateGroupAccountDecisionPolicyResponse

func (x *MsgUpdateGroupAccountDecisionPolicyResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupAccountDecisionPolicyResponse)(x)
}

func (x *MsgUpdateGroupAccountDecisionPolicyResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateGroupAccountDecisionPolicyResponse_messageType fastReflection_MsgUpdateGroupAccountDecisionPolicyResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateGroupAccountDecisionPolicyResponse_messageType{}

type fastReflection_MsgUpdateGroupAccountDecisionPolicyResponse_messageType struct{}

func (x fastReflection_MsgUpdateGroupAccountDecisionPolicyResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupAccountDecisionPolicyResponse)(nil)
}
func (x fastReflection_MsgUpdateGroupAccountDecisionPolicyResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupAccountDecisionPolicyResponse)
}
func (x fastReflection_MsgUpdateGroupAccountDecisionPolicyResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupAccountDecisionPolicyResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicyResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupAccountDecisionPolicyResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicyResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateGroupAccountDecisionPolicyResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicyResponse) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupAccountDecisionPolicyResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicyResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateGroupAccountDecisionPolicyResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicyResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicyResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicyResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicyResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicyResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicyResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicyResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicyResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicyResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicyResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicyResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicyResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicyResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicyResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicyResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicyResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicyResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicyResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicyResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicyResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicyResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicyResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicyResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicyResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateGroupAccountDecisionPolicyResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateGroupAccountDecisionPolicyResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupAccountDecisionPolicyResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupAccountDecisionPolicyResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupAccountDecisionPolicyResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupAccountDecisionPolicyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgUpdateGroupAccountMetadata          protoreflect.MessageDescriptor
	fd_MsgUpdateGroupAccountMetadata_admin    protoreflect.FieldDescriptor
	fd_MsgUpdateGroupAccountMetadata_address  protoreflect.FieldDescriptor
	fd_MsgUpdateGroupAccountMetadata_metadata protoreflect.FieldDescriptor
)

func init() {
	file_regen_group_v1alpha1_tx_proto_init()
	md_MsgUpdateGroupAccountMetadata = File_regen_group_v1alpha1_tx_proto.Messages().ByName("MsgUpdateGroupAccountMetadata")
	fd_MsgUpdateGroupAccountMetadata_admin = md_MsgUpdateGroupAccountMetadata.Fields().ByName("admin")
	fd_MsgUpdateGroupAccountMetadata_address = md_MsgUpdateGroupAccountMetadata.Fields().ByName("address")
	fd_MsgUpdateGroupAccountMetadata_metadata = md_MsgUpdateGroupAccountMetadata.Fields().ByName("metadata")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateGroupAccountMetadata)(nil)

type fastReflection_MsgUpdateGroupAccountMetadata MsgUpdateGroupAccountMetadata

func (x *MsgUpdateGroupAccountMetadata) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupAccountMetadata)(x)
}

func (x *MsgUpdateGroupAccountMetadata) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateGroupAccountMetadata_messageType fastReflection_MsgUpdateGroupAccountMetadata_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateGroupAccountMetadata_messageType{}

type fastReflection_MsgUpdateGroupAccountMetadata_messageType struct{}

func (x fastReflection_MsgUpdateGroupAccountMetadata_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupAccountMetadata)(nil)
}
func (x fastReflection_MsgUpdateGroupAccountMetadata_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupAccountMetadata)
}
func (x fastReflection_MsgUpdateGroupAccountMetadata_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupAccountMetadata
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateGroupAccountMetadata) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupAccountMetadata
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateGroupAccountMetadata) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateGroupAccountMetadata_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateGroupAccountMetadata) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupAccountMetadata)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateGroupAccountMetadata) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateGroupAccountMetadata)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateGroupAccountMetadata) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Admin != "" {
		value := protoreflect.ValueOfString(x.Admin)
		if !f(fd_MsgUpdateGroupAccountMetadata_admin, value) {
			return
		}
	}
	if x.Address != "" {
		value := protoreflect.ValueOfString(x.Address)
		if !f(fd_MsgUpdateGroupAccountMetadata_address, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_MsgUpdateGroupAccountMetadata_metadata, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgUpdateGroupAccountMetadata) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupAccountMetadata.admin":
		return x.Admin != ""
	case "regen.group.v1alpha1.MsgUpdateGroupAccountMetadata.address":
		return x.Address != ""
	case "regen.group.v1alpha1.MsgUpdateGroupAccountMetadata.metadata":
		return len(x.Metadata) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountMetadata"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountMetadata does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAccountMetadata) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupAccountMetadata.admin":
		x.Admin = ""
	case "regen.group.v1alpha1.MsgUpdateGroupAccountMetadata.address":
		x.Address = ""
	case "regen.group.v1alpha1.MsgUpdateGroupAccountMetadata.metadata":
		x.Metadata = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountMetadata"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountMetadata does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateGroupAccountMetadata) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupAccountMetadata.admin":
		value := x.Admin
		return protoreflect.ValueOfString(value)
	case "regen.group.v1alpha1.MsgUpdateGroupAccountMetadata.address":
		value := x.Address
		return protoreflect.ValueOfString(value)
	case "regen.group.v1alpha1.MsgUpdateGroupAccountMetadata.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountMetadata"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountMetadata does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAccountMetadata) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupAccountMetadata.admin":
		x.Admin = value.Interface().(string)
	case "regen.group.v1alpha1.MsgUpdateGroupAccountMetadata.address":
		x.Address = value.Interface().(string)
	case "regen.group.v1alpha1.MsgUpdateGroupAccountMetadata.metadata":
		x.Metadata = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountMetadata"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountMetadata does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAccountMetadata) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupAccountMetadata.admin":
		panic(fmt.Errorf("field admin of message regen.group.v1alpha1.MsgUpdateGroupAccountMetadata is not mutable"))
	case "regen.group.v1alpha1.MsgUpdateGroupAccountMetadata.address":
		panic(fmt.Errorf("field address of message regen.group.v1alpha1.MsgUpdateGroupAccountMetadata is not mutable"))
	case "regen.group.v1alpha1.MsgUpdateGroupAccountMetadata.metadata":
		panic(fmt.Errorf("field metadata of message regen.group.v1alpha1.MsgUpdateGroupAccountMetadata is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountMetadata"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountMetadata does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateGroupAccountMetadata) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgUpdateGroupAccountMetadata.admin":
		return protoreflect.ValueOfString("")
	case "regen.group.v1alpha1.MsgUpdateGroupAccountMetadata.address":
		return protoreflect.ValueOfString("")
	case "regen.group.v1alpha1.MsgUpdateGroupAccountMetadata.metadata":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountMetadata"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountMetadata does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateGroupAccountMetadata) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.MsgUpdateGroupAccountMetadata", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateGroupAccountMetadata) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAccountMetadata) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgUpdateGroupAccountMetadata) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateGroupAccountMetadata) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateGroupAccountMetadata)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Admin)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Address)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Metadata)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupAccountMetadata)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Address) > 0 {
			i -= len(x.Address)
			copy(dAtA[i:], x.Address)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Address)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Admin) > 0 {
			i -= len(x.Admin)
			copy(dAtA[i:], x.Admin)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Admin)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupAccountMetadata)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupAccountMetadata: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupAccountMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Admin = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Address = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Metadata = append(x.Metadata[:0], dAtA[iNdEx:postIndex]...)
				if x.Metadata == nil {
					x.Metadata = []byte{}
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgUpdateGroupAccountMetadataResponse protoreflect.MessageDescriptor
)

func init() {
	file_regen_group_v1alpha1_tx_proto_init()
	md_MsgUpdateGroupAccountMetadataResponse = File_regen_group_v1alpha1_tx_proto.Messages().ByName("MsgUpdateGroupAccountMetadataResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgUpdateGroupAccountMetadataResponse)(nil)

type fastReflection_MsgUpdateGroupAccountMetadataResponse MsgUpdateGroupAccountMetadataResponse

func (x *MsgUpdateGroupAccountMetadataResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupAccountMetadataResponse)(x)
}

func (x *MsgUpdateGroupAccountMetadataResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgUpdateGroupAccountMetadataResponse_messageType fastReflection_MsgUpdateGroupAccountMetadataResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgUpdateGroupAccountMetadataResponse_messageType{}

type fastReflection_MsgUpdateGroupAccountMetadataResponse_messageType struct{}

func (x fastReflection_MsgUpdateGroupAccountMetadataResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgUpdateGroupAccountMetadataResponse)(nil)
}
func (x fastReflection_MsgUpdateGroupAccountMetadataResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupAccountMetadataResponse)
}
func (x fastReflection_MsgUpdateGroupAccountMetadataResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupAccountMetadataResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgUpdateGroupAccountMetadataResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgUpdateGroupAccountMetadataResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgUpdateGroupAccountMetadataResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgUpdateGroupAccountMetadataResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgUpdateGroupAccountMetadataResponse) New() protoreflect.Message {
	return new(fastReflection_MsgUpdateGroupAccountMetadataResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgUpdateGroupAccountMetadataResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgUpdateGroupAccountMetadataResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgUpdateGroupAccountMetadataResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgUpdateGroupAccountMetadataResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountMetadataResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountMetadataResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAccountMetadataResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountMetadataResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountMetadataResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgUpdateGroupAccountMetadataResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountMetadataResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountMetadataResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAccountMetadataResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountMetadataResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountMetadataResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAccountMetadataResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountMetadataResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountMetadataResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgUpdateGroupAccountMetadataResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgUpdateGroupAccountMetadataResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgUpdateGroupAccountMetadataResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgUpdateGroupAccountMetadataResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.MsgUpdateGroupAccountMetadataResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgUpdateGroupAccountMetadataResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgUpdateGroupAccountMetadataResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgUpdateGroupAccountMetadataResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgUpdateGroupAccountMetadataResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgUpdateGroupAccountMetadataResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupAccountMetadataResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgUpdateGroupAccountMetadataResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupAccountMetadataResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgUpdateGroupAccountMetadataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_MsgCreateProposal_2_list)(nil)

type _MsgCreateProposal_2_list struct {
	list *[]string
}

func (x *_MsgCreateProposal_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgCreateProposal_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfString((*x.list)[i])
}

func (x *_MsgCreateProposal_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_MsgCreateProposal_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgCreateProposal_2_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message MsgCreateProposal at list field Proposers as it is not of Message kind"))
}

func (x *_MsgCreateProposal_2_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_MsgCreateProposal_2_list) NewElement() protoreflect.Value {
	v := ""
	return protoreflect.ValueOfString(v)
}

func (x *_MsgCreateProposal_2_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_MsgCreateProposal_4_list)(nil)

type _MsgCreateProposal_4_list struct {
	list *[]*anypb.Any
}

func (x *_MsgCreateProposal_4_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgCreateProposal_4_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_MsgCreateProposal_4_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*anypb.Any)
	(*x.list)[i] = concreteValue
}

func (x *_MsgCreateProposal_4_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*anypb.Any)
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgCreateProposal_4_list) AppendMutable() protoreflect.Value {
	v := new(anypb.Any)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgCreateProposal_4_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_MsgCreateProposal_4_list) NewElement() protoreflect.Value {
	v := new(anypb.Any)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgCreateProposal_4_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgCreateProposal           protoreflect.MessageDescriptor
	fd_MsgCreateProposal_address   protoreflect.FieldDescriptor
	fd_MsgCreateProposal_proposers protoreflect.FieldDescriptor
	fd_MsgCreateProposal_metadata  protoreflect.FieldDescriptor
	fd_MsgCreateProposal_msgs      protoreflect.FieldDescriptor
	fd_MsgCreateProposal_exec      protoreflect.FieldDescriptor
)

func init() {
	file_regen_group_v1alpha1_tx_proto_init()
	md_MsgCreateProposal = File_regen_group_v1alpha1_tx_proto.Messages().ByName("MsgCreateProposal")
	fd_MsgCreateProposal_address = md_MsgCreateProposal.Fields().ByName("address")
	fd_MsgCreateProposal_proposers = md_MsgCreateProposal.Fields().ByName("proposers")
	fd_MsgCreateProposal_metadata = md_MsgCreateProposal.Fields().ByName("metadata")
	fd_MsgCreateProposal_msgs = md_MsgCreateProposal.Fields().ByName("msgs")
	fd_MsgCreateProposal_exec = md_MsgCreateProposal.Fields().ByName("exec")
}

var _ protoreflect.Message = (*fastReflection_MsgCreateProposal)(nil)

type fastReflection_MsgCreateProposal MsgCreateProposal

func (x *MsgCreateProposal) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCreateProposal)(x)
}

func (x *MsgCreateProposal) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCreateProposal_messageType fastReflection_MsgCreateProposal_messageType
var _ protoreflect.MessageType = fastReflection_MsgCreateProposal_messageType{}

type fastReflection_MsgCreateProposal_messageType struct{}

func (x fastReflection_MsgCreateProposal_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCreateProposal)(nil)
}
func (x fastReflection_MsgCreateProposal_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCreateProposal)
}
func (x fastReflection_MsgCreateProposal_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateProposal
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCreateProposal) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateProposal
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCreateProposal) Type() protoreflect.MessageType {
	return _fastReflection_MsgCreateProposal_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCreateProposal) New() protoreflect.Message {
	return new(fastReflection_MsgCreateProposal)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCreateProposal) Interface() protoreflect.ProtoMessage {
	return (*MsgCreateProposal)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCreateProposal) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Address != "" {
		value := protoreflect.ValueOfString(x.Address)
		if !f(fd_MsgCreateProposal_address, value) {
			return
		}
	}
	if len(x.Proposers) != 0 {
		value := protoreflect.ValueOfList(&_MsgCreateProposal_2_list{list: &x.Proposers})
		if !f(fd_MsgCreateProposal_proposers, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_MsgCreateProposal_metadata, value) {
			return
		}
	}
	if len(x.Msgs) != 0 {
		value := protoreflect.ValueOfList(&_MsgCreateProposal_4_list{list: &x.Msgs})
		if !f(fd_MsgCreateProposal_msgs, value) {
			return
		}
	}
	if x.Exec != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.Exec))
		if !f(fd_MsgCreateProposal_exec, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgCreateProposal) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgCreateProposal.address":
		return x.Address != ""
	case "regen.group.v1alpha1.MsgCreateProposal.proposers":
		return len(x.Proposers) != 0
	case "regen.group.v1alpha1.MsgCreateProposal.metadata":
		return len(x.Metadata) != 0
	case "regen.group.v1alpha1.MsgCreateProposal.msgs":
		return len(x.Msgs) != 0
	case "regen.group.v1alpha1.MsgCreateProposal.exec":
		return x.Exec != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateProposal"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateProposal does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateProposal) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgCreateProposal.address":
		x.Address = ""
	case "regen.group.v1alpha1.MsgCreateProposal.proposers":
		x.Proposers = nil
	case "regen.group.v1alpha1.MsgCreateProposal.metadata":
		x.Metadata = nil
	case "regen.group.v1alpha1.MsgCreateProposal.msgs":
		x.Msgs = nil
	case "regen.group.v1alpha1.MsgCreateProposal.exec":
		x.Exec = 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateProposal"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateProposal does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCreateProposal) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.group.v1alpha1.MsgCreateProposal.address":
		value := x.Address
		return protoreflect.ValueOfString(value)
	case "regen.group.v1alpha1.MsgCreateProposal.proposers":
		if len(x.Proposers) == 0 {
			return protoreflect.ValueOfList(&_MsgCreateProposal_2_list{})
		}
		listValue := &_MsgCreateProposal_2_list{list: &x.Proposers}
		return protoreflect.ValueOfList(listValue)
	case "regen.group.v1alpha1.MsgCreateProposal.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	case "regen.group.v1alpha1.MsgCreateProposal.msgs":
		if len(x.Msgs) == 0 {
			return protoreflect.ValueOfList(&_MsgCreateProposal_4_list{})
		}
		listValue := &_MsgCreateProposal_4_list{list: &x.Msgs}
		return protoreflect.ValueOfList(listValue)
	case "regen.group.v1alpha1.MsgCreateProposal.exec":
		value := x.Exec
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateProposal"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateProposal does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateProposal) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgCreateProposal.address":
		x.Address = value.Interface().(string)
	case "regen.group.v1alpha1.MsgCreateProposal.proposers":
		lv := value.List()
		clv := lv.(*_MsgCreateProposal_2_list)
		x.Proposers = *clv.list
	case "regen.group.v1alpha1.MsgCreateProposal.metadata":
		x.Metadata = value.Bytes()
	case "regen.group.v1alpha1.MsgCreateProposal.msgs":
		lv := value.List()
		clv := lv.(*_MsgCreateProposal_4_list)
		x.Msgs = *clv.list
	case "regen.group.v1alpha1.MsgCreateProposal.exec":
		x.Exec = (Exec)(value.Enum())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateProposal"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateProposal does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateProposal) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgCreateProposal.proposers":
		if x.Proposers == nil {
			x.Proposers = []string{}
		}
		value := &_MsgCreateProposal_2_list{list: &x.Proposers}
		return protoreflect.ValueOfList(value)
	case "regen.group.v1alpha1.MsgCreateProposal.msgs":
		if x.Msgs == nil {
			x.Msgs = []*anypb.Any{}
		}
		value := &_MsgCreateProposal_4_list{list: &x.Msgs}
		return protoreflect.ValueOfList(value)
	case "regen.group.v1alpha1.MsgCreateProposal.address":
		panic(fmt.Errorf("field address of message regen.group.v1alpha1.MsgCreateProposal is not mutable"))
	case "regen.group.v1alpha1.MsgCreateProposal.metadata":
		panic(fmt.Errorf("field metadata of message regen.group.v1alpha1.MsgCreateProposal is not mutable"))
	case "regen.group.v1alpha1.MsgCreateProposal.exec":
		panic(fmt.Errorf("field exec of message regen.group.v1alpha1.MsgCreateProposal is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateProposal"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateProposal does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCreateProposal) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgCreateProposal.address":
		return protoreflect.ValueOfString("")
	case "regen.group.v1alpha1.MsgCreateProposal.proposers":
		list := []string{}
		return protoreflect.ValueOfList(&_MsgCreateProposal_2_list{list: &list})
	case "regen.group.v1alpha1.MsgCreateProposal.metadata":
		return protoreflect.ValueOfBytes(nil)
	case "regen.group.v1alpha1.MsgCreateProposal.msgs":
		list := []*anypb.Any{}
		return protoreflect.ValueOfList(&_MsgCreateProposal_4_list{list: &list})
	case "regen.group.v1alpha1.MsgCreateProposal.exec":
		return protoreflect.ValueOfEnum(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateProposal"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateProposal does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCreateProposal) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.MsgCreateProposal", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCreateProposal) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateProposal) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgCreateProposal) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCreateProposal) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCreateProposal)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Address)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Proposers) > 0 {
			for _, s := range x.Proposers {
				l = len(s)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		l = len(x.Metadata)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Msgs) > 0 {
			for _, e := range x.Msgs {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.Exec != 0 {
			n += 1 + runtime.Sov(uint64(x.Exec))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgCreateProposal)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Exec != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Exec))
			i--
			dAtA[i] = 0x28
		}
		if len(x.Msgs) > 0 {
			for iNdEx := len(x.Msgs) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Msgs[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x22
			}
		}
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Proposers) > 0 {
			for iNdEx := len(x.Proposers) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.Proposers[iNdEx])
				copy(dAtA[i:], x.Proposers[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Proposers[iNdEx])))
				i--
				dAtA[i] = 0x12
			}
		}
		if len(x.Address) > 0 {
			i -= len(x.Address)
			copy(dAtA[i:], x.Address)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Address)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgCreateProposal)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateProposal: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateProposal: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Address = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Proposers", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Proposers = append(x.Proposers, string(dAtA[iNdEx:postIndex]))
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Metadata = append(x.Metadata[:0], dAtA[iNdEx:postIndex]...)
				if x.Metadata == nil {
					x.Metadata = []byte{}
				}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Msgs", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Msgs = append(x.Msgs, &anypb.Any{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Msgs[len(x.Msgs)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Exec", wireType)
				}
				x.Exec = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Exec |= Exec(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgCreateProposalResponse             protoreflect.MessageDescriptor
	fd_MsgCreateProposalResponse_proposal_id protoreflect.FieldDescriptor
)

func init() {
	file_regen_group_v1alpha1_tx_proto_init()
	md_MsgCreateProposalResponse = File_regen_group_v1alpha1_tx_proto.Messages().ByName("MsgCreateProposalResponse")
	fd_MsgCreateProposalResponse_proposal_id = md_MsgCreateProposalResponse.Fields().ByName("proposal_id")
}

var _ protoreflect.Message = (*fastReflection_MsgCreateProposalResponse)(nil)

type fastReflection_MsgCreateProposalResponse MsgCreateProposalResponse

func (x *MsgCreateProposalResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgCreateProposalResponse)(x)
}

func (x *MsgCreateProposalResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgCreateProposalResponse_messageType fastReflection_MsgCreateProposalResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgCreateProposalResponse_messageType{}

type fastReflection_MsgCreateProposalResponse_messageType struct{}

func (x fastReflection_MsgCreateProposalResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgCreateProposalResponse)(nil)
}
func (x fastReflection_MsgCreateProposalResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgCreateProposalResponse)
}
func (x fastReflection_MsgCreateProposalResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateProposalResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgCreateProposalResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgCreateProposalResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgCreateProposalResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgCreateProposalResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgCreateProposalResponse) New() protoreflect.Message {
	return new(fastReflection_MsgCreateProposalResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgCreateProposalResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgCreateProposalResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgCreateProposalResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ProposalId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.ProposalId)
		if !f(fd_MsgCreateProposalResponse_proposal_id, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgCreateProposalResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgCreateProposalResponse.proposal_id":
		return x.ProposalId != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateProposalResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateProposalResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateProposalResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgCreateProposalResponse.proposal_id":
		x.ProposalId = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateProposalResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateProposalResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgCreateProposalResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.group.v1alpha1.MsgCreateProposalResponse.proposal_id":
		value := x.ProposalId
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateProposalResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateProposalResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateProposalResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgCreateProposalResponse.proposal_id":
		x.ProposalId = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateProposalResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateProposalResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateProposalResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgCreateProposalResponse.proposal_id":
		panic(fmt.Errorf("field proposal_id of message regen.group.v1alpha1.MsgCreateProposalResponse is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateProposalResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateProposalResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgCreateProposalResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgCreateProposalResponse.proposal_id":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgCreateProposalResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgCreateProposalResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgCreateProposalResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.MsgCreateProposalResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgCreateProposalResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgCreateProposalResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgCreateProposalResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgCreateProposalResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgCreateProposalResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.ProposalId != 0 {
			n += 1 + runtime.Sov(uint64(x.ProposalId))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgCreateProposalResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.ProposalId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ProposalId))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgCreateProposalResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateProposalResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgCreateProposalResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
				}
				x.ProposalId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ProposalId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgVote             protoreflect.MessageDescriptor
	fd_MsgVote_proposal_id protoreflect.FieldDescriptor
	fd_MsgVote_voter       protoreflect.FieldDescriptor
	fd_MsgVote_choice      protoreflect.FieldDescriptor
	fd_MsgVote_metadata    protoreflect.FieldDescriptor
	fd_MsgVote_exec        protoreflect.FieldDescriptor
)

func init() {
	file_regen_group_v1alpha1_tx_proto_init()
	md_MsgVote = File_regen_group_v1alpha1_tx_proto.Messages().ByName("MsgVote")
	fd_MsgVote_proposal_id = md_MsgVote.Fields().ByName("proposal_id")
	fd_MsgVote_voter = md_MsgVote.Fields().ByName("voter")
	fd_MsgVote_choice = md_MsgVote.Fields().ByName("choice")
	fd_MsgVote_metadata = md_MsgVote.Fields().ByName("metadata")
	fd_MsgVote_exec = md_MsgVote.Fields().ByName("exec")
}

var _ protoreflect.Message = (*fastReflection_MsgVote)(nil)

type fastReflection_MsgVote MsgVote

func (x *MsgVote) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgVote)(x)
}

func (x *MsgVote) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[18]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgVote_messageType fastReflection_MsgVote_messageType
var _ protoreflect.MessageType = fastReflection_MsgVote_messageType{}

type fastReflection_MsgVote_messageType struct{}

func (x fastReflection_MsgVote_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgVote)(nil)
}
func (x fastReflection_MsgVote_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgVote)
}
func (x fastReflection_MsgVote_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgVote
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgVote) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgVote
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgVote) Type() protoreflect.MessageType {
	return _fastReflection_MsgVote_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgVote) New() protoreflect.Message {
	return new(fastReflection_MsgVote)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgVote) Interface() protoreflect.ProtoMessage {
	return (*MsgVote)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgVote) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ProposalId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.ProposalId)
		if !f(fd_MsgVote_proposal_id, value) {
			return
		}
	}
	if x.Voter != "" {
		value := protoreflect.ValueOfString(x.Voter)
		if !f(fd_MsgVote_voter, value) {
			return
		}
	}
	if x.Choice != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.Choice))
		if !f(fd_MsgVote_choice, value) {
			return
		}
	}
	if len(x.Metadata) != 0 {
		value := protoreflect.ValueOfBytes(x.Metadata)
		if !f(fd_MsgVote_metadata, value) {
			return
		}
	}
	if x.Exec != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.Exec))
		if !f(fd_MsgVote_exec, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgVote) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgVote.proposal_id":
		return x.ProposalId != uint64(0)
	case "regen.group.v1alpha1.MsgVote.voter":
		return x.Voter != ""
	case "regen.group.v1alpha1.MsgVote.choice":
		return x.Choice != 0
	case "regen.group.v1alpha1.MsgVote.metadata":
		return len(x.Metadata) != 0
	case "regen.group.v1alpha1.MsgVote.exec":
		return x.Exec != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgVote"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgVote does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgVote) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgVote.proposal_id":
		x.ProposalId = uint64(0)
	case "regen.group.v1alpha1.MsgVote.voter":
		x.Voter = ""
	case "regen.group.v1alpha1.MsgVote.choice":
		x.Choice = 0
	case "regen.group.v1alpha1.MsgVote.metadata":
		x.Metadata = nil
	case "regen.group.v1alpha1.MsgVote.exec":
		x.Exec = 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgVote"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgVote does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgVote) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.group.v1alpha1.MsgVote.proposal_id":
		value := x.ProposalId
		return protoreflect.ValueOfUint64(value)
	case "regen.group.v1alpha1.MsgVote.voter":
		value := x.Voter
		return protoreflect.ValueOfString(value)
	case "regen.group.v1alpha1.MsgVote.choice":
		value := x.Choice
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	case "regen.group.v1alpha1.MsgVote.metadata":
		value := x.Metadata
		return protoreflect.ValueOfBytes(value)
	case "regen.group.v1alpha1.MsgVote.exec":
		value := x.Exec
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgVote"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgVote does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgVote) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgVote.proposal_id":
		x.ProposalId = value.Uint()
	case "regen.group.v1alpha1.MsgVote.voter":
		x.Voter = value.Interface().(string)
	case "regen.group.v1alpha1.MsgVote.choice":
		x.Choice = (Choice)(value.Enum())
	case "regen.group.v1alpha1.MsgVote.metadata":
		x.Metadata = value.Bytes()
	case "regen.group.v1alpha1.MsgVote.exec":
		x.Exec = (Exec)(value.Enum())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgVote"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgVote does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgVote) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgVote.proposal_id":
		panic(fmt.Errorf("field proposal_id of message regen.group.v1alpha1.MsgVote is not mutable"))
	case "regen.group.v1alpha1.MsgVote.voter":
		panic(fmt.Errorf("field voter of message regen.group.v1alpha1.MsgVote is not mutable"))
	case "regen.group.v1alpha1.MsgVote.choice":
		panic(fmt.Errorf("field choice of message regen.group.v1alpha1.MsgVote is not mutable"))
	case "regen.group.v1alpha1.MsgVote.metadata":
		panic(fmt.Errorf("field metadata of message regen.group.v1alpha1.MsgVote is not mutable"))
	case "regen.group.v1alpha1.MsgVote.exec":
		panic(fmt.Errorf("field exec of message regen.group.v1alpha1.MsgVote is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgVote"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgVote does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgVote) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgVote.proposal_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.group.v1alpha1.MsgVote.voter":
		return protoreflect.ValueOfString("")
	case "regen.group.v1alpha1.MsgVote.choice":
		return protoreflect.ValueOfEnum(0)
	case "regen.group.v1alpha1.MsgVote.metadata":
		return protoreflect.ValueOfBytes(nil)
	case "regen.group.v1alpha1.MsgVote.exec":
		return protoreflect.ValueOfEnum(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgVote"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgVote does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgVote) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.MsgVote", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgVote) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgVote) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgVote) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgVote) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgVote)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.ProposalId != 0 {
			n += 1 + runtime.Sov(uint64(x.ProposalId))
		}
		l = len(x.Voter)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Choice != 0 {
			n += 1 + runtime.Sov(uint64(x.Choice))
		}
		l = len(x.Metadata)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Exec != 0 {
			n += 1 + runtime.Sov(uint64(x.Exec))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgVote)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Exec != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Exec))
			i--
			dAtA[i] = 0x28
		}
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
			i--
			dAtA[i] = 0x22
		}
		if x.Choice != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Choice))
			i--
			dAtA[i] = 0x18
		}
		if len(x.Voter) > 0 {
			i -= len(x.Voter)
			copy(dAtA[i:], x.Voter)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Voter)))
			i--
			dAtA[i] = 0x12
		}
		if x.ProposalId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ProposalId))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgVote)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgVote: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgVote: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
				}
				x.ProposalId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ProposalId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Voter", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Voter = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Choice", wireType)
				}
				x.Choice = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Choice |= Choice(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Metadata = append(x.Metadata[:0], dAtA[iNdEx:postIndex]...)
				if x.Metadata == nil {
					x.Metadata = []byte{}
				}
				iNdEx = postIndex
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Exec", wireType)
				}
				x.Exec = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Exec |= Exec(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgVoteResponse protoreflect.MessageDescriptor
)

func init() {
	file_regen_group_v1alpha1_tx_proto_init()
	md_MsgVoteResponse = File_regen_group_v1alpha1_tx_proto.Messages().ByName("MsgVoteResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgVoteResponse)(nil)

type fastReflection_MsgVoteResponse MsgVoteResponse

func (x *MsgVoteResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgVoteResponse)(x)
}

func (x *MsgVoteResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[19]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgVoteResponse_messageType fastReflection_MsgVoteResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgVoteResponse_messageType{}

type fastReflection_MsgVoteResponse_messageType struct{}

func (x fastReflection_MsgVoteResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgVoteResponse)(nil)
}
func (x fastReflection_MsgVoteResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgVoteResponse)
}
func (x fastReflection_MsgVoteResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgVoteResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgVoteResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgVoteResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgVoteResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgVoteResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgVoteResponse) New() protoreflect.Message {
	return new(fastReflection_MsgVoteResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgVoteResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgVoteResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgVoteResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgVoteResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgVoteResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgVoteResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgVoteResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgVoteResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgVoteResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgVoteResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgVoteResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgVoteResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgVoteResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgVoteResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgVoteResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgVoteResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgVoteResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgVoteResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgVoteResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgVoteResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgVoteResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgVoteResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.MsgVoteResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgVoteResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgVoteResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgVoteResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgVoteResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgVoteResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgVoteResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgVoteResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgVoteResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgVoteResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgExec             protoreflect.MessageDescriptor
	fd_MsgExec_proposal_id protoreflect.FieldDescriptor
	fd_MsgExec_signer      protoreflect.FieldDescriptor
)

func init() {
	file_regen_group_v1alpha1_tx_proto_init()
	md_MsgExec = File_regen_group_v1alpha1_tx_proto.Messages().ByName("MsgExec")
	fd_MsgExec_proposal_id = md_MsgExec.Fields().ByName("proposal_id")
	fd_MsgExec_signer = md_MsgExec.Fields().ByName("signer")
}

var _ protoreflect.Message = (*fastReflection_MsgExec)(nil)

type fastReflection_MsgExec MsgExec

func (x *MsgExec) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgExec)(x)
}

func (x *MsgExec) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[20]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgExec_messageType fastReflection_MsgExec_messageType
var _ protoreflect.MessageType = fastReflection_MsgExec_messageType{}

type fastReflection_MsgExec_messageType struct{}

func (x fastReflection_MsgExec_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgExec)(nil)
}
func (x fastReflection_MsgExec_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgExec)
}
func (x fastReflection_MsgExec_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgExec
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgExec) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgExec
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgExec) Type() protoreflect.MessageType {
	return _fastReflection_MsgExec_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgExec) New() protoreflect.Message {
	return new(fastReflection_MsgExec)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgExec) Interface() protoreflect.ProtoMessage {
	return (*MsgExec)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgExec) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ProposalId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.ProposalId)
		if !f(fd_MsgExec_proposal_id, value) {
			return
		}
	}
	if x.Signer != "" {
		value := protoreflect.ValueOfString(x.Signer)
		if !f(fd_MsgExec_signer, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgExec) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgExec.proposal_id":
		return x.ProposalId != uint64(0)
	case "regen.group.v1alpha1.MsgExec.signer":
		return x.Signer != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgExec"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgExec does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgExec) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgExec.proposal_id":
		x.ProposalId = uint64(0)
	case "regen.group.v1alpha1.MsgExec.signer":
		x.Signer = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgExec"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgExec does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgExec) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.group.v1alpha1.MsgExec.proposal_id":
		value := x.ProposalId
		return protoreflect.ValueOfUint64(value)
	case "regen.group.v1alpha1.MsgExec.signer":
		value := x.Signer
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgExec"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgExec does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgExec) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgExec.proposal_id":
		x.ProposalId = value.Uint()
	case "regen.group.v1alpha1.MsgExec.signer":
		x.Signer = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgExec"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgExec does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgExec) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgExec.proposal_id":
		panic(fmt.Errorf("field proposal_id of message regen.group.v1alpha1.MsgExec is not mutable"))
	case "regen.group.v1alpha1.MsgExec.signer":
		panic(fmt.Errorf("field signer of message regen.group.v1alpha1.MsgExec is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgExec"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgExec does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgExec) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.MsgExec.proposal_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.group.v1alpha1.MsgExec.signer":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgExec"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgExec does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgExec) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.MsgExec", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgExec) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgExec) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgExec) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgExec) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgExec)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.ProposalId != 0 {
			n += 1 + runtime.Sov(uint64(x.ProposalId))
		}
		l = len(x.Signer)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgExec)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Signer) > 0 {
			i -= len(x.Signer)
			copy(dAtA[i:], x.Signer)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Signer)))
			i--
			dAtA[i] = 0x12
		}
		if x.ProposalId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ProposalId))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgExec)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgExec: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgExec: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
				}
				x.ProposalId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ProposalId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Signer = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgExecResponse protoreflect.MessageDescriptor
)

func init() {
	file_regen_group_v1alpha1_tx_proto_init()
	md_MsgExecResponse = File_regen_group_v1alpha1_tx_proto.Messages().ByName("MsgExecResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgExecResponse)(nil)

type fastReflection_MsgExecResponse MsgExecResponse

func (x *MsgExecResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgExecResponse)(x)
}

func (x *MsgExecResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[21]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgExecResponse_messageType fastReflection_MsgExecResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgExecResponse_messageType{}

type fastReflection_MsgExecResponse_messageType struct{}

func (x fastReflection_MsgExecResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgExecResponse)(nil)
}
func (x fastReflection_MsgExecResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgExecResponse)
}
func (x fastReflection_MsgExecResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgExecResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgExecResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgExecResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgExecResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgExecResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgExecResponse) New() protoreflect.Message {
	return new(fastReflection_MsgExecResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgExecResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgExecResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgExecResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgExecResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgExecResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgExecResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgExecResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgExecResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgExecResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgExecResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgExecResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgExecResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgExecResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgExecResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgExecResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgExecResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgExecResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgExecResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgExecResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.MsgExecResponse"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.MsgExecResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgExecResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.MsgExecResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgExecResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgExecResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgExecResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgExecResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgExecResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgExecResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgExecResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgExecResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgExecResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: regen/group/v1alpha1/tx.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Exec defines modes of execution of a proposal on creation or on new vote.
type Exec int32

const (
	// An empty value means that there should be a separate
	// MsgExec request for the proposal to execute.
	Exec_EXEC_UNSPECIFIED Exec = 0
	// Try to execute the proposal immediately.
	// If the proposal is not allowed per the DecisionPolicy,
	// the proposal will still be open and could
	// be executed at a later point.
	Exec_EXEC_TRY Exec = 1
)

// Enum value maps for Exec.
var (
	Exec_name = map[int32]string{
		0: "EXEC_UNSPECIFIED",
		1: "EXEC_TRY",
	}
	Exec_value = map[string]int32{
		"EXEC_UNSPECIFIED": 0,
		"EXEC_TRY":         1,
	}
)

func (x Exec) Enum() *Exec {
	p := new(Exec)
	*p = x
	return p
}

func (x Exec) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Exec) Descriptor() protoreflect.EnumDescriptor {
	return file_regen_group_v1alpha1_tx_proto_enumTypes[0].Descriptor()
}

func (Exec) Type() protoreflect.EnumType {
	return &file_regen_group_v1alpha1_tx_proto_enumTypes[0]
}

func (x Exec) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Exec.Descriptor instead.
func (Exec) EnumDescriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_tx_proto_rawDescGZIP(), []int{0}
}

// MsgCreateGroup is the Msg/CreateGroup request type.
type MsgCreateGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// admin is the account address of the group admin.
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// members defines the group members.
	Members []*Member `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	// metadata is any arbitrary metadata to attached to the group.
	Metadata []byte `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *MsgCreateGroup) Reset() {
	*x = MsgCreateGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateGroup) ProtoMessage() {}

// Deprecated: Use MsgCreateGroup.ProtoReflect.Descriptor instead.
func (*MsgCreateGroup) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_tx_proto_rawDescGZIP(), []int{0}
}

func (x *MsgCreateGroup) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *MsgCreateGroup) GetMembers() []*Member {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *MsgCreateGroup) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// MsgCreateGroupResponse is the Msg/CreateGroup response type.
type MsgCreateGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// group_id is the unique ID of the newly created group.
	GroupId uint64 `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *MsgCreateGroupResponse) Reset() {
	*x = MsgCreateGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateGroupResponse) ProtoMessage() {}

// Deprecated: Use MsgCreateGroupResponse.ProtoReflect.Descriptor instead.
func (*MsgCreateGroupResponse) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_tx_proto_rawDescGZIP(), []int{1}
}

func (x *MsgCreateGroupResponse) GetGroupId() uint64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

// MsgUpdateGroupMembers is the Msg/UpdateGroupMembers request type.
type MsgUpdateGroupMembers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// admin is the account address of the group admin.
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// group_id is the unique ID of the group.
	GroupId uint64 `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// member_updates is the list of members to update,
	// set weight to 0 to remove a member.
	MemberUpdates []*Member `protobuf:"bytes,3,rep,name=member_updates,json=memberUpdates,proto3" json:"member_updates,omitempty"`
}

func (x *MsgUpdateGroupMembers) Reset() {
	*x = MsgUpdateGroupMembers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateGroupMembers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateGroupMembers) ProtoMessage() {}

// Deprecated: Use MsgUpdateGroupMembers.ProtoReflect.Descriptor instead.
func (*MsgUpdateGroupMembers) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_tx_proto_rawDescGZIP(), []int{2}
}

func (x *MsgUpdateGroupMembers) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *MsgUpdateGroupMembers) GetGroupId() uint64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *MsgUpdateGroupMembers) GetMemberUpdates() []*Member {
	if x != nil {
		return x.MemberUpdates
	}
	return nil
}

// MsgUpdateGroupMembersResponse is the Msg/UpdateGroupMembers response type.
type MsgUpdateGroupMembersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgUpdateGroupMembersResponse) Reset() {
	*x = MsgUpdateGroupMembersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateGroupMembersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateGroupMembersResponse) ProtoMessage() {}

// Deprecated: Use MsgUpdateGroupMembersResponse.ProtoReflect.Descriptor instead.
func (*MsgUpdateGroupMembersResponse) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_tx_proto_rawDescGZIP(), []int{3}
}

// MsgUpdateGroupAdmin is the Msg/UpdateGroupAdmin request type.
type MsgUpdateGroupAdmin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// admin is the current account address of the group admin.
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// group_id is the unique ID of the group.
	GroupId uint64 `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// new_admin is the group new admin account address.
	NewAdmin string `protobuf:"bytes,3,opt,name=new_admin,json=newAdmin,proto3" json:"new_admin,omitempty"`
}

func (x *MsgUpdateGroupAdmin) Reset() {
	*x = MsgUpdateGroupAdmin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateGroupAdmin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateGroupAdmin) ProtoMessage() {}

// Deprecated: Use MsgUpdateGroupAdmin.ProtoReflect.Descriptor instead.
func (*MsgUpdateGroupAdmin) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_tx_proto_rawDescGZIP(), []int{4}
}

func (x *MsgUpdateGroupAdmin) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *MsgUpdateGroupAdmin) GetGroupId() uint64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *MsgUpdateGroupAdmin) GetNewAdmin() string {
	if x != nil {
		return x.NewAdmin
	}
	return ""
}

// MsgUpdateGroupAdminResponse is the Msg/UpdateGroupAdmin response type.
type MsgUpdateGroupAdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgUpdateGroupAdminResponse) Reset() {
	*x = MsgUpdateGroupAdminResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateGroupAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateGroupAdminResponse) ProtoMessage() {}

// Deprecated: Use MsgUpdateGroupAdminResponse.ProtoReflect.Descriptor instead.
func (*MsgUpdateGroupAdminResponse) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_tx_proto_rawDescGZIP(), []int{5}
}

// MsgUpdateGroupMetadata is the Msg/UpdateGroupMetadata request type.
type MsgUpdateGroupMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// admin is the account address of the group admin.
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// group_id is the unique ID of the group.
	GroupId uint64 `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// metadata is the updated group's metadata.
	Metadata []byte `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *MsgUpdateGroupMetadata) Reset() {
	*x = MsgUpdateGroupMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateGroupMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateGroupMetadata) ProtoMessage() {}

// Deprecated: Use MsgUpdateGroupMetadata.ProtoReflect.Descriptor instead.
func (*MsgUpdateGroupMetadata) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_tx_proto_rawDescGZIP(), []int{6}
}

func (x *MsgUpdateGroupMetadata) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *MsgUpdateGroupMetadata) GetGroupId() uint64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *MsgUpdateGroupMetadata) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// MsgUpdateGroupMetadataResponse is the Msg/UpdateGroupMetadata response type.
type MsgUpdateGroupMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgUpdateGroupMetadataResponse) Reset() {
	*x = MsgUpdateGroupMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateGroupMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateGroupMetadataResponse) ProtoMessage() {}

// Deprecated: Use MsgUpdateGroupMetadataResponse.ProtoReflect.Descriptor instead.
func (*MsgUpdateGroupMetadataResponse) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_tx_proto_rawDescGZIP(), []int{7}
}

// MsgCreateGroupAccount is the Msg/CreateGroupAccount request type.
type MsgCreateGroupAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// admin is the account address of the group admin.
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// group_id is the unique ID of the group.
	GroupId uint64 `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// metadata is any arbitrary metadata to attached to the group account.
	Metadata []byte `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// decision_policy specifies the group account's decision policy.
	DecisionPolicy *anypb.Any `protobuf:"bytes,4,opt,name=decision_policy,json=decisionPolicy,proto3" json:"decision_policy,omitempty"`
}

func (x *MsgCreateGroupAccount) Reset() {
	*x = MsgCreateGroupAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateGroupAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateGroupAccount) ProtoMessage() {}

// Deprecated: Use MsgCreateGroupAccount.ProtoReflect.Descriptor instead.
func (*MsgCreateGroupAccount) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_tx_proto_rawDescGZIP(), []int{8}
}

func (x *MsgCreateGroupAccount) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *MsgCreateGroupAccount) GetGroupId() uint64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *MsgCreateGroupAccount) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *MsgCreateGroupAccount) GetDecisionPolicy() *anypb.Any {
	if x != nil {
		return x.DecisionPolicy
	}
	return nil
}

// MsgCreateGroupAccountResponse is the Msg/CreateGroupAccount response type.
type MsgCreateGroupAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// address is the account address of the newly created group account.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *MsgCreateGroupAccountResponse) Reset() {
	*x = MsgCreateGroupAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateGroupAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateGroupAccountResponse) ProtoMessage() {}

// Deprecated: Use MsgCreateGroupAccountResponse.ProtoReflect.Descriptor instead.
func (*MsgCreateGroupAccountResponse) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_tx_proto_rawDescGZIP(), []int{9}
}

func (x *MsgCreateGroupAccountResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

// MsgUpdateGroupAccountAdmin is the Msg/UpdateGroupAccountAdmin request type.
type MsgUpdateGroupAccountAdmin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// admin is the account address of the group admin.
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// address is the group account address.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// new_admin is the new group account admin.
	NewAdmin string `protobuf:"bytes,3,opt,name=new_admin,json=newAdmin,proto3" json:"new_admin,omitempty"`
}

func (x *MsgUpdateGroupAccountAdmin) Reset() {
	*x = MsgUpdateGroupAccountAdmin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateGroupAccountAdmin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateGroupAccountAdmin) ProtoMessage() {}

// Deprecated: Use MsgUpdateGroupAccountAdmin.ProtoReflect.Descriptor instead.
func (*MsgUpdateGroupAccountAdmin) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_tx_proto_rawDescGZIP(), []int{10}
}

func (x *MsgUpdateGroupAccountAdmin) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *MsgUpdateGroupAccountAdmin) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *MsgUpdateGroupAccountAdmin) GetNewAdmin() string {
	if x != nil {
		return x.NewAdmin
	}
	return ""
}

// MsgUpdateGroupAccountAdminResponse is the Msg/UpdateGroupAccountAdmin response type.
type MsgUpdateGroupAccountAdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgUpdateGroupAccountAdminResponse) Reset() {
	*x = MsgUpdateGroupAccountAdminResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateGroupAccountAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateGroupAccountAdminResponse) ProtoMessage() {}

// Deprecated: Use MsgUpdateGroupAccountAdminResponse.ProtoReflect.Descriptor instead.
func (*MsgUpdateGroupAccountAdminResponse) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_tx_proto_rawDescGZIP(), []int{11}
}

// MsgUpdateGroupAccountDecisionPolicy is the Msg/UpdateGroupAccountDecisionPolicy request type.
type MsgUpdateGroupAccountDecisionPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// admin is the account address of the group admin.
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// address is the group account address.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// decision_policy is the updated group account decision policy.
	DecisionPolicy *anypb.Any `protobuf:"bytes,3,opt,name=decision_policy,json=decisionPolicy,proto3" json:"decision_policy,omitempty"`
}

func (x *MsgUpdateGroupAccountDecisionPolicy) Reset() {
	*x = MsgUpdateGroupAccountDecisionPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateGroupAccountDecisionPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateGroupAccountDecisionPolicy) ProtoMessage() {}

// Deprecated: Use MsgUpdateGroupAccountDecisionPolicy.ProtoReflect.Descriptor instead.
func (*MsgUpdateGroupAccountDecisionPolicy) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_tx_proto_rawDescGZIP(), []int{12}
}

func (x *MsgUpdateGroupAccountDecisionPolicy) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *MsgUpdateGroupAccountDecisionPolicy) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *MsgUpdateGroupAccountDecisionPolicy) GetDecisionPolicy() *anypb.Any {
	if x != nil {
		return x.DecisionPolicy
	}
	return nil
}

// MsgUpdateGroupAccountDecisionPolicyResponse is the Msg/UpdateGroupAccountDecisionPolicy response type.
type MsgUpdateGroupAccountDecisionPolicyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgUpdateGroupAccountDecisionPolicyResponse) Reset() {
	*x = MsgUpdateGroupAccountDecisionPolicyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateGroupAccountDecisionPolicyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateGroupAccountDecisionPolicyResponse) ProtoMessage() {}

// Deprecated: Use MsgUpdateGroupAccountDecisionPolicyResponse.ProtoReflect.Descriptor instead.
func (*MsgUpdateGroupAccountDecisionPolicyResponse) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_tx_proto_rawDescGZIP(), []int{13}
}

// MsgUpdateGroupAccountMetadata is the Msg/UpdateGroupAccountMetadata request type.
type MsgUpdateGroupAccountMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// admin is the account address of the group admin.
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// address is the group account address.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// metadata is the updated group account metadata.
	Metadata []byte `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *MsgUpdateGroupAccountMetadata) Reset() {
	*x = MsgUpdateGroupAccountMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateGroupAccountMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateGroupAccountMetadata) ProtoMessage() {}

// Deprecated: Use MsgUpdateGroupAccountMetadata.ProtoReflect.Descriptor instead.
func (*MsgUpdateGroupAccountMetadata) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_tx_proto_rawDescGZIP(), []int{14}
}

func (x *MsgUpdateGroupAccountMetadata) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *MsgUpdateGroupAccountMetadata) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *MsgUpdateGroupAccountMetadata) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// MsgUpdateGroupAccountMetadataResponse is the Msg/UpdateGroupAccountMetadata response type.
type MsgUpdateGroupAccountMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgUpdateGroupAccountMetadataResponse) Reset() {
	*x = MsgUpdateGroupAccountMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdateGroupAccountMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdateGroupAccountMetadataResponse) ProtoMessage() {}

// Deprecated: Use MsgUpdateGroupAccountMetadataResponse.ProtoReflect.Descriptor instead.
func (*MsgUpdateGroupAccountMetadataResponse) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_tx_proto_rawDescGZIP(), []int{15}
}

// MsgCreateProposal is the Msg/CreateProposal request type.
type MsgCreateProposal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// address is the group account address.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// proposers are the account addresses of the proposers.
	// Proposers signatures will be counted as yes votes.
	Proposers []string `protobuf:"bytes,2,rep,name=proposers,proto3" json:"proposers,omitempty"`
	// metadata is any arbitrary metadata to attached to the proposal.
	Metadata []byte `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// msgs is a list of Msgs that will be executed if the proposal passes.
	Msgs []*anypb.Any `protobuf:"bytes,4,rep,name=msgs,proto3" json:"msgs,omitempty"`
	// exec defines the mode of execution of the proposal,
	// whether it should be executed immediately on creation or not.
	// If so, proposers signatures are considered as Yes votes.
	Exec Exec `protobuf:"varint,5,opt,name=exec,proto3,enum=regen.group.v1alpha1.Exec" json:"exec,omitempty"`
}

func (x *MsgCreateProposal) Reset() {
	*x = MsgCreateProposal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateProposal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateProposal) ProtoMessage() {}

// Deprecated: Use MsgCreateProposal.ProtoReflect.Descriptor instead.
func (*MsgCreateProposal) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_tx_proto_rawDescGZIP(), []int{16}
}

func (x *MsgCreateProposal) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *MsgCreateProposal) GetProposers() []string {
	if x != nil {
		return x.Proposers
	}
	return nil
}

func (x *MsgCreateProposal) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *MsgCreateProposal) GetMsgs() []*anypb.Any {
	if x != nil {
		return x.Msgs
	}
	return nil
}

func (x *MsgCreateProposal) GetExec() Exec {
	if x != nil {
		return x.Exec
	}
	return Exec_EXEC_UNSPECIFIED
}

// MsgCreateProposalResponse is the Msg/CreateProposal response type.
type MsgCreateProposalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// proposal is the unique ID of the proposal.
	ProposalId uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
}

func (x *MsgCreateProposalResponse) Reset() {
	*x = MsgCreateProposalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateProposalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateProposalResponse) ProtoMessage() {}

// Deprecated: Use MsgCreateProposalResponse.ProtoReflect.Descriptor instead.
func (*MsgCreateProposalResponse) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_tx_proto_rawDescGZIP(), []int{17}
}

func (x *MsgCreateProposalResponse) GetProposalId() uint64 {
	if x != nil {
		return x.ProposalId
	}
	return 0
}

// MsgVote is the Msg/Vote request type.
type MsgVote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// proposal is the unique ID of the proposal.
	ProposalId uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	// voter is the voter account address.
	Voter string `protobuf:"bytes,2,opt,name=voter,proto3" json:"voter,omitempty"`
	// choice is the voter's choice on the proposal.
	Choice Choice `protobuf:"varint,3,opt,name=choice,proto3,enum=regen.group.v1alpha1.Choice" json:"choice,omitempty"`
	// metadata is any arbitrary metadata to attached to the vote.
	Metadata []byte `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// exec defines whether the proposal should be executed
	// immediately after voting or not.
	Exec Exec `protobuf:"varint,5,opt,name=exec,proto3,enum=regen.group.v1alpha1.Exec" json:"exec,omitempty"`
}

func (x *MsgVote) Reset() {
	*x = MsgVote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[18]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgVote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgVote) ProtoMessage() {}

// Deprecated: Use MsgVote.ProtoReflect.Descriptor instead.
func (*MsgVote) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_tx_proto_rawDescGZIP(), []int{18}
}

func (x *MsgVote) GetProposalId() uint64 {
	if x != nil {
		return x.ProposalId
	}
	return 0
}

func (x *MsgVote) GetVoter() string {
	if x != nil {
		return x.Voter
	}
	return ""
}

func (x *MsgVote) GetChoice() Choice {
	if x != nil {
		return x.Choice
	}
	return Choice_CHOICE_UNSPECIFIED
}

func (x *MsgVote) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *MsgVote) GetExec() Exec {
	if x != nil {
		return x.Exec
	}
	return Exec_EXEC_UNSPECIFIED
}

// MsgVoteResponse is the Msg/Vote response type.
type MsgVoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgVoteResponse) Reset() {
	*x = MsgVoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[19]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgVoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgVoteResponse) ProtoMessage() {}

// Deprecated: Use MsgVoteResponse.ProtoReflect.Descriptor instead.
func (*MsgVoteResponse) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_tx_proto_rawDescGZIP(), []int{19}
}

// MsgExec is the Msg/Exec request type.
type MsgExec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// proposal is the unique ID of the proposal.
	ProposalId uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	// signer is the account address used to execute the proposal.
	Signer string `protobuf:"bytes,2,opt,name=signer,proto3" json:"signer,omitempty"`
}

func (x *MsgExec) Reset() {
	*x = MsgExec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[20]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgExec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgExec) ProtoMessage() {}

// Deprecated: Use MsgExec.ProtoReflect.Descriptor instead.
func (*MsgExec) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_tx_proto_rawDescGZIP(), []int{20}
}

func (x *MsgExec) GetProposalId() uint64 {
	if x != nil {
		return x.ProposalId
	}
	return 0
}

func (x *MsgExec) GetSigner() string {
	if x != nil {
		return x.Signer
	}
	return ""
}

// MsgExecResponse is the Msg/Exec request type.
type MsgExecResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgExecResponse) Reset() {
	*x = MsgExecResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_tx_proto_msgTypes[21]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgExecResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgExecResponse) ProtoMessage() {}

// Deprecated: Use MsgExecResponse.ProtoReflect.Descriptor instead.
func (*MsgExecResponse) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_tx_proto_rawDescGZIP(), []int{21}
}

var File_regen_group_v1alpha1_tx_proto protoreflect.FileDescriptor

var file_regen_group_v1alpha1_tx_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x14, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x80, 0x01, 0x0a, 0x0e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x3c, 0x0a, 0x07,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x04, 0xc8, 0xde, 0x1f,
	0x00, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x33, 0x0a, 0x16, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x93, 0x01, 0x0a, 0x15,
	0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x49, 0x0a, 0x0e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x04, 0xc8, 0xde,
	0x1f, 0x00, 0x52, 0x0d, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x73, 0x22, 0x1f, 0x0a, 0x1d, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x63, 0x0a, 0x13, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12,
	0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65,
	0x77, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x65, 0x77, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x1d, 0x0a, 0x1b, 0x4d, 0x73, 0x67, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x65, 0x0a, 0x16, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x20, 0x0a,
	0x1e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0xbd, 0x01, 0x0a, 0x15, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12,
	0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x51, 0x0a, 0x0f, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x42, 0x12, 0xca, 0xb4, 0x2d, 0x0e, 0x44, 0x65, 0x63, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0e, 0x64, 0x65, 0x63, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x3a, 0x04, 0x88, 0xa0, 0x1f, 0x00, 0x22,
	0x39, 0x0a, 0x1d, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x69, 0x0a, 0x1a, 0x4d, 0x73,
	0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x5f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x65, 0x77,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x24, 0x0a, 0x22, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xae, 0x01, 0x0a, 0x23,
	0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x51, 0x0a, 0x0f, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x42, 0x12, 0xca, 0xb4, 0x2d, 0x0e, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0e, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x3a, 0x04, 0x88, 0xa0, 0x1f, 0x00, 0x22, 0x2d, 0x0a, 0x2b,
	0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6b, 0x0a, 0x1d, 0x4d,
	0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x27, 0x0a, 0x25, 0x4d, 0x73, 0x67, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0xc7, 0x01, 0x0a, 0x11, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x04, 0x6d,
	0x73, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x04, 0x6d, 0x73, 0x67, 0x73, 0x12, 0x2e, 0x0a, 0x04, 0x65, 0x78, 0x65, 0x63, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x52,
	0x04, 0x65, 0x78, 0x65, 0x63, 0x3a, 0x04, 0x88, 0xa0, 0x1f, 0x00, 0x22, 0x3c, 0x0a, 0x19, 0x4d,
	0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x49, 0x64, 0x22, 0xc2, 0x01, 0x0a, 0x07, 0x4d, 0x73,
	0x67, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x06,
	0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x72,
	0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x06, 0x63, 0x68, 0x6f, 0x69,
	0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2e,
	0x0a, 0x04, 0x65, 0x78, 0x65, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x72,
	0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x52, 0x04, 0x65, 0x78, 0x65, 0x63, 0x22, 0x11,
	0x0a, 0x0f, 0x4d, 0x73, 0x67, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x42, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x45, 0x78, 0x65, 0x63, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x69, 0x67, 0x6e, 0x65, 0x72, 0x22, 0x11, 0x0a, 0x0f, 0x4d, 0x73, 0x67, 0x45, 0x78, 0x65, 0x63,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x2a, 0x0a, 0x04, 0x45, 0x78, 0x65, 0x63,
	0x12, 0x14, 0x0a, 0x10, 0x45, 0x58, 0x45, 0x43, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x58, 0x45, 0x43, 0x5f, 0x54,
	0x52, 0x59, 0x10, 0x01, 0x32, 0x89, 0x0a, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x61, 0x0a, 0x0b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x24, 0x2e, 0x72, 0x65,
	0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x1a, 0x2c, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x76, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x2b, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x1a, 0x33, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x70, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x29, 0x2e, 0x72, 0x65,
	0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x31, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73,
	0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x79, 0x0a, 0x13, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x2c, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x34,
	0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x76, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2b, 0x2e, 0x72, 0x65, 0x67,
	0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x33, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d,
	0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x85, 0x01, 0x0a,
	0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x30, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e,
	0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x38, 0x2e, 0x72, 0x65, 0x67,
	0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0xa0, 0x01, 0x0a, 0x20, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x63, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x39, 0x2e, 0x72, 0x65, 0x67, 0x65,
	0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x1a, 0x41, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x8e, 0x01, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x33, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73,
	0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x2e, 0x72, 0x65,
	0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12, 0x27, 0x2e, 0x72, 0x65, 0x67,
	0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x61, 0x6c, 0x1a, 0x2f, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x04, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x72,
	0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x56, 0x6f, 0x74, 0x65, 0x1a, 0x25, 0x2e, 0x72, 0x65,
	0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4c, 0x0a, 0x04, 0x45, 0x78, 0x65, 0x63, 0x12, 0x1d, 0x2e, 0x72, 0x65, 0x67,
	0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x4d, 0x73, 0x67, 0x45, 0x78, 0x65, 0x63, 0x1a, 0x25, 0x2e, 0x72, 0x65, 0x67, 0x65,
	0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x4d, 0x73, 0x67, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0xe3, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x07, 0x54,
	0x78, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x52, 0x47, 0x58, 0xaa, 0x02, 0x14, 0x52,
	0x65, 0x67, 0x65, 0x6e, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xca, 0x02, 0x14, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x5c, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x20, 0x52, 0x65, 0x67,
	0x65, 0x6e, 0x5c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x16,
	0x52, 0x65, 0x67, 0x65, 0x6e, 0x3a, 0x3a, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x3a, 0x3a, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_regen_group_v1alpha1_tx_proto_rawDescOnce sync.Once
	file_regen_group_v1alpha1_tx_proto_rawDescData = file_regen_group_v1alpha1_tx_proto_rawDesc
)

func file_regen_group_v1alpha1_tx_proto_rawDescGZIP() []byte {
	file_regen_group_v1alpha1_tx_proto_rawDescOnce.Do(func() {
		file_regen_group_v1alpha1_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_regen_group_v1alpha1_tx_proto_rawDescData)
	})
	return file_regen_group_v1alpha1_tx_proto_rawDescData
}

var file_regen_group_v1alpha1_tx_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_regen_group_v1alpha1_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 22)
var file_regen_group_v1alpha1_tx_proto_goTypes = []interface{}{
	(Exec)(0),                                           // 0: regen.group.v1alpha1.Exec
	(*MsgCreateGroup)(nil),                              // 1: regen.group.v1alpha1.MsgCreateGroup
	(*MsgCreateGroupResponse)(nil),                      // 2: regen.group.v1alpha1.MsgCreateGroupResponse
	(*MsgUpdateGroupMembers)(nil),                       // 3: regen.group.v1alpha1.MsgUpdateGroupMembers
	(*MsgUpdateGroupMembersResponse)(nil),               // 4: regen.group.v1alpha1.MsgUpdateGroupMembersResponse
	(*MsgUpdateGroupAdmin)(nil),                         // 5: regen.group.v1alpha1.MsgUpdateGroupAdmin
	(*MsgUpdateGroupAdminResponse)(nil),                 // 6: regen.group.v1alpha1.MsgUpdateGroupAdminResponse
	(*MsgUpdateGroupMetadata)(nil),                      // 7: regen.group.v1alpha1.MsgUpdateGroupMetadata
	(*MsgUpdateGroupMetadataResponse)(nil),              // 8: regen.group.v1alpha1.MsgUpdateGroupMetadataResponse
	(*MsgCreateGroupAccount)(nil),                       // 9: regen.group.v1alpha1.MsgCreateGroupAccount
	(*MsgCreateGroupAccountResponse)(nil),               // 10: regen.group.v1alpha1.MsgCreateGroupAccountResponse
	(*MsgUpdateGroupAccountAdmin)(nil),                  // 11: regen.group.v1alpha1.MsgUpdateGroupAccountAdmin
	(*MsgUpdateGroupAccountAdminResponse)(nil),          // 12: regen.group.v1alpha1.MsgUpdateGroupAccountAdminResponse
	(*MsgUpdateGroupAccountDecisionPolicy)(nil),         // 13: regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy
	(*MsgUpdateGroupAccountDecisionPolicyResponse)(nil), // 14: regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicyResponse
	(*MsgUpdateGroupAccountMetadata)(nil),               // 15: regen.group.v1alpha1.MsgUpdateGroupAccountMetadata
	(*MsgUpdateGroupAccountMetadataResponse)(nil),       // 16: regen.group.v1alpha1.MsgUpdateGroupAccountMetadataResponse
	(*MsgCreateProposal)(nil),                           // 17: regen.group.v1alpha1.MsgCreateProposal
	(*MsgCreateProposalResponse)(nil),                   // 18: regen.group.v1alpha1.MsgCreateProposalResponse
	(*MsgVote)(nil),                                     // 19: regen.group.v1alpha1.MsgVote
	(*MsgVoteResponse)(nil),                             // 20: regen.group.v1alpha1.MsgVoteResponse
	(*MsgExec)(nil),                                     // 21: regen.group.v1alpha1.MsgExec
	(*MsgExecResponse)(nil),                             // 22: regen.group.v1alpha1.MsgExecResponse
	(*Member)(nil),                                      // 23: regen.group.v1alpha1.Member
	(*anypb.Any)(nil),                                   // 24: google.protobuf.Any
	(Choice)(0),                                         // 25: regen.group.v1alpha1.Choice
}
var file_regen_group_v1alpha1_tx_proto_depIdxs = []int32{
	23, // 0: regen.group.v1alpha1.MsgCreateGroup.members:type_name -> regen.group.v1alpha1.Member
	23, // 1: regen.group.v1alpha1.MsgUpdateGroupMembers.member_updates:type_name -> regen.group.v1alpha1.Member
	24, // 2: regen.group.v1alpha1.MsgCreateGroupAccount.decision_policy:type_name -> google.protobuf.Any
	24, // 3: regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy.decision_policy:type_name -> google.protobuf.Any
	24, // 4: regen.group.v1alpha1.MsgCreateProposal.msgs:type_name -> google.protobuf.Any
	0,  // 5: regen.group.v1alpha1.MsgCreateProposal.exec:type_name -> regen.group.v1alpha1.Exec
	25, // 6: regen.group.v1alpha1.MsgVote.choice:type_name -> regen.group.v1alpha1.Choice
	0,  // 7: regen.group.v1alpha1.MsgVote.exec:type_name -> regen.group.v1alpha1.Exec
	1,  // 8: regen.group.v1alpha1.Msg.CreateGroup:input_type -> regen.group.v1alpha1.MsgCreateGroup
	3,  // 9: regen.group.v1alpha1.Msg.UpdateGroupMembers:input_type -> regen.group.v1alpha1.MsgUpdateGroupMembers
	5,  // 10: regen.group.v1alpha1.Msg.UpdateGroupAdmin:input_type -> regen.group.v1alpha1.MsgUpdateGroupAdmin
	7,  // 11: regen.group.v1alpha1.Msg.UpdateGroupMetadata:input_type -> regen.group.v1alpha1.MsgUpdateGroupMetadata
	9,  // 12: regen.group.v1alpha1.Msg.CreateGroupAccount:input_type -> regen.group.v1alpha1.MsgCreateGroupAccount
	11, // 13: regen.group.v1alpha1.Msg.UpdateGroupAccountAdmin:input_type -> regen.group.v1alpha1.MsgUpdateGroupAccountAdmin
	13, // 14: regen.group.v1alpha1.Msg.UpdateGroupAccountDecisionPolicy:input_type -> regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicy
	15, // 15: regen.group.v1alpha1.Msg.UpdateGroupAccountMetadata:input_type -> regen.group.v1alpha1.MsgUpdateGroupAccountMetadata
	17, // 16: regen.group.v1alpha1.Msg.CreateProposal:input_type -> regen.group.v1alpha1.MsgCreateProposal
	19, // 17: regen.group.v1alpha1.Msg.Vote:input_type -> regen.group.v1alpha1.MsgVote
	21, // 18: regen.group.v1alpha1.Msg.Exec:input_type -> regen.group.v1alpha1.MsgExec
	2,  // 19: regen.group.v1alpha1.Msg.CreateGroup:output_type -> regen.group.v1alpha1.MsgCreateGroupResponse
	4,  // 20: regen.group.v1alpha1.Msg.UpdateGroupMembers:output_type -> regen.group.v1alpha1.MsgUpdateGroupMembersResponse
	6,  // 21: regen.group.v1alpha1.Msg.UpdateGroupAdmin:output_type -> regen.group.v1alpha1.MsgUpdateGroupAdminResponse
	8,  // 22: regen.group.v1alpha1.Msg.UpdateGroupMetadata:output_type -> regen.group.v1alpha1.MsgUpdateGroupMetadataResponse
	10, // 23: regen.group.v1alpha1.Msg.CreateGroupAccount:output_type -> regen.group.v1alpha1.MsgCreateGroupAccountResponse
	12, // 24: regen.group.v1alpha1.Msg.UpdateGroupAccountAdmin:output_type -> regen.group.v1alpha1.MsgUpdateGroupAccountAdminResponse
	14, // 25: regen.group.v1alpha1.Msg.UpdateGroupAccountDecisionPolicy:output_type -> regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicyResponse
	16, // 26: regen.group.v1alpha1.Msg.UpdateGroupAccountMetadata:output_type -> regen.group.v1alpha1.MsgUpdateGroupAccountMetadataResponse
	18, // 27: regen.group.v1alpha1.Msg.CreateProposal:output_type -> regen.group.v1alpha1.MsgCreateProposalResponse
	20, // 28: regen.group.v1alpha1.Msg.Vote:output_type -> regen.group.v1alpha1.MsgVoteResponse
	22, // 29: regen.group.v1alpha1.Msg.Exec:output_type -> regen.group.v1alpha1.MsgExecResponse
	19, // [19:30] is the sub-list for method output_type
	8,  // [8:19] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_regen_group_v1alpha1_tx_proto_init() }
func file_regen_group_v1alpha1_tx_proto_init() {
	if File_regen_group_v1alpha1_tx_proto != nil {
		return
	}
	file_regen_group_v1alpha1_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_regen_group_v1alpha1_tx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateGroup); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_group_v1alpha1_tx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateGroupResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_group_v1alpha1_tx_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateGroupMembers); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_group_v1alpha1_tx_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateGroupMembersResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_group_v1alpha1_tx_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateGroupAdmin); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_group_v1alpha1_tx_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateGroupAdminResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_group_v1alpha1_tx_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateGroupMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_group_v1alpha1_tx_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateGroupMetadataResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_group_v1alpha1_tx_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateGroupAccount); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_group_v1alpha1_tx_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateGroupAccountResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_group_v1alpha1_tx_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateGroupAccountAdmin); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_group_v1alpha1_tx_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateGroupAccountAdminResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_group_v1alpha1_tx_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateGroupAccountDecisionPolicy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_group_v1alpha1_tx_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateGroupAccountDecisionPolicyResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_group_v1alpha1_tx_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateGroupAccountMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_group_v1alpha1_tx_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdateGroupAccountMetadataResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_group_v1alpha1_tx_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateProposal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_group_v1alpha1_tx_proto_msgTypes[17].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateProposalResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_group_v1alpha1_tx_proto_msgTypes[18].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgVote); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_group_v1alpha1_tx_proto_msgTypes[19].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgVoteResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_group_v1alpha1_tx_proto_msgTypes[20].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgExec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_group_v1alpha1_tx_proto_msgTypes[21].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgExecResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_regen_group_v1alpha1_tx_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   22,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_regen_group_v1alpha1_tx_proto_goTypes,
		DependencyIndexes: file_regen_group_v1alpha1_tx_proto_depIdxs,
		EnumInfos:         file_regen_group_v1alpha1_tx_proto_enumTypes,
		MessageInfos:      file_regen_group_v1alpha1_tx_proto_msgTypes,
	}.Build()
	File_regen_group_v1alpha1_tx_proto = out.File
	file_regen_group_v1alpha1_tx_proto_rawDesc = nil
	file_regen_group_v1alpha1_tx_proto_goTypes = nil
	file_regen_group_v1alpha1_tx_proto_depIdxs = nil
}