import { txClient, queryClient, MissingWalletError , registry} from './module'

import { ActionParam } from "./module/types/nftoracle/action_request"
import { ActionRequest } from "./module/types/nftoracle/action_request"
import { ActionSignature } from "./module/types/nftoracle/action_signature"
import { OriginContractParam } from "./module/types/nftoracle/collection_owner_request"
import { OriginContractInfo } from "./module/types/nftoracle/collection_owner_request"
import { CollectionOwnerRequest } from "./module/types/nftoracle/collection_owner_request"
import { CollectionOwnerSignature } from "./module/types/nftoracle/collection_owner_signature"
import { MintRequest } from "./module/types/nftoracle/mint_request"
import { Trait } from "./module/types/nftoracle/opensea"
import { OracleConfig } from "./module/types/nftoracle/oracle_config"
import { Params } from "./module/types/nftoracle/params"
import { NftOriginData } from "./module/types/nftoracle/request"
import { DataHash } from "./module/types/nftoracle/request"


export { ActionParam, ActionRequest, ActionSignature, OriginContractParam, OriginContractInfo, CollectionOwnerRequest, CollectionOwnerSignature, MintRequest, Trait, OracleConfig, Params, NftOriginData, DataHash };

async function initTxClient(vuexGetters) {
	return await txClient(vuexGetters['common/wallet/signer'], {
		addr: vuexGetters['common/env/apiTendermint']
	})
}

async function initQueryClient(vuexGetters) {
	return await queryClient({
		addr: vuexGetters['common/env/apiCosmos']
	})
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

function getStructure(template) {
	let structure = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field: any = {}
		field.name = key
		field.type = typeof value
		structure.fields.push(field)
	}
	return structure
}

const getDefaultState = () => {
	return {
				Params: {},
				MintRequest: {},
				MintRequestAll: {},
				ActionRequest: {},
				ActionRequestAll: {},
				CollectionOwnerRequest: {},
				CollectionOwnerRequestAll: {},
				OracleConfig: {},
				
				_Structure: {
						ActionParam: getStructure(ActionParam.fromPartial({})),
						ActionRequest: getStructure(ActionRequest.fromPartial({})),
						ActionSignature: getStructure(ActionSignature.fromPartial({})),
						OriginContractParam: getStructure(OriginContractParam.fromPartial({})),
						OriginContractInfo: getStructure(OriginContractInfo.fromPartial({})),
						CollectionOwnerRequest: getStructure(CollectionOwnerRequest.fromPartial({})),
						CollectionOwnerSignature: getStructure(CollectionOwnerSignature.fromPartial({})),
						MintRequest: getStructure(MintRequest.fromPartial({})),
						Trait: getStructure(Trait.fromPartial({})),
						OracleConfig: getStructure(OracleConfig.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						NftOriginData: getStructure(NftOriginData.fromPartial({})),
						DataHash: getStructure(DataHash.fromPartial({})),
						
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
				getMintRequest: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.MintRequest[JSON.stringify(params)] ?? {}
		},
				getMintRequestAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.MintRequestAll[JSON.stringify(params)] ?? {}
		},
				getActionRequest: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ActionRequest[JSON.stringify(params)] ?? {}
		},
				getActionRequestAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ActionRequestAll[JSON.stringify(params)] ?? {}
		},
				getCollectionOwnerRequest: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.CollectionOwnerRequest[JSON.stringify(params)] ?? {}
		},
				getCollectionOwnerRequestAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.CollectionOwnerRequestAll[JSON.stringify(params)] ?? {}
		},
				getOracleConfig: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.OracleConfig[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: thesixnetwork.sixnft.nftoracle initialized!')
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
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryMintRequest({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryMintRequest( key.id)).data
				
					
				commit('QUERY', { query: 'MintRequest', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryMintRequest', payload: { options: { all }, params: {...key},query }})
				return getters['getMintRequest']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryMintRequest API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryMintRequestAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryMintRequestAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryMintRequestAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'MintRequestAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryMintRequestAll', payload: { options: { all }, params: {...key},query }})
				return getters['getMintRequestAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryMintRequestAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryActionRequest({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryActionRequest( key.id)).data
				
					
				commit('QUERY', { query: 'ActionRequest', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryActionRequest', payload: { options: { all }, params: {...key},query }})
				return getters['getActionRequest']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryActionRequest API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryActionRequestAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryActionRequestAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryActionRequestAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ActionRequestAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryActionRequestAll', payload: { options: { all }, params: {...key},query }})
				return getters['getActionRequestAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryActionRequestAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryCollectionOwnerRequest({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryCollectionOwnerRequest( key.id)).data
				
					
				commit('QUERY', { query: 'CollectionOwnerRequest', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryCollectionOwnerRequest', payload: { options: { all }, params: {...key},query }})
				return getters['getCollectionOwnerRequest']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryCollectionOwnerRequest API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryCollectionOwnerRequestAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryCollectionOwnerRequestAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryCollectionOwnerRequestAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'CollectionOwnerRequestAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryCollectionOwnerRequestAll', payload: { options: { all }, params: {...key},query }})
				return getters['getCollectionOwnerRequestAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryCollectionOwnerRequestAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryOracleConfig({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryOracleConfig()).data
				
					
				commit('QUERY', { query: 'OracleConfig', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryOracleConfig', payload: { options: { all }, params: {...key},query }})
				return getters['getOracleConfig']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryOracleConfig API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgSubmitActionResponse({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSubmitActionResponse(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSubmitActionResponse:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSubmitActionResponse:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateVerifyCollectionOwnerRequest({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateVerifyCollectionOwnerRequest(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateVerifyCollectionOwnerRequest:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateVerifyCollectionOwnerRequest:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSubmitMintResponse({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSubmitMintResponse(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSubmitMintResponse:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSubmitMintResponse:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSubmitVerifyCollectionOwner({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSubmitVerifyCollectionOwner(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSubmitVerifyCollectionOwner:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSubmitVerifyCollectionOwner:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateActionRequest({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateActionRequest(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateActionRequest:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateActionRequest:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateMintRequest({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateMintRequest(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateMintRequest:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateMintRequest:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSetMinimumConfirmation({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetMinimumConfirmation(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetMinimumConfirmation:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSetMinimumConfirmation:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgSubmitActionResponse({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSubmitActionResponse(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSubmitActionResponse:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSubmitActionResponse:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateVerifyCollectionOwnerRequest({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateVerifyCollectionOwnerRequest(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateVerifyCollectionOwnerRequest:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateVerifyCollectionOwnerRequest:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSubmitMintResponse({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSubmitMintResponse(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSubmitMintResponse:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSubmitMintResponse:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSubmitVerifyCollectionOwner({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSubmitVerifyCollectionOwner(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSubmitVerifyCollectionOwner:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSubmitVerifyCollectionOwner:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateActionRequest({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateActionRequest(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateActionRequest:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateActionRequest:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateMintRequest({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateMintRequest(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateMintRequest:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateMintRequest:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSetMinimumConfirmation({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetMinimumConfirmation(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetMinimumConfirmation:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSetMinimumConfirmation:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
