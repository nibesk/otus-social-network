<template>
    <ul class="list-group">
        <li v-for="friend in getAvailableFriends" class="list-group-item available_friend_item">
            <div>
                <p class="mb-0">{{ friend.name}} ({{ friend.email }})</p>
                <p class="mb-0">From {{ friend.city }}</p>
                <p class="mb-0">Age {{ friend.age }}</p>
            </div>
            <div >
                <button v-if="!friend.is_friend" @click="addFriend(friend)" class="btn btn-success">Add friend</button>
                <button v-else class="btn btn-danger" @click="deleteFriend(friend)">Remove</button>
            </div>
        </li>
    </ul>
</template>

<script>
import {mapGetters, mapActions, mapMutations} from 'vuex'

export default {
    name: `FriendsList`,

    computed: {
        ...mapGetters('user', ['getUser']),
        ...mapGetters('friends', ['getAvailableFriends']),
    },

    async mounted() {
        await this.apiAvailableGetFriends();
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
