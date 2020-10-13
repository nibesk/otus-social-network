import Vuex from 'vuex'
import Vue from 'vue'

import user from './user'
import friends from './friends'
import system from './system'
import ws from './ws'

Vue.use(Vuex)

export default new Vuex.Store({
    modules: {
        user,
        friends,
        system,
        ws
    }
})
