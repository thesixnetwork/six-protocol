import { txClient, queryClient, MissingWalletError , registry} from './module'

import { Attestation } from "./module/types/gravity/v1/attestation"
import { ERC20Token } from "./module/types/gravity/v1/attestation"
import { EventObservation } from "./module/types/gravity/v1/attestation"
import { EventInvalidSendToCosmosReceiver } from "./module/types/gravity/v1/attestation"
import { EventSendToCosmos } from "./module/types/gravity/v1/attestation"
import { EventSendToCosmosLocal } from "./module/types/gravity/v1/attestation"
import { EventSendToCosmosPendingIbcAutoForward } from "./module/types/gravity/v1/attestation"
import { EventSendToCosmosExecutedIbcAutoForward } from "./module/types/gravity/v1/attestation"
import { OutgoingTxBatch } from "./module/types/gravity/v1/batch"
import { OutgoingTransferTx } from "./module/types/gravity/v1/batch"
import { OutgoingLogicCall } from "./module/types/gravity/v1/batch"
import { EventOutgoingBatchCanceled } from "./module/types/gravity/v1/batch"
import { EventOutgoingBatch } from "./module/types/gravity/v1/batch"
import { Params } from "./module/types/gravity/v1/genesis"
import { GravityNonces } from "./module/types/gravity/v1/genesis"
import { EventSetOperatorAddress } from "./module/types/gravity/v1/msgs"
import { EventValsetConfirmKey } from "./module/types/gravity/v1/msgs"
import { EventBatchCreated } from "./module/types/gravity/v1/msgs"
import { EventBatchConfirmKey } from "./module/types/gravity/v1/msgs"
import { EventBatchSendToEthClaim } from "./module/types/gravity/v1/msgs"
import { EventClaim } from "./module/types/gravity/v1/msgs"
import { EventBadSignatureEvidence } from "./module/types/gravity/v1/msgs"
import { EventERC20DeployedClaim } from "./module/types/gravity/v1/msgs"
import { EventValsetUpdatedClaim } from "./module/types/gravity/v1/msgs"
import { EventMultisigUpdateRequest } from "./module/types/gravity/v1/msgs"
import { EventOutgoingLogicCallCanceled } from "./module/types/gravity/v1/msgs"
import { EventSignatureSlashing } from "./module/types/gravity/v1/msgs"
import { EventOutgoingTxId } from "./module/types/gravity/v1/msgs"
import { IDSet } from "./module/types/gravity/v1/pool"
import { BatchFees } from "./module/types/gravity/v1/pool"
import { EventWithdrawalReceived } from "./module/types/gravity/v1/pool"
import { EventWithdrawCanceled } from "./module/types/gravity/v1/pool"
import { BridgeValidator } from "./module/types/gravity/v1/types"
import { Valset } from "./module/types/gravity/v1/types"
import { LastObservedEthereumBlockHeight } from "./module/types/gravity/v1/types"
import { ERC20ToDenom } from "./module/types/gravity/v1/types"
import { UnhaltBridgeProposal } from "./module/types/gravity/v1/types"
import { AirdropProposal } from "./module/types/gravity/v1/types"
import { IBCMetadataProposal } from "./module/types/gravity/v1/types"
import { PendingIbcAutoForward } from "./module/types/gravity/v1/types"


