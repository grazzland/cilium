// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/http_status.proto

/*
Package envoy_type is a generated protocol buffer package.

It is generated from these files:
	envoy/type/http_status.proto
	envoy/type/percent.proto
	envoy/type/range.proto

It has these top-level messages:
	HttpStatus
	Percent
	FractionalPercent
	Int64Range
	DoubleRange
*/
package envoy_type

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/lyft/protoc-gen-validate/validate"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// HTTP response codes supported in Envoy.
// For more details: http://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
type StatusCode int32

const (
	// Empty - This code not part of the HTTP status code specification, but it is needed for proto
	// `enum` type.
	StatusCode_Empty                         StatusCode = 0
	StatusCode_Continue                      StatusCode = 100
	StatusCode_OK                            StatusCode = 200
	StatusCode_Created                       StatusCode = 201
	StatusCode_Accepted                      StatusCode = 202
	StatusCode_NonAuthoritativeInformation   StatusCode = 203
	StatusCode_NoContent                     StatusCode = 204
	StatusCode_ResetContent                  StatusCode = 205
	StatusCode_PartialContent                StatusCode = 206
	StatusCode_MultiStatus                   StatusCode = 207
	StatusCode_AlreadyReported               StatusCode = 208
	StatusCode_IMUsed                        StatusCode = 226
	StatusCode_MultipleChoices               StatusCode = 300
	StatusCode_MovedPermanently              StatusCode = 301
	StatusCode_Found                         StatusCode = 302
	StatusCode_SeeOther                      StatusCode = 303
	StatusCode_NotModified                   StatusCode = 304
	StatusCode_UseProxy                      StatusCode = 305
	StatusCode_TemporaryRedirect             StatusCode = 307
	StatusCode_PermanentRedirect             StatusCode = 308
	StatusCode_BadRequest                    StatusCode = 400
	StatusCode_Unauthorized                  StatusCode = 401
	StatusCode_PaymentRequired               StatusCode = 402
	StatusCode_Forbidden                     StatusCode = 403
	StatusCode_NotFound                      StatusCode = 404
	StatusCode_MethodNotAllowed              StatusCode = 405
	StatusCode_NotAcceptable                 StatusCode = 406
	StatusCode_ProxyAuthenticationRequired   StatusCode = 407
	StatusCode_RequestTimeout                StatusCode = 408
	StatusCode_Conflict                      StatusCode = 409
	StatusCode_Gone                          StatusCode = 410
	StatusCode_LengthRequired                StatusCode = 411
	StatusCode_PreconditionFailed            StatusCode = 412
	StatusCode_PayloadTooLarge               StatusCode = 413
	StatusCode_URITooLong                    StatusCode = 414
	StatusCode_UnsupportedMediaType          StatusCode = 415
	StatusCode_RangeNotSatisfiable           StatusCode = 416
	StatusCode_ExpectationFailed             StatusCode = 417
	StatusCode_MisdirectedRequest            StatusCode = 421
	StatusCode_UnprocessableEntity           StatusCode = 422
	StatusCode_Locked                        StatusCode = 423
	StatusCode_FailedDependency              StatusCode = 424
	StatusCode_UpgradeRequired               StatusCode = 426
	StatusCode_PreconditionRequired          StatusCode = 428
	StatusCode_TooManyRequests               StatusCode = 429
	StatusCode_RequestHeaderFieldsTooLarge   StatusCode = 431
	StatusCode_InternalServerError           StatusCode = 500
	StatusCode_NotImplemented                StatusCode = 501
	StatusCode_BadGateway                    StatusCode = 502
	StatusCode_ServiceUnavailable            StatusCode = 503
	StatusCode_GatewayTimeout                StatusCode = 504
	StatusCode_HTTPVersionNotSupported       StatusCode = 505
	StatusCode_VariantAlsoNegotiates         StatusCode = 506
	StatusCode_InsufficientStorage           StatusCode = 507
	StatusCode_LoopDetected                  StatusCode = 508
	StatusCode_NotExtended                   StatusCode = 510
	StatusCode_NetworkAuthenticationRequired StatusCode = 511
)

