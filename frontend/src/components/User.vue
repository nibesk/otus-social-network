<template>
    <div class="card mb-2" v-if="null !== getUser">
        <img src="social.jpg" alt="Image" class="card-img-top">

        <div class="card-header">
            <h4 class="mb-0">Hi, {{ getUser.name }} {{ getUser.surname }}</h4>
        </div>

        <div class="card-body">
            <div class="card-text">
                <p><b>Age:</b> {{ getUser.age }}</p>
                <p><b>City:</b> {{ getUser.city }}</p>
                <p><b>Email:</b> {{ getUser.email }}</p>
                <p><b>Interests:</b> {{ getUser.interests }}</p>
            </div>
        </div>
    </div>
</template>

<script>
import {mapGetters} from 'vuex'

export default {
    name: `User`,

    computed: {
        ...mapGetters('user', ['getUser']),
    },
}
</script>

<style>

</style>
