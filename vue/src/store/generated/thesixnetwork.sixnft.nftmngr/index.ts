import { Client, registry, MissingWalletError } from 'thesixnetwork-six-protocol-client-ts'

import { Action } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixnft.nftmngr/types"
import { ActionByRefId } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixnft.nftmngr/types"
import { DefaultMintValue } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixnft.nftmngr/types"
import { AttributeDefinition } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixnft.nftmngr/types"
import { DisplayOption } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixnft.nftmngr/types"
import { NftAttributeValue } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixnft.nftmngr/types"
import { NumberAttributeValue } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixnft.nftmngr/types"
import { StringAttributeValue } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixnft.nftmngr/types"
import { BooleanAttributeValue } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixnft.nftmngr/types"
import { FloatAttributeValue } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixnft.nftmngr/types"
import { NftData } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixnft.nftmngr/types"
import { NFTSchema } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixnft.nftmngr/types"
import { OnChainData } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixnft.nftmngr/types"
import { OnOffSwitch } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixnft.nftmngr/types"
import { OpenseaDisplayOption } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixnft.nftmngr/types"
import { Organization } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixnft.nftmngr/types"
import { OriginData } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixnft.nftmngr/types"
import { Params } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixnft.nftmngr/types"
import { Status } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixnft.nftmngr/types"
import { OpenseaAttribute } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixnft.nftmngr/types"
import { UpdatedOpenseaAttributes } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixnft.nftmngr/types"
import { UpdatedOriginData } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixnft.nftmngr/types"


export { Action, ActionByRefId, DefaultMintValue, AttributeDefinition, DisplayOption, NftAttributeValue, NumberAttributeValue, StringAttributeValue, BooleanAttributeValue, FloatAttributeValue, NftData, NFTSchema, OnChainData, OnOffSwitch, OpenseaDisplayOption, Organization, OriginData, Params, Status, OpenseaAttribute, UpdatedOpenseaAttributes, UpdatedOriginData };

function initClient(vuexGetters) {
	return new Client(vuexGetters['common/env/getEnv'], vuexGetters['common/wallet/signer'])
}

