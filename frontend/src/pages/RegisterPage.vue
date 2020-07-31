<template>
  <div class="row centered">
    <form class="form-signin text-center">
      <h1 class="h3 mb-3 font-weight-normal">Sign up</h1>

      <input v-model="name" type="text" class="form-control" placeholder="Name" required="" autofocus="">

      <input v-model="email" type="email" class="form-control" placeholder="Email address" required="" autofocus="">

      <input v-model="password" type="password" class="form-control mt-2" placeholder="Password" required="">

      <input v-model="confirmPassword" type="password" class="form-control mt-2" placeholder="Confirm password" required="">

      <button @click.prevent="submit" class="btn btn-lg btn-primary btn-block mt-2" type="submit">Submit</button>
    </form>
  </div>
</template>

<script>
  import { mapActions, mapGetters } from 'vuex'
  export default {
      data: () => ({
          email: null,
          name: null,
          password: null,
          confirmPassword: null,
      }),

      methods: {
          ...mapActions('user', ['register']),

          async submit() {

              const responseMessage = await this.register({
                  email: this.email,
                  name: this.name,
                  password: this.password,
                  confirmPassword: this.confirmPassword
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
