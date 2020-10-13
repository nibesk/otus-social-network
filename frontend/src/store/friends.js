import {httpRequest, RequestMessage} from '../api/http'
import {routes} from "../config/routes";
import _ from 'lodash'

export default {
    namespaced: true,

    state: () => ({
        friends: [],
        availableFriends: [],
        lastViewedUserId: null,
        lastLoadedAvailableFriendsBatch: [],
    }),

    actions: {
        async apiGetFriends({commit}) {
            const {responseMessage} = await httpRequest.get(routes.service_users.friends);

            if (responseMessage.status) {
                const {users} = responseMessage.data;
                commit('SET_FRIENDS', users);
            }

            return responseMessage
        },

        async apiAvailableGetFriends({commit, state}, payload) {
            let queryParams = {};
            if (null !== state.lastViewedUserId) {
                queryParams.lastViewedUserId = state.lastViewedUserId
            }

            if (`undefined` !== typeof payload) {
                queryParams = {...queryParams, ...payload}
            }

            const {responseMessage} = await httpRequest.get(`${routes.service_users.availableFriends}?${new URLSearchParams(queryParams)}`);

            if (responseMessage.status) {
                const {users} = responseMessage.data;
                commit('PUSH_AVAILABLE_FRIENDS', users);
                commit('SET_LAST_LOADED_AVAILABLE_FRIENDS_BATCH', users);
            }

            return responseMessage
        },

        async getUserById({commit}, userId) {
            const {responseMessage} = await httpRequest.get(routes.service_users.getUserById(userId));

            if (responseMessage.status) {
                const {user} = responseMessage.data;

                return user;
            }

            return null;
        },

        async apiDeleteFriend({commit}, payload) {
            const {responseMessage} = await httpRequest.post(routes.service_users.removeFriends, new RequestMessage(payload));

            if (responseMessage.status) {
                const {user} = responseMessage.data;
                commit('REMOVE_USER_FROM_FRIENDS', user);
                commit('CHANGE_AVAILABLE_FRIEND_STATUS', {user: user, status: false});
            }

            return responseMessage
        },

        async apiAddFriend({commit}, payload) {
            const {responseMessage} = await httpRequest.post(routes.service_users.friends, new RequestMessage(payload));

            if (responseMessage.status) {
                const {user} = responseMessage.data;
                commit('ADD_FRIEND', user);
                commit('CHANGE_AVAILABLE_FRIEND_STATUS', {user: user});
            }

            return responseMessage
        },
    },

    mutations: {
        SET_LAST_LOADED_AVAILABLE_FRIENDS_BATCH(state, lastLoadedAvailableFriendsBatch) {
            state.lastLoadedAvailableFriendsBatch = lastLoadedAvailableFriendsBatch;
            if (lastLoadedAvailableFriendsBatch.length > 0) {
                const lastUserInBatch = _.last(lastLoadedAvailableFriendsBatch);
                console.log(lastUserInBatch);
                state.lastViewedUserId = lastUserInBatch.user_id;
            }
        },
        RESET_AVAILABLE_FRIENDS(state) {
            state.lastLoadedAvailableFriendsBatch = [];
            state.lastViewedUserId = null;
            state.availableFriends = [];
        },
        SET_FRIENDS(state, friends) {
            state.friends = friends
        },
        ADD_FRIEND(state, friend) {
            state.friends.push(friend);
        },
        REMOVE_USER_FROM_FRIENDS(state, user) {
            state.friends = state.friends.filter((friend) => user.user_id !== friend.user_id);
        },
        PUSH_AVAILABLE_FRIENDS(state, availableFriends) {
            state.availableFriends.push(...availableFriends)
        },
        CHANGE_AVAILABLE_FRIEND_STATUS(state, payload) {
            state.availableFriends.forEach((friend) => {
                if (payload.user.user_id === friend.user_id) {
                    friend.is_friend = typeof payload.status === 'undefined' ? true : payload.status;
                }
            })
        }
    },

    getters: {
        getFriends: state => state.friends,
        getAvailableFriends: state => state.availableFriends,
        getLastViewedUserId: state => state.lastViewedUserId,
        getLastLoadedAvailableFriendsBatch: state => state.lastLoadedAvailableFriendsBatch,
    }
}
