import Vue from 'vue';
import App from './App.vue';
import axios from 'axios';
import router from './router';
import store from './store';


Vue.config.productionTip = false

// golangのパス
axios.defaults.baseURL = process.env.VUE_APP_API_ENDPOINT;


// axios.defaults.withCredentials = true;


axios.interceptors.request.use(
    config => {
        return config;
    },
    error => {
        // "error"の場合は"catch"に行くようにする
        return Promise.reject(error);
    }
);
// サーバーに送った後
axios.interceptors.response.use(
    config => {
        return config;
    },
    error => {
        // "error"の場合は"catch"に行くように
        return Promise.reject(error);
    }
);

store.dispatch('token/autoLogin').then(() => {
    new Vue({
        router,
        store,
        render: h => h(App),
    }).$mount('#app')
});
