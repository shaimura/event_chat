<template>
    <div>
        <h1>マイページ</h1>
        <br>
        <h2>名前</h2>
        <br>
        <h3>{{ user.Username }}</h3>
    </div>
</template>


<script>
import axios from "axios";

export default {
    props: ["id"],
    data() {
        return {
            user: {}
        }
    },
    created() {
        console.log(this.id)
        axios.get('/getuser/' + this.id).then(response => {
            console.log(response);
            this.user = response.data;
        }).catch(error => {
            console.log(error);
        })
        console.log(this.user);
    }
}
</script>

