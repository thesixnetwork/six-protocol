import { txClient, queryClient, MissingWalletError , registry} from './module'

import { ActionParameter } from "./module/types/nftoracle/action_request"
import { ActionOracleParam } from "./module/types/nftoracle/action_request"
import { ActionOracleRequest } from "./module/types/nftoracle/action_request"
import { ActionSignature } from "./module/types/nftoracle/action_signature"
import { ActionSigner } from "./module/types/nftoracle/action_signer"
import { SetSignerSignature } from "./module/types/nftoracle/action_signer"
import { SetSignerParams } from "./module/types/nftoracle/action_signer"
import { ActionSignerConfig } from "./module/types/nftoracle/action_signer_config"
import { _SetSignerParams } from "./module/types/nftoracle/binded_signer"
import { BindedSigner } from "./module/types/nftoracle/binded_signer"
import { OriginContractParam } from "./module/types/nftoracle/collection_owner_request"
import { CollectionOwnerRequest } from "./module/types/nftoracle/collection_owner_request"
import { OriginContractInfo } from "./module/types/nftoracle/collection_owner_request"
import { CollectionOwnerSignature } from "./module/types/nftoracle/collection_owner_signature"
import { MintRequest } from "./module/types/nftoracle/mint_request"
import { Trait } from "./module/types/nftoracle/opensea"
import { OracleConfig } from "./module/types/nftoracle/oracle_config"
import { Params } from "./module/types/nftoracle/params"
import { NftOriginData } from "./module/types/nftoracle/request"
import { DataHash } from "./module/types/nftoracle/request"
import { SyncActionSigner } from "./module/types/nftoracle/sync_action_signer"
import { ParameterSyncSignerByOracle } from "./module/types/nftoracle/sync_action_signer"
import { ContractInfoHash } from "./module/types/nftoracle/sync_action_signer"


