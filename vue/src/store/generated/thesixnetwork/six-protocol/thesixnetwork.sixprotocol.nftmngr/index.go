import { txClient, queryClient, MissingWalletError , registry} from './module'

import { ActionParams } from "./module/types/nftmngr/action"
import { Action } from "./module/types/nftmngr/action"
import { ActionByRefId } from "./module/types/nftmngr/action_by_ref_id"
import { ActionExecutor } from "./module/types/nftmngr/action_executor"
import { ActionExecutorBySchema } from "./module/types/nftmngr/action_executor"
import { ActionOfSchema } from "./module/types/nftmngr/action_of_schema"
import { DefaultMintValue } from "./module/types/nftmngr/attribute_definition"
import { AttributeDefinition } from "./module/types/nftmngr/attribute_definition"
import { AttributeOfSchema } from "./module/types/nftmngr/attribute_of_schema"
import { DisplayOption } from "./module/types/nftmngr/display_option"
import { ExecutorOfSchema } from "./module/types/nftmngr/executor_of_schema"
import { LockSchemaFee } from "./module/types/nftmngr/lock_schema_fee"
import { MapTokenToMinter } from "./module/types/nftmngr/metadata_creator"
import { MetadataCreator } from "./module/types/nftmngr/metadata_creator"
import { NftAttributeValue } from "./module/types/nftmngr/nft_attribute_value"
import { NumberAttributeValue } from "./module/types/nftmngr/nft_attribute_value"
import { StringAttributeValue } from "./module/types/nftmngr/nft_attribute_value"
import { BooleanAttributeValue } from "./module/types/nftmngr/nft_attribute_value"
import { FloatAttributeValue } from "./module/types/nftmngr/nft_attribute_value"
import { NftCollection } from "./module/types/nftmngr/nft_collection"
import { NftData } from "./module/types/nftmngr/nft_data"
import { NFTFeeBalance } from "./module/types/nftmngr/nft_fee_balance"
import { FeeDistribution } from "./module/types/nftmngr/nft_fee_config"
import { FeeConfig } from "./module/types/nftmngr/nft_fee_config"
import { NFTFeeConfig } from "./module/types/nftmngr/nft_fee_config"
import { NFTSchema } from "./module/types/nftmngr/nft_schema"
import { NFTSchemaINPUT } from "./module/types/nftmngr/nft_schema"
import { NFTSchemaQueryResult } from "./module/types/nftmngr/nft_schema"
import { NFTSchemaByContract } from "./module/types/nftmngr/nft_schema_by_contract"
import { FlagStatus } from "./module/types/nftmngr/on_chain_data"
import { OnChainData } from "./module/types/nftmngr/on_chain_data"
import { OnChainDataResult } from "./module/types/nftmngr/on_chain_data"
import { OpenseaDisplayOption } from "./module/types/nftmngr/opensea_display_option"
import { Organization } from "./module/types/nftmngr/organization"
import { OriginData } from "./module/types/nftmngr/origin_data"
import { Params } from "./module/types/nftmngr/params"
import { SchemaAttribute } from "./module/types/nftmngr/schema_attribute"
import { SchemaAttributeValue } from "./module/types/nftmngr/schema_attribute"
import { OpenseaAttribute } from "./module/types/nftmngr/tx"
import { UpdatedOpenseaAttributes } from "./module/types/nftmngr/tx"
import { UpdatedOriginData } from "./module/types/nftmngr/tx"
import { ActionParameter } from "./module/types/nftmngr/tx"
import { MsgUpdateActionExecutorResponse } from "./module/types/nftmngr/tx"
import { MsgCreateVirtualActionResponse } from "./module/types/nftmngr/tx"
import { MsgUpdateVirtualActionResponse } from "./module/types/nftmngr/tx"
import { MsgDeleteVirtualActionResponse } from "./module/types/nftmngr/tx"
import { TokenIdMap } from "./module/types/nftmngr/tx"
import { VirtualAction } from "./module/types/nftmngr/virtual_action"
import { VirtualSchemaProposal } from "./module/types/nftmngr/virtual_schema"
import { VirtualSchemaProposalRequest } from "./module/types/nftmngr/virtual_schema"
import { VirtualSchema } from "./module/types/nftmngr/virtual_schema"
import { VirtualSchemaRegistry } from "./module/types/nftmngr/virtual_schema"
import { VirtualSchemaRegistryRequest } from "./module/types/nftmngr/virtual_schema"
import { ActiveVirtualSchemaProposal } from "./module/types/nftmngr/virtual_schema"
import { InactiveVirtualSchemaProposal } from "./module/types/nftmngr/virtual_schema"


