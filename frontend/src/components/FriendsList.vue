<template>
    <ul class="list-group">
        <li v-for="friend in getFriends" class="list-group-item friend_item">
            <div>
                <p class="mb-0 friend__name">{{ friend.name }} {{ friend.surname }} ({{ friend.city }})</p>
                <p class="mb-0 friend__email">{{ friend.age }} years old</p>
                <p class="mb-0 friend__email">{{ friend.email }}</p>
            </div>
            <div style="display: flex">
                <router-link class="btn btn-secondary" :to="{name: 'chat', params: {userId: friend.user_id}}" style="margin-right: 10px">💬️</router-link>
                <button @click="deleteFriend(friend)" class="btn btn-danger">Remove</button>
            </div>
        </li>
    </ul>
</template>

<script>
    import {mapGetters, mapActions} from 'vuex'

    export default {
        name: `FriendsList`,

        computed: {
            ...mapGetters('user', ['getUser']),
            ...mapGetters('friends', ['getFriends']),
        },

        async mounted() {
            await this.apiGetFriends()
        },

        methods: {
            ...mapActions('friends', ['apiGetFriends', 'apiDeleteFriend']),

            async deleteFriend(user) {
                const responseMessage = await this.apiDeleteFriend({
                    friend_user_id: user.user_id,
                });

                if (responseMessage.status) {
                    this.$toast.success('Friend has been removed');
                }
            }
        }
    }
</script>

<style scoped lang="less">
.friend_item {
    display: flex;
    justify-content: space-between;
    align-items: center;
}
</style>
