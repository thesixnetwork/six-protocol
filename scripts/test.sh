
POINT=$(sixd query nftmngr show-nft-data six-protocol.test_v071 0 --output json | jq '.nftData.onchain_attributes[] | select(.name=="points").number_attribute_value.value') 

POINT2=$(echo $POINT | sed "s/\"//g")

if [ "${POINT2}" == "323" ]; then
	echo "Yes"
fi
