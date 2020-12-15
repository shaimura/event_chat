<template>
    <div>
        <h1>トップ画面</h1>
        <h3>{{ hello }}</h3>
        <h3>{{ day }}</h3>
        

        <h2>ユーザー 一覧</h2>
        <div v-for="user in users" :key="user.ID">
            <h3>{{ user.ID }}</h3>
            <h3>{{ user.Username }}</h3>
            <router-link @click="userchatroom(user.ID)" :to="curentusername + '/' + curentuserid + '/' + user.Username + '/' + user.ID + '/chatroom'">チャットルーム</router-link>
            <button @click="userchatroom(user.ID)">テスト</button>
            <hr>
        </div>

    </div>
</template>


<script>
import axios from "axios";


export default {
    data() {
        return {
            day: "",
            hello: "",
            curentuserid: localStorage.getItem('userID'),
            curentusername: localStorage.getItem('userName'),
            users: []
        }
    },
    created() {
        console.log(axios.defaults.baseURL);
        axios.get('/test').then(response => {
            console.log(response)
            this.hello = response.data.hello;
            this.day = response.data.today;
        }).catch(error => {
            console.log(error);
        });

        const param = new URLSearchParams();
        param.append('id',this.curentuserid);
        axios.post('/getusers', param).then(response => {
            this.users = response.data;
        }).catch(error => {
            console.log(error);
        });
    },
    methods: {
        userchatroom(id) {
            const params = new URLSearchParams();
            params.append('curentuserid', this.curentuserid);
            params.append('seconduserid', id);
            console.log(params)
            axios.post('/userchatroom', params).then(response => {
                console.log(response);
            }).catch(error => {
                console.log(error);
            })
        }
    }
}
</script>