package tracing

import (
	"errors"

	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Code generated by chromedp-gen. DO NOT EDIT.

// MemoryDumpConfig configuration for memory dump. Used only when
// "memory-infra" category is enabled.
type MemoryDumpConfig struct{}

// TraceConfig [no description].
type TraceConfig struct {
	RecordMode           RecordMode        `json:"recordMode,omitempty"`           // Controls how the trace buffer stores data.
	EnableSampling       bool              `json:"enableSampling,omitempty"`       // Turns on JavaScript stack sampling.
	EnableSystrace       bool              `json:"enableSystrace,omitempty"`       // Turns on system tracing.
	EnableArgumentFilter bool              `json:"enableArgumentFilter,omitempty"` // Turns on argument filter.
	IncludedCategories   []string          `json:"includedCategories,omitempty"`   // Included category filters.
	ExcludedCategories   []string          `json:"excludedCategories,omitempty"`   // Excluded category filters.
	SyntheticDelays      []string          `json:"syntheticDelays,omitempty"`      // Configuration to synthesize the delays in tracing.
	MemoryDumpConfig     *MemoryDumpConfig `json:"memoryDumpConfig,omitempty"`     // Configuration for memory dump triggers. Used only when "memory-infra" category is enabled.
}

// StreamCompression compression type to use for traces returned via streams.
type StreamCompression string

// String returns the StreamCompression as string value.
func (t StreamCompression) String() string {
	return string(t)
}

// StreamCompression values.
const (
	StreamCompressionNone StreamCompression = "none"
	StreamCompressionGzip StreamCompression = "gzip"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t StreamCompression) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t StreamCompression) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *StreamCompression) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch StreamCompression(in.String()) {
	case StreamCompressionNone:
		*t = StreamCompressionNone
	case StreamCompressionGzip:
		*t = StreamCompressionGzip

	default:
		in.AddError(errors.New("unknown StreamCompression value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *StreamCompression) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// RecordMode controls how the trace buffer stores data.
type RecordMode string

// String returns the RecordMode as string value.
func (t RecordMode) String() string {
	return string(t)
}

// RecordMode values.
const (
	RecordModeRecordUntilFull        RecordMode = "recordUntilFull"
	RecordModeRecordContinuously     RecordMode = "recordContinuously"
	RecordModeRecordAsMuchAsPossible RecordMode = "recordAsMuchAsPossible"
	RecordModeEchoToConsole          RecordMode = "echoToConsole"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t RecordMode) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t RecordMode) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *RecordMode) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch RecordMode(in.String()) {
	case RecordModeRecordUntilFull:
		*t = RecordModeRecordUntilFull
	case RecordModeRecordContinuously:
		*t = RecordModeRecordContinuously
	case RecordModeRecordAsMuchAsPossible:
		*t = RecordModeRecordAsMuchAsPossible
	case RecordModeEchoToConsole:
		*t = RecordModeEchoToConsole

	default:
		in.AddError(errors.New("unknown RecordMode value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *RecordMode) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// TransferMode whether to report trace events as series of dataCollected
// events or to save trace to a stream (defaults to ReportEvents).
type TransferMode string

// String returns the TransferMode as string value.
func (t TransferMode) String() string {
	return string(t)
}

// TransferMode values.
const (
	TransferModeReportEvents   TransferMode = "ReportEvents"
	TransferModeReturnAsStream TransferMode = "ReturnAsStream"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t TransferMode) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t TransferMode) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *TransferMode) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch TransferMode(in.String()) {
	case TransferModeReportEvents:
		*t = TransferModeReportEvents
	case TransferModeReturnAsStream:
		*t = TransferModeReturnAsStream

	default:
		in.AddError(errors.New("unknown TransferMode value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *TransferMode) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