function mergeResults(value, next_values) {
	for (let prop of Object.keys(next_values)) {
		if (Array.isArray(next_values[prop])) {
			value[prop]=[...value[prop], ...next_values[prop]]
		}else{
			value[prop]=next_values[prop]
		}
	}
	return value
}

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	let structure: {fields: Field[]} = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
const getDefaultState = () => {
	return {
				Params: {},
				NFTSchema: {},
				NFTSchemaAll: {},
				NftData: {},
				NftDataAll: {},
				ActionByRefId: {},
				ActionByRefIdAll: {},
				Organization: {},
				OrganizationAll: {},
				
				_Structure: {
						Action: getStructure(Action.fromPartial({})),
						ActionByRefId: getStructure(ActionByRefId.fromPartial({})),
						DefaultMintValue: getStructure(DefaultMintValue.fromPartial({})),
						AttributeDefinition: getStructure(AttributeDefinition.fromPartial({})),
						DisplayOption: getStructure(DisplayOption.fromPartial({})),
						NftAttributeValue: getStructure(NftAttributeValue.fromPartial({})),
						NumberAttributeValue: getStructure(NumberAttributeValue.fromPartial({})),
						StringAttributeValue: getStructure(StringAttributeValue.fromPartial({})),
						BooleanAttributeValue: getStructure(BooleanAttributeValue.fromPartial({})),
						FloatAttributeValue: getStructure(FloatAttributeValue.fromPartial({})),
						NftData: getStructure(NftData.fromPartial({})),
						NFTSchema: getStructure(NFTSchema.fromPartial({})),
						OnChainData: getStructure(OnChainData.fromPartial({})),
						OnOffSwitch: getStructure(OnOffSwitch.fromPartial({})),
						OpenseaDisplayOption: getStructure(OpenseaDisplayOption.fromPartial({})),
						Organization: getStructure(Organization.fromPartial({})),
						OriginData: getStructure(OriginData.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						Status: getStructure(Status.fromPartial({})),
						OpenseaAttribute: getStructure(OpenseaAttribute.fromPartial({})),
						UpdatedOpenseaAttributes: getStructure(UpdatedOpenseaAttributes.fromPartial({})),
						UpdatedOriginData: getStructure(UpdatedOriginData.fromPartial({})),
						
		},
		_Registry: registry,
		_Subscriptions: new Set(),
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(JSON.stringify(subscription))
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(JSON.stringify(subscription))
		}
	},
	getters: {
				getParams: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Params[JSON.stringify(params)] ?? {}
		},
				getNFTSchema: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.NFTSchema[JSON.stringify(params)] ?? {}
		},
				getNFTSchemaAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.NFTSchemaAll[JSON.stringify(params)] ?? {}
		},
				getNftData: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.NftData[JSON.stringify(params)] ?? {}
		},
				getNftDataAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.NftDataAll[JSON.stringify(params)] ?? {}
		},
				getActionByRefId: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ActionByRefId[JSON.stringify(params)] ?? {}
		},
				getActionByRefIdAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ActionByRefIdAll[JSON.stringify(params)] ?? {}
		},
				getOrganization: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Organization[JSON.stringify(params)] ?? {}
		},
				getOrganizationAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.OrganizationAll[JSON.stringify(params)] ?? {}
		},
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: thesixnetwork.sixnft.nftmngr initialized!')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach(async (subscription) => {
				try {
					const sub=JSON.parse(subscription)
					await dispatch(sub.action, sub.payload)
				}catch(e) {
					throw new Error('Subscriptions: ' + e.message)
				}
			})
		},
		
		
		
		 		
		
		
		async QueryParams({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ThesixnetworkSixnftNftmngr.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryNFTSchema({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ThesixnetworkSixnftNftmngr.query.queryNFTSchema( key.code)).data
				
					
				commit('QUERY', { query: 'NFTSchema', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryNFTSchema', payload: { options: { all }, params: {...key},query }})
				return getters['getNFTSchema']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryNFTSchema API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryNFTSchemaAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ThesixnetworkSixnftNftmngr.query.queryNFTSchemaAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ThesixnetworkSixnftNftmngr.query.queryNFTSchemaAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'NFTSchemaAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryNFTSchemaAll', payload: { options: { all }, params: {...key},query }})
				return getters['getNFTSchemaAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryNFTSchemaAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryNftData({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ThesixnetworkSixnftNftmngr.query.queryNftData( key.nftSchemaCode,  key.tokenId)).data
				
					
				commit('QUERY', { query: 'NftData', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryNftData', payload: { options: { all }, params: {...key},query }})
				return getters['getNftData']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryNftData API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryNftDataAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ThesixnetworkSixnftNftmngr.query.queryNftDataAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ThesixnetworkSixnftNftmngr.query.queryNftDataAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'NftDataAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryNftDataAll', payload: { options: { all }, params: {...key},query }})
				return getters['getNftDataAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryNftDataAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryActionByRefId({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ThesixnetworkSixnftNftmngr.query.queryActionByRefId( key.refId)).data
				
					
				commit('QUERY', { query: 'ActionByRefId', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryActionByRefId', payload: { options: { all }, params: {...key},query }})
				return getters['getActionByRefId']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryActionByRefId API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryActionByRefIdAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ThesixnetworkSixnftNftmngr.query.queryActionByRefIdAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ThesixnetworkSixnftNftmngr.query.queryActionByRefIdAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ActionByRefIdAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryActionByRefIdAll', payload: { options: { all }, params: {...key},query }})
				return getters['getActionByRefIdAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryActionByRefIdAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryOrganization({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ThesixnetworkSixnftNftmngr.query.queryOrganization( key.name)).data
				
					
				commit('QUERY', { query: 'Organization', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryOrganization', payload: { options: { all }, params: {...key},query }})
				return getters['getOrganization']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryOrganization API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryOrganizationAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ThesixnetworkSixnftNftmngr.query.queryOrganizationAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ThesixnetworkSixnftNftmngr.query.queryOrganizationAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'OrganizationAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryOrganizationAll', payload: { options: { all }, params: {...key},query }})
				return getters['getOrganizationAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryOrganizationAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgSetNFTAttribute({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ThesixnetworkSixnftNftmngr.tx.sendMsgSetNFTAttribute({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetNFTAttribute:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSetNFTAttribute:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgPerformActionByAdmin({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ThesixnetworkSixnftNftmngr.tx.sendMsgPerformActionByAdmin({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgPerformActionByAdmin:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgPerformActionByAdmin:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgAddTokenAttribute({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ThesixnetworkSixnftNftmngr.tx.sendMsgAddTokenAttribute({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddTokenAttribute:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgAddTokenAttribute:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgAddAttribute({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ThesixnetworkSixnftNftmngr.tx.sendMsgAddAttribute({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddAttribute:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgAddAttribute:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateNFTSchema({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ThesixnetworkSixnftNftmngr.tx.sendMsgCreateNFTSchema({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateNFTSchema:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateNFTSchema:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateMetadata({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ThesixnetworkSixnftNftmngr.tx.sendMsgCreateMetadata({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateMetadata:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateMetadata:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgAddAction({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ThesixnetworkSixnftNftmngr.tx.sendMsgAddAction({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddAction:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgAddAction:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgSetNFTAttribute({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ThesixnetworkSixnftNftmngr.tx.msgSetNFTAttribute({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetNFTAttribute:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSetNFTAttribute:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgPerformActionByAdmin({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ThesixnetworkSixnftNftmngr.tx.msgPerformActionByAdmin({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgPerformActionByAdmin:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgPerformActionByAdmin:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgAddTokenAttribute({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ThesixnetworkSixnftNftmngr.tx.msgAddTokenAttribute({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddTokenAttribute:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgAddTokenAttribute:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgAddAttribute({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ThesixnetworkSixnftNftmngr.tx.msgAddAttribute({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddAttribute:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgAddAttribute:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateNFTSchema({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ThesixnetworkSixnftNftmngr.tx.msgCreateNFTSchema({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateNFTSchema:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateNFTSchema:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateMetadata({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ThesixnetworkSixnftNftmngr.tx.msgCreateMetadata({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateMetadata:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateMetadata:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgAddAction({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ThesixnetworkSixnftNftmngr.tx.msgAddAction({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddAction:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgAddAction:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
