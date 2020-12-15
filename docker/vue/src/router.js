import Vue from 'vue';
import Router from 'vue-router';
import store from './store';
// import axios from 'axios';

const Top = () => import( /* webpackChunkName: "Top" */ './views/Top.vue');
const Signup = () => import( /* webpackChunkName: "Signup" */ './views/Signup.vue');
const Login = () => import( /* webpackChunkName: "Login" */ './views/Login.vue');
const User = () => import( /* webpackChunkName: "User" */ './views/User.vue');
const Userchat = () => import( /* webpackChunkName: "Userchat" */ './views/Userchat.vue');

Vue.use(Router);


export default new Router({
    mode: 'history',
    routes: [{
        path: '/',
        component: Top,
        beforeEnter(to, from, next) {
            if (store.getters['token/idToken']) {
                next();
            } else {
                next({
                    path: '/login',
                    query: { redirect: to.fullPath}
                })
            }
        }
    },
        {
            path: '/signup',
            component: Signup
        },
        {
            path: '/login',
            component: Login
        },
        {
            path: '/user/:id',
            component: User,
            props: true
        },
        {
            path: '/:firstuser/:firstid/:seconduser/:secondid/chatroom',
            component: Userchat,
            props: true
        },
        {
        path: '*',
        redirect: '/'
        }
    ]
});