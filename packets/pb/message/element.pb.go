// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/message/element.proto

package message

import (
	proto "github.com/RomiChan/protobuf/proto"
)

type Elem struct {
	Text                  *Text                  `protobuf:"bytes,1,opt"`
	Face                  *Face                  `protobuf:"bytes,2,opt"`
	OnlineImage           *OnlineImage           `protobuf:"bytes,3,opt"`
	NotOnlineImage        *NotOnlineImage        `protobuf:"bytes,4,opt"`
	TransElem             *TransElem             `protobuf:"bytes,5,opt"`
	MarketFace            *MarketFace            `protobuf:"bytes,6,opt"`
	CustomFace            *CustomFace            `protobuf:"bytes,8,opt"`
	ElemFlags2            *ElemFlags2            `protobuf:"bytes,9,opt"`
	RichMsg               *RichMsg               `protobuf:"bytes,12,opt"`
	GroupFile             *GroupFile             `protobuf:"bytes,13,opt"`
	ExtraInfo             *ExtraInfo             `protobuf:"bytes,16,opt"`
	VideoFile             *VideoFile             `protobuf:"bytes,19,opt"`
	AnonymousGroupMessage *AnonymousGroupMessage `protobuf:"bytes,21,opt"`
	QQWalletMsg           *QQWalletMsg           `protobuf:"bytes,24,opt"`
	CustomElem            *CustomElem            `protobuf:"bytes,31,opt"`
	GeneralFlags          *GeneralFlags          `protobuf:"bytes,37,opt"`
	SrcMsg                *SrcMsg                `protobuf:"bytes,45,opt"`
	LightAppElem          *LightAppElem          `protobuf:"bytes,51,opt"`
	CommonElem            *CommonElem            `protobuf:"bytes,53,opt"`
	_                     [0]func()
}

type AnonymousGroupMessage struct {
	Flags        int32  `protobuf:"varint,1,opt"`
	AnonId       []byte `protobuf:"bytes,2,opt"`
	AnonNick     []byte `protobuf:"bytes,3,opt"`
	HeadPortrait int32  `protobuf:"varint,4,opt"`
	ExpireTime   int32  `protobuf:"varint,5,opt"`
	BubbleId     int32  `protobuf:"varint,6,opt"`
	RankColor    []byte `protobuf:"bytes,7,opt"`
}

type CommonElem struct {
	ServiceType  int32  `protobuf:"varint,1,opt"`
	PbElem       []byte `protobuf:"bytes,2,opt"`
	BusinessType uint32 `protobuf:"varint,3,opt"`
}

type CustomElem struct {
	Desc     []byte `protobuf:"bytes,1,opt"`
	Data     []byte `protobuf:"bytes,2,opt"`
	EnumType int32  `protobuf:"varint,3,opt"`
	Ext      []byte `protobuf:"bytes,4,opt"`
	Sound    []byte `protobuf:"bytes,5,opt"`
}

type CustomFace struct {
	Guid        []byte               `protobuf:"bytes,1,opt"`
	FilePath    string               `protobuf:"bytes,2,opt"`
	Shortcut    string               `protobuf:"bytes,3,opt"`
	Buffer      []byte               `protobuf:"bytes,4,opt"`
	Flag        []byte               `protobuf:"bytes,5,opt"`
	OldData     []byte               `protobuf:"bytes,6,opt"`
	FileId      uint32               `protobuf:"varint,7,opt"`
	ServerIp    proto.Option[int32]  `protobuf:"varint,8,opt"`
	ServerPort  proto.Option[int32]  `protobuf:"varint,9,opt"`
	FileType    int32                `protobuf:"varint,10,opt"`
	Signature   []byte               `protobuf:"bytes,11,opt"`
	Useful      int32                `protobuf:"varint,12,opt"`
	Md5         []byte               `protobuf:"bytes,13,opt"`
	ThumbUrl    string               `protobuf:"bytes,14,opt"`
	BigUrl      string               `protobuf:"bytes,15,opt"`
	OrigUrl     string               `protobuf:"bytes,16,opt"`
	BizType     int32                `protobuf:"varint,17,opt"`
	RepeatIndex int32                `protobuf:"varint,18,opt"`
	RepeatImage int32                `protobuf:"varint,19,opt"`
	ImageType   int32                `protobuf:"varint,20,opt"`
	Index       int32                `protobuf:"varint,21,opt"`
	Width       int32                `protobuf:"varint,22,opt"`
	Height      int32                `protobuf:"varint,23,opt"`
	Source      int32                `protobuf:"varint,24,opt"`
	Size        uint32               `protobuf:"varint,25,opt"`
	Origin      int32                `protobuf:"varint,26,opt"`
	ThumbWidth  proto.Option[int32]  `protobuf:"varint,27,opt"`
	ThumbHeight proto.Option[int32]  `protobuf:"varint,28,opt"`
	ShowLen     int32                `protobuf:"varint,29,opt"`
	DownloadLen int32                `protobuf:"varint,30,opt"`
	X400Url     proto.Option[string] `protobuf:"bytes,31,opt"`
	X400Width   int32                `protobuf:"varint,32,opt"`
	X400Height  int32                `protobuf:"varint,33,opt"`
	PbReserve   []byte               `protobuf:"bytes,34,opt"`
}

