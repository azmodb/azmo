package pb

import "errors"

var (
	errRequestsRequired   = errors.New("requires at least 1 request")
	errInvalidRequestType = errors.New("invalid request type")
	errInvalidTxnID       = errors.New("invalid txnid")
	errFromRangeRequired  = errors.New("from range required")
	errKeyRequired        = errors.New("key required")
)

func (m *GenericRequest) Verify() error {
	if m.Type < GenericRequest_DeleteRequest || m.Type > GenericRequest_PutRequest {
		return errInvalidRequestType
	}
	if m.Num <= 0 {
		return errInvalidTxnID
	}
	if len(m.Key) <= 0 {
		return errKeyRequired
	}
	return nil
}

func (m *TxnRequest) Verify() error {
	if len(m.Requests) <= 0 {
		return errRequestsRequired
	}
	for _, req := range m.Requests {
		if err := req.Verify(); err != nil {
			return err
		}
	}
	return nil
}

func (m *RangeRequest) Verify() error {
	if len(m.From) <= 0 {
		return errFromRangeRequired
	}
	return nil
}

func (m *DeleteRequest) Verify() error {
	if len(m.Key) <= 0 {
		return errKeyRequired
	}
	return nil
}

func (m *PutRequest) Verify() error {
	if len(m.Key) <= 0 {
		return errKeyRequired
	}
	return nil
}

func (m *GetRequest) Verify() error {
	if len(m.Key) <= 0 {
		return errKeyRequired
	}
	return nil
}
