<template>
    <form class="text-center">
        <h1 class="h3 mb-3 font-weight-normal">Sign up</h1>

        <div class="row justify-content-center">
            <div class="col-md-3">
                <input v-model="name" type="text" class="form-control" placeholder="Name" required="" autofocus="">

                <input v-model="surname" type="text" class="form-control mt-2" placeholder="Surname" required="" autofocus="">

                <input v-model.number="age" type="number" class="form-control mt-2" placeholder="Age" required="" autofocus="">

                <input v-model="city" type="text" class="form-control mt-2" placeholder="City" required="" autofocus="">

                <b-form-select v-model="sex" class="mt-2" :options="sexOptions">
                    <template v-slot:first>
                        <b-form-select-option :value="null" disabled>Sex</b-form-select-option>
                    </template>
                </b-form-select>

                <b-form-textarea
                    id="textarea"
                    v-model="interests"
                    placeholder="Interests"
                    class="mt-2"
                    rows="3"
                    max-rows="6"
                ></b-form-textarea>
            </div>

            <div class="col-md-3">
                <input v-model="email" type="email" class="form-control" placeholder="Email address" required="" autofocus="">

                <input v-model="password" type="password" class="form-control mt-2" placeholder="Password" required="">

                <input v-model="confirmPassword" type="password" class="form-control mt-2" placeholder="Confirm password" required="">

                <button @click.prevent="submit" class="btn btn-lg btn-primary btn-block mt-2" type="submit">Submit</button>
            </div>
        </div>
    </form>
</template>

<script>
    import {mapActions} from 'vuex'
    import crypto from 'crypto-js'

    export default {
        name: `RegisterPage`,

        data: () => ({
            email: null,
            name: null,
            surname: null,
            age: null,
            city: null,
            interests: null,
            sex: null,
            password: null,
            confirmPassword: null,

            sexOptions: [
                {value: 1, text: 'Male'},
                {value: 2, text: 'Female'},
                {value: 3, text: 'Other'},
            ]
        }),

        methods: {
            ...mapActions('user', ['register']),

            async submit() {
                const pwd = crypto.MD5(this.password).toString();

                const responseMessage = await this.register({
                    email: this.email,
                    name: this.name,
                    password: pwd,
                    confirm_password: pwd,
                    surname: this.surname,
                    age: this.age,
                    city: this.city,
                    interests: this.interests,
                    sex: this.sex,
                });

                if (responseMessage.status) {
                    this.$toast.success('Register success');
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