type ElemFlags2 struct {
	ColorTextId      uint32               `protobuf:"varint,1,opt"`
	MsgId            uint64               `protobuf:"varint,2,opt"`
	WhisperSessionId uint32               `protobuf:"varint,3,opt"`
	PttChangeBit     uint32               `protobuf:"varint,4,opt"`
	VipStatus        uint32               `protobuf:"varint,5,opt"`
	CompatibleId     uint32               `protobuf:"varint,6,opt"`
	Insts            []*Instance          `protobuf:"bytes,7,rep"`
	MsgRptCnt        uint32               `protobuf:"varint,8,opt"`
	SrcInst          *Instance            `protobuf:"bytes,9,opt"`
	Longtitude       uint32               `protobuf:"varint,10,opt"`
	Latitude         uint32               `protobuf:"varint,11,opt"`
	CustomFont       uint32               `protobuf:"varint,12,opt"`
	PcSupportDef     *PcSupportDef        `protobuf:"bytes,13,opt"`
	CrmFlags         proto.Option[uint32] `protobuf:"varint,14,opt"`
}

type PcSupportDef struct {
	PcPtlBegin     uint32   `protobuf:"varint,1,opt"`
	PcPtlEnd       uint32   `protobuf:"varint,2,opt"`
	MacPtlBegin    uint32   `protobuf:"varint,3,opt"`
	MacPtlEnd      uint32   `protobuf:"varint,4,opt"`
	PtlsSupport    []int32  `protobuf:"varint,5,rep"`
	PtlsNotSupport []uint32 `protobuf:"varint,6,rep"`
}

type Instance struct {
	AppId  uint32 `protobuf:"varint,1,opt"`
	InstId uint32 `protobuf:"varint,2,opt"`
	_      [0]func()
}

type ExtraInfo struct {
	Nick          []byte `protobuf:"bytes,1,opt"`
	GroupCard     []byte `protobuf:"bytes,2,opt"`
	Level         int32  `protobuf:"varint,3,opt"`
	Flags         int32  `protobuf:"varint,4,opt"`
	GroupMask     int32  `protobuf:"varint,5,opt"`
	MsgTailId     int32  `protobuf:"varint,6,opt"`
	SenderTitle   []byte `protobuf:"bytes,7,opt"`
	ApnsTips      []byte `protobuf:"bytes,8,opt"`
	Uin           uint64 `protobuf:"varint,9,opt"`
	MsgStateFlag  int32  `protobuf:"varint,10,opt"`
	ApnsSoundType int32  `protobuf:"varint,11,opt"`
	NewGroupFlag  int32  `protobuf:"varint,12,opt"`
}

type Face struct {
	Index proto.Option[int32] `protobuf:"varint,1,opt"`
	Old   []byte              `protobuf:"bytes,2,opt"`
	Buf   []byte              `protobuf:"bytes,11,opt"`
}

type GeneralFlags struct {
	BubbleDiyTextId     int32                `protobuf:"varint,1,opt"`
	GroupFlagNew        int32                `protobuf:"varint,2,opt"`
	Uin                 uint64               `protobuf:"varint,3,opt"`
	RpId                []byte               `protobuf:"bytes,4,opt"`
	PrpFold             int32                `protobuf:"varint,5,opt"`
	LongTextFlag        int32                `protobuf:"varint,6,opt"`
	LongTextResId       proto.Option[string] `protobuf:"bytes,7,opt"`
	GroupType           int32                `protobuf:"varint,8,opt"`
	ToUinFlag           int32                `protobuf:"varint,9,opt"`
	GlamourLevel        int32                `protobuf:"varint,10,opt"`
	MemberLevel         int32                `protobuf:"varint,11,opt"`
	GroupRankSeq        uint64               `protobuf:"varint,12,opt"`
	OlympicTorch        int32                `protobuf:"varint,13,opt"`
	BabyqGuideMsgCookie []byte               `protobuf:"bytes,14,opt"`
	Uin32ExpertFlag     int32                `protobuf:"varint,15,opt"`
	BubbleSubId         int32                `protobuf:"varint,16,opt"`
	PendantId           uint64               `protobuf:"varint,17,opt"`
	RpIndex             []byte               `protobuf:"bytes,18,opt"`
	PbReserve           []byte               `protobuf:"bytes,19,opt"`
}

