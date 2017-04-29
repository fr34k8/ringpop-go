// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package pingpong

import (
	"bytes"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type PingPong interface {
	// Parameters:
	//  - Request
	Ping(request *Ping) (r *Pong, err error)
}

type PingPongClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewPingPongClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *PingPongClient {
	return &PingPongClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewPingPongClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *PingPongClient {
	return &PingPongClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - Request
func (p *PingPongClient) Ping(request *Ping) (r *Pong, err error) {
	if err = p.sendPing(request); err != nil {
		return
	}
	return p.recvPing()
}

func (p *PingPongClient) sendPing(request *Ping) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("Ping", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := PingArgs{
		Request: request,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *PingPongClient) recvPing() (value *Pong, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error1 error
		error1, err = error0.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error1
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "Ping failed: out of sequence response")
		return
	}
	result := PingResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.PingError != nil {
		err = result.PingError
		return
	}
	value = result.GetSuccess()
	return
}

type PingPongProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      PingPong
}

func (p *PingPongProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *PingPongProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *PingPongProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewPingPongProcessor(handler PingPong) *PingPongProcessor {

	self2 := &PingPongProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self2.processorMap["Ping"] = &pingPongProcessorPing{handler: handler}
	return self2
}

func (p *PingPongProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x3 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x3.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x3

}

type pingPongProcessorPing struct {
	handler PingPong
}

func (p *pingPongProcessorPing) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := PingArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("Ping", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := PingResult{}
	var retval *Pong
	var err2 error
	if retval, err2 = p.handler.Ping(args.Request); err2 != nil {
		switch v := err2.(type) {
		case *PingError:
			result.PingError = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Ping: "+err2.Error())
			oprot.WriteMessageBegin("Ping", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("Ping", thrift.REPLY, seqId); err2 != nil {
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

type PingArgs struct {
	Request *Ping `thrift:"request,1" json:"request"`
}

func NewPingArgs() *PingArgs {
	return &PingArgs{}
}

var PingArgs_Request_DEFAULT *Ping

func (p *PingArgs) GetRequest() *Ping {
	if !p.IsSetRequest() {
		return PingArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *PingArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *PingArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *PingArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Request = &Ping{}
	if err := p.Request.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Request, err)
	}
	return nil
}

func (p *PingArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Ping_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *PingArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("request", thrift.STRUCT, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:request: %s", p, err)
	}
	if err := p.Request.Write(oprot); err != nil {
		return fmt.Errorf("%T error writing struct: %s", p.Request, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:request: %s", p, err)
	}
	return err
}

func (p *PingArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PingArgs(%+v)", *p)
}

type PingResult struct {
	Success   *Pong      `thrift:"success,0" json:"success"`
	PingError *PingError `thrift:"pingError,1" json:"pingError"`
}

func NewPingResult() *PingResult {
	return &PingResult{}
}

var PingResult_Success_DEFAULT *Pong

func (p *PingResult) GetSuccess() *Pong {
	if !p.IsSetSuccess() {
		return PingResult_Success_DEFAULT
	}
	return p.Success
}

var PingResult_PingError_DEFAULT *PingError

func (p *PingResult) GetPingError() *PingError {
	if !p.IsSetPingError() {
		return PingResult_PingError_DEFAULT
	}
	return p.PingError
}
func (p *PingResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PingResult) IsSetPingError() bool {
	return p.PingError != nil
}

func (p *PingResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.ReadField0(iprot); err != nil {
				return err
			}
		case 1:
			if err := p.ReadField1(iprot); err != nil {
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *PingResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &Pong{}
	if err := p.Success.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Success, err)
	}
	return nil
}

func (p *PingResult) ReadField1(iprot thrift.TProtocol) error {
	p.PingError = &PingError{}
	if err := p.PingError.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.PingError, err)
	}
	return nil
}

func (p *PingResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Ping_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *PingResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Success, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 0:success: %s", p, err)
		}
	}
	return err
}

func (p *PingResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetPingError() {
		if err := oprot.WriteFieldBegin("pingError", thrift.STRUCT, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:pingError: %s", p, err)
		}
		if err := p.PingError.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.PingError, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:pingError: %s", p, err)
		}
	}
	return err
}

func (p *PingResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PingResult(%+v)", *p)
}
