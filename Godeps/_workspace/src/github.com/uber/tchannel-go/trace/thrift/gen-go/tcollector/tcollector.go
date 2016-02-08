// Autogenerated by Thrift Compiler (1.0.0-dev)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package tcollector

import (
	"bytes"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type TCollector interface {
	// Parameters:
	//  - Span
	Submit(span *Span) (r *Response, err error)
	// Parameters:
	//  - Spans
	SubmitBatch(spans []*Span) (r []*Response, err error)
	// Parameters:
	//  - ServiceName
	GetSamplingStrategy(serviceName string) (r *SamplingStrategyResponse, err error)
}

type TCollectorClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewTCollectorClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *TCollectorClient {
	return &TCollectorClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewTCollectorClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *TCollectorClient {
	return &TCollectorClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - Span
func (p *TCollectorClient) Submit(span *Span) (r *Response, err error) {
	if err = p.sendSubmit(span); err != nil {
		return
	}
	return p.recvSubmit()
}

func (p *TCollectorClient) sendSubmit(span *Span) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("submit", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := TCollectorSubmitArgs{
		Span: span,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *TCollectorClient) recvSubmit() (value *Response, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "submit" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "submit failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "submit failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error2 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error3 error
		error3, err = error2.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error3
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "submit failed: invalid message type")
		return
	}
	result := TCollectorSubmitResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - Spans
func (p *TCollectorClient) SubmitBatch(spans []*Span) (r []*Response, err error) {
	if err = p.sendSubmitBatch(spans); err != nil {
		return
	}
	return p.recvSubmitBatch()
}

func (p *TCollectorClient) sendSubmitBatch(spans []*Span) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("submitBatch", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := TCollectorSubmitBatchArgs{
		Spans: spans,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *TCollectorClient) recvSubmitBatch() (value []*Response, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "submitBatch" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "submitBatch failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "submitBatch failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error4 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error5 error
		error5, err = error4.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error5
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "submitBatch failed: invalid message type")
		return
	}
	result := TCollectorSubmitBatchResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - ServiceName
func (p *TCollectorClient) GetSamplingStrategy(serviceName string) (r *SamplingStrategyResponse, err error) {
	if err = p.sendGetSamplingStrategy(serviceName); err != nil {
		return
	}
	return p.recvGetSamplingStrategy()
}

func (p *TCollectorClient) sendGetSamplingStrategy(serviceName string) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("getSamplingStrategy", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := TCollectorGetSamplingStrategyArgs{
		ServiceName: serviceName,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *TCollectorClient) recvGetSamplingStrategy() (value *SamplingStrategyResponse, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "getSamplingStrategy" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "getSamplingStrategy failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "getSamplingStrategy failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error6 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error7 error
		error7, err = error6.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error7
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "getSamplingStrategy failed: invalid message type")
		return
	}
	result := TCollectorGetSamplingStrategyResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type TCollectorProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      TCollector
}

func (p *TCollectorProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *TCollectorProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *TCollectorProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewTCollectorProcessor(handler TCollector) *TCollectorProcessor {

	self8 := &TCollectorProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self8.processorMap["submit"] = &tCollectorProcessorSubmit{handler: handler}
	self8.processorMap["submitBatch"] = &tCollectorProcessorSubmitBatch{handler: handler}
	self8.processorMap["getSamplingStrategy"] = &tCollectorProcessorGetSamplingStrategy{handler: handler}
	return self8
}

func (p *TCollectorProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x9 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x9.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x9

}

type tCollectorProcessorSubmit struct {
	handler TCollector
}

func (p *tCollectorProcessorSubmit) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := TCollectorSubmitArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("submit", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := TCollectorSubmitResult{}
	var retval *Response
	var err2 error
	if retval, err2 = p.handler.Submit(args.Span); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing submit: "+err2.Error())
		oprot.WriteMessageBegin("submit", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("submit", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type tCollectorProcessorSubmitBatch struct {
	handler TCollector
}

func (p *tCollectorProcessorSubmitBatch) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := TCollectorSubmitBatchArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("submitBatch", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := TCollectorSubmitBatchResult{}
	var retval []*Response
	var err2 error
	if retval, err2 = p.handler.SubmitBatch(args.Spans); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing submitBatch: "+err2.Error())
		oprot.WriteMessageBegin("submitBatch", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("submitBatch", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type tCollectorProcessorGetSamplingStrategy struct {
	handler TCollector
}

func (p *tCollectorProcessorGetSamplingStrategy) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := TCollectorGetSamplingStrategyArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getSamplingStrategy", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := TCollectorGetSamplingStrategyResult{}
	var retval *SamplingStrategyResponse
	var err2 error
	if retval, err2 = p.handler.GetSamplingStrategy(args.ServiceName); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getSamplingStrategy: "+err2.Error())
		oprot.WriteMessageBegin("getSamplingStrategy", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getSamplingStrategy", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Span
type TCollectorSubmitArgs struct {
	Span *Span `thrift:"span,1" json:"span"`
}

func NewTCollectorSubmitArgs() *TCollectorSubmitArgs {
	return &TCollectorSubmitArgs{}
}

var TCollectorSubmitArgs_Span_DEFAULT *Span

func (p *TCollectorSubmitArgs) GetSpan() *Span {
	if !p.IsSetSpan() {
		return TCollectorSubmitArgs_Span_DEFAULT
	}
	return p.Span
}
func (p *TCollectorSubmitArgs) IsSetSpan() bool {
	return p.Span != nil
}

func (p *TCollectorSubmitArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TCollectorSubmitArgs) readField1(iprot thrift.TProtocol) error {
	p.Span = &Span{}
	if err := p.Span.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Span), err)
	}
	return nil
}

func (p *TCollectorSubmitArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("submit_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TCollectorSubmitArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("span", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:span: ", p), err)
	}
	if err := p.Span.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Span), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:span: ", p), err)
	}
	return err
}

func (p *TCollectorSubmitArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TCollectorSubmitArgs(%+v)", *p)
}

// Attributes:
//  - Success
type TCollectorSubmitResult struct {
	Success *Response `thrift:"success,0" json:"success,omitempty"`
}

func NewTCollectorSubmitResult() *TCollectorSubmitResult {
	return &TCollectorSubmitResult{}
}

var TCollectorSubmitResult_Success_DEFAULT *Response

func (p *TCollectorSubmitResult) GetSuccess() *Response {
	if !p.IsSetSuccess() {
		return TCollectorSubmitResult_Success_DEFAULT
	}
	return p.Success
}
func (p *TCollectorSubmitResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *TCollectorSubmitResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TCollectorSubmitResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &Response{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *TCollectorSubmitResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("submit_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TCollectorSubmitResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *TCollectorSubmitResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TCollectorSubmitResult(%+v)", *p)
}

// Attributes:
//  - Spans
type TCollectorSubmitBatchArgs struct {
	Spans []*Span `thrift:"spans,1" json:"spans"`
}

func NewTCollectorSubmitBatchArgs() *TCollectorSubmitBatchArgs {
	return &TCollectorSubmitBatchArgs{}
}

func (p *TCollectorSubmitBatchArgs) GetSpans() []*Span {
	return p.Spans
}
func (p *TCollectorSubmitBatchArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TCollectorSubmitBatchArgs) readField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*Span, 0, size)
	p.Spans = tSlice
	for i := 0; i < size; i++ {
		_elem10 := &Span{}
		if err := _elem10.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem10), err)
		}
		p.Spans = append(p.Spans, _elem10)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *TCollectorSubmitBatchArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("submitBatch_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TCollectorSubmitBatchArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("spans", thrift.LIST, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:spans: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Spans)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Spans {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:spans: ", p), err)
	}
	return err
}

func (p *TCollectorSubmitBatchArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TCollectorSubmitBatchArgs(%+v)", *p)
}

// Attributes:
//  - Success
type TCollectorSubmitBatchResult struct {
	Success []*Response `thrift:"success,0" json:"success,omitempty"`
}

func NewTCollectorSubmitBatchResult() *TCollectorSubmitBatchResult {
	return &TCollectorSubmitBatchResult{}
}

var TCollectorSubmitBatchResult_Success_DEFAULT []*Response

func (p *TCollectorSubmitBatchResult) GetSuccess() []*Response {
	return p.Success
}
func (p *TCollectorSubmitBatchResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *TCollectorSubmitBatchResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TCollectorSubmitBatchResult) readField0(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*Response, 0, size)
	p.Success = tSlice
	for i := 0; i < size; i++ {
		_elem11 := &Response{}
		if err := _elem11.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem11), err)
		}
		p.Success = append(p.Success, _elem11)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *TCollectorSubmitBatchResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("submitBatch_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TCollectorSubmitBatchResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.LIST, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Success)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Success {
			if err := v.Write(oprot); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *TCollectorSubmitBatchResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TCollectorSubmitBatchResult(%+v)", *p)
}

// Attributes:
//  - ServiceName
type TCollectorGetSamplingStrategyArgs struct {
	ServiceName string `thrift:"serviceName,1" json:"serviceName"`
}

func NewTCollectorGetSamplingStrategyArgs() *TCollectorGetSamplingStrategyArgs {
	return &TCollectorGetSamplingStrategyArgs{}
}

func (p *TCollectorGetSamplingStrategyArgs) GetServiceName() string {
	return p.ServiceName
}
func (p *TCollectorGetSamplingStrategyArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TCollectorGetSamplingStrategyArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ServiceName = v
	}
	return nil
}

func (p *TCollectorGetSamplingStrategyArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getSamplingStrategy_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TCollectorGetSamplingStrategyArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("serviceName", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:serviceName: ", p), err)
	}
	if err := oprot.WriteString(string(p.ServiceName)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.serviceName (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:serviceName: ", p), err)
	}
	return err
}

func (p *TCollectorGetSamplingStrategyArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TCollectorGetSamplingStrategyArgs(%+v)", *p)
}

// Attributes:
//  - Success
type TCollectorGetSamplingStrategyResult struct {
	Success *SamplingStrategyResponse `thrift:"success,0" json:"success,omitempty"`
}

func NewTCollectorGetSamplingStrategyResult() *TCollectorGetSamplingStrategyResult {
	return &TCollectorGetSamplingStrategyResult{}
}

var TCollectorGetSamplingStrategyResult_Success_DEFAULT *SamplingStrategyResponse

func (p *TCollectorGetSamplingStrategyResult) GetSuccess() *SamplingStrategyResponse {
	if !p.IsSetSuccess() {
		return TCollectorGetSamplingStrategyResult_Success_DEFAULT
	}
	return p.Success
}
func (p *TCollectorGetSamplingStrategyResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *TCollectorGetSamplingStrategyResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TCollectorGetSamplingStrategyResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &SamplingStrategyResponse{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *TCollectorGetSamplingStrategyResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getSamplingStrategy_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TCollectorGetSamplingStrategyResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *TCollectorGetSamplingStrategyResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TCollectorGetSamplingStrategyResult(%+v)", *p)
}
