// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.0--rc2
// source: api/proto/card/card.proto

package card

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateCardRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CustomerId    uint32                 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	CardNumber    string                 `protobuf:"bytes,2,opt,name=card_number,json=cardNumber,proto3" json:"card_number,omitempty"`
	CardType      string                 `protobuf:"bytes,3,opt,name=card_type,json=cardType,proto3" json:"card_type,omitempty"`
	ExpiryDate    string                 `protobuf:"bytes,4,opt,name=expiry_date,json=expiryDate,proto3" json:"expiry_date,omitempty"`
	CreditLimit   float32                `protobuf:"fixed32,5,opt,name=credit_limit,json=creditLimit,proto3" json:"credit_limit,omitempty"`
	Balance       float32                `protobuf:"fixed32,6,opt,name=balance,proto3" json:"balance,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCardRequest) Reset() {
	*x = CreateCardRequest{}
	mi := &file_api_proto_card_card_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCardRequest) ProtoMessage() {}

func (x *CreateCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_card_card_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCardRequest.ProtoReflect.Descriptor instead.
func (*CreateCardRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_card_card_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCardRequest) GetCustomerId() uint32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *CreateCardRequest) GetCardNumber() string {
	if x != nil {
		return x.CardNumber
	}
	return ""
}

func (x *CreateCardRequest) GetCardType() string {
	if x != nil {
		return x.CardType
	}
	return ""
}

func (x *CreateCardRequest) GetExpiryDate() string {
	if x != nil {
		return x.ExpiryDate
	}
	return ""
}

func (x *CreateCardRequest) GetCreditLimit() float32 {
	if x != nil {
		return x.CreditLimit
	}
	return 0
}

func (x *CreateCardRequest) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type CreateCardResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId    uint32                 `protobuf:"varint,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	CardNumber    string                 `protobuf:"bytes,3,opt,name=card_number,json=cardNumber,proto3" json:"card_number,omitempty"`
	CardType      string                 `protobuf:"bytes,4,opt,name=card_type,json=cardType,proto3" json:"card_type,omitempty"`
	ExpiryDate    string                 `protobuf:"bytes,5,opt,name=expiry_date,json=expiryDate,proto3" json:"expiry_date,omitempty"`
	CreditLimit   float32                `protobuf:"fixed32,6,opt,name=credit_limit,json=creditLimit,proto3" json:"credit_limit,omitempty"`
	Balance       float32                `protobuf:"fixed32,7,opt,name=balance,proto3" json:"balance,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCardResponse) Reset() {
	*x = CreateCardResponse{}
	mi := &file_api_proto_card_card_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCardResponse) ProtoMessage() {}

func (x *CreateCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_card_card_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCardResponse.ProtoReflect.Descriptor instead.
func (*CreateCardResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_card_card_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCardResponse) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateCardResponse) GetCustomerId() uint32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *CreateCardResponse) GetCardNumber() string {
	if x != nil {
		return x.CardNumber
	}
	return ""
}

func (x *CreateCardResponse) GetCardType() string {
	if x != nil {
		return x.CardType
	}
	return ""
}

func (x *CreateCardResponse) GetExpiryDate() string {
	if x != nil {
		return x.ExpiryDate
	}
	return ""
}

func (x *CreateCardResponse) GetCreditLimit() float32 {
	if x != nil {
		return x.CreditLimit
	}
	return 0
}

func (x *CreateCardResponse) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type GetCardRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCardRequest) Reset() {
	*x = GetCardRequest{}
	mi := &file_api_proto_card_card_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCardRequest) ProtoMessage() {}

func (x *GetCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_card_card_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCardRequest.ProtoReflect.Descriptor instead.
func (*GetCardRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_card_card_proto_rawDescGZIP(), []int{2}
}

func (x *GetCardRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCardResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId    uint32                 `protobuf:"varint,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	CardNumber    string                 `protobuf:"bytes,3,opt,name=card_number,json=cardNumber,proto3" json:"card_number,omitempty"`
	CardType      string                 `protobuf:"bytes,4,opt,name=card_type,json=cardType,proto3" json:"card_type,omitempty"`
	ExpiryDate    string                 `protobuf:"bytes,5,opt,name=expiry_date,json=expiryDate,proto3" json:"expiry_date,omitempty"`
	CreditLimit   float32                `protobuf:"fixed32,6,opt,name=credit_limit,json=creditLimit,proto3" json:"credit_limit,omitempty"`
	Balance       float32                `protobuf:"fixed32,7,opt,name=balance,proto3" json:"balance,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCardResponse) Reset() {
	*x = GetCardResponse{}
	mi := &file_api_proto_card_card_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCardResponse) ProtoMessage() {}

func (x *GetCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_card_card_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCardResponse.ProtoReflect.Descriptor instead.
func (*GetCardResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_card_card_proto_rawDescGZIP(), []int{3}
}

func (x *GetCardResponse) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetCardResponse) GetCustomerId() uint32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *GetCardResponse) GetCardNumber() string {
	if x != nil {
		return x.CardNumber
	}
	return ""
}

func (x *GetCardResponse) GetCardType() string {
	if x != nil {
		return x.CardType
	}
	return ""
}

func (x *GetCardResponse) GetExpiryDate() string {
	if x != nil {
		return x.ExpiryDate
	}
	return ""
}

func (x *GetCardResponse) GetCreditLimit() float32 {
	if x != nil {
		return x.CreditLimit
	}
	return 0
}

func (x *GetCardResponse) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type UpdateCardRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId    uint32                 `protobuf:"varint,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	CardNumber    string                 `protobuf:"bytes,3,opt,name=card_number,json=cardNumber,proto3" json:"card_number,omitempty"`
	CardType      string                 `protobuf:"bytes,4,opt,name=card_type,json=cardType,proto3" json:"card_type,omitempty"`
	ExpiryDate    string                 `protobuf:"bytes,5,opt,name=expiry_date,json=expiryDate,proto3" json:"expiry_date,omitempty"`
	CreditLimit   float32                `protobuf:"fixed32,6,opt,name=credit_limit,json=creditLimit,proto3" json:"credit_limit,omitempty"`
	Balance       float32                `protobuf:"fixed32,7,opt,name=balance,proto3" json:"balance,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateCardRequest) Reset() {
	*x = UpdateCardRequest{}
	mi := &file_api_proto_card_card_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCardRequest) ProtoMessage() {}

func (x *UpdateCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_card_card_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCardRequest.ProtoReflect.Descriptor instead.
func (*UpdateCardRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_card_card_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCardRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateCardRequest) GetCustomerId() uint32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *UpdateCardRequest) GetCardNumber() string {
	if x != nil {
		return x.CardNumber
	}
	return ""
}

func (x *UpdateCardRequest) GetCardType() string {
	if x != nil {
		return x.CardType
	}
	return ""
}

func (x *UpdateCardRequest) GetExpiryDate() string {
	if x != nil {
		return x.ExpiryDate
	}
	return ""
}

func (x *UpdateCardRequest) GetCreditLimit() float32 {
	if x != nil {
		return x.CreditLimit
	}
	return 0
}

func (x *UpdateCardRequest) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type UpdateCardResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId    uint32                 `protobuf:"varint,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	CardNumber    string                 `protobuf:"bytes,3,opt,name=card_number,json=cardNumber,proto3" json:"card_number,omitempty"`
	CardType      string                 `protobuf:"bytes,4,opt,name=card_type,json=cardType,proto3" json:"card_type,omitempty"`
	ExpiryDate    string                 `protobuf:"bytes,5,opt,name=expiry_date,json=expiryDate,proto3" json:"expiry_date,omitempty"`
	CreditLimit   float32                `protobuf:"fixed32,6,opt,name=credit_limit,json=creditLimit,proto3" json:"credit_limit,omitempty"`
	Balance       float32                `protobuf:"fixed32,7,opt,name=balance,proto3" json:"balance,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateCardResponse) Reset() {
	*x = UpdateCardResponse{}
	mi := &file_api_proto_card_card_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCardResponse) ProtoMessage() {}

func (x *UpdateCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_card_card_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCardResponse.ProtoReflect.Descriptor instead.
func (*UpdateCardResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_card_card_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateCardResponse) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateCardResponse) GetCustomerId() uint32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *UpdateCardResponse) GetCardNumber() string {
	if x != nil {
		return x.CardNumber
	}
	return ""
}

func (x *UpdateCardResponse) GetCardType() string {
	if x != nil {
		return x.CardType
	}
	return ""
}

func (x *UpdateCardResponse) GetExpiryDate() string {
	if x != nil {
		return x.ExpiryDate
	}
	return ""
}

func (x *UpdateCardResponse) GetCreditLimit() float32 {
	if x != nil {
		return x.CreditLimit
	}
	return 0
}

func (x *UpdateCardResponse) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type DeleteCardRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCardRequest) Reset() {
	*x = DeleteCardRequest{}
	mi := &file_api_proto_card_card_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCardRequest) ProtoMessage() {}

func (x *DeleteCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_card_card_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCardRequest.ProtoReflect.Descriptor instead.
func (*DeleteCardRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_card_card_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteCardRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteCardResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCardResponse) Reset() {
	*x = DeleteCardResponse{}
	mi := &file_api_proto_card_card_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCardResponse) ProtoMessage() {}

func (x *DeleteCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_card_card_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCardResponse.ProtoReflect.Descriptor instead.
func (*DeleteCardResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_card_card_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteCardResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ListCardsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCardsRequest) Reset() {
	*x = ListCardsRequest{}
	mi := &file_api_proto_card_card_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCardsRequest) ProtoMessage() {}

func (x *ListCardsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_card_card_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCardsRequest.ProtoReflect.Descriptor instead.
func (*ListCardsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_card_card_proto_rawDescGZIP(), []int{8}
}

type ListCardsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Cards         []*GetCardResponse     `protobuf:"bytes,1,rep,name=cards,proto3" json:"cards,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCardsResponse) Reset() {
	*x = ListCardsResponse{}
	mi := &file_api_proto_card_card_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCardsResponse) ProtoMessage() {}

