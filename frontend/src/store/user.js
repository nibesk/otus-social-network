import {httpRequest, RequestMessage} from '../api/http'
import {routes} from "../router/routes";

export default {
    namespaced: true,

    state: () => ({
        user: null
    }),

    actions: {
        async checkAuth({commit}) {
            const {responseMessage} = await httpRequest.get(routes.api.getUser);

            if (responseMessage.status) {
                const {user} = responseMessage.data;
                commit('SET_USER', user);
            }

            return responseMessage
        },

        async login({commit}, payload) {
            const {responseMessage} = await httpRequest.post(routes.api.login, new RequestMessage(payload));

            if (responseMessage.status) {
                const {user} = responseMessage.data;
                commit('SET_USER', user);
            }

            return responseMessage
        },

        async register({commit}, payload) {
            const {responseMessage} = await httpRequest.post(routes.api.register, new RequestMessage(payload));

            if (responseMessage.status) {
                const {user} = responseMessage.data;
                commit('SET_USER', user);
            }

            return responseMessage
        },

        async logout({commit}) {
            const {responseMessage} = await httpRequest.post(routes.api.logout);

            if (responseMessage.status) {
                commit('CLEAR_USER');
            }

            return responseMessage
        },
    },

    mutations: {
        SET_USER(state, user) {
            state.user = user
        },
        CLEAR_USER(state) {
            state.user = null
        },
    },

    getters: {
        getUser: state => state.user,
    }
}