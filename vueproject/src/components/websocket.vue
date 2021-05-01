<template>
    <div>
        <el-button type="primary">发送</el-button>
    </div>
</template>

<script>
export default{
    name:'websocket',
    data(){
        return {
            // path:'ws://47.99.242.48:7777/ws',
            path:"ws://127.0.0.1:61/ws?uid=2"+"&to_uid=1",
            socket:{}
        }
    },
    methods:{
        init:function(){
            this.socket = new WebSocket(this.path);
            this.socket.onopen=this.open;
            this.socket.onerror=this.error;
            this.socket.onmessage=this.message;

        },
        error:function(event){
            console.log(event);
        },
        open:function(){
            this.socket.send('hi');
            
            console.log('连接成功');
        },
        message:function(event){
            console.log(event);
            console.log(event.data);
        },
        close:function(event){
            console.log(event)
            console.log('连接关闭');
        }
    },
    mounted(){
        this.init();
    },
    destroyed(){
        this.socket.onclose=this.close;
    }
}
</script>

<style>

</style>