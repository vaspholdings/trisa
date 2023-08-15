// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: trisa/data/generic/v1beta1/transaction.proto

package generic

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Generic Transaction message for TRISA transaction payloads. The goal of this payload
// is to provide enough information to link Travel Rule Compliance information in the
// identity payload with a transaction on the blockchain or network. All fields are
// optional, this message serves as a convienience for parsing transaction payloads.
type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Txid        string  `protobuf:"bytes,1,opt,name=txid,proto3" json:"txid,omitempty"`                            // a transaction ID unique to the chain/network
	Originator  string  `protobuf:"bytes,2,opt,name=originator,proto3" json:"originator,omitempty"`                // crypto address of the originator
	Beneficiary string  `protobuf:"bytes,3,opt,name=beneficiary,proto3" json:"beneficiary,omitempty"`              // crypto address of the beneficiary
	Amount      float64 `protobuf:"fixed64,4,opt,name=amount,proto3" json:"amount,omitempty"`                      // amount of the transaction
	Network     string  `protobuf:"bytes,5,opt,name=network,proto3" json:"network,omitempty"`                      // the chain/network of the transaction
	Timestamp   string  `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                  // RFC 3339 timestamp of the transaction
	ExtraJson   string  `protobuf:"bytes,7,opt,name=extra_json,json=extraJson,proto3" json:"extra_json,omitempty"` // any extra data as a JSON formatted object
	AssetType   string  `protobuf:"bytes,8,opt,name=asset_type,json=assetType,proto3" json:"asset_type,omitempty"` // the type of virtual asset for mult-asset chains
	Tag         string  `protobuf:"bytes,9,opt,name=tag,proto3" json:"tag,omitempty"`                              // optional memo/destination-tag required by some ledgers to identify transactions
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trisa_data_generic_v1beta1_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_trisa_data_generic_v1beta1_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_trisa_data_generic_v1beta1_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetTxid() string {
	if x != nil {
		return x.Txid
	}
	return ""
}

func (x *Transaction) GetOriginator() string {
	if x != nil {
		return x.Originator
	}
	return ""
}

func (x *Transaction) GetBeneficiary() string {
	if x != nil {
		return x.Beneficiary
	}
	return ""
}

func (x *Transaction) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Transaction) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Transaction) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *Transaction) GetExtraJson() string {
	if x != nil {
		return x.ExtraJson
	}
	return ""
}

func (x *Transaction) GetAssetType() string {
	if x != nil {
		return x.AssetType
	}
	return ""
}

func (x *Transaction) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

// A control flow message to support asynchronous TRISA transfers. Pending messages can
// be returned as an intermediate response during a compliance transfer if further
// processing is required before a response can be sent. The Pending message should
// provide information to the originator about when they can expect a response via the
// reply_not_before and reply_not_after timestamps. The Pending message should also
// provide collation information such as the envelope_id and original transaction so
// that the response message can be matched to the original request.
type Pending struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnvelopeId     string       `protobuf:"bytes,1,opt,name=envelope_id,json=envelopeId,proto3" json:"envelope_id,omitempty"`               // the TRISA envelope ID that refers to the compliance communication
	ReceivedBy     string       `protobuf:"bytes,2,opt,name=received_by,json=receivedBy,proto3" json:"received_by,omitempty"`               // the name of the recipient or recipient VASP
	ReceivedAt     string       `protobuf:"bytes,3,opt,name=received_at,json=receivedAt,proto3" json:"received_at,omitempty"`               // the RFC3339 formatted timestamp when the request was received
	Message        string       `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`                                       // an optional message to respond with to the counterparty
	ReplyNotAfter  string       `protobuf:"bytes,5,opt,name=reply_not_after,json=replyNotAfter,proto3" json:"reply_not_after,omitempty"`    // the RFC3339 formatted timestamp when the response will be returned by
	ReplyNotBefore string       `protobuf:"bytes,6,opt,name=reply_not_before,json=replyNotBefore,proto3" json:"reply_not_before,omitempty"` // the RFC339 formatted timestamp that the response will not be sent before
	ExtraJson      string       `protobuf:"bytes,7,opt,name=extra_json,json=extraJson,proto3" json:"extra_json,omitempty"`                  // any extra data as a JSON formatted object
	Transaction    *Transaction `protobuf:"bytes,15,opt,name=transaction,proto3" json:"transaction,omitempty"`                              // the original transaction for reference
}

func (x *Pending) Reset() {
	*x = Pending{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trisa_data_generic_v1beta1_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pending) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pending) ProtoMessage() {}

func (x *Pending) ProtoReflect() protoreflect.Message {
	mi := &file_trisa_data_generic_v1beta1_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pending.ProtoReflect.Descriptor instead.
func (*Pending) Descriptor() ([]byte, []int) {
	return file_trisa_data_generic_v1beta1_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *Pending) GetEnvelopeId() string {
	if x != nil {
		return x.EnvelopeId
	}
	return ""
}

func (x *Pending) GetReceivedBy() string {
	if x != nil {
		return x.ReceivedBy
	}
	return ""
}

func (x *Pending) GetReceivedAt() string {
	if x != nil {
		return x.ReceivedAt
	}
	return ""
}

func (x *Pending) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Pending) GetReplyNotAfter() string {
	if x != nil {
		return x.ReplyNotAfter
	}
	return ""
}

func (x *Pending) GetReplyNotBefore() string {
	if x != nil {
		return x.ReplyNotBefore
	}
	return ""
}

func (x *Pending) GetExtraJson() string {
	if x != nil {
		return x.ExtraJson
	}
	return ""
}

func (x *Pending) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

// DEPRECATED: the ConfirmationReceipt message was deemed to cause confusion as it could
// signal that the confirmation payload was received and the transfer is concluded. Use
// the Pending message instead to signal that this is a control flow message.
type ConfirmationReceipt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnvelopeId       string `protobuf:"bytes,1,opt,name=envelope_id,json=envelopeId,proto3" json:"envelope_id,omitempty"`                    // the TRISA envelope ID for reference
	ReceivedBy       string `protobuf:"bytes,2,opt,name=received_by,json=receivedBy,proto3" json:"received_by,omitempty"`                    // name of the recipient or recipient VASP
	ReceivedAt       string `protobuf:"bytes,3,opt,name=received_at,json=receivedAt,proto3" json:"received_at,omitempty"`                    // RFC 3339 timestamp of the receipt of request
	Message          string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`                                            // A generic message to respond with
	ResponseRequired bool   `protobuf:"varint,5,opt,name=response_required,json=responseRequired,proto3" json:"response_required,omitempty"` // If the message requires an additional transfer message
}

