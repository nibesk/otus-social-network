<template>
    <ul class="list-group">
        <li v-for="friend in getAvailableFriends" class="list-group-item available_friend_item">
            <div>
                <p class="mb-0 friend__name">{{ friend.name }} {{ friend.surname }} ({{ friend.city }})</p>
                <p class="mb-0 friend__email">{{ friend.age }} years old</p>
                <p class="mb-0 friend__email">{{ friend.email }}</p>
            </div>
            <div >
                <button v-if="!friend.is_friend" @click="addFriend(friend)" class="btn btn-success">Add friend</button>
                <button v-else class="btn btn-danger" @click="deleteFriend(friend)">Remove</button>
            </div>
        </li>

        <infinite-loading @infinite="infiniteHandler"></infinite-loading>
    </ul>
</template>

<script>
import {mapGetters, mapActions, mapMutations} from 'vuex'

export default {
    name: `FriendsList`,

    computed: {
        ...mapGetters('user', ['getUser']),
        ...mapGetters('friends', ['getAvailableFriends', 'getLastLoadedAvailableFriendsBatch']),
    },

    methods: {
        ...mapActions('friends', ['apiAvailableGetFriends', 'apiAddFriend', 'apiDeleteFriend']),
        ...mapMutations('system', ['LOADER_ACTIVE', 'LOADER_DISABLE']),

        async addFriend(user) {
            this.LOADER_ACTIVE();

            const responseMessage = await this.apiAddFriend({
                friend_user_id: user.user_id,
            });

            if (responseMessage.status) {
                this.$toast.success('Friend has been added');
            }

            this.LOADER_DISABLE();
        },

        async deleteFriend(user) {
            this.LOADER_ACTIVE();

            const responseMessage = await this.apiDeleteFriend({
                friend_user_id: user.user_id,
            });

            if (responseMessage.status) {
                this.$toast.success('Friend has been removed');
            }

            this.LOADER_DISABLE();
        },

        async infiniteHandler($state) {
            await this.apiAvailableGetFriends();

            if (this.getLastLoadedAvailableFriendsBatch.length > 0) {
                $state.loaded();
            } else {
                $state.complete()
            }
        }
    }
}
</script>

<style scoped lang="less">
.available_friend_item {
    display: flex;
    justify-content: space-between;
    align-items: center;
}
</style>
