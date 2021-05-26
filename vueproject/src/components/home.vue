<template>
    <div class="mainbox">
        <div style="height: 50px"></div>
        <el-col >
        <el-container >
            <el-aside width="200px">
                <!--页面左侧在线用户栏-->
                <el-menu :default-openeds="['1']" background-color="rgb(76 99 123)" class="list">
                    <el-submenu index="1">
                        <template slot="title" >
                            <i class="el-icon-user-solid" ></i>
                            <span style="color:white">在线用户</span>
                        </template> 
                        <el-menu-item-group>
                            <el-menu-item  v-for="item in userlist" v-bind:key="item.index" style="color:white">
                                <!--<router-link :to="{name:'userhome',params:{username:item}}" class="el-icon-user"></router-link>-->
                                  {{ item }} 
                            </el-menu-item>
    
                        </el-menu-item-group>
                    </el-submenu>
                </el-menu>
            </el-aside>
            <el-container>
                <el-main id='main'>
                    <div v-for="item,index in messagelist" :key="item.index">
                        <div v-bind:style="[showlist[index] ? right_style : no_style]">
                            <el-tag effect="Plain">{{ item.content }}</el-tag><span> : {{ item.sender }}</span>
                        </div>
                        <div v-bind:style="[!showlist[index] ? left_style : no_style]">
                            <span>{{ item.sender }} : </span><el-tag effect='Plain'>{{ item.content }}</el-tag>
                        </div>    
                    </div>    
                </el-main> 
                <el-footer height="150px">
                    <div style="height:110px; width:700px; ">
                        <el-input type="textarea" v-model="content" @keyup.enter.native="putmessage" ></el-input>
                    </div>
                    <div class="buttonbox">
                        <el-button type="primary" size="mini" @click.native="putmessage" style="background-color: #516479;border-color: #516479;">发送</el-button>
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
            content:'',
            messagelist:[],
            userlist:[],
            // path:"ws://127.0.0.1:60/ws",
            path:"ws://47.99.242.48:60/ws",
            socket:{},
            right_style:{
                clear:'both'
            },
            left_style:{
                clear:'both',
                float:'left'
            },
            no_style:{
                display:'none'
            },
            showlist:[]
        }
    },
    methods:{
        initws:function(){
            this.socket = new WebSocket(this.path);
            this.socket.onopen=this.onopen;
            this.socket.onerror=this.onerror;
            this.socket.onmessage=this.onmessage;
        },
        getCookie:function(){
            // console.log(this.$cookies.get("name"));
            return this.$cookies.get("name");
        },
        onopen:function(){
            console.log('连接成功');
            console.log(this.$store.state.username)
            // this.socket.send({username:this.$store.state.username})
            // var username=this.$store.state.username;
            var username=this.getCookie();
            var json=JSON.stringify(username);
            this.socket.send(json);
        },
        onerror:function(event){
            console.log(event);
        },
        onmessage:function(event){
            // console.log(event.data);
            var messageobj = JSON.parse(event.data);
            if(messageobj.sender == 'avalon'){
                console.log(messageobj)
                // var user = messageobj.content.split(':')[1];
                this.$message(messageobj.content);
                // console.log(messageobj);
                
            }
            else if(messageobj.sender == 'Invisible Air'){
                console.log(messageobj.content)
                this.userlist.push(messageobj.content);
            }
            else {
                this.messagelist.push(messageobj);
                if(document.hidden){
                    document.getElementsByTagName('title')[0].innerText='you have a message';
                }else {
                    document.getElementsByTagName('title')[0].innerText='vueproject';
                }
                if(messageobj.sender == this.$store.state.username){
                   this.showlist.push(true);
                } else {
                    this.showlist.push(false);
                }
            }
        },
        putmessage:function(){
            var messageobj={sender:this.getCookie(), content: this.content};
            var json = JSON.stringify(messageobj);
            this.socket.send(json);
            this.content = '';
        }
    },
    mounted(){
        // this.$ajax.post('http://127.0.0.1:60/test',{username:'zyb'})
        // .then((response) => {
        //     console.log(response);
        // },(err) => {
        //     console.log(err);
        // });
        if(this.getCookie()==null){
            //在这里填入弹出窗口让用户登录，点击确定跳转到登录页面

            console.log("fuck!");
            window.stop();
            document.execCommand("Stop");
        }
        this.initws();
        document.getElementsByTagName('title')[0].innerText='vueproject';
    }
}
</script>

<style>
/* body{
    background-color: rgba(0, 0, 0, .15);
} */
.mainbox {
    margin: 0 0 0 295px;
}
.el-container {
    height: 600px;
    width: 900px;
}
.el-footer {
    
    /* text-align: center; */
    background-color: white;
    /* box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1); */
    border: none;
    padding: 0%;
    line-height: 60px;
}
.list.el-menu {
    border: none;
}
#main {
    background-color: #516479;
    text-align: right;
    line-height: 60px;
}
.list {
    height: 600px;
}
.el-textarea__inner {
    resize: none;
    /* border-color: #516479; */
    border: none;
    width:700px;
    height:100px;
    padding: 2px;
    border-radius: 0px
}
.buttonbox {
    line-height: 0%;
    text-align: right;
}
.el-button {
    margin: 0 10px 0 0;
}

</style>