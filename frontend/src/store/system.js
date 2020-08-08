export default {
    namespaced: true,

    state: () => ({
        loader: false
    }),

    mutations: {
        LOADER_ACTIVE(state) {
            state.loader = true
        },
        LOADER_DISABLE(state) {
            state.loader = false
        },
    },

    getters: {
        getLoader: state => state.loader,
    }
}
