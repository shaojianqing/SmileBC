package p2p

import (
	"fmt"

	"github.com/shaojianqing/smilebc/common"
)

func (p *Peer) PerformHandshake(network, version string, genesis common.Hash) error {
	state := &PeerState{
		Network: network,
		Version: version,
		Genesis: genesis,
	}

	message := &Message{MessageType: PeerStateReq, MessageBody: state}
	if err := p.WriteMessage(message); err != nil {
		return err
	}

	message, err := p.ReadMessage()
	if err != nil {
		return err
	}

	if message.MessageType != PeerStateResp {
		return fmt.Errorf("message type is not correct")
	}

	newState := message.MessageBody.(*PeerState)
	if newState.Network != state.Network {
		return fmt.Errorf("network does not match")
	}
	if newState.Version != state.Version {
		return fmt.Errorf("version does not match")
	}
	if newState.Genesis != state.Genesis {
		return fmt.Errorf("genesis does not match")
	}
	return nil
}
