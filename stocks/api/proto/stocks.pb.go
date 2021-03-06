// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/stocks.proto

package stocks

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PopularRequest struct {
	Sector               string   `protobuf:"bytes,1,opt,name=sector,proto3" json:"sector,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PopularRequest) Reset()         { *m = PopularRequest{} }
func (m *PopularRequest) String() string { return proto.CompactTextString(m) }
func (*PopularRequest) ProtoMessage()    {}
func (*PopularRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f15b1b797e78b2ce, []int{0}
}

func (m *PopularRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PopularRequest.Unmarshal(m, b)
}
func (m *PopularRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PopularRequest.Marshal(b, m, deterministic)
}
func (m *PopularRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PopularRequest.Merge(m, src)
}
func (m *PopularRequest) XXX_Size() int {
	return xxx_messageInfo_PopularRequest.Size(m)
}
func (m *PopularRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PopularRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PopularRequest proto.InternalMessageInfo

func (m *PopularRequest) GetSector() string {
	if m != nil {
		return m.Sector
	}
	return ""
}

type SearchRequest struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f15b1b797e78b2ce, []int{1}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SearchRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ListResponse struct {
	Stocks               []*Stock `protobuf:"bytes,1,rep,name=stocks,proto3" json:"stocks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f15b1b797e78b2ce, []int{2}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetStocks() []*Stock {
	if m != nil {
		return m.Stocks
	}
	return nil
}

