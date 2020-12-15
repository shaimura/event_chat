<template>
    <div>
        <h2>ログイン画面</h2>
        <label for="name">ユーザー名</label>
        <input type="name" id="name" v-model="name">
        <br><br>
        <label for="email">メールアドレス</label>
        <input type="email" id="email" v-model="email">
        <br><br>
        <label for="password">パスワード</label>
        <input type="password" id="password" v-model="password" maxlength="8" pattern="^[0-9A-Za-z]+$">
        <p>※パスワードは半角英数字4〜8文字で入力してください</p>

        <br><br>
        <button @click="login">ログイン</button>
    </div>
</template>

<script>
import { mapActions } from 'vuex';

export default {
    data() {
        return {
            name: "",
            password: "",
            email: "",
            message: ""
        }
    },
    methods: {
        ...mapActions('token', 'signin'),
        login() {
            const params = new URLSearchParams();
            params.append('username', this.name);
            params.append('password', this.password);
            params.append('email', this.email);
            this.$store.dispatch('token/signin', params);
            this.name = "";
            this.password = "";
            this.email = "";

        }
    }
}
</script>