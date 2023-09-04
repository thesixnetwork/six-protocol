import json

def generate_nft_data(token_id, all_attributes):
    with open("./nft-data.json", "r") as nft_data_file:
        nft_data = json.load(nft_data_file)
        
        # Find the token with matching tokenId
        for item in all_attributes:
            if item["tokenId"] == int(token_id):
                # Replace token_id with the actual token ID
                nft_data["token_id"] = str(item["tokenId"])
                attributes = nft_data["origin_attributes"]
                
                # Map attribute values
                attribute_mapping = {k.lower(): k for k in item.keys() if k != "tokenId"}
                for attribute in attributes:
                    attribute_name = attribute["name"]
                    attribute_key = attribute_name.lower()
                    if attribute_key in attribute_mapping:
                        original_attribute_name = attribute_mapping[attribute_key]
                        attribute_value = item[original_attribute_name]
                        attribute["string_attribute_value"]["value"] = attribute_value
                break
        
        output_filename = f"{token_id}.json"
        with open(output_filename, "w") as outfile:
            json.dump(nft_data, outfile, indent=4)

# Load all-attribute.json
with open("./all-objects-VIP.json", "r") as all_attributes_file:
    all_attributes = json.load(all_attributes_file)
    
    # Generate nft-data.json for each token ID
    for item in all_attributes:
        token_id = str(item["tokenId"])
        generate_nft_data(token_id, all_attributes)