type Request struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_f15b1b797e78b2ce, []int{3}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type Stock struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Symbol               string   `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Exchange             string   `protobuf:"bytes,4,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Type                 string   `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Region               string   `protobuf:"bytes,6,opt,name=region,proto3" json:"region,omitempty"`
	Currency             string   `protobuf:"bytes,7,opt,name=currency,proto3" json:"currency,omitempty"`
	ProfilePictureUrl    string   `protobuf:"bytes,8,opt,name=profile_picture_url,json=profilePictureUrl,proto3" json:"profile_picture_url,omitempty"`
	Following            bool     `protobuf:"varint,9,opt,name=following,proto3" json:"following,omitempty"`
	MarketCap            float32  `protobuf:"fixed32,11,opt,name=market_cap,json=marketCap,proto3" json:"market_cap,omitempty"`
	Week_52High          float32  `protobuf:"fixed32,12,opt,name=week_52_high,json=week52High,proto3" json:"week_52_high,omitempty"`
	Week_52Low           float32  `protobuf:"fixed32,13,opt,name=week_52_low,json=week52Low,proto3" json:"week_52_low,omitempty"`
	Avg_10Volume         float32  `protobuf:"fixed32,14,opt,name=avg_10_volume,json=avg10Volume,proto3" json:"avg_10_volume,omitempty"`
	TtmEps               float32  `protobuf:"fixed32,15,opt,name=ttm_eps,json=ttmEps,proto3" json:"ttm_eps,omitempty"`
	TtmDividendRate      float32  `protobuf:"fixed32,16,opt,name=ttm_dividend_rate,json=ttmDividendRate,proto3" json:"ttm_dividend_rate,omitempty"`
	DividendYield        float32  `protobuf:"fixed32,17,opt,name=dividend_yield,json=dividendYield,proto3" json:"dividend_yield,omitempty"`
	NextDividendDate     string   `protobuf:"bytes,18,opt,name=next_dividend_date,json=nextDividendDate,proto3" json:"next_dividend_date,omitempty"`
	ExDividendDate       string   `protobuf:"bytes,19,opt,name=ex_dividend_date,json=exDividendDate,proto3" json:"ex_dividend_date,omitempty"`
	NextEarningsDate     string   `protobuf:"bytes,20,opt,name=next_earnings_date,json=nextEarningsDate,proto3" json:"next_earnings_date,omitempty"`
	PeRatio              float32  `protobuf:"fixed32,21,opt,name=pe_ratio,json=peRatio,proto3" json:"pe_ratio,omitempty"`
	Beta                 float32  `protobuf:"fixed32,22,opt,name=beta,proto3" json:"beta,omitempty"`
	YtdChangePercent     float32  `protobuf:"fixed32,23,opt,name=ytd_change_percent,json=ytdChangePercent,proto3" json:"ytd_change_percent,omitempty"`
	Month_1ChangePercent float32  `protobuf:"fixed32,24,opt,name=month_1_change_percent,json=month1ChangePercent,proto3" json:"month_1_change_percent,omitempty"`
	Day_5ChangePercent   float32  `protobuf:"fixed32,25,opt,name=day_5_change_percent,json=day5ChangePercent,proto3" json:"day_5_change_percent,omitempty"`
	PrevDayOpen          float32  `protobuf:"fixed32,26,opt,name=prev_day_open,json=prevDayOpen,proto3" json:"prev_day_open,omitempty"`
	PrevDayClose         float32  `protobuf:"fixed32,27,opt,name=prev_day_close,json=prevDayClose,proto3" json:"prev_day_close,omitempty"`
	PrevDayHigh          float32  `protobuf:"fixed32,28,opt,name=prev_day_high,json=prevDayHigh,proto3" json:"prev_day_high,omitempty"`
	PrevDayLow           float32  `protobuf:"fixed32,29,opt,name=prev_day_low,json=prevDayLow,proto3" json:"prev_day_low,omitempty"`
	PrevDayVolume        float32  `protobuf:"fixed32,30,opt,name=prev_day_volume,json=prevDayVolume,proto3" json:"prev_day_volume,omitempty"`
	CurrentPrice         float32  `protobuf:"fixed32,31,opt,name=current_price,json=currentPrice,proto3" json:"current_price,omitempty"`
	Sector               string   `protobuf:"bytes,32,opt,name=sector,proto3" json:"sector,omitempty"`
	Industry             string   `protobuf:"bytes,33,opt,name=industry,proto3" json:"industry,omitempty"`
	Website              string   `protobuf:"bytes,34,opt,name=website,proto3" json:"website,omitempty"`
	Description          string   `protobuf:"bytes,35,opt,name=description,proto3" json:"description,omitempty"`
	Color                string   `protobuf:"bytes,36,opt,name=color,proto3" json:"color,omitempty"`
	Summary              string   `protobuf:"bytes,37,opt,name=summary,proto3" json:"summary,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Stock) Reset()         { *m = Stock{} }
func (m *Stock) String() string { return proto.CompactTextString(m) }
func (*Stock) ProtoMessage()    {}
func (*Stock) Descriptor() ([]byte, []int) {
	return fileDescriptor_f15b1b797e78b2ce, []int{4}
}

func (m *Stock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stock.Unmarshal(m, b)
}
func (m *Stock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stock.Marshal(b, m, deterministic)
}
func (m *Stock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stock.Merge(m, src)
}
func (m *Stock) XXX_Size() int {
	return xxx_messageInfo_Stock.Size(m)
}
func (m *Stock) XXX_DiscardUnknown() {
	xxx_messageInfo_Stock.DiscardUnknown(m)
}

var xxx_messageInfo_Stock proto.InternalMessageInfo

func (m *Stock) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Stock) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Stock) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Stock) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *Stock) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Stock) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *Stock) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *Stock) GetProfilePictureUrl() string {
	if m != nil {
		return m.ProfilePictureUrl
	}
	return ""
}

func (m *Stock) GetFollowing() bool {
	if m != nil {
		return m.Following
	}
	return false
}

func (m *Stock) GetMarketCap() float32 {
	if m != nil {
		return m.MarketCap
	}
	return 0
}

func (m *Stock) GetWeek_52High() float32 {
	if m != nil {
		return m.Week_52High
	}
	return 0
}

func (m *Stock) GetWeek_52Low() float32 {
	if m != nil {
		return m.Week_52Low
	}
	return 0
}

func (m *Stock) GetAvg_10Volume() float32 {
	if m != nil {
		return m.Avg_10Volume
	}
	return 0
}

func (m *Stock) GetTtmEps() float32 {
	if m != nil {
		return m.TtmEps
	}
	return 0
}

func (m *Stock) GetTtmDividendRate() float32 {
	if m != nil {
		return m.TtmDividendRate
	}
	return 0
}

func (m *Stock) GetDividendYield() float32 {
	if m != nil {
		return m.DividendYield
	}
	return 0
}

func (m *Stock) GetNextDividendDate() string {
	if m != nil {
		return m.NextDividendDate
	}
	return ""
}

func (m *Stock) GetExDividendDate() string {
	if m != nil {
		return m.ExDividendDate
	}
	return ""
}

func (m *Stock) GetNextEarningsDate() string {
	if m != nil {
		return m.NextEarningsDate
	}
	return ""
}

func (m *Stock) GetPeRatio() float32 {
	if m != nil {
		return m.PeRatio
	}
	return 0
}

func (m *Stock) GetBeta() float32 {
	if m != nil {
		return m.Beta
	}
	return 0
}

func (m *Stock) GetYtdChangePercent() float32 {
	if m != nil {
		return m.YtdChangePercent
	}
	return 0
}

func (m *Stock) GetMonth_1ChangePercent() float32 {
	if m != nil {
		return m.Month_1ChangePercent
	}
	return 0
}

func (m *Stock) GetDay_5ChangePercent() float32 {
	if m != nil {
		return m.Day_5ChangePercent
	}
	return 0
}

func (m *Stock) GetPrevDayOpen() float32 {
	if m != nil {
		return m.PrevDayOpen
	}
	return 0
}

func (m *Stock) GetPrevDayClose() float32 {
	if m != nil {
		return m.PrevDayClose
	}
	return 0
}

func (m *Stock) GetPrevDayHigh() float32 {
	if m != nil {
		return m.PrevDayHigh
	}
	return 0
}

func (m *Stock) GetPrevDayLow() float32 {
	if m != nil {
		return m.PrevDayLow
	}
	return 0
}

func (m *Stock) GetPrevDayVolume() float32 {
	if m != nil {
		return m.PrevDayVolume
	}
	return 0
}

func (m *Stock) GetCurrentPrice() float32 {
	if m != nil {
		return m.CurrentPrice
	}
	return 0
}

func (m *Stock) GetSector() string {
	if m != nil {
		return m.Sector
	}
	return ""
}

func (m *Stock) GetIndustry() string {
	if m != nil {
		return m.Industry
	}
	return ""
}

func (m *Stock) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *Stock) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Stock) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func (m *Stock) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_f15b1b797e78b2ce, []int{5}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Response struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Stock                *Stock   `protobuf:"bytes,2,opt,name=stock,proto3" json:"stock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_f15b1b797e78b2ce, []int{6}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Response) GetStock() *Stock {
	if m != nil {
		return m.Stock
	}
	return nil
}

