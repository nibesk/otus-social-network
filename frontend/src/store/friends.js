import {httpRequest, RequestMessage} from '../api/http'
import {routes} from "../router/routes";

export default {
    namespaced: true,

    state: () => ({
        friends: [],
        availableFriends: []
    }),

    actions: {
        async apiGetFriends({commit}) {
            const {responseMessage} = await httpRequest.get(routes.api.friends);

            if (responseMessage.status) {
                const {users} = responseMessage.data;
                commit('SET_FRIENDS', users);
            }

            return responseMessage
        },

        async apiAvailableGetFriends({commit}) {
            const {responseMessage} = await httpRequest.get(routes.api.availableFriends);

            if (responseMessage.status) {
                const {users} = responseMessage.data;
                commit('SET_AVAILABLE_FRIENDS', users);
            }

            return responseMessage
        },

        async apiDeleteFriend({commit}, payload) {
            const {responseMessage} = await httpRequest.post(routes.api.removeFriends, new RequestMessage(payload));

            if (responseMessage.status) {
                const {user} = responseMessage.data;
                commit('REMOVE_USER_FROM_FRIENDS', user);
                commit('CHANGE_AVAILABLE_FRIEND_STATUS', {user: user, status: false});
            }

            return responseMessage
        },

        async apiAddFriend({commit}, payload) {
            const {responseMessage} = await httpRequest.post(routes.api.friends, new RequestMessage(payload));

            if (responseMessage.status) {
                const {user} = responseMessage.data;
                commit('ADD_FRIEND', user);
                commit('CHANGE_AVAILABLE_FRIEND_STATUS', {user: user});
            }

            return responseMessage
        },
    },

    mutations: {
        SET_FRIENDS(state, friends) {
            state.friends = friends
        },
        ADD_FRIEND(state, friend) {
            state.friends.push(friend);
        },
        REMOVE_USER_FROM_FRIENDS(state, user) {
            state.friends = state.friends.filter((friend) => user.user_id !== friend.user_id);
        },
        SET_AVAILABLE_FRIENDS(state, availableFriends) {
            state.availableFriends = availableFriends
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
    }
}