var StatusCode_name = map[int32]string{
	0:   "Empty",
	100: "Continue",
	200: "OK",
	201: "Created",
	202: "Accepted",
	203: "NonAuthoritativeInformation",
	204: "NoContent",
	205: "ResetContent",
	206: "PartialContent",
	207: "MultiStatus",
	208: "AlreadyReported",
	226: "IMUsed",
	300: "MultipleChoices",
	301: "MovedPermanently",
	302: "Found",
	303: "SeeOther",
	304: "NotModified",
	305: "UseProxy",
	307: "TemporaryRedirect",
	308: "PermanentRedirect",
	400: "BadRequest",
	401: "Unauthorized",
	402: "PaymentRequired",
	403: "Forbidden",
	404: "NotFound",
	405: "MethodNotAllowed",
	406: "NotAcceptable",
	407: "ProxyAuthenticationRequired",
	408: "RequestTimeout",
	409: "Conflict",
	410: "Gone",
	411: "LengthRequired",
	412: "PreconditionFailed",
	413: "PayloadTooLarge",
	414: "URITooLong",
	415: "UnsupportedMediaType",
	416: "RangeNotSatisfiable",
	417: "ExpectationFailed",
	421: "MisdirectedRequest",
	422: "UnprocessableEntity",
	423: "Locked",
	424: "FailedDependency",
	426: "UpgradeRequired",
	428: "PreconditionRequired",
	429: "TooManyRequests",
	431: "RequestHeaderFieldsTooLarge",
	500: "InternalServerError",
	501: "NotImplemented",
	502: "BadGateway",
	503: "ServiceUnavailable",
	504: "GatewayTimeout",
	505: "HTTPVersionNotSupported",
	506: "VariantAlsoNegotiates",
	507: "InsufficientStorage",
	508: "LoopDetected",
	510: "NotExtended",
	511: "NetworkAuthenticationRequired",
}
var StatusCode_value = map[string]int32{
	"Empty":                         0,
	"Continue":                      100,
	"OK":                            200,
	"Created":                       201,
	"Accepted":                      202,
	"NonAuthoritativeInformation":   203,
	"NoContent":                     204,
	"ResetContent":                  205,
	"PartialContent":                206,
	"MultiStatus":                   207,
	"AlreadyReported":               208,
	"IMUsed":                        226,
	"MultipleChoices":               300,
	"MovedPermanently":              301,
	"Found":                         302,
	"SeeOther":                      303,
	"NotModified":                   304,
	"UseProxy":                      305,
	"TemporaryRedirect":             307,
	"PermanentRedirect":             308,
	"BadRequest":                    400,
	"Unauthorized":                  401,
	"PaymentRequired":               402,
	"Forbidden":                     403,
	"NotFound":                      404,
	"MethodNotAllowed":              405,
	"NotAcceptable":                 406,
	"ProxyAuthenticationRequired":   407,
	"RequestTimeout":                408,
	"Conflict":                      409,
	"Gone":                          410,
	"LengthRequired":                411,
	"PreconditionFailed":            412,
	"PayloadTooLarge":               413,
	"URITooLong":                    414,
	"UnsupportedMediaType":          415,
	"RangeNotSatisfiable":           416,
	"ExpectationFailed":             417,
	"MisdirectedRequest":            421,
	"UnprocessableEntity":           422,
	"Locked":                        423,
	"FailedDependency":              424,
	"UpgradeRequired":               426,
	"PreconditionRequired":          428,
	"TooManyRequests":               429,
	"RequestHeaderFieldsTooLarge":   431,
	"InternalServerError":           500,
	"NotImplemented":                501,
	"BadGateway":                    502,
	"ServiceUnavailable":            503,
	"GatewayTimeout":                504,
	"HTTPVersionNotSupported":       505,
	"VariantAlsoNegotiates":         506,
	"InsufficientStorage":           507,
	"LoopDetected":                  508,
	"NotExtended":                   510,
	"NetworkAuthenticationRequired": 511,
}

