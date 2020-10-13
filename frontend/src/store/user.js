import {httpRequest, RequestMessage} from '../api/http'
import {routes} from "../config/routes";
import {localStorageGet, localStorageSet, localStorageDelete} from "utils"
import {default as globals} from "utils"

const AUTH_TOKEN_KEY = 'token';

export default {
    namespaced: true,

    state: () => ({
        user: null,
        token: null,
        authCheckedTtl: 0
    }),

    actions: {
        async checkAuth({commit, getters}) {
            if (getters.isAuthChecked) {
                return null;
            }

            const token = localStorageGet(AUTH_TOKEN_KEY);
            if (null !== token) {
                commit('SET_TOKEN', token);
            }

            const {responseMessage} = await httpRequest.get(routes.service_users.getUser);

            if (responseMessage.status) {
                const {user} = responseMessage.data;
                commit('SET_USER', user);
            }

            commit('SET_CHECK_AUTH');

            return responseMessage
        },

        async login({commit, state}, payload) {
            const {responseMessage} = await httpRequest.post(routes.service_users.login, new RequestMessage(payload));

            if (responseMessage.status) {
                const {user} = responseMessage.data;
                commit('SET_USER', user);
                commit('SET_TOKEN', user.token.String);
                localStorageSet(AUTH_TOKEN_KEY, state.token)
            }

            return responseMessage
        },

        async register({commit, state}, payload) {
            const {responseMessage} = await httpRequest.post(routes.service_users.register, new RequestMessage(payload));

            if (responseMessage.status) {
                const {user} = responseMessage.data;
                commit('SET_USER', user);
                commit('SET_TOKEN', user.token.String);
                localStorageSet(AUTH_TOKEN_KEY, state.token)
            }

            return responseMessage
        },

        async logout({commit}) {
            const {responseMessage} = await httpRequest.post(routes.service_users.logout);

            if (responseMessage.status) {
                commit('CLEAR_USER');
            }

            return responseMessage
        },
    },

    mutations: {
        SET_USER(state, user) {
            state.user = user;
            state.isAuthChecked = true;
        },
        SET_CHECK_AUTH(state) {
            state.authCheckedTtl = Date.now();
        },
        CLEAR_USER(state) {
            state.user = null;
            state.token = null;
            localStorageDelete('token')
        },
        SET_TOKEN(state, token) {
            state.token = token;
        }
    },

    getters: {
        getUser: state => state.user,
        getToken: state => state.token,
        isAuthChecked: state => {
            return Date.now() < state.authCheckedTtl + globals.MINUTE * globals.MILLISECONDS
        }
    }
}
