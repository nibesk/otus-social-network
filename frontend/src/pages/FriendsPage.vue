<template>
    <div class="row">
        <div class="col-md-12">
            <available-friends-list/>
        </div>
    </div>
</template>

<script>
import AvailableFriendsList from "components/AvailableFriendsList";
import {mapActions, mapMutations} from "vuex";

export default {
    name: `FriendsPage`,
    components: {
        AvailableFriendsList
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
