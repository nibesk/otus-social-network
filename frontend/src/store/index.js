import Vuex from 'vuex'
import Vue from 'vue'

import user from './user'
import friends from './friends'
import system from './system'

Vue.use(Vuex)

export default new Vuex.Store({
    modules: {
        user,
        friends,
        system
    }
})
