package types

const (
	EventTypesConvertCoinToWei   = "convert_coin_to_asix"
	EventTypesConvertCoinToMicro = "convert_coin_to_usix"
	EventTypesSentToCrossChain   = "sent_to_cross_chain"

	AttributeKeyDestChain   = "to_chain"
	AttributeKeyDestAddress = "to_address"
	AttributeKeyEvmSender   = "evm_sender"
	AttributeKeyMemo        = "memo"
	AttributeKeyAmount      = "amount"
)
