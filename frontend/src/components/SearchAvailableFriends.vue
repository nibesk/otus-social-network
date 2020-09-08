<template>
    <div class="mb-4">
        <div class="card">
            <h5 class="card-header">Search friends</h5>
            <div class="card-body">
                <p class="card-text">Fill name and Surname to find new friends</p>
                <form>
                    <input v-model="name" name="search_name" type="text" class="form-control mb-2" placeholder="Name..." required>
                    <input v-model="surname" name="search_surname" type="text" class="form-control mb-2" placeholder="Surname..." required>
                    <button @click.prevent="search" class="btn btn-primary btn-block" type="submit">Search</button>
                </form>
            </div>
        </div>
    </div>
</template>

<script>
import {mapGetters, mapActions, mapMutations} from 'vuex'

export default {
    name: `SearchAvailableFriends`,

    data: () => ({
        name: null,
        surname:null
    }),

    methods: {
        ...mapActions('friends', ['apiAvailableGetFriends']),
        ...mapMutations('friends', ['RESET_AVAILABLE_FRIENDS']),
        ...mapMutations('system', ['LOADER_ACTIVE', 'LOADER_DISABLE']),

        async search() {
            if (!this.name || !this.surname) {
                this.$toast.error('Name and Surname should be filled both!');

                return;
            }

            this.LOADER_ACTIVE();
            this.RESET_AVAILABLE_FRIENDS();

            await this.apiAvailableGetFriends({
                name: this.name,
                surname: this.surname,
            });

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
