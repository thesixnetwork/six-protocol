import { ethers } from "ethers";
import NFTMngrABI from "../../out/INFTManager.sol/INFTMNGR.json";
import { provider, wallet } from "../client";
import dotenv from "dotenv";
dotenv.config();

import exmapleSchema from "../../resources/nft-schema.json"
import ORIGINAL_METADATA from "../../resources/nft-data.json";

let schema_name = exmapleSchema.code;
const split_schema = schema_name.split(".")
const _name = split_schema[1]
const org_name = process.env.ORG_NAME
let schemaCode: string;
schemaCode = `${org_name}.${_name}`;

const VIP_TIER_START = 1;
const VIP_TIER_END = 10;

const SPEAKER_TIER_START = 11;
const SPEAKER_TIER_END = 20;

const INVESTOR_TIER_START = 21;
const INVESTOR_TIER_END = 30;

const PARTNER_TIER_START = 31;
const PARTNER_TIER_END = 40;

const MEDIA_TIER_START = 41;
const MEDIA_TIER_END = 50;

const EXHIBITOR_TIER_START = 51;
const EXHIBITOR_TIER_END = 60;

const ORGANIZER_TIER_START = 61;
const ORGANIZER_TIER_END = 70;

const STAFF_TIER_START = 71;
const STAFF_TIER_END = 80;

const ATTENDEE_TIER_START = 81;
const ATTENDEE_TIER_END = 100;

const TIER_RANGE = {
  VIP: [VIP_TIER_START, VIP_TIER_END],
  Speaker: [SPEAKER_TIER_START, SPEAKER_TIER_END],
  Investor: [INVESTOR_TIER_START, INVESTOR_TIER_END],
  Partner: [PARTNER_TIER_START, PARTNER_TIER_END],
  Media: [MEDIA_TIER_START, MEDIA_TIER_END],
  Exhibitor: [EXHIBITOR_TIER_START, EXHIBITOR_TIER_END],
  Organizer: [ORGANIZER_TIER_START, ORGANIZER_TIER_END],
  Staff: [STAFF_TIER_START, STAFF_TIER_END],
  Attendee: [ATTENDEE_TIER_START, ATTENDEE_TIER_END],
};


console.log(schemaCode);

exmapleSchema.code = schemaCode;
ORIGINAL_METADATA.nft_schema_code = schemaCode;

async function main() {
    
  const precompileContract = "0x0000000000000000000000000000000000001055";
  const tokenID = "2"
  const nftmngr = new ethers.Contract(
    precompileContract,
    NFTMngrABI.abi,
    provider,
  );

  const ORIGINAL_METADATA = await generateNFTData(tokenID);
  ORIGINAL_METADATA.token_id = tokenID;
  let encodeBase64Metadata = Buffer.from(
    JSON.stringify(ORIGINAL_METADATA),
  ).toString("base64");

  // @ts-ignore
  const tx = await nftmngr.connect(wallet).createMetadata(schemaCode,tokenID, encodeBase64Metadata)

  console.log("Transaction hash:", tx.hash);

  // Wait for the transaction to be mined
  const receipt = await tx.wait();

  // Get gas used
  const gasUsed = receipt.gasUsed;

  console.log(`Gas used: ${gasUsed.toString()}`);
  console.log("TX Info", tx);
}


async function generateNFTData(tokenId: string): Promise<any> {
    const nftData: any = ORIGINAL_METADATA;
    // Find the tier of the token ID
    let tier = "";
    let tier_count: number;
    for (const key in TIER_RANGE) {
      const range = TIER_RANGE[key];
      tier_count = range[1] - range[0];
      if (parseInt(tokenId) >= range[0] && parseInt(tokenId) <= range[1]) {
        tier = key;
        break;
      }
    }
    if (tier === "") {
      throw new Error("Invalid token ID");
    }
  
    for (let i = 0; i <= tier_count; i++) {
      // Replace token_id with the actual token ID
      nftData.token_id = String(tokenId);
      const onchain_attributes = nftData.onchain_attributes;
      for (const attribute of onchain_attributes) {
        const attributeKey = attribute.name.toLowerCase();
        // find after_party_claimed attribute and set to false
        if (tier === "VIP" || tier === "Speaker" || tier === "Investor") {
          nftData.origin_image = `https://storage.googleapis.com/techsauce2024/nonalumni/afterparty.png`;
          if (attributeKey === "after_party_claimed") {
            attribute.boolean_attribute_value.value = true;
          }
        } else {
          nftData.origin_image = `https://storage.googleapis.com/techsauce2024/nonalumni/normal.png`;
          if (attributeKey === "after_party_claimed") {
            attribute.boolean_attribute_value.value = false;
          }
        }
        if (attributeKey === "tier") {
          attribute.string_attribute_value.value = tier;
        }
      }
    }
    return nftData;
  }

main()
  .then(() => {
    console.log("Transaction complete.");
  })
  .catch((err) => {
    console.error(err);
  });
