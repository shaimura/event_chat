<template>
    <div class="chat">
        <h2 class="chat__ttl">チャット画面</h2>

        <div class="chat__index">
            <div v-for="message in messages" :key="message.Message.ID">
                <div  :class="[message.User.Username == username ? 'curentuser': 'secconduser']">
                    <div>
                        <div :class="[message.User.Username == username ? 'orange': 'blue']">{{ message.User.Username }}</div>
                        <div :class="[message.User.Username == username ? 'chat__usermessage': 'chat__seccondmessage']">{{ message.Message.Message }}</div>
                    </div>
                </div>
            </div>

            <div id="chat_messages" v-html="chatContent"></div>
        </div>

        <div class="chat__input_container">
            <input class="chat__input" type="text" v-model="newMsg" @keyup.enter="send">
            <button @click="send">Send</button>
        </div>
    </div>
</template>



<script>
import axios from "axios";

export default {
    props: ["firstuser","firstid","seconduser","secondid"],
    data() {
        return {
            ws: null,
            newMsg: '',
            chatContent: '',
            username: localStorage.getItem('userName'),
            userid: localStorage.getItem('userID'),
            messages: [],
            userchatroom: {},
            userchatroomid: "",
            reciveuserid: ""
        }
    },
    created() {
        if(this.firstid == this.userid){
            this.reciveuserid = this.secondid
        } else {
            this.reciveuserid = this.firstid
        }

        const params = new URLSearchParams();
        params.append('curentuserid',this.firstid);
        params.append('seconduserid',this.secondid);
        axios.post('/userchatroom', params).then(response => {
            this.userchatroom = response.data;
            this.userchatroomid = this.userchatroom.ID;
            console.log(this.userchatroom);
            axios.get('/getusermessage/' + this.userchatroomid).then(response => {
                this.messages = response.data
            }).catch(error => {
                console.log(error);
            });
        }).catch(error => {
            console.log(error);
        });

        


        var self = this;
        // this.ws = new WebSocket('ws://' + axios.defaults.baseURL + '/ws');
        // this.ws = new WebSocket('ws://' + window.location.host + '/ws');
        this.ws = new WebSocket('ws://localhost:8888/ws');
        console.log(this.ws);
        this.ws.addEventListener('message', function(e) {
            var msg = JSON.parse(e.data);
            let color = '';
            let senduser = '';
            let chatmessage = '';
            console.log(msg)
            // if((self.firstid == (msg.Senduserid || msg.Reciveuserid)) && (self.secondid == (msg.Senduserid || msg.Reciveuserid))){
            // if((self.firstid == msg.Reciveuserid || self.firstid == msg.Senduserid) && (self.secondid == msg.Reciveuserid || self.secondid == msg.Senduserid)){
            if(self.userchatroomid == msg.userchatroomid){
                if(msg.username === self.username) {
                    color = 'orange';
                    senduser = 'curentuser';
                    chatmessage = 'chat__usermessage';

                }else{
                    color = 'blue'; 
                    senduser = 'secconduser';
                    chatmessage = 'chat__seccondmessage';
                }
                self.chatContent += '<div class="' + senduser + '">'
                + '<div class="chip ' + color + '">'
                + msg.username + '</div>'
                + '<div class="' + chatmessage + '">' 
                + msg.message
                + '</div>'
                + '</div>';
    
                var element = document.getElementById('chat_messages');
                element.scrollTop = element.scrollHeight;
            }
        });
    },
    methods: {
        send() {
            let p = document.getElementsByTagName('<p>');
            let message = p.innerHTML = this.newMsg;
            if(this.newMsg != '') {
                this.ws.send(
                    JSON.stringify({
                        userchatroomid: this.userchatroomid,
                        username: this.username,
                        message: message,
                    })
                );
                this.newMsg = '';
            }else{
                return
            }

            const messageparams = new URLSearchParams();
            messageparams.append('userid', this.userid);
            messageparams.append('userchatroomid', this.userchatroomid);
            messageparams.append('message', message);

            axios.post('/sendusermessage', messageparams)


        }
    }
}
</script>

<style>

.chat {
    padding: 40px 120px 0;
}

.chat__ttl{
    font-size: 32px;
    margin-bottom: 40px;
    text-align: center;
}

.orange {
    color: orange;
    margin-bottom: 10px;
}

.blue {
    color: blue;
    margin-bottom: 10px;
}

.none{
    display: none;
}



.chat__input{
    display: block;
    padding: 10px;
    border: 2px solid #aaa;
    border-radius: 5px;
    font-size: 16px;
    outline: none;
    text-align: left;
    margin: 10px auto;
}

.chat__index{
    position: relative;
    margin-bottom: 40px;
}



.curentuser {
    display: flex;
    justify-content: flex-end;
    margin: 5px 0;
    clear: both;
}

.secconduser {
    position: relative;
    display: flex;
    justify-content: flex-start;
    margin-bottom: 5px;
    clear: both;
}

.chat__usermessage,
.chat__seccondmessage {
    margin-bottom: 20px;
    padding: 10px;
    border: 1px solid #eee;
}

.chat__input_container {
    margin-bottom: 20px;
}

</style>