<template>
    <div>
        <div style="height: 50px"></div>
        <el-col :offset="4">
        <el-container>
            <el-aside width="200px">
                <el-menu :default-openeds="['1']" background-color="#7eb2e6" class="list">
                    <el-submenu index="1">
                        <template slot="title">
                            <span>我的好友</span>
                        </template> 
                        <el-menu-item-group>
                            <el-menu-item v-for="item in friends" v-bind:key="item.index"><router-link :to="{name:'userhome',params:{username:item}}" class="el-icon-user"></router-link>  {{ item }} </el-menu-item>
    
                        </el-menu-item-group>
                    </el-submenu>
                </el-menu>
            </el-aside>
            <el-container>
                <el-main id='main'>
                    <div v-for="item in messagelist" :key="item.index">
                        <el-tag effect="Plain">{{ item.message }}</el-tag><span> : {{ item.username }}</span>    
                    </div>    
                </el-main> 
                <el-footer height="150px">
                    <div style="height:110px; width:700px;">
                        <el-input type="textarea" v-model="message" @keyup.enter.native="putmessage"></el-input>
                    </div>
                    <div class="buttonbox">
                        <el-button type="primary" size="mini" @click.native="putmessage">发送</el-button>
                    </div>
                   
                </el-footer>
                <!-- <JwChat-index :taleList="list" @enter="binEnter" v-model="message" :showRightBox="true" scrollType="noscroll"/> -->
            </el-container>
        </el-container>
        </el-col>
    </div>
</template>

<script>

export default{
    name:'home',
    data(){
        return{
            friends:['zyb','cw'],
            message:'',
            messagelist:[],
            path:"ws://127.0.0.1:61/ws",
            socket:{}
        }
    },
    methods:{
        initws:function(){
            this.socket = new WebSocket(this.path);
            this.socket.onopen=this.onopen;
            this.socket.onerror=this.onerror;
            this.socket.onmessage=this.onmessage;
        },
        onopen:function(){
            console.log('连接成功');
            this.socket.send({username:this.$store.state.username})
        },
        onerror:function(event){
            console.log(event);
        },
        onmessage:function(event){
            console.log(event.data);
        },
        putmessage:function(){
            var messageobj={username:this.$store.state.username, message: this.message};
            this.messagelist.push(messageobj);
            this.message='';
        }
    },
    mounted(){
        this.initws();
    }
}
</script>

<style>
.el-container {
    height: 600px;
    width: 900px;
}
.el-footer {
    
    /* text-align: center; */
    /* background-color: #99CCFF; */
    /* box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1); */
    border: none;
    padding: 0%;
    line-height: 60px;
}
.list.el-menu {
    border: none;
}
#main {
    background-color: #99CCFF;
    /* background-color: #7eb2e6; */
    text-align: right;
    
    /* background-color: #3c9dc0; */
    line-height: 60px;
}
.list {
    height: 600px;
}
.el-textarea__inner {
    resize: none;
    border: none;
    width:700px;
    height:100px;
    /* background-color: -internal-light-dark(rgb(255, 255, 255), rgb(255, 255, 255)); */
    padding: 2px;
}
.buttonbox {
    line-height: 0%;
    text-align: right;
}

</style>