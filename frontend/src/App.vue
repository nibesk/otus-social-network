<template>
    <div id="app" class="app">
        <header-bar></header-bar>
        <div class="page container">
            <loading :active.sync="getLoader"></loading>
            <router-view></router-view>
        </div>
    </div>
</template>

<script>
import HeaderBar from 'components/HeaderBar'
import {mapGetters, mapMutations} from "vuex";

export default {
    components: {
        HeaderBar
    },
    name: 'app',
    computed: {
        ...mapGetters('system', ['getLoader']),
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
</style>
