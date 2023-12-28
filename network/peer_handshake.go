package p2p

const (
	AuthRequestLength  = 1024
	AuthResponseLength = 1024
)

type AuthRequest struct {
	remoteNodeID  NodeID        `json:"remoteNodeID"`
	aesEncryptKey AESEncryptKey `json:"shareEncryptKey"`
}

type AuthResponse struct {
	selfNodeID   NodeID `json:"selfNodeID"`
	remoteNodeID NodeID `json:"remoteNodeID"`
	echoMessage  string `json:"echoMessage"`
}

func NewAuthRequest(remoteNodeID NodeID, aesEncryptKey AESEncryptKey) *AuthRequest {
	return &AuthRequest{
		remoteNodeID:  remoteNodeID,
		aesEncryptKey: aesEncryptKey,
	}
}

func NewAuthResponse(selfNodeID, remoteNodeID NodeID, echoMessage string) *AuthResponse {
	return &AuthResponse{
		selfNodeID:   selfNodeID,
		remoteNodeID: remoteNodeID,
		echoMessage:  echoMessage,
	}
}
