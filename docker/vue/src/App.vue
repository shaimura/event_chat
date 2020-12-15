<template>
  <div id="app">
    <header></header>
      <!-- リダイレクト時にクエリにデータを渡しているので、それがあるかないかで判断する -->
    <template v-if="$route.query.redirect">
      <p class="redirect-alert">ログインしてください</p><br>
    </template>
    <template v-if="isAuthenticated">
      <div class="header">
        <router-link class="header__item" to="/">トップ画面へ</router-link>
        <router-link class="header__item" :to="'/user/' + userid">ユーザーページへ</router-link>
        <p class="header__item" @click="vuelogout">ログアウト</p>
      </div>
    </template>
    <template v-if="!isAuthenticated">
      <div class="header">
        <router-link class="header__item" to="/signup">ユーザー登録画面へ</router-link>
        <router-link class="header__item" to="/login">ログイン画面へ</router-link>
      </div>
    </template>
    <!-- <img alt="Vue logo" src="./assets/logo.png"> -->
    <router-view></router-view>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex';

export default {
  data() {
    return {
      userid: localStorage.getItem('userID')
    }
  },
    computed: {
        ...mapGetters('token', ['idToken']),
        isAuthenticated() {
            return this.$store.getters['token/idToken'] !== null;
        }
    },
    methods: {
        ...mapActions('token', 'logout'),
        vuelogout() {
            this.$store.dispatch('token/logout');
        }
    }
}


</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

.header{
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
  width: 80%;
  margin: auto;
}

.header__item{
  margin-left: 20px;
  cursor: pointer;
}

.redirect-alert{
  color: #f00;
  display: inline-block;
}
</style>
