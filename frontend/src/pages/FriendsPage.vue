<template>
    <div class="row">
        <div class="col-md-12">
            <available-friends-list/>
        </div>
    </div>
</template>

<script>
import AvailableFriendsList from "components/AvailableFriendsList";

export default {
    name: `FriendsPage`,
    components: {
        AvailableFriendsList
    },
}
</script>

<style>
</style>
