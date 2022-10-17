import { Client, registry, MissingWalletError } from 'thesixnetwork-six-protocol-client-ts'

import { Burn } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixprotocol.tokenmngr/types"
import { Mintperm } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixprotocol.tokenmngr/types"
import { Options } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixprotocol.tokenmngr/types"
import { TokenmngrPacketData } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixprotocol.tokenmngr/types"
import { NoData } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixprotocol.tokenmngr/types"
import { Params } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixprotocol.tokenmngr/types"
import { Token } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixprotocol.tokenmngr/types"
import { DenomUnit } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixprotocol.tokenmngr/types"
import { Metadata } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixprotocol.tokenmngr/types"
import { TokenBurn } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixprotocol.tokenmngr/types"


export { Burn, Mintperm, Options, TokenmngrPacketData, NoData, Params, Token, DenomUnit, Metadata, TokenBurn };

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
				Token: {},
				TokenAll: {},
				Mintperm: {},
				MintpermAll: {},
				Options: {},
				Burns: {},
				TokenBurn: {},
				TokenBurnAll: {},
				
				_Structure: {
						Burn: getStructure(Burn.fromPartial({})),
						Mintperm: getStructure(Mintperm.fromPartial({})),
						Options: getStructure(Options.fromPartial({})),
						TokenmngrPacketData: getStructure(TokenmngrPacketData.fromPartial({})),
						NoData: getStructure(NoData.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						Token: getStructure(Token.fromPartial({})),
						DenomUnit: getStructure(DenomUnit.fromPartial({})),
						Metadata: getStructure(Metadata.fromPartial({})),
						TokenBurn: getStructure(TokenBurn.fromPartial({})),
						
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
				getToken: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Token[JSON.stringify(params)] ?? {}
		},
				getTokenAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.TokenAll[JSON.stringify(params)] ?? {}
		},
				getMintperm: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Mintperm[JSON.stringify(params)] ?? {}
		},
				getMintpermAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.MintpermAll[JSON.stringify(params)] ?? {}
		},
				getOptions: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Options[JSON.stringify(params)] ?? {}
		},
				getBurns: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Burns[JSON.stringify(params)] ?? {}
		},
				getTokenBurn: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.TokenBurn[JSON.stringify(params)] ?? {}
		},
				getTokenBurnAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.TokenBurnAll[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: thesixnetwork.sixprotocol.tokenmngr initialized!')
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
				let value= (await client.ThesixnetworkSixprotocolTokenmngr.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryToken({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ThesixnetworkSixprotocolTokenmngr.query.queryToken( key.name)).data
				
					
				commit('QUERY', { query: 'Token', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryToken', payload: { options: { all }, params: {...key},query }})
				return getters['getToken']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryToken API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryTokenAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ThesixnetworkSixprotocolTokenmngr.query.queryTokenAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ThesixnetworkSixprotocolTokenmngr.query.queryTokenAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'TokenAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryTokenAll', payload: { options: { all }, params: {...key},query }})
				return getters['getTokenAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryTokenAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryMintperm({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ThesixnetworkSixprotocolTokenmngr.query.queryMintperm( key.token,  key.address)).data
				
					
				commit('QUERY', { query: 'Mintperm', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryMintperm', payload: { options: { all }, params: {...key},query }})
				return getters['getMintperm']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryMintperm API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryMintpermAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ThesixnetworkSixprotocolTokenmngr.query.queryMintpermAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ThesixnetworkSixprotocolTokenmngr.query.queryMintpermAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'MintpermAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryMintpermAll', payload: { options: { all }, params: {...key},query }})
				return getters['getMintpermAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryMintpermAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryOptions({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ThesixnetworkSixprotocolTokenmngr.query.queryOptions()).data
				
					
				commit('QUERY', { query: 'Options', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryOptions', payload: { options: { all }, params: {...key},query }})
				return getters['getOptions']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryOptions API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryBurns({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ThesixnetworkSixprotocolTokenmngr.query.queryBurns(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ThesixnetworkSixprotocolTokenmngr.query.queryBurns({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'Burns', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryBurns', payload: { options: { all }, params: {...key},query }})
				return getters['getBurns']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryBurns API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryTokenBurn({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ThesixnetworkSixprotocolTokenmngr.query.queryTokenBurn( key.token)).data
				
					
				commit('QUERY', { query: 'TokenBurn', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryTokenBurn', payload: { options: { all }, params: {...key},query }})
				return getters['getTokenBurn']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryTokenBurn API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryTokenBurnAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ThesixnetworkSixprotocolTokenmngr.query.queryTokenBurnAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ThesixnetworkSixprotocolTokenmngr.query.queryTokenBurnAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'TokenBurnAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryTokenBurnAll', payload: { options: { all }, params: {...key},query }})
				return getters['getTokenBurnAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryTokenBurnAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgMint({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ThesixnetworkSixprotocolTokenmngr.tx.sendMsgMint({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgMint:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgMint:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeleteToken({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ThesixnetworkSixprotocolTokenmngr.tx.sendMsgDeleteToken({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteToken:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeleteToken:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeleteOptions({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ThesixnetworkSixprotocolTokenmngr.tx.sendMsgDeleteOptions({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteOptions:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeleteOptions:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgBurn({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ThesixnetworkSixprotocolTokenmngr.tx.sendMsgBurn({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgBurn:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgBurn:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateOptions({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ThesixnetworkSixprotocolTokenmngr.tx.sendMsgUpdateOptions({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateOptions:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateOptions:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateMintperm({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ThesixnetworkSixprotocolTokenmngr.tx.sendMsgCreateMintperm({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateMintperm:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateMintperm:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateMintperm({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ThesixnetworkSixprotocolTokenmngr.tx.sendMsgUpdateMintperm({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateMintperm:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateMintperm:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeleteMintperm({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ThesixnetworkSixprotocolTokenmngr.tx.sendMsgDeleteMintperm({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteMintperm:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeleteMintperm:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateOptions({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ThesixnetworkSixprotocolTokenmngr.tx.sendMsgCreateOptions({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateOptions:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateOptions:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateToken({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ThesixnetworkSixprotocolTokenmngr.tx.sendMsgCreateToken({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateToken:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateToken:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateToken({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ThesixnetworkSixprotocolTokenmngr.tx.sendMsgUpdateToken({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateToken:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateToken:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgMint({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ThesixnetworkSixprotocolTokenmngr.tx.msgMint({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgMint:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgMint:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDeleteToken({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ThesixnetworkSixprotocolTokenmngr.tx.msgDeleteToken({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteToken:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeleteToken:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDeleteOptions({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ThesixnetworkSixprotocolTokenmngr.tx.msgDeleteOptions({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteOptions:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeleteOptions:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgBurn({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ThesixnetworkSixprotocolTokenmngr.tx.msgBurn({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgBurn:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgBurn:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateOptions({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ThesixnetworkSixprotocolTokenmngr.tx.msgUpdateOptions({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateOptions:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateOptions:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateMintperm({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ThesixnetworkSixprotocolTokenmngr.tx.msgCreateMintperm({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateMintperm:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateMintperm:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateMintperm({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ThesixnetworkSixprotocolTokenmngr.tx.msgUpdateMintperm({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateMintperm:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateMintperm:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDeleteMintperm({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ThesixnetworkSixprotocolTokenmngr.tx.msgDeleteMintperm({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteMintperm:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeleteMintperm:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateOptions({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ThesixnetworkSixprotocolTokenmngr.tx.msgCreateOptions({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateOptions:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateOptions:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateToken({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ThesixnetworkSixprotocolTokenmngr.tx.msgCreateToken({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateToken:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateToken:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateToken({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ThesixnetworkSixprotocolTokenmngr.tx.msgUpdateToken({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateToken:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateToken:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
