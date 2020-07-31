<template>
    <div class="d-flex flex-column flex-md-row align-items-center p-3 px-md-4 mb-3 bg-white border-bottom box-shadow">
        <router-link class="my-0 mr-md-auto text-dark font-weight-normal" :to="$routes.index">Home</router-link>
        <template v-if="null !== getUser">
            <nav class="my-2 my-md-0 mr-md-3">
                <a @click.prevent="" class="p-2 text-dark" href="#">{{getUser.name}}</a>
            </nav>
            <a @click.prevent="doLogout" class="btn btn-warning" href="#">Log out</a>
        </template>
        <template v-else>
            <router-link class="btn btn-outline-primary" :to="$routes.register">Sign up</router-link>
            <router-link class="btn btn-outline-secondary ml-2" :to="$routes.login">Log in</router-link>
        </template>
    </div>
</template>

<script>
    import {mapGetters, mapActions} from 'vuex'

    export default {
        computed: {
            ...mapGetters('user', ['getUser']),
        },
        methods: {
            ...mapActions('user', ['logout']),

            async doLogout () {
                await this.logout();
                this.$router.push({name: 'index'});
            }
        }
    }
</script>

<style>

</style>