type GroupFile struct {
	Filename    []byte `protobuf:"bytes,1,opt"`
	FileSize    uint64 `protobuf:"varint,2,opt"`
	FileId      []byte `protobuf:"bytes,3,opt"`
	BatchId     []byte `protobuf:"bytes,4,opt"`
	FileKey     []byte `protobuf:"bytes,5,opt"`
	Mark        []byte `protobuf:"bytes,6,opt"`
	Sequence    uint64 `protobuf:"varint,7,opt"`
	BatchItemId []byte `protobuf:"bytes,8,opt"`
	FeedMsgTime int32  `protobuf:"varint,9,opt"`
	PbReserve   []byte `protobuf:"bytes,10,opt"`
}

type LightAppElem struct {
	Data     []byte `protobuf:"bytes,1,opt"`
	MsgResid []byte `protobuf:"bytes,2,opt"`
}

type MarketFace struct {
	FaceName    []byte `protobuf:"bytes,1,opt"`
	ItemType    int32  `protobuf:"varint,2,opt"`
	FaceInfo    int32  `protobuf:"varint,3,opt"`
	FaceId      []byte `protobuf:"bytes,4,opt"`
	TabId       int32  `protobuf:"varint,5,opt"`
	SubType     int32  `protobuf:"varint,6,opt"`
	Key         []byte `protobuf:"bytes,7,opt"`
	Param       []byte `protobuf:"bytes,8,opt"`
	MediaType   int32  `protobuf:"varint,9,opt"`
	ImageWidth  int32  `protobuf:"varint,10,opt"`
	ImageHeight int32  `protobuf:"varint,11,opt"`
	Mobileparam []byte `protobuf:"bytes,12,opt"`
	PbReserve   []byte `protobuf:"bytes,13,opt"`
}

type NotOnlineImage struct {
	FilePath       string `protobuf:"bytes,1,opt"`
	FileLen        uint32 `protobuf:"varint,2,opt"`
	DownloadPath   string `protobuf:"bytes,3,opt"`
	OldVerSendFile []byte `protobuf:"bytes,4,opt"`
	ImgType        int32  `protobuf:"varint,5,opt"`
	PreviewsImage  []byte `protobuf:"bytes,6,opt"`
	PicMd5         []byte `protobuf:"bytes,7,opt"`
	PicHeight      uint32 `protobuf:"varint,8,opt"`
	PicWidth       uint32 `protobuf:"varint,9,opt"`
	ResId          string `protobuf:"bytes,10,opt"`
	Flag           []byte `protobuf:"bytes,11,opt"`
	ThumbUrl       string `protobuf:"bytes,12,opt"`
	Original       int32  `protobuf:"varint,13,opt"`
	BigUrl         string `protobuf:"bytes,14,opt"`
	OrigUrl        string `protobuf:"bytes,15,opt"`
	BizType        int32  `protobuf:"varint,16,opt"`
	Result         int32  `protobuf:"varint,17,opt"`
	Index          int32  `protobuf:"varint,18,opt"`
	OpFaceBuf      []byte `protobuf:"bytes,19,opt"`
	OldPicMd5      bool   `protobuf:"varint,20,opt"`
	ThumbWidth     int32  `protobuf:"varint,21,opt"`
	ThumbHeight    int32  `protobuf:"varint,22,opt"`
	FileId         int32  `protobuf:"varint,23,opt"`
	ShowLen        uint32 `protobuf:"varint,24,opt"`
	DownloadLen    uint32 `protobuf:"varint,25,opt"`
	PbRes          []byte `protobuf:"bytes,26,opt"`
}

type OnlineImage struct {
	Guid           []byte `protobuf:"bytes,1,opt"`
	FilePath       []byte `protobuf:"bytes,2,opt"`
	OldVerSendFile []byte `protobuf:"bytes,3,opt"`
}

