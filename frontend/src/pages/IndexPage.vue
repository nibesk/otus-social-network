<template>
  <div class='row justify-content-center align-content-center text-center index'>
      <div class="col-md-6">
          <h1 class="cover-heading">Social network</h1>
          <p class="lead">
              Social networking is the use of Internet-based social media sites to stay connected with friends, family, colleagues, customers, or clients
          </p>
          <p class="lead">
              <router-link class="btn btn-lg btn-secondary" :to="$routes.register">Sign up</router-link>
          </p>
      </div>
  </div>
</template>

<script>
  import { mapGetters, mapActions } from 'vuex'
  export default {
      name: `IndexPage`,
  }
</script>

<style scoped lang="less">
.index {
    height: 80vh;
}
</style>
