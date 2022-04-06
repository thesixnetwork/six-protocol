/* eslint-disable */
import { Params } from "../protocoladmin/params";
import { Group } from "../protocoladmin/group";
import { Admin } from "../protocoladmin/admin";
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "thesixnetwork.sixprotocol.protocoladmin";
const baseGenesisState = { portId: "" };
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        if (message.params !== undefined) {
            Params.encode(message.params, writer.uint32(10).fork()).ldelim();
        }
        if (message.portId !== "") {
            writer.uint32(18).string(message.portId);
        }
        for (const v of message.groupList) {
            Group.encode(v, writer.uint32(26).fork()).ldelim();
        }
        for (const v of message.adminList) {
            Admin.encode(v, writer.uint32(34).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.groupList = [];
        message.adminList = [];
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
                    message.groupList.push(Group.decode(reader, reader.uint32()));
                    break;
                case 4:
                    message.adminList.push(Admin.decode(reader, reader.uint32()));
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
        message.groupList = [];
        message.adminList = [];
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
        if (object.groupList !== undefined && object.groupList !== null) {
            for (const e of object.groupList) {
                message.groupList.push(Group.fromJSON(e));
            }
        }
        if (object.adminList !== undefined && object.adminList !== null) {
            for (const e of object.adminList) {
                message.adminList.push(Admin.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.params !== undefined &&
            (obj.params = message.params ? Params.toJSON(message.params) : undefined);
        message.portId !== undefined && (obj.portId = message.portId);
        if (message.groupList) {
            obj.groupList = message.groupList.map((e) => e ? Group.toJSON(e) : undefined);
        }
        else {
            obj.groupList = [];
        }
        if (message.adminList) {
            obj.adminList = message.adminList.map((e) => e ? Admin.toJSON(e) : undefined);
        }
        else {
            obj.adminList = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.groupList = [];
        message.adminList = [];
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
        if (object.groupList !== undefined && object.groupList !== null) {
            for (const e of object.groupList) {
                message.groupList.push(Group.fromPartial(e));
            }
        }
        if (object.adminList !== undefined && object.adminList !== null) {
            for (const e of object.adminList) {
                message.adminList.push(Admin.fromPartial(e));
            }
        }
        return message;
    },
};
