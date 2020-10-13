<template>
    <div :class="chatUser ? `chat` : `chat chat__no-user`">

        <div class="card w-100 mb-3" >
            <div class="card-header">
                <span class="font-weight-bold" v-if="chatUser">
                    {{ chatUser.name }} {{ chatUser.surname }} <small>({{ chatUser.email }})</small>
                </span>
                <template v-else>
                    Wrong user
                </template>
            </div>
            <div id="chat__main" class="card-body chat__main">
                <p v-if="0 === messages.length" class="card-text">
                    You have no messages with this person yet...
                </p>

                <ul class="list-unstyled mb-0">
                    <li class="chat__main-item mb-3" v-for="item in messages">
                        <div v-if="item.user_id !== getUser.user_id" class="media">
                            <div class="">
                                <div class="bg-light rounded py-2 px-3 mb-2">
                                    <p class="text-small mb-0 text-muted">{{ item.message }}</p>
                                </div>
                                <p class="small text-muted mb-0">{{ new Date(item.timestamp * 1e3).toUTCString() }}</p>
                            </div>
                        </div>

                        <div v-else class="media chat__main-item_self-msg">
                            <div>
                                <div class="bg-primary rounded py-2 px-3 mb-2">
                                    <p class="text-small mb-0 text-white">{{ item.message }}</p>
                                </div>
                                <p class="small text-muted mb-0">{{ new Date(item.timestamp * 1e3).toUTCString() }}</p>
                            </div>
                        </div>
                    </li>
                </ul>
            </div>

            <div class="card-footer">
                <form @submit.prevent="submit" class="chat__footer">
                    <div class="row">
                        <div class="col-md-10">
                            <textarea
                                rows="3"
                                v-model="msg"
                                type="text"
                                class="form-control"
                                placeholder="write something..."
                                required=""
                                autofocus=""
                                @keydown.enter.exact.prevent
                                @keyup.enter.exact="submit"

                                :disabled="!chatUser"
                            />
                        </div>
                        <div class="col-md-2">
                            <button
                                class="btn btn-primary btn-block chat__send-btn"
                                type="submit"
                                :disabled="!chatUser"
                            >
                                send
                            </button>
                        </div>
                    </div>
                </form>
            </div>
        </div>
    </div>
</template>

<script>
import {mapActions, mapGetters, mapMutations} from 'vuex'

    export default {
        name: `ChatPage`,

        data: () => ({
            chatUser: null,
            msg: "",
            messages: []
        }),

        computed: {
            ...mapGetters('user', ['getUser']),
        },

        created() {
            this.$store.watch(
                (state, getters) => getters['ws/getMessages'],
                (val) => {
                    this.messages = val[this.$route.params.userId];

                    this.$nextTick(() => {
                        let container = this.$el.querySelector("#chat__main");
                        container.scrollTop = container.scrollHeight + 100;
                    });
                }, {
                    deep: true
                }
            );
        },

        async mounted() {
            this.LOADER_ACTIVE();
            this.chatUser = await this.getUserById(this.$route.params.userId);
            await this.loadMessagesWithUser(this.$route.params.userId);
            this.LOADER_DISABLE();

            if (!this.chatUser) {
                this.$toast.error('User for chat is not found!');
            }
        },

        methods: {
            ...mapActions('ws', ['sendMessage', 'loadMessagesWithUser']),
            ...mapActions('friends', ['getUserById']),
            ...mapMutations('system', ['LOADER_DISABLE', 'LOADER_ACTIVE']),

            async submit() {
                this.sendMessage({
                    message: this.msg,
                    fromUserId: this.getUser.user_id,
                    toUserId: this.$route.params.userId
                });
                this.msg = null;
            },
        }
    }
</script>

<style>
    .chat {
    }

    .chat__send-btn {

    }

    .chat__main-item_self-msg {
        justify-content: flex-end;
    }

    .chat__main-item:last-child {
        margin-bottom: unset !important;
    }

    .chat__footer {
    }

    .chat__main {
        height: 600px;
        overflow: auto;
    }

    .chat__no-user {
        background-color: rgba(245, 245, 245, 1);
        opacity: .4;
    }
</style>
