<template>
    <div>
        <h1>ユーザー登録画面</h1>
        <form>
            <label for="name">ユーザー名</label>
            <input type="name" id="name" v-model="name">
            <br><br>
            <label for="email">メールアドレス</label>
            <input type="email" id="email" v-model="email">
            <br><br>
            <label for="password">パスワード</label>
            <input type="password" id="password" v-model="password">
            <p>※パスワードは半角英数字4〜8桁で入力してください</p>
            <br><br>

            <button @click="signup">登録</button>
        </form>
    </div>
</template>


<script>
import { mapActions } from 'vuex';

export default {
    data() {
        return {
            name: "",
            password: "",
            email: ""
        }
    },
    methods: {
        ...mapActions('token', 'signup', { root: true }),
        signup() {
            const params = new URLSearchParams();
            params.append('username', this.name);
            params.append('password', this.password);
            params.append('email', this.email);
            this.$store.dispatch('token/signup', params);
            this.name = "";
            this.password = "";
            this.email = "";
        }
    }
}
</script>