func (x *ListCardsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_card_card_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCardsResponse.ProtoReflect.Descriptor instead.
func (*ListCardsResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_card_card_proto_rawDescGZIP(), []int{9}
}

func (x *ListCardsResponse) GetCards() []*GetCardResponse {
	if x != nil {
		return x.Cards
	}
	return nil
}

type GetCustomerCardsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CustomerId    uint32                 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCustomerCardsRequest) Reset() {
	*x = GetCustomerCardsRequest{}
	mi := &file_api_proto_card_card_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCustomerCardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerCardsRequest) ProtoMessage() {}

func (x *GetCustomerCardsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_card_card_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerCardsRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerCardsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_card_card_proto_rawDescGZIP(), []int{10}
}

func (x *GetCustomerCardsRequest) GetCustomerId() uint32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

type GetCustomerCardsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Cards         []*GetCardResponse     `protobuf:"bytes,1,rep,name=cards,proto3" json:"cards,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCustomerCardsResponse) Reset() {
	*x = GetCustomerCardsResponse{}
	mi := &file_api_proto_card_card_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCustomerCardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerCardsResponse) ProtoMessage() {}

func (x *GetCustomerCardsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_card_card_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerCardsResponse.ProtoReflect.Descriptor instead.
func (*GetCustomerCardsResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_card_card_proto_rawDescGZIP(), []int{11}
}

func (x *GetCustomerCardsResponse) GetCards() []*GetCardResponse {
	if x != nil {
		return x.Cards
	}
	return nil
}

type AddCardRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CustomerId    uint32                 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	CardNumber    string                 `protobuf:"bytes,2,opt,name=card_number,json=cardNumber,proto3" json:"card_number,omitempty"`
	CardType      string                 `protobuf:"bytes,3,opt,name=card_type,json=cardType,proto3" json:"card_type,omitempty"`
	ExpiryDate    string                 `protobuf:"bytes,4,opt,name=expiry_date,json=expiryDate,proto3" json:"expiry_date,omitempty"`
	Cvv           string                 `protobuf:"bytes,5,opt,name=cvv,proto3" json:"cvv,omitempty"`
	CreditLimit   float32                `protobuf:"fixed32,6,opt,name=credit_limit,json=creditLimit,proto3" json:"credit_limit,omitempty"`
	Balance       float32                `protobuf:"fixed32,7,opt,name=balance,proto3" json:"balance,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddCardRequest) Reset() {
	*x = AddCardRequest{}
	mi := &file_api_proto_card_card_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCardRequest) ProtoMessage() {}