func (x *ConfirmationReceipt) Reset() {
	*x = ConfirmationReceipt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trisa_data_generic_v1beta1_transaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmationReceipt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmationReceipt) ProtoMessage() {}

func (x *ConfirmationReceipt) ProtoReflect() protoreflect.Message {
	mi := &file_trisa_data_generic_v1beta1_transaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmationReceipt.ProtoReflect.Descriptor instead.
func (*ConfirmationReceipt) Descriptor() ([]byte, []int) {
	return file_trisa_data_generic_v1beta1_transaction_proto_rawDescGZIP(), []int{2}
}

func (x *ConfirmationReceipt) GetEnvelopeId() string {
	if x != nil {
		return x.EnvelopeId
	}
	return ""
}

func (x *ConfirmationReceipt) GetReceivedBy() string {
	if x != nil {
		return x.ReceivedBy
	}
	return ""
}

func (x *ConfirmationReceipt) GetReceivedAt() string {
	if x != nil {
		return x.ReceivedAt
	}
	return ""
}

func (x *ConfirmationReceipt) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ConfirmationReceipt) GetResponseRequired() bool {
	if x != nil {
		return x.ResponseRequired
	}
	return false
}

var File_trisa_data_generic_v1beta1_transaction_proto protoreflect.FileDescriptor

