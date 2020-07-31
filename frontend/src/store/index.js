import Vuex from 'vuex'
import Vue from 'vue'
import user from './user'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    user
  }
})