type QQWalletMsg struct {
	Type *QQWalletAioBody `protobuf:"bytes,1,opt"`
	_    [0]func()
}

type QQWalletAioBody struct {
	SendUin     uint64           `protobuf:"varint,1,opt"`
	Sender      *QQWalletAioElem `protobuf:"bytes,2,opt"`
	Receiver    *QQWalletAioElem `protobuf:"bytes,3,opt"`
	ChannelId   int32            `protobuf:"zigzag32,4,opt"`
	TemplateId  int32            `protobuf:"zigzag32,5,opt"`
	Resend      uint32           `protobuf:"varint,6,opt"`
	MsgPriority uint32           `protobuf:"varint,7,opt"`
	RedType     int32            `protobuf:"zigzag32,8,opt"`
	BillNo      []byte           `protobuf:"bytes,9,opt"`
	AuthKey     []byte           `protobuf:"bytes,10,opt"`
	SessionType int32            `protobuf:"zigzag32,11,opt"`
	MsgType     int32            `protobuf:"zigzag32,12,opt"`
	EnvelOpeId  int32            `protobuf:"zigzag32,13,opt"`
	Name        []byte           `protobuf:"bytes,14,opt"`
	ConfType    int32            `protobuf:"zigzag32,15,opt"`
	MsgFrom     int32            `protobuf:"zigzag32,16,opt"`
	PcBody      []byte           `protobuf:"bytes,17,opt"`
	Index       []byte           `protobuf:"bytes,18,opt"`
	RedChannel  uint32           `protobuf:"varint,19,opt"`
	GrapUin     uint64           `protobuf:"varint,20,opt"`
	PbReserve   []byte           `protobuf:"bytes,21,opt"`
}

type QQWalletAioElem struct {
	Background      uint32 `protobuf:"varint,1,opt"`
	Icon            uint32 `protobuf:"varint,2,opt"`
	Title           string `protobuf:"bytes,3,opt"`
	Subtitle        string `protobuf:"bytes,4,opt"`
	Content         string `protobuf:"bytes,5,opt"`
	LinkUrl         []byte `protobuf:"bytes,6,opt"`
	BlackStripe     []byte `protobuf:"bytes,7,opt"`
	Notice          []byte `protobuf:"bytes,8,opt"`
	TitleColor      uint32 `protobuf:"varint,9,opt"`
	SubtitleColor   uint32 `protobuf:"varint,10,opt"`
	ActionsPriority []byte `protobuf:"bytes,11,opt"`
	JumpUrl         []byte `protobuf:"bytes,12,opt"`
	NativeIos       []byte `protobuf:"bytes,13,opt"`
	NativeAndroid   []byte `protobuf:"bytes,14,opt"`
	IconUrl         []byte `protobuf:"bytes,15,opt"`
	ContentColor    uint32 `protobuf:"varint,16,opt"`
	ContentBgColor  uint32 `protobuf:"varint,17,opt"`
	AioImageLeft    []byte `protobuf:"bytes,18,opt"`
	AioImageRight   []byte `protobuf:"bytes,19,opt"`
	CftImage        []byte `protobuf:"bytes,20,opt"`
	PbReserve       []byte `protobuf:"bytes,21,opt"`
}

type RedBagInfo struct {
	RedBagType proto.Option[uint32] `protobuf:"varint,1,opt"`
	_          [0]func()
}

type RichMsg struct {
	Template1 []byte               `protobuf:"bytes,1,opt"`
	ServiceId proto.Option[int32]  `protobuf:"varint,2,opt"`
	MsgResId  []byte               `protobuf:"bytes,3,opt"`
	Rand      proto.Option[int32]  `protobuf:"varint,4,opt"`
	Seq       proto.Option[uint32] `protobuf:"varint,5,opt"`
}

type SrcMsg struct {
	OrigSeqs  []uint32             `protobuf:"varint,1,rep"`
	SenderUin uint64               `protobuf:"varint,2,opt"`
	Time      proto.Option[int32]  `protobuf:"varint,3,opt"`
	Flag      proto.Option[int32]  `protobuf:"varint,4,opt"`
	Elems     []*Elem              `protobuf:"bytes,5,rep"`
	Type      proto.Option[int32]  `protobuf:"varint,6,opt"`
	RichMsg   []byte               `protobuf:"bytes,7,opt"`
	PbReserve []byte               `protobuf:"bytes,8,opt"`
	SourceMsg []byte               `protobuf:"bytes,9,opt"`
	ToUin     proto.Option[uint64] `protobuf:"varint,10,opt"`
	TroopName []byte               `protobuf:"bytes,11,opt"`
}