var file_trisa_data_generic_v1beta1_transaction_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x74, 0x72, 0x69, 0x73, 0x61, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a,
	0x74, 0x72, 0x69, 0x73, 0x61, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x69, 0x63, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x22, 0x83, 0x02, 0x0a, 0x0b, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x78,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x69, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x20,
	0x0a, 0x0b, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69, 0x61, 0x72, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x74, 0x72, 0x61, 0x4a, 0x73, 0x6f, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67,
	0x22, 0xc2, 0x02, 0x0a, 0x07, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x1f, 0x0a, 0x0b,
	0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x72, 0x65, 0x70,
	0x6c, 0x79, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x4e, 0x6f, 0x74, 0x41, 0x66, 0x74, 0x65,
	0x72, 0x12, 0x28, 0x0a, 0x10, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x62,
	0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x70,
	0x6c, 0x79, 0x4e, 0x6f, 0x74, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65,
	0x78, 0x74, 0x72, 0x61, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x65, 0x78, 0x74, 0x72, 0x61, 0x4a, 0x73, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x0b, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x74, 0x72, 0x69, 0x73, 0x61, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xbf, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x42, 0x79, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x69, 0x73, 0x61, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x2f, 0x74, 0x72, 0x69, 0x73, 0x61, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x72, 0x69, 0x73,
	0x61, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trisa_data_generic_v1beta1_transaction_proto_rawDescOnce sync.Once
	file_trisa_data_generic_v1beta1_transaction_proto_rawDescData = file_trisa_data_generic_v1beta1_transaction_proto_rawDesc
)

func file_trisa_data_generic_v1beta1_transaction_proto_rawDescGZIP() []byte {
	file_trisa_data_generic_v1beta1_transaction_proto_rawDescOnce.Do(func() {
		file_trisa_data_generic_v1beta1_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_trisa_data_generic_v1beta1_transaction_proto_rawDescData)
	})
	return file_trisa_data_generic_v1beta1_transaction_proto_rawDescData
}

var file_trisa_data_generic_v1beta1_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_trisa_data_generic_v1beta1_transaction_proto_goTypes = []interface{}{
	(*Transaction)(nil),         // 0: trisa.data.generic.v1beta1.Transaction
	(*Pending)(nil),             // 1: trisa.data.generic.v1beta1.Pending
	(*ConfirmationReceipt)(nil), // 2: trisa.data.generic.v1beta1.ConfirmationReceipt
}
var file_trisa_data_generic_v1beta1_transaction_proto_depIdxs = []int32{
	0, // 0: trisa.data.generic.v1beta1.Pending.transaction:type_name -> trisa.data.generic.v1beta1.Transaction
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_trisa_data_generic_v1beta1_transaction_proto_init() }
func file_trisa_data_generic_v1beta1_transaction_proto_init() {
	if File_trisa_data_generic_v1beta1_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trisa_data_generic_v1beta1_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_trisa_data_generic_v1beta1_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pending); i {
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
		file_trisa_data_generic_v1beta1_transaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmationReceipt); i {
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
			RawDescriptor: file_trisa_data_generic_v1beta1_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_trisa_data_generic_v1beta1_transaction_proto_goTypes,
		DependencyIndexes: file_trisa_data_generic_v1beta1_transaction_proto_depIdxs,
		MessageInfos:      file_trisa_data_generic_v1beta1_transaction_proto_msgTypes,
	}.Build()
	File_trisa_data_generic_v1beta1_transaction_proto = out.File
	file_trisa_data_generic_v1beta1_transaction_proto_rawDesc = nil
	file_trisa_data_generic_v1beta1_transaction_proto_goTypes = nil
	file_trisa_data_generic_v1beta1_transaction_proto_depIdxs = nil
}
