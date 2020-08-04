<template>
<div id="app" class="app">
  <header-bar></header-bar>
  <div class="page container">
    <router-view></router-view>
  </div>
</div>
</template>

<script>
import HeaderBar from 'components/HeaderBar'
import {mapGetters, mapState} from "vuex";
export default {
  components: {
    HeaderBar
  },
  name: 'app',
  computed: {
    ...mapState('user', ['user']),
  },
}
</script>

<style>
  body {
    font-family: "HelveticaNeue", "Helvetica Neue", Helvetica, Arial, sans-serif;
    color: #333;
  }
  html, body {
    height: 100%;
  }
  .centered {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
  }
</style>
