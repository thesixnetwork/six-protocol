/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "thesixnetwork.sixprotocol.tokenmngr";
const baseOptions = { defaultMintee: "" };
export const Options = {
    encode(message, writer = Writer.create()) {
        if (message.defaultMintee !== "") {
            writer.uint32(10).string(message.defaultMintee);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseOptions };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.defaultMintee = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseOptions };
        if (object.defaultMintee !== undefined && object.defaultMintee !== null) {
            message.defaultMintee = String(object.defaultMintee);
        }
        else {
            message.defaultMintee = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.defaultMintee !== undefined &&
            (obj.defaultMintee = message.defaultMintee);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseOptions };
        if (object.defaultMintee !== undefined && object.defaultMintee !== null) {
            message.defaultMintee = object.defaultMintee;
        }
        else {
            message.defaultMintee = "";
        }
        return message;
    },
};