export { Attestation, ERC20Token, EventObservation, EventInvalidSendToCosmosReceiver, EventSendToCosmos, EventSendToCosmosLocal, EventSendToCosmosPendingIbcAutoForward, EventSendToCosmosExecutedIbcAutoForward, OutgoingTxBatch, OutgoingTransferTx, OutgoingLogicCall, EventOutgoingBatchCanceled, EventOutgoingBatch, Params, GravityNonces, EventSetOperatorAddress, EventValsetConfirmKey, EventBatchCreated, EventBatchConfirmKey, EventBatchSendToEthClaim, EventClaim, EventBadSignatureEvidence, EventERC20DeployedClaim, EventValsetUpdatedClaim, EventMultisigUpdateRequest, EventOutgoingLogicCallCanceled, EventSignatureSlashing, EventOutgoingTxId, IDSet, BatchFees, EventWithdrawalReceived, EventWithdrawCanceled, BridgeValidator, Valset, LastObservedEthereumBlockHeight, ERC20ToDenom, UnhaltBridgeProposal, AirdropProposal, IBCMetadataProposal, PendingIbcAutoForward };

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
				ValsetConfirm: {},
				SendToEth: {},
				RequestBatch: {},
				ConfirmBatch: {},
				ConfirmLogicCall: {},
				SendToCosmosClaim: {},
				ExecuteIbcAutoForwards: {},
				BatchSendToEthClaim: {},
				ValsetUpdateClaim: {},
				ERC20DeployedClaim: {},
				LogicCallExecutedClaim: {},
				SetOrchestratorAddress: {},
				CancelSendToEth: {},
				SubmitBadSignatureEvidence: {},
				Params: {},
				CurrentValset: {},
				ValsetRequest: {},
				ValsetConfirm: {},
				ValsetConfirmsByNonce: {},
				LastValsetRequests: {},
				LastPendingValsetRequestByAddr: {},
				LastPendingBatchRequestByAddr: {},
				LastPendingLogicCallByAddr: {},
				LastEventNonceByAddr: {},
				BatchFees: {},
				OutgoingTxBatches: {},
				OutgoingLogicCalls: {},
				BatchRequestByNonce: {},
				BatchConfirms: {},
				LogicConfirms: {},
				ERC20ToDenom: {},
				DenomToERC20: {},
				GetAttestations: {},
				GetDelegateKeyByValidator: {},
				GetDelegateKeyByEth: {},
				GetDelegateKeyByOrchestrator: {},
				GetPendingSendToEth: {},
				GetPendingIbcAutoForwards: {},
				
				_Structure: {
						Attestation: getStructure(Attestation.fromPartial({})),
						ERC20Token: getStructure(ERC20Token.fromPartial({})),
						EventObservation: getStructure(EventObservation.fromPartial({})),
						EventInvalidSendToCosmosReceiver: getStructure(EventInvalidSendToCosmosReceiver.fromPartial({})),
						EventSendToCosmos: getStructure(EventSendToCosmos.fromPartial({})),
						EventSendToCosmosLocal: getStructure(EventSendToCosmosLocal.fromPartial({})),
						EventSendToCosmosPendingIbcAutoForward: getStructure(EventSendToCosmosPendingIbcAutoForward.fromPartial({})),
						EventSendToCosmosExecutedIbcAutoForward: getStructure(EventSendToCosmosExecutedIbcAutoForward.fromPartial({})),
						OutgoingTxBatch: getStructure(OutgoingTxBatch.fromPartial({})),
						OutgoingTransferTx: getStructure(OutgoingTransferTx.fromPartial({})),
						OutgoingLogicCall: getStructure(OutgoingLogicCall.fromPartial({})),
						EventOutgoingBatchCanceled: getStructure(EventOutgoingBatchCanceled.fromPartial({})),
						EventOutgoingBatch: getStructure(EventOutgoingBatch.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						GravityNonces: getStructure(GravityNonces.fromPartial({})),
						EventSetOperatorAddress: getStructure(EventSetOperatorAddress.fromPartial({})),
						EventValsetConfirmKey: getStructure(EventValsetConfirmKey.fromPartial({})),
						EventBatchCreated: getStructure(EventBatchCreated.fromPartial({})),
						EventBatchConfirmKey: getStructure(EventBatchConfirmKey.fromPartial({})),
						EventBatchSendToEthClaim: getStructure(EventBatchSendToEthClaim.fromPartial({})),
						EventClaim: getStructure(EventClaim.fromPartial({})),
						EventBadSignatureEvidence: getStructure(EventBadSignatureEvidence.fromPartial({})),
						EventERC20DeployedClaim: getStructure(EventERC20DeployedClaim.fromPartial({})),
						EventValsetUpdatedClaim: getStructure(EventValsetUpdatedClaim.fromPartial({})),
						EventMultisigUpdateRequest: getStructure(EventMultisigUpdateRequest.fromPartial({})),
						EventOutgoingLogicCallCanceled: getStructure(EventOutgoingLogicCallCanceled.fromPartial({})),
						EventSignatureSlashing: getStructure(EventSignatureSlashing.fromPartial({})),
						EventOutgoingTxId: getStructure(EventOutgoingTxId.fromPartial({})),
						IDSet: getStructure(IDSet.fromPartial({})),
						BatchFees: getStructure(BatchFees.fromPartial({})),
						EventWithdrawalReceived: getStructure(EventWithdrawalReceived.fromPartial({})),
						EventWithdrawCanceled: getStructure(EventWithdrawCanceled.fromPartial({})),
						BridgeValidator: getStructure(BridgeValidator.fromPartial({})),
						Valset: getStructure(Valset.fromPartial({})),
						LastObservedEthereumBlockHeight: getStructure(LastObservedEthereumBlockHeight.fromPartial({})),
						ERC20ToDenom: getStructure(ERC20ToDenom.fromPartial({})),
						UnhaltBridgeProposal: getStructure(UnhaltBridgeProposal.fromPartial({})),
						AirdropProposal: getStructure(AirdropProposal.fromPartial({})),
						IBCMetadataProposal: getStructure(IBCMetadataProposal.fromPartial({})),
						PendingIbcAutoForward: getStructure(PendingIbcAutoForward.fromPartial({})),
						
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
				getValsetConfirm: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ValsetConfirm[JSON.stringify(params)] ?? {}
		},
				getSendToEth: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.SendToEth[JSON.stringify(params)] ?? {}
		},
				getRequestBatch: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.RequestBatch[JSON.stringify(params)] ?? {}
		},
				getConfirmBatch: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ConfirmBatch[JSON.stringify(params)] ?? {}
		},
				getConfirmLogicCall: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ConfirmLogicCall[JSON.stringify(params)] ?? {}
		},
				getSendToCosmosClaim: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.SendToCosmosClaim[JSON.stringify(params)] ?? {}
		},
				getExecuteIbcAutoForwards: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ExecuteIbcAutoForwards[JSON.stringify(params)] ?? {}
		},
				getBatchSendToEthClaim: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.BatchSendToEthClaim[JSON.stringify(params)] ?? {}
		},
				getValsetUpdateClaim: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ValsetUpdateClaim[JSON.stringify(params)] ?? {}
		},
				getERC20DeployedClaim: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ERC20DeployedClaim[JSON.stringify(params)] ?? {}
		},
				getLogicCallExecutedClaim: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.LogicCallExecutedClaim[JSON.stringify(params)] ?? {}
		},
				getSetOrchestratorAddress: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.SetOrchestratorAddress[JSON.stringify(params)] ?? {}
		},
				getCancelSendToEth: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.CancelSendToEth[JSON.stringify(params)] ?? {}
		},
				getSubmitBadSignatureEvidence: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.SubmitBadSignatureEvidence[JSON.stringify(params)] ?? {}
		},
				getParams: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Params[JSON.stringify(params)] ?? {}
		},
				getCurrentValset: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.CurrentValset[JSON.stringify(params)] ?? {}
		},
				getValsetRequest: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ValsetRequest[JSON.stringify(params)] ?? {}
		},
				getValsetConfirm: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ValsetConfirm[JSON.stringify(params)] ?? {}
		},
				getValsetConfirmsByNonce: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ValsetConfirmsByNonce[JSON.stringify(params)] ?? {}
		},
				getLastValsetRequests: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.LastValsetRequests[JSON.stringify(params)] ?? {}
		},
				getLastPendingValsetRequestByAddr: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.LastPendingValsetRequestByAddr[JSON.stringify(params)] ?? {}
		},
				getLastPendingBatchRequestByAddr: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.LastPendingBatchRequestByAddr[JSON.stringify(params)] ?? {}
		},
				getLastPendingLogicCallByAddr: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.LastPendingLogicCallByAddr[JSON.stringify(params)] ?? {}
		},
				getLastEventNonceByAddr: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.LastEventNonceByAddr[JSON.stringify(params)] ?? {}
		},
				getBatchFees: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.BatchFees[JSON.stringify(params)] ?? {}
		},
				getOutgoingTxBatches: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.OutgoingTxBatches[JSON.stringify(params)] ?? {}
		},
				getOutgoingLogicCalls: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.OutgoingLogicCalls[JSON.stringify(params)] ?? {}
		},
				getBatchRequestByNonce: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.BatchRequestByNonce[JSON.stringify(params)] ?? {}
		},
				getBatchConfirms: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.BatchConfirms[JSON.stringify(params)] ?? {}
		},
				getLogicConfirms: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.LogicConfirms[JSON.stringify(params)] ?? {}
		},
				getERC20ToDenom: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ERC20ToDenom[JSON.stringify(params)] ?? {}
		},
				getDenomToERC20: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.DenomToERC20[JSON.stringify(params)] ?? {}
		},
				getGetAttestations: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetAttestations[JSON.stringify(params)] ?? {}
		},
				getGetDelegateKeyByValidator: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetDelegateKeyByValidator[JSON.stringify(params)] ?? {}
		},
				getGetDelegateKeyByEth: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetDelegateKeyByEth[JSON.stringify(params)] ?? {}
		},
				getGetDelegateKeyByOrchestrator: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetDelegateKeyByOrchestrator[JSON.stringify(params)] ?? {}
		},
				getGetPendingSendToEth: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetPendingSendToEth[JSON.stringify(params)] ?? {}
		},
				getGetPendingIbcAutoForwards: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetPendingIbcAutoForwards[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: gravity.v1 initialized!')
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
		
		
		
		 		
		
		
		async MsgValsetConfirm({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.msgValsetConfirm(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.msgValsetConfirm({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ValsetConfirm', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'MsgValsetConfirm', payload: { options: { all }, params: {...key},query }})
				return getters['getValsetConfirm']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:MsgValsetConfirm API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async MsgSendToEth({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.msgSendToEth(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.msgSendToEth({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'SendToEth', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'MsgSendToEth', payload: { options: { all }, params: {...key},query }})
				return getters['getSendToEth']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:MsgSendToEth API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async MsgRequestBatch({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.msgRequestBatch(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.msgRequestBatch({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'RequestBatch', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'MsgRequestBatch', payload: { options: { all }, params: {...key},query }})
				return getters['getRequestBatch']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:MsgRequestBatch API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async MsgConfirmBatch({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.msgConfirmBatch(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.msgConfirmBatch({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ConfirmBatch', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'MsgConfirmBatch', payload: { options: { all }, params: {...key},query }})
				return getters['getConfirmBatch']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:MsgConfirmBatch API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async MsgConfirmLogicCall({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.msgConfirmLogicCall(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.msgConfirmLogicCall({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ConfirmLogicCall', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'MsgConfirmLogicCall', payload: { options: { all }, params: {...key},query }})
				return getters['getConfirmLogicCall']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:MsgConfirmLogicCall API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async MsgSendToCosmosClaim({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.msgSendToCosmosClaim(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.msgSendToCosmosClaim({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'SendToCosmosClaim', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'MsgSendToCosmosClaim', payload: { options: { all }, params: {...key},query }})
				return getters['getSendToCosmosClaim']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:MsgSendToCosmosClaim API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async MsgExecuteIbcAutoForwards({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.msgExecuteIbcAutoForwards(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.msgExecuteIbcAutoForwards({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ExecuteIbcAutoForwards', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'MsgExecuteIbcAutoForwards', payload: { options: { all }, params: {...key},query }})
				return getters['getExecuteIbcAutoForwards']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:MsgExecuteIbcAutoForwards API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async MsgBatchSendToEthClaim({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.msgBatchSendToEthClaim(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.msgBatchSendToEthClaim({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'BatchSendToEthClaim', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'MsgBatchSendToEthClaim', payload: { options: { all }, params: {...key},query }})
				return getters['getBatchSendToEthClaim']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:MsgBatchSendToEthClaim API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async MsgValsetUpdateClaim({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.msgValsetUpdateClaim(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.msgValsetUpdateClaim({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ValsetUpdateClaim', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'MsgValsetUpdateClaim', payload: { options: { all }, params: {...key},query }})
				return getters['getValsetUpdateClaim']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:MsgValsetUpdateClaim API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async MsgERC20DeployedClaim({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.msgERC20DeployedClaim(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.msgERC20DeployedClaim({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ERC20DeployedClaim', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'MsgERC20DeployedClaim', payload: { options: { all }, params: {...key},query }})
				return getters['getERC20DeployedClaim']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:MsgERC20DeployedClaim API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async MsgLogicCallExecutedClaim({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.msgLogicCallExecutedClaim(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.msgLogicCallExecutedClaim({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'LogicCallExecutedClaim', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'MsgLogicCallExecutedClaim', payload: { options: { all }, params: {...key},query }})
				return getters['getLogicCallExecutedClaim']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:MsgLogicCallExecutedClaim API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async MsgSetOrchestratorAddress({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.msgSetOrchestratorAddress(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.msgSetOrchestratorAddress({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'SetOrchestratorAddress', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'MsgSetOrchestratorAddress', payload: { options: { all }, params: {...key},query }})
				return getters['getSetOrchestratorAddress']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:MsgSetOrchestratorAddress API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async MsgCancelSendToEth({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.msgCancelSendToEth(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.msgCancelSendToEth({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'CancelSendToEth', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'MsgCancelSendToEth', payload: { options: { all }, params: {...key},query }})
				return getters['getCancelSendToEth']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:MsgCancelSendToEth API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async MsgSubmitBadSignatureEvidence({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.msgSubmitBadSignatureEvidence(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.msgSubmitBadSignatureEvidence({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'SubmitBadSignatureEvidence', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'MsgSubmitBadSignatureEvidence', payload: { options: { all }, params: {...key},query }})
				return getters['getSubmitBadSignatureEvidence']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:MsgSubmitBadSignatureEvidence API Node Unavailable. Could not perform query: ' + e.message)
				
			}
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
		
		
		
		
		 		
		
		
		async QueryCurrentValset({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryCurrentValset()).data
				
					
				commit('QUERY', { query: 'CurrentValset', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryCurrentValset', payload: { options: { all }, params: {...key},query }})
				return getters['getCurrentValset']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryCurrentValset API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryValsetRequest({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryValsetRequest(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryValsetRequest({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ValsetRequest', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryValsetRequest', payload: { options: { all }, params: {...key},query }})
				return getters['getValsetRequest']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryValsetRequest API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryValsetConfirm({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryValsetConfirm(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryValsetConfirm({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ValsetConfirm', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryValsetConfirm', payload: { options: { all }, params: {...key},query }})
				return getters['getValsetConfirm']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryValsetConfirm API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryValsetConfirmsByNonce({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryValsetConfirmsByNonce( key.nonce)).data
				
					
				commit('QUERY', { query: 'ValsetConfirmsByNonce', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryValsetConfirmsByNonce', payload: { options: { all }, params: {...key},query }})
				return getters['getValsetConfirmsByNonce']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryValsetConfirmsByNonce API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryLastValsetRequests({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryLastValsetRequests()).data
				
					
				commit('QUERY', { query: 'LastValsetRequests', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryLastValsetRequests', payload: { options: { all }, params: {...key},query }})
				return getters['getLastValsetRequests']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryLastValsetRequests API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryLastPendingValsetRequestByAddr({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryLastPendingValsetRequestByAddr(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryLastPendingValsetRequestByAddr({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'LastPendingValsetRequestByAddr', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryLastPendingValsetRequestByAddr', payload: { options: { all }, params: {...key},query }})
				return getters['getLastPendingValsetRequestByAddr']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryLastPendingValsetRequestByAddr API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryLastPendingBatchRequestByAddr({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryLastPendingBatchRequestByAddr(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryLastPendingBatchRequestByAddr({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'LastPendingBatchRequestByAddr', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryLastPendingBatchRequestByAddr', payload: { options: { all }, params: {...key},query }})
				return getters['getLastPendingBatchRequestByAddr']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryLastPendingBatchRequestByAddr API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryLastPendingLogicCallByAddr({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryLastPendingLogicCallByAddr( key.address)).data
				
					
				commit('QUERY', { query: 'LastPendingLogicCallByAddr', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryLastPendingLogicCallByAddr', payload: { options: { all }, params: {...key},query }})
				return getters['getLastPendingLogicCallByAddr']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryLastPendingLogicCallByAddr API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryLastEventNonceByAddr({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryLastEventNonceByAddr( key.address)).data
				
					
				commit('QUERY', { query: 'LastEventNonceByAddr', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryLastEventNonceByAddr', payload: { options: { all }, params: {...key},query }})
				return getters['getLastEventNonceByAddr']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryLastEventNonceByAddr API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryBatchFees({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryBatchFees()).data
				
					
				commit('QUERY', { query: 'BatchFees', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryBatchFees', payload: { options: { all }, params: {...key},query }})
				return getters['getBatchFees']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryBatchFees API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryOutgoingTxBatches({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryOutgoingTxBatches()).data
				
					
				commit('QUERY', { query: 'OutgoingTxBatches', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryOutgoingTxBatches', payload: { options: { all }, params: {...key},query }})
				return getters['getOutgoingTxBatches']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryOutgoingTxBatches API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryOutgoingLogicCalls({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryOutgoingLogicCalls()).data
				
					
				commit('QUERY', { query: 'OutgoingLogicCalls', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryOutgoingLogicCalls', payload: { options: { all }, params: {...key},query }})
				return getters['getOutgoingLogicCalls']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryOutgoingLogicCalls API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryBatchRequestByNonce({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryBatchRequestByNonce(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryBatchRequestByNonce({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'BatchRequestByNonce', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryBatchRequestByNonce', payload: { options: { all }, params: {...key},query }})
				return getters['getBatchRequestByNonce']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryBatchRequestByNonce API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryBatchConfirms({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryBatchConfirms(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryBatchConfirms({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'BatchConfirms', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryBatchConfirms', payload: { options: { all }, params: {...key},query }})
				return getters['getBatchConfirms']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryBatchConfirms API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryLogicConfirms({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryLogicConfirms(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryLogicConfirms({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'LogicConfirms', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryLogicConfirms', payload: { options: { all }, params: {...key},query }})
				return getters['getLogicConfirms']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryLogicConfirms API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryERC20ToDenom({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryERC20ToDenom(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryERC20ToDenom({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ERC20ToDenom', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryERC20ToDenom', payload: { options: { all }, params: {...key},query }})
				return getters['getERC20ToDenom']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryERC20ToDenom API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryDenomToERC20({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryDenomToERC20(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryDenomToERC20({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'DenomToERC20', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryDenomToERC20', payload: { options: { all }, params: {...key},query }})
				return getters['getDenomToERC20']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryDenomToERC20 API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetAttestations({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryGetAttestations(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryGetAttestations({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'GetAttestations', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetAttestations', payload: { options: { all }, params: {...key},query }})
				return getters['getGetAttestations']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryGetAttestations API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetDelegateKeyByValidator({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryGetDelegateKeyByValidator(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryGetDelegateKeyByValidator({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'GetDelegateKeyByValidator', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetDelegateKeyByValidator', payload: { options: { all }, params: {...key},query }})
				return getters['getGetDelegateKeyByValidator']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryGetDelegateKeyByValidator API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetDelegateKeyByEth({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryGetDelegateKeyByEth(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryGetDelegateKeyByEth({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'GetDelegateKeyByEth', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetDelegateKeyByEth', payload: { options: { all }, params: {...key},query }})
				return getters['getGetDelegateKeyByEth']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryGetDelegateKeyByEth API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetDelegateKeyByOrchestrator({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryGetDelegateKeyByOrchestrator(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryGetDelegateKeyByOrchestrator({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'GetDelegateKeyByOrchestrator', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetDelegateKeyByOrchestrator', payload: { options: { all }, params: {...key},query }})
				return getters['getGetDelegateKeyByOrchestrator']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryGetDelegateKeyByOrchestrator API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetPendingSendToEth({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryGetPendingSendToEth(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryGetPendingSendToEth({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'GetPendingSendToEth', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetPendingSendToEth', payload: { options: { all }, params: {...key},query }})
				return getters['getGetPendingSendToEth']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryGetPendingSendToEth API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetPendingIbcAutoForwards({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryGetPendingIbcAutoForwards(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryGetPendingIbcAutoForwards({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'GetPendingIbcAutoForwards', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetPendingIbcAutoForwards', payload: { options: { all }, params: {...key},query }})
				return getters['getGetPendingIbcAutoForwards']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryGetPendingIbcAutoForwards API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgRequestBatch({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgRequestBatch(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRequestBatch:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRequestBatch:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgConfirmLogicCall({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgConfirmLogicCall(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgConfirmLogicCall:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgConfirmLogicCall:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgValsetUpdatedClaim({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgValsetUpdatedClaim(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgValsetUpdatedClaim:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgValsetUpdatedClaim:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSetOrchestratorAddress({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetOrchestratorAddress(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetOrchestratorAddress:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSetOrchestratorAddress:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSendToEth({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSendToEth(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSendToEth:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSendToEth:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgValsetConfirm({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgValsetConfirm(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgValsetConfirm:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgValsetConfirm:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSendToCosmosClaim({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSendToCosmosClaim(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSendToCosmosClaim:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSendToCosmosClaim:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgERC20DeployedClaim({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgERC20DeployedClaim(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgERC20DeployedClaim:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgERC20DeployedClaim:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSubmitBadSignatureEvidence({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSubmitBadSignatureEvidence(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSubmitBadSignatureEvidence:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSubmitBadSignatureEvidence:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgBatchSendToEthClaim({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgBatchSendToEthClaim(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgBatchSendToEthClaim:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgBatchSendToEthClaim:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgConfirmBatch({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgConfirmBatch(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgConfirmBatch:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgConfirmBatch:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCancelSendToEth({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCancelSendToEth(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCancelSendToEth:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCancelSendToEth:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgLogicCallExecutedClaim({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgLogicCallExecutedClaim(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgLogicCallExecutedClaim:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgLogicCallExecutedClaim:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgRequestBatch({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgRequestBatch(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRequestBatch:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgRequestBatch:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgConfirmLogicCall({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgConfirmLogicCall(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgConfirmLogicCall:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgConfirmLogicCall:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgValsetUpdatedClaim({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgValsetUpdatedClaim(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgValsetUpdatedClaim:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgValsetUpdatedClaim:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSetOrchestratorAddress({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSetOrchestratorAddress(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetOrchestratorAddress:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSetOrchestratorAddress:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSendToEth({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSendToEth(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSendToEth:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSendToEth:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgValsetConfirm({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgValsetConfirm(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgValsetConfirm:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgValsetConfirm:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSendToCosmosClaim({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSendToCosmosClaim(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSendToCosmosClaim:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSendToCosmosClaim:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgERC20DeployedClaim({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgERC20DeployedClaim(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgERC20DeployedClaim:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgERC20DeployedClaim:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSubmitBadSignatureEvidence({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSubmitBadSignatureEvidence(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSubmitBadSignatureEvidence:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSubmitBadSignatureEvidence:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgBatchSendToEthClaim({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgBatchSendToEthClaim(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgBatchSendToEthClaim:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgBatchSendToEthClaim:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgConfirmBatch({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgConfirmBatch(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgConfirmBatch:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgConfirmBatch:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCancelSendToEth({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCancelSendToEth(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCancelSendToEth:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCancelSendToEth:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgLogicCallExecutedClaim({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgLogicCallExecutedClaim(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgLogicCallExecutedClaim:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgLogicCallExecutedClaim:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
