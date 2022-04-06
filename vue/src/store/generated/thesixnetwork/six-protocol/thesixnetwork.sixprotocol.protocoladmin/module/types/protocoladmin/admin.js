/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "thesixnetwork.sixprotocol.protocoladmin";
const baseAdmin = { group: "", admin: "" };
export const Admin = {
    encode(message, writer = Writer.create()) {
        if (message.group !== "") {
            writer.uint32(10).string(message.group);
        }
        if (message.admin !== "") {
            writer.uint32(18).string(message.admin);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseAdmin };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.group = reader.string();
                    break;
                case 2:
                    message.admin = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseAdmin };
        if (object.group !== undefined && object.group !== null) {
            message.group = String(object.group);
        }
        else {
            message.group = "";
        }
        if (object.admin !== undefined && object.admin !== null) {
            message.admin = String(object.admin);
        }
        else {
            message.admin = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.group !== undefined && (obj.group = message.group);
        message.admin !== undefined && (obj.admin = message.admin);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseAdmin };
        if (object.group !== undefined && object.group !== null) {
            message.group = object.group;
        }
        else {
            message.group = "";
        }
        if (object.admin !== undefined && object.admin !== null) {
            message.admin = object.admin;
        }
        else {
            message.admin = "";
        }
        return message;
    },
};
