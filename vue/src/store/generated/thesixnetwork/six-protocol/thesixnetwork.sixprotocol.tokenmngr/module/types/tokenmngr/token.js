/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "thesixnetwork.sixprotocol.tokenmngr";
const baseToken = {
    name: "",
    base: "",
    maxSupply: 0,
    mintee: "",
    creator: "",
};
export const Token = {
    encode(message, writer = Writer.create()) {
        if (message.name !== "") {
            writer.uint32(10).string(message.name);
        }
        if (message.base !== "") {
            writer.uint32(18).string(message.base);
        }
        if (message.maxSupply !== 0) {
            writer.uint32(24).uint64(message.maxSupply);
        }
        if (message.mintee !== "") {
            writer.uint32(34).string(message.mintee);
        }
        if (message.creator !== "") {
            writer.uint32(42).string(message.creator);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseToken };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.name = reader.string();
                    break;
                case 2:
                    message.base = reader.string();
                    break;
                case 3:
                    message.maxSupply = longToNumber(reader.uint64());
                    break;
                case 4:
                    message.mintee = reader.string();
                    break;
                case 5:
                    message.creator = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseToken };
        if (object.name !== undefined && object.name !== null) {
            message.name = String(object.name);
        }
        else {
            message.name = "";
        }
        if (object.base !== undefined && object.base !== null) {
            message.base = String(object.base);
        }
        else {
            message.base = "";
        }
        if (object.maxSupply !== undefined && object.maxSupply !== null) {
            message.maxSupply = Number(object.maxSupply);
        }
        else {
            message.maxSupply = 0;
        }
        if (object.mintee !== undefined && object.mintee !== null) {
            message.mintee = String(object.mintee);
        }
        else {
            message.mintee = "";
        }
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.name !== undefined && (obj.name = message.name);
        message.base !== undefined && (obj.base = message.base);
        message.maxSupply !== undefined && (obj.maxSupply = message.maxSupply);
        message.mintee !== undefined && (obj.mintee = message.mintee);
        message.creator !== undefined && (obj.creator = message.creator);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseToken };
        if (object.name !== undefined && object.name !== null) {
            message.name = object.name;
        }
        else {
            message.name = "";
        }
        if (object.base !== undefined && object.base !== null) {
            message.base = object.base;
        }
        else {
            message.base = "";
        }
        if (object.maxSupply !== undefined && object.maxSupply !== null) {
            message.maxSupply = object.maxSupply;
        }
        else {
            message.maxSupply = 0;
        }
        if (object.mintee !== undefined && object.mintee !== null) {
            message.mintee = object.mintee;
        }
        else {
            message.mintee = "";
        }
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        return message;
    },
};
const baseDenomUnit = { denom: "", exponent: 0, aliases: "" };
export const DenomUnit = {
    encode(message, writer = Writer.create()) {
        if (message.denom !== "") {
            writer.uint32(10).string(message.denom);
        }
        if (message.exponent !== 0) {
            writer.uint32(16).uint32(message.exponent);
        }
        for (const v of message.aliases) {
            writer.uint32(26).string(v);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseDenomUnit };
        message.aliases = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.denom = reader.string();
                    break;
                case 2:
                    message.exponent = reader.uint32();
                    break;
                case 3:
                    message.aliases.push(reader.string());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseDenomUnit };
        message.aliases = [];
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = String(object.denom);
        }
        else {
            message.denom = "";
        }
        if (object.exponent !== undefined && object.exponent !== null) {
            message.exponent = Number(object.exponent);
        }
        else {
            message.exponent = 0;
        }
        if (object.aliases !== undefined && object.aliases !== null) {
            for (const e of object.aliases) {
                message.aliases.push(String(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.denom !== undefined && (obj.denom = message.denom);
        message.exponent !== undefined && (obj.exponent = message.exponent);
        if (message.aliases) {
            obj.aliases = message.aliases.map((e) => e);
        }
        else {
            obj.aliases = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseDenomUnit };
        message.aliases = [];
        if (object.denom !== undefined && object.denom !== null) {
            message.denom = object.denom;
        }
        else {
            message.denom = "";
        }
        if (object.exponent !== undefined && object.exponent !== null) {
            message.exponent = object.exponent;
        }
        else {
            message.exponent = 0;
        }
        if (object.aliases !== undefined && object.aliases !== null) {
            for (const e of object.aliases) {
                message.aliases.push(e);
            }
        }
        return message;
    },
};
const baseMetadata = {
    description: "",
    base: "",
    display: "",
    name: "",
    symbol: "",
};
export const Metadata = {
    encode(message, writer = Writer.create()) {
        if (message.description !== "") {
            writer.uint32(10).string(message.description);
        }
        for (const v of message.denomUnits) {
            DenomUnit.encode(v, writer.uint32(18).fork()).ldelim();
        }
        if (message.base !== "") {
            writer.uint32(26).string(message.base);
        }
        if (message.display !== "") {
            writer.uint32(34).string(message.display);
        }
        if (message.name !== "") {
            writer.uint32(42).string(message.name);
        }
        if (message.symbol !== "") {
            writer.uint32(50).string(message.symbol);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMetadata };
        message.denomUnits = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.description = reader.string();
                    break;
                case 2:
                    message.denomUnits.push(DenomUnit.decode(reader, reader.uint32()));
                    break;
                case 3:
                    message.base = reader.string();
                    break;
                case 4:
                    message.display = reader.string();
                    break;
                case 5:
                    message.name = reader.string();
                    break;
                case 6:
                    message.symbol = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMetadata };
        message.denomUnits = [];
        if (object.description !== undefined && object.description !== null) {
            message.description = String(object.description);
        }
        else {
            message.description = "";
        }
        if (object.denomUnits !== undefined && object.denomUnits !== null) {
            for (const e of object.denomUnits) {
                message.denomUnits.push(DenomUnit.fromJSON(e));
            }
        }
        if (object.base !== undefined && object.base !== null) {
            message.base = String(object.base);
        }
        else {
            message.base = "";
        }
        if (object.display !== undefined && object.display !== null) {
            message.display = String(object.display);
        }
        else {
            message.display = "";
        }
        if (object.name !== undefined && object.name !== null) {
            message.name = String(object.name);
        }
        else {
            message.name = "";
        }
        if (object.symbol !== undefined && object.symbol !== null) {
            message.symbol = String(object.symbol);
        }
        else {
            message.symbol = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.description !== undefined &&
            (obj.description = message.description);
        if (message.denomUnits) {
            obj.denomUnits = message.denomUnits.map((e) => e ? DenomUnit.toJSON(e) : undefined);
        }
        else {
            obj.denomUnits = [];
        }
        message.base !== undefined && (obj.base = message.base);
        message.display !== undefined && (obj.display = message.display);
        message.name !== undefined && (obj.name = message.name);
        message.symbol !== undefined && (obj.symbol = message.symbol);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMetadata };
        message.denomUnits = [];
        if (object.description !== undefined && object.description !== null) {
            message.description = object.description;
        }
        else {
            message.description = "";
        }
        if (object.denomUnits !== undefined && object.denomUnits !== null) {
            for (const e of object.denomUnits) {
                message.denomUnits.push(DenomUnit.fromPartial(e));
            }
        }
        if (object.base !== undefined && object.base !== null) {
            message.base = object.base;
        }
        else {
            message.base = "";
        }
        if (object.display !== undefined && object.display !== null) {
            message.display = object.display;
        }
        else {
            message.display = "";
        }
        if (object.name !== undefined && object.name !== null) {
            message.name = object.name;
        }
        else {
            message.name = "";
        }
        if (object.symbol !== undefined && object.symbol !== null) {
            message.symbol = object.symbol;
        }
        else {
            message.symbol = "";
        }
        return message;
    },
};
var globalThis = (() => {
    if (typeof globalThis !== "undefined")
        return globalThis;
    if (typeof self !== "undefined")
        return self;
    if (typeof window !== "undefined")
        return window;
    if (typeof global !== "undefined")
        return global;
    throw "Unable to locate global object";
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