func (x *AddCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_card_card_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCardRequest.ProtoReflect.Descriptor instead.
func (*AddCardRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_card_card_proto_rawDescGZIP(), []int{12}
}

func (x *AddCardRequest) GetCustomerId() uint32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *AddCardRequest) GetCardNumber() string {
	if x != nil {
		return x.CardNumber
	}
	return ""
}

func (x *AddCardRequest) GetCardType() string {
	if x != nil {
		return x.CardType
	}
	return ""
}

func (x *AddCardRequest) GetExpiryDate() string {
	if x != nil {
		return x.ExpiryDate
	}
	return ""
}

func (x *AddCardRequest) GetCvv() string {
	if x != nil {
		return x.Cvv
	}
	return ""
}

func (x *AddCardRequest) GetCreditLimit() float32 {
	if x != nil {
		return x.CreditLimit
	}
	return 0
}

func (x *AddCardRequest) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type AddCardResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddCardResponse) Reset() {
	*x = AddCardResponse{}
	mi := &file_api_proto_card_card_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCardResponse) ProtoMessage() {}

func (x *AddCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_card_card_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCardResponse.ProtoReflect.Descriptor instead.
func (*AddCardResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_card_card_proto_rawDescGZIP(), []int{13}
}

func (x *AddCardResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type RemoveCardRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CustomerId    uint32                 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	CardNumber    string                 `protobuf:"bytes,2,opt,name=card_number,json=cardNumber,proto3" json:"card_number,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveCardRequest) Reset() {
	*x = RemoveCardRequest{}
	mi := &file_api_proto_card_card_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveCardRequest) ProtoMessage() {}

func (x *RemoveCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_card_card_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveCardRequest.ProtoReflect.Descriptor instead.
func (*RemoveCardRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_card_card_proto_rawDescGZIP(), []int{14}
}

func (x *RemoveCardRequest) GetCustomerId() uint32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *RemoveCardRequest) GetCardNumber() string {
	if x != nil {
		return x.CardNumber
	}
	return ""
}

type RemoveCardResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveCardResponse) Reset() {
	*x = RemoveCardResponse{}
	mi := &file_api_proto_card_card_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveCardResponse) ProtoMessage() {}

func (x *RemoveCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_card_card_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveCardResponse.ProtoReflect.Descriptor instead.
func (*RemoveCardResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_card_card_proto_rawDescGZIP(), []int{15}
}

func (x *RemoveCardResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_api_proto_card_card_proto protoreflect.FileDescriptor

const file_api_proto_card_card_proto_rawDesc = "" +
	"\n" +
	"\x19api/proto/card/card.proto\x12\x04card\"\xd0\x01\n" +
	"\x11CreateCardRequest\x12\x1f\n" +
	"\vcustomer_id\x18\x01 \x01(\rR\n" +
	"customerId\x12\x1f\n" +
	"\vcard_number\x18\x02 \x01(\tR\n" +
	"cardNumber\x12\x1b\n" +
	"\tcard_type\x18\x03 \x01(\tR\bcardType\x12\x1f\n" +
	"\vexpiry_date\x18\x04 \x01(\tR\n" +
	"expiryDate\x12!\n" +
	"\fcredit_limit\x18\x05 \x01(\x02R\vcreditLimit\x12\x18\n" +
	"\abalance\x18\x06 \x01(\x02R\abalance\"\xe1\x01\n" +
	"\x12CreateCardResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\x12\x1f\n" +
	"\vcustomer_id\x18\x02 \x01(\rR\n" +
	"customerId\x12\x1f\n" +
	"\vcard_number\x18\x03 \x01(\tR\n" +
	"cardNumber\x12\x1b\n" +
	"\tcard_type\x18\x04 \x01(\tR\bcardType\x12\x1f\n" +
	"\vexpiry_date\x18\x05 \x01(\tR\n" +
	"expiryDate\x12!\n" +
	"\fcredit_limit\x18\x06 \x01(\x02R\vcreditLimit\x12\x18\n" +
	"\abalance\x18\a \x01(\x02R\abalance\" \n" +
	"\x0eGetCardRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\"\xde\x01\n" +
	"\x0fGetCardResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\x12\x1f\n" +
	"\vcustomer_id\x18\x02 \x01(\rR\n" +
	"customerId\x12\x1f\n" +
	"\vcard_number\x18\x03 \x01(\tR\n" +
	"cardNumber\x12\x1b\n" +
	"\tcard_type\x18\x04 \x01(\tR\bcardType\x12\x1f\n" +
	"\vexpiry_date\x18\x05 \x01(\tR\n" +
	"expiryDate\x12!\n" +
	"\fcredit_limit\x18\x06 \x01(\x02R\vcreditLimit\x12\x18\n" +
	"\abalance\x18\a \x01(\x02R\abalance\"\xe0\x01\n" +
	"\x11UpdateCardRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\x12\x1f\n" +
	"\vcustomer_id\x18\x02 \x01(\rR\n" +
	"customerId\x12\x1f\n" +
	"\vcard_number\x18\x03 \x01(\tR\n" +
	"cardNumber\x12\x1b\n" +
	"\tcard_type\x18\x04 \x01(\tR\bcardType\x12\x1f\n" +
	"\vexpiry_date\x18\x05 \x01(\tR\n" +
	"expiryDate\x12!\n" +
	"\fcredit_limit\x18\x06 \x01(\x02R\vcreditLimit\x12\x18\n" +
	"\abalance\x18\a \x01(\x02R\abalance\"\xe1\x01\n" +
	"\x12UpdateCardResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\x12\x1f\n" +
	"\vcustomer_id\x18\x02 \x01(\rR\n" +
	"customerId\x12\x1f\n" +
	"\vcard_number\x18\x03 \x01(\tR\n" +
	"cardNumber\x12\x1b\n" +
	"\tcard_type\x18\x04 \x01(\tR\bcardType\x12\x1f\n" +
	"\vexpiry_date\x18\x05 \x01(\tR\n" +
	"expiryDate\x12!\n" +
	"\fcredit_limit\x18\x06 \x01(\x02R\vcreditLimit\x12\x18\n" +
	"\abalance\x18\a \x01(\x02R\abalance\"#\n" +
	"\x11DeleteCardRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\".\n" +
	"\x12DeleteCardResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"\x12\n" +
	"\x10ListCardsRequest\"@\n" +
	"\x11ListCardsResponse\x12+\n" +
	"\x05cards\x18\x01 \x03(\v2\x15.card.GetCardResponseR\x05cards\":\n" +
	"\x17GetCustomerCardsRequest\x12\x1f\n" +
	"\vcustomer_id\x18\x01 \x01(\rR\n" +
	"customerId\"G\n" +
	"\x18GetCustomerCardsResponse\x12+\n" +
	"\x05cards\x18\x01 \x03(\v2\x15.card.GetCardResponseR\x05cards\"\xdf\x01\n" +
	"\x0eAddCardRequest\x12\x1f\n" +
	"\vcustomer_id\x18\x01 \x01(\rR\n" +
	"customerId\x12\x1f\n" +
	"\vcard_number\x18\x02 \x01(\tR\n" +
	"cardNumber\x12\x1b\n" +
	"\tcard_type\x18\x03 \x01(\tR\bcardType\x12\x1f\n" +
	"\vexpiry_date\x18\x04 \x01(\tR\n" +
	"expiryDate\x12\x10\n" +
	"\x03cvv\x18\x05 \x01(\tR\x03cvv\x12!\n" +
	"\fcredit_limit\x18\x06 \x01(\x02R\vcreditLimit\x12\x18\n" +
	"\abalance\x18\a \x01(\x02R\abalance\"+\n" +
	"\x0fAddCardResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"U\n" +
	"\x11RemoveCardRequest\x12\x1f\n" +
	"\vcustomer_id\x18\x01 \x01(\rR\n" +
	"customerId\x12\x1f\n" +
	"\vcard_number\x18\x02 \x01(\tR\n" +
	"cardNumber\".\n" +
	"\x12RemoveCardResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess2\x92\x04\n" +
	"\vCardService\x12?\n" +
	"\n" +
	"CreateCard\x12\x17.card.CreateCardRequest\x1a\x18.card.CreateCardResponse\x126\n" +
	"\aGetCard\x12\x14.card.GetCardRequest\x1a\x15.card.GetCardResponse\x12?\n" +
	"\n" +
	"UpdateCard\x12\x17.card.UpdateCardRequest\x1a\x18.card.UpdateCardResponse\x12?\n" +
	"\n" +
	"DeleteCard\x12\x17.card.DeleteCardRequest\x1a\x18.card.DeleteCardResponse\x12<\n" +
	"\tListCards\x12\x16.card.ListCardsRequest\x1a\x17.card.ListCardsResponse\x12Q\n" +
	"\x10GetCustomerCards\x12\x1d.card.GetCustomerCardsRequest\x1a\x1e.card.GetCustomerCardsResponse\x126\n" +
	"\aAddCard\x12\x14.card.AddCardRequest\x1a\x15.card.AddCardResponse\x12?\n" +
	"\n" +
	"RemoveCard\x12\x17.card.RemoveCardRequest\x1a\x18.card.RemoveCardResponseB\x15Z\x13govo/api/proto/cardb\x06proto3"

var (
	file_api_proto_card_card_proto_rawDescOnce sync.Once
	file_api_proto_card_card_proto_rawDescData []byte
)

func file_api_proto_card_card_proto_rawDescGZIP() []byte {
	file_api_proto_card_card_proto_rawDescOnce.Do(func() {
		file_api_proto_card_card_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_proto_card_card_proto_rawDesc), len(file_api_proto_card_card_proto_rawDesc)))
	})
	return file_api_proto_card_card_proto_rawDescData
}

var file_api_proto_card_card_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_api_proto_card_card_proto_goTypes = []any{
	(*CreateCardRequest)(nil),        // 0: card.CreateCardRequest
	(*CreateCardResponse)(nil),       // 1: card.CreateCardResponse
	(*GetCardRequest)(nil),           // 2: card.GetCardRequest
	(*GetCardResponse)(nil),          // 3: card.GetCardResponse
	(*UpdateCardRequest)(nil),        // 4: card.UpdateCardRequest
	(*UpdateCardResponse)(nil),       // 5: card.UpdateCardResponse
	(*DeleteCardRequest)(nil),        // 6: card.DeleteCardRequest
	(*DeleteCardResponse)(nil),       // 7: card.DeleteCardResponse
	(*ListCardsRequest)(nil),         // 8: card.ListCardsRequest
	(*ListCardsResponse)(nil),        // 9: card.ListCardsResponse
	(*GetCustomerCardsRequest)(nil),  // 10: card.GetCustomerCardsRequest
	(*GetCustomerCardsResponse)(nil), // 11: card.GetCustomerCardsResponse
	(*AddCardRequest)(nil),           // 12: card.AddCardRequest
	(*AddCardResponse)(nil),          // 13: card.AddCardResponse
	(*RemoveCardRequest)(nil),        // 14: card.RemoveCardRequest
	(*RemoveCardResponse)(nil),       // 15: card.RemoveCardResponse
}
var file_api_proto_card_card_proto_depIdxs = []int32{
	3,  // 0: card.ListCardsResponse.cards:type_name -> card.GetCardResponse
	3,  // 1: card.GetCustomerCardsResponse.cards:type_name -> card.GetCardResponse
	0,  // 2: card.CardService.CreateCard:input_type -> card.CreateCardRequest
	2,  // 3: card.CardService.GetCard:input_type -> card.GetCardRequest
	4,  // 4: card.CardService.UpdateCard:input_type -> card.UpdateCardRequest
	6,  // 5: card.CardService.DeleteCard:input_type -> card.DeleteCardRequest
	8,  // 6: card.CardService.ListCards:input_type -> card.ListCardsRequest
	10, // 7: card.CardService.GetCustomerCards:input_type -> card.GetCustomerCardsRequest
	12, // 8: card.CardService.AddCard:input_type -> card.AddCardRequest
	14, // 9: card.CardService.RemoveCard:input_type -> card.RemoveCardRequest
	1,  // 10: card.CardService.CreateCard:output_type -> card.CreateCardResponse
	3,  // 11: card.CardService.GetCard:output_type -> card.GetCardResponse
	5,  // 12: card.CardService.UpdateCard:output_type -> card.UpdateCardResponse
	7,  // 13: card.CardService.DeleteCard:output_type -> card.DeleteCardResponse
	9,  // 14: card.CardService.ListCards:output_type -> card.ListCardsResponse
	11, // 15: card.CardService.GetCustomerCards:output_type -> card.GetCustomerCardsResponse
	13, // 16: card.CardService.AddCard:output_type -> card.AddCardResponse
	15, // 17: card.CardService.RemoveCard:output_type -> card.RemoveCardResponse
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_api_proto_card_card_proto_init() }
func file_api_proto_card_card_proto_init() {
	if File_api_proto_card_card_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_proto_card_card_proto_rawDesc), len(file_api_proto_card_card_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_card_card_proto_goTypes,
		DependencyIndexes: file_api_proto_card_card_proto_depIdxs,
		MessageInfos:      file_api_proto_card_card_proto_msgTypes,
	}.Build()
	File_api_proto_card_card_proto = out.File
	file_api_proto_card_card_proto_goTypes = nil
	file_api_proto_card_card_proto_depIdxs = nil
}