func (x StatusCode) String() string {
	return proto.EnumName(StatusCode_name, int32(x))
}
func (StatusCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// HTTP status.
type HttpStatus struct {
	// Supplies HTTP response code.
	Code StatusCode `protobuf:"varint,1,opt,name=code,enum=envoy.type.StatusCode" json:"code,omitempty"`
}

func (m *HttpStatus) Reset()                    { *m = HttpStatus{} }
func (m *HttpStatus) String() string            { return proto.CompactTextString(m) }
func (*HttpStatus) ProtoMessage()               {}
func (*HttpStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HttpStatus) GetCode() StatusCode {
	if m != nil {
		return m.Code
	}
	return StatusCode_Empty
}

func init() {
	proto.RegisterType((*HttpStatus)(nil), "envoy.type.HttpStatus")
	proto.RegisterEnum("envoy.type.StatusCode", StatusCode_name, StatusCode_value)
}

func init() { proto.RegisterFile("envoy/type/http_status.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 894 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x49, 0x6f, 0x1c, 0x45,
	0x14, 0x4e, 0x4f, 0xc5, 0x49, 0x5c, 0x71, 0x9c, 0x97, 0x8a, 0x13, 0x87, 0x10, 0x24, 0x2b, 0x27,
	0xc4, 0xc1, 0x96, 0xe0, 0x0f, 0xe0, 0x38, 0x76, 0x6c, 0xe1, 0x99, 0x8c, 0x66, 0xc9, 0x15, 0x95,
	0xbb, 0xde, 0xcc, 0x94, 0xd2, 0x53, 0xaf, 0x53, 0xfd, 0x7a, 0xec, 0xe6, 0xc8, 0x2f, 0x60, 0xdf,
	0xd7, 0x03, 0x8b, 0x50, 0x42, 0x40, 0xc0, 0x85, 0x13, 0x47, 0x76, 0xf8, 0x0d, 0xdc, 0xb8, 0xb3,
	0x83, 0x00, 0x55, 0xcd, 0x92, 0x5c, 0x72, 0xab, 0x7a, 0xf5, 0x96, 0x6f, 0x79, 0xdd, 0xf2, 0x02,
	0xba, 0x11, 0x55, 0x6b, 0x5c, 0xe5, 0xb8, 0x36, 0x60, 0xce, 0x1f, 0x2f, 0x58, 0x73, 0x59, 0xac,
	0xe6, 0x9e, 0x98, 0x94, 0x8c, 0xaf, 0xab, 0xe1, 0xf5, 0xfc, 0xf2, 0x48, 0x67, 0xd6, 0x68, 0xc6,
	0xb5, 0xe9, 0x61, 0x9c, 0x74, 0xb1, 0x21, 0xe5, 0x36, 0x73, 0xde, 0x8e, 0x85, 0xea, 0x51, 0x79,
	0x38, 0x25, 0x83, 0xe7, 0x92, 0x95, 0xe4, 0xc1, 0xc5, 0x87, 0xcf, 0xae, 0xde, 0xe9, 0xb0, 0x3a,
	0xce, 0xd8, 0x20, 0x83, 0x97, 0x96, 0x3e, 0xff, 0xf9, 0x0b, 0x31, 0xf7, 0x64, 0x52, 0x5b, 0x39,
	0x34, 0x3d, 0x41, 0xd2, 0x8a, 0x95, 0x0f, 0x7d, 0x36, 0x2f, 0xe5, 0x9d, 0x54, 0x35, 0x2f, 0xe7,
	0x36, 0x87, 0x39, 0x57, 0x70, 0x48, 0x2d, 0xc8, 0x63, 0x1b, 0xe4, 0xd8, 0xba, 0x12, 0xc1, 0xa8,
	0xa3, 0xb2, 0x76, 0xf5, 0x31, 0xf8, 0x32, 0x51, 0x0b, 0xf2, 0xe8, 0x86, 0x47, 0xcd, 0x68, 0xe0,
	0xab, 0x44, 0x9d, 0x90, 0xc7, 0xd6, 0xd3, 0x14, 0xf3, 0x70, 0xfd, 0x3a, 0x51, 0x2b, 0xf2, 0xfe,
	0x06, 0xb9, 0xf5, 0x92, 0x07, 0xe4, 0x2d, 0x6b, 0xb6, 0x23, 0xdc, 0x71, 0x3d, 0xf2, 0x43, 0xcd,
	0x96, 0x1c, 0x7c, 0x93, 0xa8, 0x45, 0x39, 0xdf, 0xa0, 0xd0, 0x17, 0x1d, 0xc3, 0xb7, 0x89, 0x3a,
	0x25, 0x17, 0x5a, 0x58, 0x20, 0x4f, 0x43, 0xdf, 0x25, 0xea, 0xb4, 0x5c, 0x6c, 0x6a, 0xcf, 0x56,
	0x67, 0xd3, 0xe0, 0xf7, 0x89, 0x02, 0x79, 0xbc, 0x5e, 0x66, 0x6c, 0xc7, 0x58, 0xe1, 0x87, 0x44,
	0x2d, 0xc9, 0x93, 0xeb, 0x99, 0x47, 0x6d, 0xaa, 0x16, 0xe6, 0xe4, 0x03, 0x82, 0x1f, 0x13, 0x75,
	0x5c, 0x1e, 0xd9, 0xa9, 0x77, 0x0b, 0x34, 0xf0, 0x53, 0x4c, 0x89, 0x45, 0x79, 0x86, 0x1b, 0x03,
	0xb2, 0x29, 0x16, 0x70, 0xb3, 0xa6, 0xce, 0x48, 0xa8, 0xd3, 0x08, 0x4d, 0x13, 0xfd, 0x50, 0x3b,
	0x74, 0x9c, 0x55, 0x70, 0xab, 0xa6, 0xa4, 0x9c, 0xdb, 0xa2, 0xd2, 0x19, 0xf8, 0xb0, 0x16, 0x68,
	0xb5, 0x11, 0xaf, 0xf2, 0x00, 0x3d, 0xdc, 0xae, 0x85, 0xe1, 0x0d, 0xe2, 0x3a, 0x19, 0xdb, 0xb3,
	0x68, 0xe0, 0xa3, 0x98, 0xd0, 0x2d, 0xb0, 0xe9, 0xe9, 0xa0, 0x82, 0x8f, 0x6b, 0xea, 0xac, 0x3c,
	0xd5, 0xc1, 0x61, 0x4e, 0x5e, 0xfb, 0xaa, 0x85, 0xc6, 0x7a, 0x4c, 0x19, 0x3e, 0x89, 0xf1, 0xd9,
	0x94, 0x59, 0xfc, 0xd3, 0x9a, 0x3a, 0x29, 0xe5, 0x25, 0x6d, 0x5a, 0x78, 0xa3, 0xc4, 0x82, 0xe1,
	0x29, 0x11, 0x64, 0xe8, 0x3a, 0x3d, 0xd6, 0xed, 0x09, 0x34, 0xf0, 0xb4, 0x08, 0xe0, 0x9b, 0xba,
	0x1a, 0xc6, 0xca, 0x1b, 0xa5, 0xf5, 0x68, 0xe0, 0x19, 0x11, 0xf4, 0xdb, 0x22, 0xbf, 0x67, 0x8d,
	0x41, 0x07, 0xcf, 0x8a, 0x00, 0xa4, 0x41, 0x3c, 0x06, 0xfe, 0x9c, 0x88, 0xdc, 0x90, 0x07, 0x64,
	0x1a, 0xc4, 0xeb, 0x59, 0x46, 0xfb, 0x68, 0xe0, 0x79, 0xa1, 0x94, 0x3c, 0x11, 0x02, 0xd1, 0x29,
	0xbd, 0x97, 0x21, 0xbc, 0x20, 0x82, 0x57, 0x11, 0x7f, 0x70, 0x0b, 0x1d, 0xdb, 0x34, 0x7a, 0x34,
	0x9b, 0xf5, 0xa2, 0x08, 0x46, 0x4c, 0x20, 0x76, 0xec, 0x10, 0xa9, 0x64, 0x78, 0x29, 0x0e, 0xdc,
	0x20, 0xd7, 0xcb, 0x6c, 0xca, 0xf0, 0xb2, 0x50, 0xf3, 0xf2, 0xf0, 0x15, 0x72, 0x08, 0xaf, 0xc4,
	0xf4, 0x5d, 0x74, 0x7d, 0x1e, 0xcc, 0x7a, 0xbc, 0x2a, 0xd4, 0xb2, 0x54, 0x4d, 0x8f, 0x29, 0x39,
	0x63, 0x43, 0xfb, 0x2d, 0x6d, 0x33, 0x34, 0xf0, 0xda, 0x94, 0x5e, 0x46, 0xda, 0x74, 0x88, 0x76,
	0xb5, 0xef, 0x23, 0xbc, 0x2e, 0x82, 0x30, 0xdd, 0xd6, 0x4e, 0x88, 0x90, 0xeb, 0xc3, 0x1b, 0x42,
	0xdd, 0x27, 0x97, 0xba, 0xae, 0x28, 0xf3, 0xb1, 0xc3, 0x75, 0x34, 0x56, 0x77, 0xaa, 0x1c, 0xe1,
	0x4d, 0xa1, 0xce, 0xc9, 0xd3, 0x2d, 0xed, 0xfa, 0xd8, 0x20, 0x6e, 0x6b, 0xb6, 0x45, 0xcf, 0x46,
	0x6a, 0x6f, 0x89, 0x20, 0xfb, 0xe6, 0x41, 0x8e, 0x29, 0xeb, 0xbb, 0x66, 0xbe, 0x1d, 0xc1, 0xd4,
	0x6d, 0x31, 0xb6, 0x01, 0x67, 0xf2, 0xbf, 0x13, 0x5b, 0x75, 0x5d, 0xee, 0x29, 0xc5, 0xa2, 0x08,
	0x4d, 0x36, 0x1d, 0x5b, 0xae, 0xe0, 0x5d, 0x11, 0xf6, 0x69, 0x97, 0xd2, 0xeb, 0x68, 0xe0, 0xbd,
	0xa8, 0xee, 0xb8, 0xd9, 0x65, 0xcc, 0xd1, 0x19, 0x74, 0x69, 0x05, 0xef, 0x47, 0x2a, 0xdd, 0xbc,
	0xef, 0xb5, 0xc1, 0x19, 0xf3, 0x0f, 0x22, 0xf2, 0xbb, 0x99, 0xcf, 0x9e, 0x6e, 0xc6, 0x82, 0x0e,
	0x51, 0x5d, 0xbb, 0x6a, 0x82, 0xa1, 0x80, 0x5b, 0xd1, 0x90, 0xc9, 0x75, 0x1b, 0xb5, 0x41, 0xbf,
	0x65, 0x31, 0x33, 0xc5, 0x4c, 0x9d, 0xdb, 0x11, 0xe6, 0x8e, 0x63, 0xf4, 0x4e, 0x67, 0x6d, 0xf4,
	0x23, 0xf4, 0x9b, 0xde, 0x93, 0x87, 0x5f, 0xa2, 0xf6, 0x0d, 0xe2, 0x9d, 0x61, 0x9e, 0x61, 0xd8,
	0x18, 0x34, 0xf0, 0xab, 0x98, 0x6c, 0xd9, 0x15, 0xcd, 0xb8, 0xaf, 0x2b, 0xf8, 0x2d, 0xf2, 0x0f,
	0x75, 0x36, 0xc5, 0xae, 0xd3, 0x23, 0x6d, 0xb3, 0x28, 0xd8, 0xef, 0xb1, 0x7c, 0x92, 0x36, 0x75,
	0xfa, 0x0f, 0xa1, 0x2e, 0xc8, 0xe5, 0xed, 0x4e, 0xa7, 0x79, 0x0d, 0x7d, 0x61, 0xc9, 0x05, 0x95,
	0xa7, 0x36, 0xc0, 0x9f, 0x42, 0x9d, 0x97, 0x67, 0xae, 0x69, 0x6f, 0xb5, 0xe3, 0xf5, 0xac, 0xa0,
	0x06, 0xf6, 0x89, 0xad, 0x66, 0x2c, 0xe0, 0xaf, 0x09, 0xce, 0xa2, 0xec, 0xf5, 0x6c, 0x6a, 0xd1,
	0x71, 0x9b, 0xc9, 0xeb, 0x3e, 0xc2, 0xdf, 0x71, 0xcf, 0x77, 0x89, 0xf2, 0xcb, 0xc8, 0xd1, 0x02,
	0xf8, 0x47, 0x4c, 0x3e, 0xae, 0xcd, 0x03, 0x0e, 0x8a, 0x1a, 0xf8, 0x57, 0xa8, 0x8b, 0xf2, 0x81,
	0x06, 0xf2, 0x3e, 0xf9, 0xeb, 0xf7, 0xd8, 0xcd, 0xff, 0xc4, 0xde, 0x91, 0xf8, 0x3b, 0x7c, 0xe4,
	0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x9b, 0x1d, 0xc7, 0x53, 0x05, 0x00, 0x00,
}
