package ethsecp256k1

type PrivKey struct {
	// key is the private key in byte form
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	
}
