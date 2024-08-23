package oidb

import (
	"github.com/LagrangeDev/LagrangeGo/client/entity"
	"github.com/LagrangeDev/LagrangeGo/client/packets/pb/service/oidb"
	"github.com/LagrangeDev/LagrangeGo/utils"
)

func BuildFetchGroupSystemMessagesReq(count uint32) (*OidbPacket, error) {
	body := &oidb.OidbSvcTrpcTcp0X10C0_1{
		Count:  count,
		Field2: 0,
	}
	return BuildOidbPacket(0x10C0, 1, body, false, false)
}

func ParseFetchGroupSystemMessagesReq(data []byte, groupUin ...uint32) ([]*entity.GroupJoinRequest, error) {
	resp, err := ParseTypedError[oidb.OidbSvcTrpcTcp0X10C0_1Response](data)
	if err != nil {
		return nil, err
	}
	requests := make([]*entity.GroupJoinRequest, 0)
	for _, r := range resp.Requests {
		if len(groupUin) > 0 && groupUin[0] != r.Group.GroupUin {
			continue
		}
		requests = append(requests, &entity.GroupJoinRequest{
			GroupUin:    r.Group.GroupUin,
			TargetUid:   r.Target.Uid,
			Sequence:    r.Sequence,
			State:       entity.EventState(r.State),
			EventType:   r.EventType,
			Comment:     r.Comment,
			InvitorUid:  utils.Ternary[string](r.Invitor != nil, r.Invitor.Uid, ""),
			OperatorUid: utils.Ternary[string](r.Operator != nil, r.Operator.Uid, ""),
		})
	}
	return requests, nil
}
