// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import ThesixnetworkSixprotocolProtocoladmin from './thesixnetwork.sixprotocol.protocoladmin'
import ThesixnetworkSixprotocolTokenmngr from './thesixnetwork.sixprotocol.tokenmngr'


export default { 
  ThesixnetworkSixprotocolProtocoladmin: load(ThesixnetworkSixprotocolProtocoladmin, 'thesixnetwork.sixprotocol.protocoladmin'),
  ThesixnetworkSixprotocolTokenmngr: load(ThesixnetworkSixprotocolTokenmngr, 'thesixnetwork.sixprotocol.tokenmngr'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}