type Preserve struct {
	MessageId      uint64 `protobuf:"varint,3,opt"`
	SenderUid      string `protobuf:"bytes,6,opt"`
	ReceiverUid    string `protobuf:"bytes,7,opt"`
	ClientSequence uint32 `protobuf:"varint,8,opt"`
	_              [0]func()
}

type Text struct {
	Str       proto.Option[string] `protobuf:"bytes,1,opt"`
	Lint      proto.Option[string] `protobuf:"bytes,2,opt"`
	Attr6Buf  []byte               `protobuf:"bytes,3,opt"`
	Attr7Buf  []byte               `protobuf:"bytes,4,opt"`
	Buf       []byte               `protobuf:"bytes,11,opt"`
	PbReserve []byte               `protobuf:"bytes,12,opt"`
}

type TransElem struct {
	ElemType  int32  `protobuf:"varint,1,opt"`
	ElemValue []byte `protobuf:"bytes,2,opt"`
}

type VideoFile struct {
	FileUuid               string   `protobuf:"bytes,1,opt"`
	FileMd5                []byte   `protobuf:"bytes,2,opt"`
	FileName               string   `protobuf:"bytes,3,opt"`
	FileFormat             int32    `protobuf:"varint,4,opt"`
	FileTime               int32    `protobuf:"varint,5,opt"`
	FileSize               int32    `protobuf:"varint,6,opt"`
	ThumbWidth             int32    `protobuf:"varint,7,opt"`
	ThumbHeight            int32    `protobuf:"varint,8,opt"`
	ThumbFileMd5           []byte   `protobuf:"bytes,9,opt"`
	Source                 []byte   `protobuf:"bytes,10,opt"`
	ThumbFileSize          int32    `protobuf:"varint,11,opt"`
	BusiType               int32    `protobuf:"varint,12,opt"`
	FromChatType           int32    `protobuf:"varint,13,opt"`
	ToChatType             int32    `protobuf:"varint,14,opt"`
	BoolSupportProgressive bool     `protobuf:"varint,15,opt"`
	FileWidth              int32    `protobuf:"varint,16,opt"`
	FileHeight             int32    `protobuf:"varint,17,opt"`
	SubBusiType            int32    `protobuf:"varint,18,opt"`
	VideoAttr              int32    `protobuf:"varint,19,opt"`
	BytesThumbFileUrls     [][]byte `protobuf:"bytes,20,rep"`
	BytesVideoFileUrls     [][]byte `protobuf:"bytes,21,rep"`
	ThumbDownloadFlag      int32    `protobuf:"varint,22,opt"`
	VideoDownloadFlag      int32    `protobuf:"varint,23,opt"`
	PbReserve              []byte   `protobuf:"bytes,24,opt"`
}

type CustomFaceExtra struct {
	Hash proto.Option[string] `protobuf:"bytes,31,opt"`
	_    [0]func()
}

type FaceExtra struct {
	FaceId proto.Option[int32] `protobuf:"varint,1,opt"`
	_      [0]func()
}

type ImageExtra struct {
	Field85 uint32 `protobuf:"varint,85,opt"`
	_       [0]func()
}

type MentionExtra struct {
	Type   proto.Option[int32]  `protobuf:"varint,3,opt"`
	Uin    proto.Option[uint32] `protobuf:"varint,4,opt"`
	Field5 proto.Option[int32]  `protobuf:"varint,5,opt"`
	Uid    proto.Option[string] `protobuf:"bytes,9,opt"`
	_      [0]func()
}

type QFaceExtra struct {
	Field1  proto.Option[string] `protobuf:"bytes,1,opt"`
	Field2  proto.Option[string] `protobuf:"bytes,2,opt"`
	FaceId  proto.Option[int32]  `protobuf:"varint,3,opt"`
	Field4  proto.Option[int32]  `protobuf:"varint,4,opt"`
	Field5  proto.Option[int32]  `protobuf:"varint,5,opt"`
	Field6  proto.Option[string] `protobuf:"bytes,6,opt"`
	Preview proto.Option[string] `protobuf:"bytes,7,opt"`
	Field9  proto.Option[int32]  `protobuf:"varint,9,opt"`
	_       [0]func()
}