export { ActionParams, Action, ActionByRefId, ActionExecutor, ActionExecutorBySchema, ActionOfSchema, DefaultMintValue, AttributeDefinition, AttributeOfSchema, DisplayOption, ExecutorOfSchema, LockSchemaFee, MapTokenToMinter, MetadataCreator, NftAttributeValue, NumberAttributeValue, StringAttributeValue, BooleanAttributeValue, FloatAttributeValue, NftCollection, NftData, NFTFeeBalance, FeeDistribution, FeeConfig, NFTFeeConfig, NFTSchema, NFTSchemaINPUT, NFTSchemaQueryResult, NFTSchemaByContract, FlagStatus, OnChainData, OnChainDataResult, OpenseaDisplayOption, Organization, OriginData, Params, SchemaAttribute, SchemaAttributeValue, OpenseaAttribute, UpdatedOpenseaAttributes, UpdatedOriginData, ActionParameter, MsgUpdateActionExecutorResponse, MsgCreateVirtualActionResponse, MsgUpdateVirtualActionResponse, MsgDeleteVirtualActionResponse, TokenIdMap, VirtualAction, VirtualSchemaProposal, VirtualSchemaProposalRequest, VirtualSchema, VirtualSchemaRegistry, VirtualSchemaRegistryRequest, ActiveVirtualSchemaProposal, InactiveVirtualSchemaProposal };

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
				NFTSchema: {},
				NFTSchemaAll: {},
				NftData: {},
				NftDataAll: {},
				ActionByRefId: {},
				ActionByRefIdAll: {},
				Organization: {},
				OrganizationAll: {},
				NftCollection: {},
				NFTSchemaByContract: {},
				NFTSchemaByContractAll: {},
				NFTFeeConfig: {},
				NFTFeeBalance: {},
				MetadataCreator: {},
				MetadataCreatorAll: {},
				ActionExecutor: {},
				ActionExecutorAll: {},
				SchemaAttribute: {},
				SchemaAttributeAll: {},
				ListAttributeBySchema: {},
				ActionOfSchema: {},
				ActionOfSchemaAll: {},
				ExecutorOfSchema: {},
				ExecutorOfSchemaAll: {},
				VirtualAction: {},
				VirtualActionAll: {},
				VirtualSchema: {},
				VirtualSchemaAll: {},
				VirtualSchemaProposal: {},
				VirtualSchemaProposalAll: {},
				ListActiveProposal: {},
				LockSchemaFee: {},
				LockSchemaFeeAll: {},
				
				_Structure: {
						ActionParams: getStructure(ActionParams.fromPartial({})),
						Action: getStructure(Action.fromPartial({})),
						ActionByRefId: getStructure(ActionByRefId.fromPartial({})),
						ActionExecutor: getStructure(ActionExecutor.fromPartial({})),
						ActionExecutorBySchema: getStructure(ActionExecutorBySchema.fromPartial({})),
						ActionOfSchema: getStructure(ActionOfSchema.fromPartial({})),
						DefaultMintValue: getStructure(DefaultMintValue.fromPartial({})),
						AttributeDefinition: getStructure(AttributeDefinition.fromPartial({})),
						AttributeOfSchema: getStructure(AttributeOfSchema.fromPartial({})),
						DisplayOption: getStructure(DisplayOption.fromPartial({})),
						ExecutorOfSchema: getStructure(ExecutorOfSchema.fromPartial({})),
						LockSchemaFee: getStructure(LockSchemaFee.fromPartial({})),
						MapTokenToMinter: getStructure(MapTokenToMinter.fromPartial({})),
						MetadataCreator: getStructure(MetadataCreator.fromPartial({})),
						NftAttributeValue: getStructure(NftAttributeValue.fromPartial({})),
						NumberAttributeValue: getStructure(NumberAttributeValue.fromPartial({})),
						StringAttributeValue: getStructure(StringAttributeValue.fromPartial({})),
						BooleanAttributeValue: getStructure(BooleanAttributeValue.fromPartial({})),
						FloatAttributeValue: getStructure(FloatAttributeValue.fromPartial({})),
						NftCollection: getStructure(NftCollection.fromPartial({})),
						NftData: getStructure(NftData.fromPartial({})),
						NFTFeeBalance: getStructure(NFTFeeBalance.fromPartial({})),
						FeeDistribution: getStructure(FeeDistribution.fromPartial({})),
						FeeConfig: getStructure(FeeConfig.fromPartial({})),
						NFTFeeConfig: getStructure(NFTFeeConfig.fromPartial({})),
						NFTSchema: getStructure(NFTSchema.fromPartial({})),
						NFTSchemaINPUT: getStructure(NFTSchemaINPUT.fromPartial({})),
						NFTSchemaQueryResult: getStructure(NFTSchemaQueryResult.fromPartial({})),
						NFTSchemaByContract: getStructure(NFTSchemaByContract.fromPartial({})),
						FlagStatus: getStructure(FlagStatus.fromPartial({})),
						OnChainData: getStructure(OnChainData.fromPartial({})),
						OnChainDataResult: getStructure(OnChainDataResult.fromPartial({})),
						OpenseaDisplayOption: getStructure(OpenseaDisplayOption.fromPartial({})),
						Organization: getStructure(Organization.fromPartial({})),
						OriginData: getStructure(OriginData.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						SchemaAttribute: getStructure(SchemaAttribute.fromPartial({})),
						SchemaAttributeValue: getStructure(SchemaAttributeValue.fromPartial({})),
						OpenseaAttribute: getStructure(OpenseaAttribute.fromPartial({})),
						UpdatedOpenseaAttributes: getStructure(UpdatedOpenseaAttributes.fromPartial({})),
						UpdatedOriginData: getStructure(UpdatedOriginData.fromPartial({})),
						ActionParameter: getStructure(ActionParameter.fromPartial({})),
						MsgUpdateActionExecutorResponse: getStructure(MsgUpdateActionExecutorResponse.fromPartial({})),
						MsgCreateVirtualActionResponse: getStructure(MsgCreateVirtualActionResponse.fromPartial({})),
						MsgUpdateVirtualActionResponse: getStructure(MsgUpdateVirtualActionResponse.fromPartial({})),
						MsgDeleteVirtualActionResponse: getStructure(MsgDeleteVirtualActionResponse.fromPartial({})),
						TokenIdMap: getStructure(TokenIdMap.fromPartial({})),
						VirtualAction: getStructure(VirtualAction.fromPartial({})),
						VirtualSchemaProposal: getStructure(VirtualSchemaProposal.fromPartial({})),
						VirtualSchemaProposalRequest: getStructure(VirtualSchemaProposalRequest.fromPartial({})),
						VirtualSchema: getStructure(VirtualSchema.fromPartial({})),
						VirtualSchemaRegistry: getStructure(VirtualSchemaRegistry.fromPartial({})),
						VirtualSchemaRegistryRequest: getStructure(VirtualSchemaRegistryRequest.fromPartial({})),
						ActiveVirtualSchemaProposal: getStructure(ActiveVirtualSchemaProposal.fromPartial({})),
						InactiveVirtualSchemaProposal: getStructure(InactiveVirtualSchemaProposal.fromPartial({})),
						
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
				getNftCollection: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.NftCollection[JSON.stringify(params)] ?? {}
		},
				getNFTSchemaByContract: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.NFTSchemaByContract[JSON.stringify(params)] ?? {}
		},
				getNFTSchemaByContractAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.NFTSchemaByContractAll[JSON.stringify(params)] ?? {}
		},
				getNFTFeeConfig: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.NFTFeeConfig[JSON.stringify(params)] ?? {}
		},
				getNFTFeeBalance: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.NFTFeeBalance[JSON.stringify(params)] ?? {}
		},
				getMetadataCreator: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.MetadataCreator[JSON.stringify(params)] ?? {}
		},
				getMetadataCreatorAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.MetadataCreatorAll[JSON.stringify(params)] ?? {}
		},
				getActionExecutor: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ActionExecutor[JSON.stringify(params)] ?? {}
		},
				getActionExecutorAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ActionExecutorAll[JSON.stringify(params)] ?? {}
		},
				getSchemaAttribute: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.SchemaAttribute[JSON.stringify(params)] ?? {}
		},
				getSchemaAttributeAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.SchemaAttributeAll[JSON.stringify(params)] ?? {}
		},
				getListAttributeBySchema: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ListAttributeBySchema[JSON.stringify(params)] ?? {}
		},
				getActionOfSchema: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ActionOfSchema[JSON.stringify(params)] ?? {}
		},
				getActionOfSchemaAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ActionOfSchemaAll[JSON.stringify(params)] ?? {}
		},
				getExecutorOfSchema: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ExecutorOfSchema[JSON.stringify(params)] ?? {}
		},
				getExecutorOfSchemaAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ExecutorOfSchemaAll[JSON.stringify(params)] ?? {}
		},
				getVirtualAction: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.VirtualAction[JSON.stringify(params)] ?? {}
		},
				getVirtualActionAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.VirtualActionAll[JSON.stringify(params)] ?? {}
		},
				getVirtualSchema: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.VirtualSchema[JSON.stringify(params)] ?? {}
		},
				getVirtualSchemaAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.VirtualSchemaAll[JSON.stringify(params)] ?? {}
		},
				getVirtualSchemaProposal: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.VirtualSchemaProposal[JSON.stringify(params)] ?? {}
		},
				getVirtualSchemaProposalAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.VirtualSchemaProposalAll[JSON.stringify(params)] ?? {}
		},
				getListActiveProposal: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ListActiveProposal[JSON.stringify(params)] ?? {}
		},
				getLockSchemaFee: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.LockSchemaFee[JSON.stringify(params)] ?? {}
		},
				getLockSchemaFeeAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.LockSchemaFeeAll[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: thesixnetwork.sixprotocol.nftmngr initialized!')
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
		
		
		
		
		 		
		
		
		async QueryNFTSchema({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryNFTSchema( key.code)).data
				
					
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
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryNFTSchemaAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryNFTSchemaAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
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
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryNftData( key.nftSchemaCode,  key.tokenId, query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryNftData( key.nftSchemaCode,  key.tokenId, {...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
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
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryNftDataAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryNftDataAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
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
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryActionByRefId( key.refId)).data
				
					
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
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryActionByRefIdAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryActionByRefIdAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
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
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryOrganization( key.name)).data
				
					
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
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryOrganizationAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryOrganizationAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'OrganizationAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryOrganizationAll', payload: { options: { all }, params: {...key},query }})
				return getters['getOrganizationAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryOrganizationAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryNftCollection({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryNftCollection( key.nftSchemaCode, query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryNftCollection( key.nftSchemaCode, {...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'NftCollection', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryNftCollection', payload: { options: { all }, params: {...key},query }})
				return getters['getNftCollection']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryNftCollection API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryNFTSchemaByContract({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryNFTSchemaByContract( key.originContractAddress)).data
				
					
				commit('QUERY', { query: 'NFTSchemaByContract', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryNFTSchemaByContract', payload: { options: { all }, params: {...key},query }})
				return getters['getNFTSchemaByContract']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryNFTSchemaByContract API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryNFTSchemaByContractAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryNFTSchemaByContractAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryNFTSchemaByContractAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'NFTSchemaByContractAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryNFTSchemaByContractAll', payload: { options: { all }, params: {...key},query }})
				return getters['getNFTSchemaByContractAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryNFTSchemaByContractAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryNFTFeeConfig({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryNFTFeeConfig()).data
				
					
				commit('QUERY', { query: 'NFTFeeConfig', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryNFTFeeConfig', payload: { options: { all }, params: {...key},query }})
				return getters['getNFTFeeConfig']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryNFTFeeConfig API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryNFTFeeBalance({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryNFTFeeBalance()).data
				
					
				commit('QUERY', { query: 'NFTFeeBalance', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryNFTFeeBalance', payload: { options: { all }, params: {...key},query }})
				return getters['getNFTFeeBalance']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryNFTFeeBalance API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryMetadataCreator({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryMetadataCreator( key.nftSchemaCode)).data
				
					
				commit('QUERY', { query: 'MetadataCreator', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryMetadataCreator', payload: { options: { all }, params: {...key},query }})
				return getters['getMetadataCreator']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryMetadataCreator API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryMetadataCreatorAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryMetadataCreatorAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryMetadataCreatorAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'MetadataCreatorAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryMetadataCreatorAll', payload: { options: { all }, params: {...key},query }})
				return getters['getMetadataCreatorAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryMetadataCreatorAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryActionExecutor({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryActionExecutor( key.nftSchemaCode,  key.executorAddress)).data
				
					
				commit('QUERY', { query: 'ActionExecutor', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryActionExecutor', payload: { options: { all }, params: {...key},query }})
				return getters['getActionExecutor']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryActionExecutor API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryActionExecutorAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryActionExecutorAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryActionExecutorAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ActionExecutorAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryActionExecutorAll', payload: { options: { all }, params: {...key},query }})
				return getters['getActionExecutorAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryActionExecutorAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QuerySchemaAttribute({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.querySchemaAttribute( key.nftSchemaCode,  key.name)).data
				
					
				commit('QUERY', { query: 'SchemaAttribute', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QuerySchemaAttribute', payload: { options: { all }, params: {...key},query }})
				return getters['getSchemaAttribute']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QuerySchemaAttribute API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QuerySchemaAttributeAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.querySchemaAttributeAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.querySchemaAttributeAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'SchemaAttributeAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QuerySchemaAttributeAll', payload: { options: { all }, params: {...key},query }})
				return getters['getSchemaAttributeAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QuerySchemaAttributeAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryListAttributeBySchema({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryListAttributeBySchema( key.nftSchemaCode)).data
				
					
				commit('QUERY', { query: 'ListAttributeBySchema', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryListAttributeBySchema', payload: { options: { all }, params: {...key},query }})
				return getters['getListAttributeBySchema']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryListAttributeBySchema API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryActionOfSchema({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryActionOfSchema( key.nftSchemaCode,  key.name)).data
				
					
				commit('QUERY', { query: 'ActionOfSchema', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryActionOfSchema', payload: { options: { all }, params: {...key},query }})
				return getters['getActionOfSchema']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryActionOfSchema API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryActionOfSchemaAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryActionOfSchemaAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryActionOfSchemaAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ActionOfSchemaAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryActionOfSchemaAll', payload: { options: { all }, params: {...key},query }})
				return getters['getActionOfSchemaAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryActionOfSchemaAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryExecutorOfSchema({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryExecutorOfSchema( key.nftSchemaCode)).data
				
					
				commit('QUERY', { query: 'ExecutorOfSchema', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryExecutorOfSchema', payload: { options: { all }, params: {...key},query }})
				return getters['getExecutorOfSchema']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryExecutorOfSchema API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryExecutorOfSchemaAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryExecutorOfSchemaAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryExecutorOfSchemaAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ExecutorOfSchemaAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryExecutorOfSchemaAll', payload: { options: { all }, params: {...key},query }})
				return getters['getExecutorOfSchemaAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryExecutorOfSchemaAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryVirtualAction({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryVirtualAction( key.nftSchemaCode, query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryVirtualAction( key.nftSchemaCode, {...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'VirtualAction', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryVirtualAction', payload: { options: { all }, params: {...key},query }})
				return getters['getVirtualAction']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryVirtualAction API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryVirtualActionAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryVirtualActionAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryVirtualActionAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'VirtualActionAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryVirtualActionAll', payload: { options: { all }, params: {...key},query }})
				return getters['getVirtualActionAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryVirtualActionAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryVirtualSchema({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryVirtualSchema( key.nftSchemaCode)).data
				
					
				commit('QUERY', { query: 'VirtualSchema', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryVirtualSchema', payload: { options: { all }, params: {...key},query }})
				return getters['getVirtualSchema']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryVirtualSchema API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryVirtualSchemaAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryVirtualSchemaAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryVirtualSchemaAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'VirtualSchemaAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryVirtualSchemaAll', payload: { options: { all }, params: {...key},query }})
				return getters['getVirtualSchemaAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryVirtualSchemaAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryVirtualSchemaProposal({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryVirtualSchemaProposal( key.index)).data
				
					
				commit('QUERY', { query: 'VirtualSchemaProposal', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryVirtualSchemaProposal', payload: { options: { all }, params: {...key},query }})
				return getters['getVirtualSchemaProposal']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryVirtualSchemaProposal API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryVirtualSchemaProposalAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryVirtualSchemaProposalAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryVirtualSchemaProposalAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'VirtualSchemaProposalAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryVirtualSchemaProposalAll', payload: { options: { all }, params: {...key},query }})
				return getters['getVirtualSchemaProposalAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryVirtualSchemaProposalAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryListActiveProposal({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryListActiveProposal()).data
				
					
				commit('QUERY', { query: 'ListActiveProposal', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryListActiveProposal', payload: { options: { all }, params: {...key},query }})
				return getters['getListActiveProposal']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryListActiveProposal API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryLockSchemaFee({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryLockSchemaFee( key.index)).data
				
					
				commit('QUERY', { query: 'LockSchemaFee', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryLockSchemaFee', payload: { options: { all }, params: {...key},query }})
				return getters['getLockSchemaFee']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryLockSchemaFee API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryLockSchemaFeeAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryLockSchemaFeeAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryLockSchemaFeeAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'LockSchemaFeeAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryLockSchemaFeeAll', payload: { options: { all }, params: {...key},query }})
				return getters['getLockSchemaFeeAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryLockSchemaFeeAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgSetMetadataFormat({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetMetadataFormat(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetMetadataFormat:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSetMetadataFormat:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateActionExecutor({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateActionExecutor(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateActionExecutor:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateActionExecutor:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgToggleAction({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgToggleAction(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgToggleAction:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgToggleAction:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgAddAttribute({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgAddAttribute(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddAttribute:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgAddAttribute:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgVoteVirtualSchemaProposal({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgVoteVirtualSchemaProposal(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgVoteVirtualSchemaProposal:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgVoteVirtualSchemaProposal:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgShowAttributes({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgShowAttributes(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgShowAttributes:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgShowAttributes:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSetAttributeOveriding({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetAttributeOveriding(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetAttributeOveriding:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSetAttributeOveriding:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSetUriRetrievalMethod({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetUriRetrievalMethod(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetUriRetrievalMethod:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSetUriRetrievalMethod:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgPerformVirtualAction({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgPerformVirtualAction(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgPerformVirtualAction:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgPerformVirtualAction:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSetFeeConfig({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetFeeConfig(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetFeeConfig:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSetFeeConfig:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSetBaseUri({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetBaseUri(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetBaseUri:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSetBaseUri:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSetOriginContract({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetOriginContract(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetOriginContract:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSetOriginContract:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgAddAction({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgAddAction(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddAction:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgAddAction:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateVirtualAction({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateVirtualAction(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateVirtualAction:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateVirtualAction:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateSchemaAttribute({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateSchemaAttribute(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateSchemaAttribute:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateSchemaAttribute:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgChangeOrgOwner({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgChangeOrgOwner(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgChangeOrgOwner:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgChangeOrgOwner:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateNFTSchema({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateNFTSchema(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
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
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateMetadata(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateMetadata:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateMetadata:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSetOriginChain({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetOriginChain(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetOriginChain:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSetOriginChain:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateAction({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateAction(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateAction:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateAction:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgPerformActionByAdmin({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgPerformActionByAdmin(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgPerformActionByAdmin:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgPerformActionByAdmin:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeleteActionExecutor({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgDeleteActionExecutor(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteActionExecutor:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeleteActionExecutor:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgChangeSchemaOwner({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgChangeSchemaOwner(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgChangeSchemaOwner:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgChangeSchemaOwner:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgProposalVirtualSchema({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgProposalVirtualSchema(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgProposalVirtualSchema:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgProposalVirtualSchema:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgResyncAttributes({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgResyncAttributes(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgResyncAttributes:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgResyncAttributes:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateVirtualAction({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateVirtualAction(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateVirtualAction:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateVirtualAction:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSetMintauth({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetMintauth(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetMintauth:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSetMintauth:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeleteVirtualAction({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgDeleteVirtualAction(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteVirtualAction:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeleteVirtualAction:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateActionExecutor({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateActionExecutor(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateActionExecutor:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateActionExecutor:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgSetMetadataFormat({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetMetadataFormat(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetMetadataFormat:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSetMetadataFormat:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateActionExecutor({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateActionExecutor(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateActionExecutor:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateActionExecutor:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgToggleAction({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgToggleAction(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgToggleAction:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgToggleAction:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgAddAttribute({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgAddAttribute(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddAttribute:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgAddAttribute:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgVoteVirtualSchemaProposal({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgVoteVirtualSchemaProposal(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgVoteVirtualSchemaProposal:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgVoteVirtualSchemaProposal:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgShowAttributes({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgShowAttributes(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgShowAttributes:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgShowAttributes:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSetAttributeOveriding({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetAttributeOveriding(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetAttributeOveriding:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSetAttributeOveriding:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSetUriRetrievalMethod({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetUriRetrievalMethod(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetUriRetrievalMethod:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSetUriRetrievalMethod:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgPerformVirtualAction({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgPerformVirtualAction(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgPerformVirtualAction:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgPerformVirtualAction:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSetFeeConfig({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetFeeConfig(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetFeeConfig:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSetFeeConfig:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSetBaseUri({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetBaseUri(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetBaseUri:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSetBaseUri:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSetOriginContract({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetOriginContract(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetOriginContract:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSetOriginContract:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgAddAction({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgAddAction(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddAction:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgAddAction:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateVirtualAction({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateVirtualAction(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateVirtualAction:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateVirtualAction:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateSchemaAttribute({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateSchemaAttribute(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateSchemaAttribute:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateSchemaAttribute:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgChangeOrgOwner({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgChangeOrgOwner(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgChangeOrgOwner:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgChangeOrgOwner:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateNFTSchema({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateNFTSchema(value)
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
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateMetadata(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateMetadata:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateMetadata:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSetOriginChain({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetOriginChain(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetOriginChain:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSetOriginChain:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateAction({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateAction(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateAction:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateAction:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgPerformActionByAdmin({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgPerformActionByAdmin(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgPerformActionByAdmin:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgPerformActionByAdmin:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDeleteActionExecutor({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgDeleteActionExecutor(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteActionExecutor:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeleteActionExecutor:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgChangeSchemaOwner({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgChangeSchemaOwner(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgChangeSchemaOwner:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgChangeSchemaOwner:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgProposalVirtualSchema({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgProposalVirtualSchema(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgProposalVirtualSchema:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgProposalVirtualSchema:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgResyncAttributes({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgResyncAttributes(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgResyncAttributes:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgResyncAttributes:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateVirtualAction({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateVirtualAction(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateVirtualAction:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateVirtualAction:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSetMintauth({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetMintauth(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetMintauth:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSetMintauth:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDeleteVirtualAction({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgDeleteVirtualAction(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteVirtualAction:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeleteVirtualAction:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateActionExecutor({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateActionExecutor(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateActionExecutor:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateActionExecutor:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
