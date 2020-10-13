import Vue from 'vue';
import wsEvents from "api/wsEvents";
import router from 'router';
import {httpRequest, RequestMessage} from "../api/http";
import {routes} from "../config/routes";

class Event {
    event = null;
    payload = null;

    constructor(event, payload) {
        this.event = event;
        this.payload = payload
    }
}

export default {
    namespaced: true,

    state: () => ({
        events: [],
        messages: {}
    }),

    actions: {
        onOpenHandler({rootGetters}) {
            console.log('ws opened');
            const token = rootGetters['user/getToken'];
            Vue.prototype.$ws.send(new Event(wsEvents.eventStartUp, {token: token}))
        },

        eventsHandler({state, commit}, event) {
            console.log(event);

            switch (event.event) {
                case wsEvents.eventError:
                    event.payload.forEach(error => Vue.$toast.error(error));
                    break;

                case wsEvents.eventMessage:
                    let {user} = event.payload;

                    if ('chat' !== router.currentRoute.name) {
                        Vue.$toast.success(`New message from ${user.name}`, {
                            onClick() {
                                console.log(`clicked msgs`);
                            }
                        });
                    }

                    commit('ADD_MESSAGE_TO_USER', {userId: user.user_id, data: event.payload});
                    break;

                default:
                    break;
            }
        },

        sendMessage({state, commit}, {message, fromUserId, toUserId}) {
            commit('ADD_MESSAGE_TO_USER', {
                userId: toUserId,
                data: {
                    message: message,
                    timestamp: Math.round(new Date().getTime() / 1000),
                    user_id: fromUserId,
                }
            });

            Vue.prototype.$ws.send(new Event(wsEvents.eventMessage, {
                message: message,
                user_id: parseInt(toUserId)
            }))
        },

        async loadMessagesWithUser({state, commit}, userId) {
            const {responseMessage} = await httpRequest.get(routes.service_chat.getMessages(userId));

            if (responseMessage.status) {
                const {messages} = responseMessage.data;
                messages.forEach((message) => {
                    commit('ADD_MESSAGE_TO_USER', {userId: userId, data: message});
                })
            }
        }
    },

    mutations: {
        ADD_MESSAGE_TO_USER(state, {userId, data}) {
            if ('undefined' === typeof state.messages[userId]) {
                Vue.set(state.messages, userId, []);
            }
            Vue.set(state.messages[userId], state.messages[userId].length, data);
        }
    },

    getters: {
        getMessages: state => state.messages,
    }
}
