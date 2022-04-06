/* eslint-disable */
import { Params } from "../tokenmngr/params";
import { Token } from "../tokenmngr/token";
import { Mintperm } from "../tokenmngr/mintperm";
import { Options } from "../tokenmngr/options";
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "thesixnetwork.sixprotocol.tokenmngr";
const baseGenesisState = { portId: "" };
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        if (message.params !== undefined) {
            Params.encode(message.params, writer.uint32(10).fork()).ldelim();
        }
        if (message.portId !== "") {
            writer.uint32(18).string(message.portId);
        }
        for (const v of message.tokenList) {
            Token.encode(v, writer.uint32(26).fork()).ldelim();
        }
        for (const v of message.mintpermList) {
            Mintperm.encode(v, writer.uint32(34).fork()).ldelim();
        }
        if (message.options !== undefined) {
            Options.encode(message.options, writer.uint32(50).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.tokenList = [];
        message.mintpermList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.params = Params.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.portId = reader.string();
                    break;
                case 3:
                    message.tokenList.push(Token.decode(reader, reader.uint32()));
                    break;
                case 4:
                    message.mintpermList.push(Mintperm.decode(reader, reader.uint32()));
                    break;
                case 6:
                    message.options = Options.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisState };
        message.tokenList = [];
        message.mintpermList = [];
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromJSON(object.params);
        }
        else {
            message.params = undefined;
        }
        if (object.portId !== undefined && object.portId !== null) {
            message.portId = String(object.portId);
        }
        else {
            message.portId = "";
        }
        if (object.tokenList !== undefined && object.tokenList !== null) {
            for (const e of object.tokenList) {
                message.tokenList.push(Token.fromJSON(e));
            }
        }
        if (object.mintpermList !== undefined && object.mintpermList !== null) {
            for (const e of object.mintpermList) {
                message.mintpermList.push(Mintperm.fromJSON(e));
            }
        }
        if (object.options !== undefined && object.options !== null) {
            message.options = Options.fromJSON(object.options);
        }
        else {
            message.options = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.params !== undefined &&
            (obj.params = message.params ? Params.toJSON(message.params) : undefined);
        message.portId !== undefined && (obj.portId = message.portId);
        if (message.tokenList) {
            obj.tokenList = message.tokenList.map((e) => e ? Token.toJSON(e) : undefined);
        }
        else {
            obj.tokenList = [];
        }
        if (message.mintpermList) {
            obj.mintpermList = message.mintpermList.map((e) => e ? Mintperm.toJSON(e) : undefined);
        }
        else {
            obj.mintpermList = [];
        }
        message.options !== undefined &&
            (obj.options = message.options
                ? Options.toJSON(message.options)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.tokenList = [];
        message.mintpermList = [];
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromPartial(object.params);
        }
        else {
            message.params = undefined;
        }
        if (object.portId !== undefined && object.portId !== null) {
            message.portId = object.portId;
        }
        else {
            message.portId = "";
        }
        if (object.tokenList !== undefined && object.tokenList !== null) {
            for (const e of object.tokenList) {
                message.tokenList.push(Token.fromPartial(e));
            }
        }
        if (object.mintpermList !== undefined && object.mintpermList !== null) {
            for (const e of object.mintpermList) {
                message.mintpermList.push(Mintperm.fromPartial(e));
            }
        }
        if (object.options !== undefined && object.options !== null) {
            message.options = Options.fromPartial(object.options);
        }
        else {
            message.options = undefined;
        }
        return message;
    },
};
