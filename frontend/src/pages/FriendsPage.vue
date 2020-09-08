<template>
    <div class="row">
        <div class="col-md-8">
            <available-friends-list/>

        </div>
        <div class="col-md-4">
            <search-available-friends/>

        </div>
    </div>
</template>

<script>
import AvailableFriendsList from "components/AvailableFriendsList";
import SearchAvailableFriends from "components/SearchAvailableFriends";
import {mapMutations} from "vuex";

export default {
    name: `FriendsPage`,
    components: {
        AvailableFriendsList,
        SearchAvailableFriends
    },

    beforeRouteLeave(to, from, next) {
        this.RESET_AVAILABLE_FRIENDS();
        next();
    },

    methods: {
        ...mapMutations('friends', ['RESET_AVAILABLE_FRIENDS']),
    }
}
</script>

<style>
</style>
