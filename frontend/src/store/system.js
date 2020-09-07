export default {
    namespaced: true,

    state: () => ({
        loader: 0
    }),

    mutations: {
        LOADER_ACTIVE(state) {
            state.loader++
        },
        LOADER_DISABLE(state) {
            if (0 === state.loader) {
                return;
            }
            state.loader--
        },
    },

    getters: {
        getLoader: state => 0 !== state.loader,
    }
}
