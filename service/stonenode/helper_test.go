package stonenode

import (
	"context"
	"testing"

	"google.golang.org/grpc"

	ptypes "github.com/bnb-chain/greenfield-storage-provider/pkg/types/v1"
	service "github.com/bnb-chain/greenfield-storage-provider/service/types/v1"
)

func setup(t *testing.T) *StoneNodeService {
	return &StoneNodeService{
		cfg: &StoneNodeConfig{
			Address:                "test1",
			StoneHubServiceAddress: "test2",
			SyncerServiceAddress:   "test3",
			StorageProvider:        "test",
			StoneJobLimit:          0,
		},
		name:       ServiceNameStoneNode,
		stoneLimit: 0,
	}
}

func mockAllocResp(objectID uint64, payloadSize uint64, redundancyType ptypes.RedundancyType) *service.StoneHubServiceAllocStoneJobResponse {
	return &service.StoneHubServiceAllocStoneJobResponse{
		TraceId: "123456",
		TxHash:  []byte("blockchain_one"),
		PieceJob: &service.PieceJob{
			TxHash:         []byte("blockchain_one"),
			ObjectId:       objectID,
			PayloadSize:    payloadSize,
			TargetIdx:      nil,
			RedundancyType: redundancyType,
		},
		ErrMessage: &service.ErrMessage{
			ErrCode: 0,
			ErrMsg:  "Success",
		},
	}
}

func dispatchPieceMap() map[string][][]byte {
	ecList1 := [][]byte{[]byte("1"), []byte("2"), []byte("3"), []byte("4"), []byte("5"), []byte("6")}
	ecList2 := [][]byte{[]byte("a"), []byte("b"), []byte("c"), []byte("d"), []byte("e"), []byte("f")}
	pMap := make(map[string][][]byte)
	pMap["123456_s0"] = ecList1
	pMap["123456_s1"] = ecList2
	return pMap
}

func dispatchSegmentMap() map[string][][]byte {
	segmentList1 := [][]byte{[]byte("10")}
	segmentList2 := [][]byte{[]byte("20")}
	segmentList3 := [][]byte{[]byte("30")}
	sMap := make(map[string][][]byte)
	sMap["789_s0"] = segmentList1
	sMap["789_s1"] = segmentList2
	sMap["789_s2"] = segmentList3
	return sMap
}

func dispatchInlineMap() map[string][][]byte {
	inlineList := [][]byte{[]byte("+")}
	iMap := make(map[string][][]byte)
	iMap["543_s0"] = inlineList
	return iMap
}

func makeStreamMock() *StreamMock {
	return &StreamMock{
		ctx:          context.Background(),
		recvToServer: make(chan *service.SyncerServiceSyncPieceRequest, 10),
	}
}

type StreamMock struct {
	grpc.ClientStream
	ctx          context.Context
	recvToServer chan *service.SyncerServiceSyncPieceRequest
}

func (m *StreamMock) Send(resp *service.SyncerServiceSyncPieceRequest) error {
	m.recvToServer <- resp
	return nil
}

func (m *StreamMock) CloseAndRecv() (*service.SyncerServiceSyncPieceResponse, error) {
	return &service.SyncerServiceSyncPieceResponse{
		TraceId: "test_traceID",
		SecondarySpInfo: &service.StorageProviderSealInfo{
			StorageProviderId: "sp1",
			PieceIdx:          1,
			PieceChecksum:     [][]byte{[]byte("1"), []byte("2"), []byte("3"), []byte("4"), []byte("5"), []byte("6")},
			IntegrityHash:     []byte("a"),
			Signature:         nil,
		},
		ErrMessage: &service.ErrMessage{
			ErrCode: service.ErrCode_ERR_CODE_SUCCESS_UNSPECIFIED,
			ErrMsg:  "Success",
		},
	}, nil
}