export { ActionParameter, ActionOracleParam, ActionOracleRequest, ActionSignature, ActionSigner, SetSignerSignature, SetSignerParams, ActionSignerConfig, _SetSignerParams, BindedSigner, OriginContractParam, CollectionOwnerRequest, OriginContractInfo, CollectionOwnerSignature, MintRequest, Trait, OracleConfig, Params, NftOriginData, DataHash, SyncActionSigner, ParameterSyncSignerByOracle, ContractInfoHash };

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
				ActionOracleRequest: {},
				ActionRequestAll: {},
				CollectionOwnerRequest: {},
				CollectionOwnerRequestAll: {},
				OracleConfig: {},
				ActionSigner: {},
				ActionSignerAll: {},
				BindedSigner: {},
				ActionSignerConfig: {},
				ActionSignerConfigAll: {},
				SyncActionSigner: {},
				SyncActionSignerAll: {},
				
				_Structure: {
						ActionParameter: getStructure(ActionParameter.fromPartial({})),
						ActionOracleParam: getStructure(ActionOracleParam.fromPartial({})),
						ActionOracleRequest: getStructure(ActionOracleRequest.fromPartial({})),
						ActionSignature: getStructure(ActionSignature.fromPartial({})),
						ActionSigner: getStructure(ActionSigner.fromPartial({})),
						SetSignerSignature: getStructure(SetSignerSignature.fromPartial({})),
						SetSignerParams: getStructure(SetSignerParams.fromPartial({})),
						ActionSignerConfig: getStructure(ActionSignerConfig.fromPartial({})),
						_SetSignerParams: getStructure(_SetSignerParams.fromPartial({})),
						BindedSigner: getStructure(BindedSigner.fromPartial({})),
						OriginContractParam: getStructure(OriginContractParam.fromPartial({})),
						CollectionOwnerRequest: getStructure(CollectionOwnerRequest.fromPartial({})),
						OriginContractInfo: getStructure(OriginContractInfo.fromPartial({})),
						CollectionOwnerSignature: getStructure(CollectionOwnerSignature.fromPartial({})),
						MintRequest: getStructure(MintRequest.fromPartial({})),
						Trait: getStructure(Trait.fromPartial({})),
						OracleConfig: getStructure(OracleConfig.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						NftOriginData: getStructure(NftOriginData.fromPartial({})),
						DataHash: getStructure(DataHash.fromPartial({})),
						SyncActionSigner: getStructure(SyncActionSigner.fromPartial({})),
						ParameterSyncSignerByOracle: getStructure(ParameterSyncSignerByOracle.fromPartial({})),
						ContractInfoHash: getStructure(ContractInfoHash.fromPartial({})),
						
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
				getActionOracleRequest: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ActionOracleRequest[JSON.stringify(params)] ?? {}
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
				getActionSigner: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ActionSigner[JSON.stringify(params)] ?? {}
		},
				getActionSignerAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ActionSignerAll[JSON.stringify(params)] ?? {}
		},
				getBindedSigner: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.BindedSigner[JSON.stringify(params)] ?? {}
		},
				getActionSignerConfig: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ActionSignerConfig[JSON.stringify(params)] ?? {}
		},
				getActionSignerConfigAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ActionSignerConfigAll[JSON.stringify(params)] ?? {}
		},
				getSyncActionSigner: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.SyncActionSigner[JSON.stringify(params)] ?? {}
		},
				getSyncActionSignerAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.SyncActionSignerAll[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: thesixnetwork.sixprotocol.nftoracle initialized!')
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
		
		
		
		
		 		
		
		
		async QueryActionOracleRequest({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryActionOracleRequest( key.id)).data
				
					
				commit('QUERY', { query: 'ActionOracleRequest', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryActionOracleRequest', payload: { options: { all }, params: {...key},query }})
				return getters['getActionOracleRequest']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryActionOracleRequest API Node Unavailable. Could not perform query: ' + e.message)
				
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
		
		
		
		
		 		
		
		
		async QueryActionSigner({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryActionSigner( key.actorAddress,  key.ownerAddress)).data
				
					
				commit('QUERY', { query: 'ActionSigner', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryActionSigner', payload: { options: { all }, params: {...key},query }})
				return getters['getActionSigner']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryActionSigner API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryActionSignerAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryActionSignerAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryActionSignerAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ActionSignerAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryActionSignerAll', payload: { options: { all }, params: {...key},query }})
				return getters['getActionSignerAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryActionSignerAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryBindedSigner({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryBindedSigner( key.ownerAddress)).data
				
					
				commit('QUERY', { query: 'BindedSigner', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryBindedSigner', payload: { options: { all }, params: {...key},query }})
				return getters['getBindedSigner']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryBindedSigner API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryActionSignerConfig({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryActionSignerConfig( key.chain)).data
				
					
				commit('QUERY', { query: 'ActionSignerConfig', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryActionSignerConfig', payload: { options: { all }, params: {...key},query }})
				return getters['getActionSignerConfig']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryActionSignerConfig API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryActionSignerConfigAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryActionSignerConfigAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryActionSignerConfigAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ActionSignerConfigAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryActionSignerConfigAll', payload: { options: { all }, params: {...key},query }})
				return getters['getActionSignerConfigAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryActionSignerConfigAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QuerySyncActionSigner({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.querySyncActionSigner( key.id)).data
				
					
				commit('QUERY', { query: 'SyncActionSigner', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QuerySyncActionSigner', payload: { options: { all }, params: {...key},query }})
				return getters['getSyncActionSigner']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QuerySyncActionSigner API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QuerySyncActionSignerAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.querySyncActionSignerAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.querySyncActionSignerAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'SyncActionSignerAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QuerySyncActionSignerAll', payload: { options: { all }, params: {...key},query }})
				return getters['getSyncActionSignerAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QuerySyncActionSignerAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgCreateSyncActionSigner({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateSyncActionSigner(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateSyncActionSigner:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateSyncActionSigner:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSubmitSyncActionSigner({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSubmitSyncActionSigner(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSubmitSyncActionSigner:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSubmitSyncActionSigner:Send Could not broadcast Tx: '+ e.message)
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
		async sendMsgDeleteActionSignerConfig({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgDeleteActionSignerConfig(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteActionSignerConfig:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeleteActionSignerConfig:Send Could not broadcast Tx: '+ e.message)
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
		async sendMsgCreateActionSigner({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateActionSigner(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateActionSigner:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateActionSigner:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeleteActionSigner({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgDeleteActionSigner(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteActionSigner:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeleteActionSigner:Send Could not broadcast Tx: '+ e.message)
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
		async sendMsgUpdateActionSignerConfig({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateActionSignerConfig(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateActionSignerConfig:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateActionSignerConfig:Send Could not broadcast Tx: '+ e.message)
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
		async sendMsgUpdateActionSigner({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateActionSigner(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateActionSigner:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateActionSigner:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateActionSignerConfig({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateActionSignerConfig(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateActionSignerConfig:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateActionSignerConfig:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgCreateSyncActionSigner({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateSyncActionSigner(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateSyncActionSigner:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateSyncActionSigner:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSubmitSyncActionSigner({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSubmitSyncActionSigner(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSubmitSyncActionSigner:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSubmitSyncActionSigner:Create Could not create message: ' + e.message)
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
		async MsgDeleteActionSignerConfig({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgDeleteActionSignerConfig(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteActionSignerConfig:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeleteActionSignerConfig:Create Could not create message: ' + e.message)
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
		async MsgCreateActionSigner({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateActionSigner(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateActionSigner:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateActionSigner:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDeleteActionSigner({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgDeleteActionSigner(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteActionSigner:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeleteActionSigner:Create Could not create message: ' + e.message)
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
		async MsgUpdateActionSignerConfig({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateActionSignerConfig(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateActionSignerConfig:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateActionSignerConfig:Create Could not create message: ' + e.message)
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
		async MsgUpdateActionSigner({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateActionSigner(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateActionSigner:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateActionSigner:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateActionSignerConfig({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateActionSignerConfig(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateActionSignerConfig:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateActionSignerConfig:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
