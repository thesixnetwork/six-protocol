import { Client, registry, MissingWalletError } from 'thesixnetwork-six-protocol-client-ts'

import { Admin } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixprotocol.protocoladmin/types"
import { Group } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixprotocol.protocoladmin/types"
import { ProtocoladminPacketData } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixprotocol.protocoladmin/types"
import { NoData } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixprotocol.protocoladmin/types"
import { Params } from "thesixnetwork-six-protocol-client-ts/thesixnetwork.sixprotocol.protocoladmin/types"


export { Admin, Group, ProtocoladminPacketData, NoData, Params };

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
				Group: {},
				GroupAll: {},
				Admin: {},
				AdminAll: {},
				ListAdminOfGroup: {},
				
				_Structure: {
						Admin: getStructure(Admin.fromPartial({})),
						Group: getStructure(Group.fromPartial({})),
						ProtocoladminPacketData: getStructure(ProtocoladminPacketData.fromPartial({})),
						NoData: getStructure(NoData.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						
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
				getGroup: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Group[JSON.stringify(params)] ?? {}
		},
				getGroupAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GroupAll[JSON.stringify(params)] ?? {}
		},
				getAdmin: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Admin[JSON.stringify(params)] ?? {}
		},
				getAdminAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.AdminAll[JSON.stringify(params)] ?? {}
		},
				getListAdminOfGroup: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ListAdminOfGroup[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: thesixnetwork.sixprotocol.protocoladmin initialized!')
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
				let value= (await client.ThesixnetworkSixprotocolProtocoladmin.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGroup({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ThesixnetworkSixprotocolProtocoladmin.query.queryGroup( key.name)).data
				
					
				commit('QUERY', { query: 'Group', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGroup', payload: { options: { all }, params: {...key},query }})
				return getters['getGroup']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryGroup API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGroupAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ThesixnetworkSixprotocolProtocoladmin.query.queryGroupAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ThesixnetworkSixprotocolProtocoladmin.query.queryGroupAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'GroupAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGroupAll', payload: { options: { all }, params: {...key},query }})
				return getters['getGroupAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryGroupAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryAdmin({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ThesixnetworkSixprotocolProtocoladmin.query.queryAdmin( key.group,  key.admin)).data
				
					
				commit('QUERY', { query: 'Admin', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryAdmin', payload: { options: { all }, params: {...key},query }})
				return getters['getAdmin']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryAdmin API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryAdminAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ThesixnetworkSixprotocolProtocoladmin.query.queryAdminAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ThesixnetworkSixprotocolProtocoladmin.query.queryAdminAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'AdminAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryAdminAll', payload: { options: { all }, params: {...key},query }})
				return getters['getAdminAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryAdminAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryListAdminOfGroup({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ThesixnetworkSixprotocolProtocoladmin.query.queryListAdminOfGroup( key.group, query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ThesixnetworkSixprotocolProtocoladmin.query.queryListAdminOfGroup( key.group, {...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ListAdminOfGroup', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryListAdminOfGroup', payload: { options: { all }, params: {...key},query }})
				return getters['getListAdminOfGroup']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryListAdminOfGroup API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgDeleteGroup({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ThesixnetworkSixprotocolProtocoladmin.tx.sendMsgDeleteGroup({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteGroup:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeleteGroup:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateGroup({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ThesixnetworkSixprotocolProtocoladmin.tx.sendMsgCreateGroup({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateGroup:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateGroup:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgAddAdminToGroup({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ThesixnetworkSixprotocolProtocoladmin.tx.sendMsgAddAdminToGroup({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddAdminToGroup:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgAddAdminToGroup:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgRemoveAdminFromGroup({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ThesixnetworkSixprotocolProtocoladmin.tx.sendMsgRemoveAdminFromGroup({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRemoveAdminFromGroup:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRemoveAdminFromGroup:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateGroup({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ThesixnetworkSixprotocolProtocoladmin.tx.sendMsgUpdateGroup({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateGroup:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateGroup:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgDeleteGroup({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ThesixnetworkSixprotocolProtocoladmin.tx.msgDeleteGroup({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteGroup:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeleteGroup:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateGroup({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ThesixnetworkSixprotocolProtocoladmin.tx.msgCreateGroup({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateGroup:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateGroup:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgAddAdminToGroup({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ThesixnetworkSixprotocolProtocoladmin.tx.msgAddAdminToGroup({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddAdminToGroup:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgAddAdminToGroup:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgRemoveAdminFromGroup({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ThesixnetworkSixprotocolProtocoladmin.tx.msgRemoveAdminFromGroup({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRemoveAdminFromGroup:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgRemoveAdminFromGroup:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateGroup({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ThesixnetworkSixprotocolProtocoladmin.tx.msgUpdateGroup({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateGroup:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateGroup:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
