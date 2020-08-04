<template>
    <div class="text-center">
        <div class="row justify-content-center">
            <div class="col-md-3">
                <form class="text-center">
                    <h1 class="h3 mb-3 font-weight-normal">Log in</h1>

                    <input v-model="email" type="email" class="form-control" placeholder="Email address" required="" autofocus="">

                    <input v-model="password" type="password" class="form-control mt-2" placeholder="Password" required="">

                    <button @click.prevent="submit" class="btn btn-lg btn-primary btn-block mt-2" type="submit">Submit</button>
                </form>
            </div>
        </div>
    </div>
</template>

<script>
    import {mapActions} from 'vuex';

    export default {
        name: `NamePage`,

        data: () => ({
            email: null,
            password: null,
        }),

        methods: {
            ...mapActions('user', ['login']),

            async submit() {
                if (!this.email || !this.password) {
                    this.$toast.error('email or password should be filled');
                    return;
                }

                const responseMessage = await this.login({email: this.email, password: this.password});
                if (responseMessage.status) {
                    this.$toast.success('Login success');
                    this.$router.push({name: 'flow'})
                }
            }
        }
    }
</script>

<style>
    .form-signin {
        width: 100%;
        max-width: 330px;
        padding: 15px;
        margin: 0 auto;
    }
</style>
