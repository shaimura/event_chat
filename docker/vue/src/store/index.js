import Vue from 'vue';
import Vuex from 'vuex';
import token from './modules/token';
import createPersistedState from "vuex-persistedstate";

Vue.use(Vuex);

export default new Vuex.Store({
    modules: {
        token: token
    },
    plugins: [createPersistedState()]
})