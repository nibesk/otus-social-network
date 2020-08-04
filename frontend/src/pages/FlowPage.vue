<template>
    <div class="row">
        <div class="col-md-4">
            <user/>
        </div>

        <div class="col-md-8">
            <div style="display: flex" class="justify-content-md-end">
                <router-link class="btn btn-info mb-2" :to="$routes.friends">Find new friends</router-link>
            </div>

            <template v-if="getFriends.length === 0">
                <div class="card">
                    <div class="card-body">
                        <h5 class="card-title">Oops! looks like your friend list is empty :(</h5>
                        <p class="card-text">Click button above to find them!</p>
                    </div>
                </div>
            </template>

            <friends-list/>
        </div>
    </div>
</template>

<script>
import FriendsList from "components/FriendsList";
import User from "components/User";
import {mapGetters} from "vuex";

export default {
    name: `FlowPage`,
    components: {
        FriendsList,
        User
    },
    computed: {
        ...mapGetters('friends', ['getFriends']),
    }
}
</script>

<style>
</style>