func init() {
	proto.RegisterType((*PopularRequest)(nil), "PopularRequest")
	proto.RegisterType((*SearchRequest)(nil), "SearchRequest")
	proto.RegisterType((*ListResponse)(nil), "ListResponse")
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Stock)(nil), "Stock")
	proto.RegisterType((*Error)(nil), "Error")
	proto.RegisterType((*Response)(nil), "Response")
}

func init() { proto.RegisterFile("proto/stocks.proto", fileDescriptor_f15b1b797e78b2ce) }

var fileDescriptor_f15b1b797e78b2ce = []byte{
	// 868 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0xdd, 0x52, 0xdb, 0x46,
	0x14, 0xae, 0x43, 0xfc, 0x77, 0xb0, 0x0d, 0x2c, 0x94, 0x6c, 0x28, 0x10, 0x47, 0x49, 0x3a, 0xee,
	0xcf, 0x38, 0x81, 0x8c, 0xaf, 0x7a, 0x09, 0xa4, 0xbd, 0xc8, 0x4c, 0x19, 0x65, 0xd2, 0x99, 0x5e,
	0x69, 0x84, 0x74, 0x90, 0x77, 0x90, 0xb4, 0xca, 0xee, 0xca, 0x46, 0x6f, 0xd5, 0x67, 0xe8, 0x93,
	0x75, 0xf6, 0xac, 0x64, 0x6c, 0xda, 0xe6, 0x6e, 0xbf, 0x9f, 0xf3, 0x49, 0x82, 0xb3, 0x9f, 0x81,
	0x15, 0x4a, 0x1a, 0xf9, 0x56, 0x1b, 0x19, 0xdd, 0xe9, 0x29, 0x01, 0x6f, 0x02, 0xa3, 0x6b, 0x59,
	0x94, 0x69, 0xa8, 0x7c, 0xfc, 0x52, 0xa2, 0x36, 0xec, 0x10, 0x3a, 0x1a, 0x23, 0x23, 0x15, 0x6f,
	0x8d, 0x5b, 0x93, 0xbe, 0x5f, 0x23, 0xef, 0x17, 0x18, 0x7e, 0xc2, 0x50, 0x45, 0xf3, 0xc6, 0x78,
	0x00, 0xed, 0x2f, 0x25, 0xaa, 0xaa, 0xf6, 0x39, 0x60, 0xd9, 0x54, 0x64, 0xc2, 0xf0, 0x27, 0xe3,
	0xd6, 0xa4, 0xed, 0x3b, 0xe0, 0x4d, 0x61, 0xf0, 0x51, 0x68, 0xe3, 0xa3, 0x2e, 0x64, 0xae, 0x91,
	0x9d, 0x42, 0xc7, 0xbd, 0x06, 0x6f, 0x8d, 0xb7, 0x26, 0xdb, 0xe7, 0x9d, 0xe9, 0x27, 0x0b, 0xfd,
	0x9a, 0xf5, 0x4e, 0xa0, 0xdb, 0x3c, 0x86, 0xc1, 0xd3, 0xb2, 0x14, 0x71, 0xfd, 0x14, 0x3a, 0x7b,
	0x7f, 0xf5, 0xa1, 0x4d, 0x03, 0xff, 0xa5, 0x5a, 0x2e, 0x0f, 0x33, 0xa4, 0x37, 0xe8, 0xfb, 0x74,
	0xa6, 0xaf, 0xaa, 0xb2, 0x1b, 0x99, 0xf2, 0xad, 0xfa, 0xab, 0x08, 0xb1, 0x23, 0xe8, 0xe1, 0x7d,
	0x34, 0x0f, 0xf3, 0x04, 0xf9, 0x53, 0x52, 0x56, 0xd8, 0xe6, 0x98, 0xaa, 0x40, 0xde, 0x76, 0x39,
	0xf6, 0x6c, 0x73, 0x14, 0x26, 0x42, 0xe6, 0xbc, 0xe3, 0x72, 0x1c, 0xb2, 0x39, 0x51, 0xa9, 0x14,
	0xe6, 0x51, 0xc5, 0xbb, 0x2e, 0xa7, 0xc1, 0x6c, 0x0a, 0xfb, 0x85, 0x92, 0xb7, 0x22, 0xc5, 0xa0,
	0x10, 0x91, 0x29, 0x15, 0x06, 0xa5, 0x4a, 0x79, 0x8f, 0x6c, 0x7b, 0xb5, 0x74, 0xed, 0x94, 0xcf,
	0x2a, 0x65, 0xc7, 0xd0, 0xbf, 0x95, 0x69, 0x2a, 0x97, 0x22, 0x4f, 0x78, 0x7f, 0xdc, 0x9a, 0xf4,
	0xfc, 0x07, 0x82, 0x9d, 0x00, 0x64, 0xa1, 0xba, 0x43, 0x13, 0x44, 0x61, 0xc1, 0xb7, 0xc7, 0xad,
	0xc9, 0x13, 0xbf, 0xef, 0x98, 0x8b, 0xb0, 0x60, 0x63, 0x18, 0x2c, 0x11, 0xef, 0x82, 0xd9, 0x79,
	0x30, 0x17, 0xc9, 0x9c, 0x0f, 0xc8, 0x00, 0x96, 0x9b, 0x9d, 0xff, 0x26, 0x92, 0x39, 0x3b, 0x85,
	0xed, 0xc6, 0x91, 0xca, 0x25, 0x1f, 0xba, 0x04, 0x67, 0xf8, 0x28, 0x97, 0xcc, 0x83, 0x61, 0xb8,
	0x48, 0x82, 0xb3, 0x77, 0xc1, 0x42, 0xa6, 0x65, 0x86, 0x7c, 0x44, 0x8e, 0xed, 0x70, 0x91, 0x9c,
	0xbd, 0xfb, 0x83, 0x28, 0xf6, 0x0c, 0xba, 0xc6, 0x64, 0x01, 0x16, 0x9a, 0xef, 0x90, 0xda, 0x31,
	0x26, 0xbb, 0x2a, 0x34, 0xfb, 0x11, 0xf6, 0xac, 0x10, 0x8b, 0x85, 0x88, 0x31, 0x8f, 0x03, 0x15,
	0x1a, 0xe4, 0xbb, 0x64, 0xd9, 0x31, 0x26, 0xbb, 0xac, 0x79, 0x3f, 0x34, 0xc8, 0xde, 0xc0, 0x68,
	0xe5, 0xab, 0x04, 0xa6, 0x31, 0xdf, 0x23, 0xe3, 0xb0, 0x61, 0xff, 0xb4, 0x24, 0xfb, 0x19, 0x58,
	0x8e, 0xf7, 0xe6, 0x21, 0x33, 0xb6, 0x99, 0x8c, 0xfe, 0x7a, 0xbb, 0x56, 0x69, 0x42, 0x2f, 0x6d,
	0xe8, 0x04, 0x76, 0xf1, 0xfe, 0x91, 0x77, 0x9f, 0xbc, 0x23, 0xbc, 0xdf, 0x70, 0x36, 0xb9, 0x18,
	0xaa, 0x5c, 0xe4, 0x89, 0x76, 0xde, 0x83, 0x87, 0xdc, 0xab, 0x5a, 0x20, 0xf7, 0x73, 0xe8, 0x15,
	0x68, 0x3f, 0x47, 0x48, 0xfe, 0x2d, 0xbd, 0x66, 0xb7, 0x40, 0xdf, 0x42, 0xbb, 0x27, 0x37, 0x68,
	0x42, 0x7e, 0x48, 0x34, 0x9d, 0x6d, 0x78, 0x65, 0xe2, 0xc0, 0x6d, 0x52, 0x50, 0xa0, 0x8a, 0x30,
	0x37, 0xfc, 0x19, 0x39, 0x76, 0x2b, 0x13, 0x5f, 0x90, 0x70, 0xed, 0x78, 0xf6, 0x1e, 0x0e, 0x33,
	0x99, 0x9b, 0x79, 0x70, 0xf6, 0x78, 0x82, 0xd3, 0xc4, 0x3e, 0xa9, 0x67, 0x9b, 0x43, 0x6f, 0xe1,
	0x20, 0x0e, 0xab, 0x60, 0xf6, 0x78, 0xe4, 0x39, 0x8d, 0xec, 0xc5, 0x61, 0x35, 0xdb, 0x1c, 0xf0,
	0x60, 0x58, 0x28, 0x5c, 0x04, 0x76, 0x4a, 0x16, 0x98, 0xf3, 0x23, 0xf7, 0x8f, 0xb5, 0xe4, 0x65,
	0x58, 0xfd, 0x5e, 0x60, 0xce, 0x5e, 0xc3, 0x68, 0xe5, 0x89, 0x52, 0xa9, 0x91, 0x7f, 0x47, 0xa6,
	0x41, 0x6d, 0xba, 0xb0, 0xdc, 0x46, 0x12, 0x6d, 0xd9, 0xf1, 0x46, 0x12, 0xad, 0xd9, 0x18, 0x06,
	0x2b, 0x8f, 0xdd, 0xb3, 0x13, 0xb7, 0x88, 0xb5, 0xc5, 0x2e, 0xda, 0xf7, 0xb0, 0xb3, 0x72, 0xd4,
	0xab, 0x76, 0xea, 0x16, 0xa0, 0x36, 0xd5, 0xcb, 0xf6, 0x0a, 0x86, 0xee, 0x2e, 0x99, 0xa0, 0x50,
	0x22, 0x42, 0xfe, 0xc2, 0xbd, 0x52, 0x4d, 0x5e, 0x5b, 0x6e, 0xad, 0xb6, 0xc6, 0xeb, 0xb5, 0x65,
	0x2f, 0xa6, 0xc8, 0xe3, 0x52, 0x1b, 0x55, 0xf1, 0x97, 0xee, 0x62, 0x36, 0x98, 0x71, 0xe8, 0x2e,
	0xf1, 0x46, 0x0b, 0x83, 0xdc, 0x23, 0xa9, 0x81, 0x6c, 0x0c, 0xdb, 0x31, 0xea, 0x48, 0x89, 0xc2,
	0xd8, 0xbb, 0xfe, 0x8a, 0xd4, 0x75, 0xca, 0xf6, 0x5c, 0x24, 0x53, 0xa9, 0xf8, 0x6b, 0xd7, 0x7e,
	0x04, 0x6c, 0xa2, 0x2e, 0xb3, 0x2c, 0x54, 0x15, 0x7f, 0xe3, 0x12, 0x6b, 0xe8, 0xcd, 0xa0, 0x7d,
	0xa5, 0x94, 0x54, 0x76, 0x5b, 0x22, 0x19, 0x23, 0x35, 0x56, 0xdb, 0xa7, 0xb3, 0x1d, 0xcb, 0x50,
	0xeb, 0x30, 0x69, 0x4a, 0xab, 0x81, 0xde, 0x07, 0xe8, 0xad, 0x4a, 0xf3, 0x18, 0xda, 0x68, 0x23,
	0x68, 0xd4, 0x76, 0x26, 0x05, 0xfa, 0x8e, 0xb4, 0x2a, 0x95, 0x27, 0x25, 0x3c, 0x34, 0xaa, 0x23,
	0xcf, 0xff, 0x6e, 0x41, 0x87, 0x08, 0xcd, 0x7e, 0x82, 0x6e, 0x5d, 0xf9, 0x6c, 0x67, 0xba, 0x59,
	0xfe, 0x47, 0xc3, 0xe9, 0x7a, 0x4d, 0x7b, 0xdf, 0xb0, 0x1f, 0xa0, 0xe3, 0x5a, 0x9f, 0x8d, 0xa6,
	0x1b, 0xf5, 0xff, 0x6f, 0xeb, 0x31, 0x6c, 0xfd, 0x8a, 0x86, 0xf5, 0xa6, 0x8d, 0xa3, 0x3f, 0x5d,
	0x53, 0x5f, 0x40, 0xe7, 0x03, 0x75, 0xd8, 0xff, 0x19, 0x5e, 0x42, 0xef, 0x73, 0x7e, 0xfb, 0x35,
	0xcb, 0x4d, 0x87, 0x7e, 0xb3, 0xde, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0x47, 0xb6, 0x9d, 0x5a,
	0xc9, 0x06, 0x00, 0x00,
}
