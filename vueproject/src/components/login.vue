<template>
    <div>
        <el-row>
            <el-col>
                <div class="top"></div>
            </el-col>
        </el-row>
        
            <div class="main_login_box">
                <div class="formtitle">
                        <i class="el-icon-s-custom"></i>
                </div>
                <div class="loginform">
                    <el-form ref="form" :model="form" label-width="80px" :rules="rules">
                        <formitem :label="label.user" :prop="prop.user" @child-event='get_user_child_data'></formitem>
                        <formitem :label="label.passwd" :prop="prop.passwd" @child-event='get_passwd_child_data'></formitem>
                        <el-form-item>
                            <el-button type="primary" @click.native="senddata('form')">登录</el-button>
                            <el-button @click.native='find'>找回密码</el-button>
                            <el-button @click.native="back">返回</el-button>
                        </el-form-item>
                    </el-form>
                </div>
            </div>
       
        <el-dialog title="提示" :visible.sync="show" width="30%" >
            <span>登录成功</span>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click.native="tohome">确定</el-button>
            </span>
        </el-dialog>
        <el-dialog title="失败" :visible.sync="errshow" width="30%">
            <span>用户名或密码错误</span>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click.native="relogin">确定</el-button>
            </span>
        </el-dialog>
    </div>
</template>

<script>
import formitem from './formitem.vue'
import {mapMutations} from 'vuex'
    export default{
        name:'login',
        data(){
            return{
                show:false,
                errshow:false,
                label:{
                    user:'用户名',
                    passwd:'密码'
                },
                prop:{
                    user:'username',
                    passwd:'password'
                },
                form:{
                    username:'',
                    password:''
                },
                rules:{
                    username:[
                        {required:true, message:'输入不能为空'},
                        {min:3, max:10, message:'用户名长度在3到10之间' }
                    ],
                    password:[
                        {required:true, message:'输入不能为空'},
                        {min:6, max:16, message:'密码长度在6到16之间'}
                    ]
                }
            }
        },
        components:{
            formitem
            
        },
        methods:{
            ...mapMutations(['get_user_data']),
            get_user_child_data:function(data){
                this.form.username = data;
            },
            get_passwd_child_data:function(data){
                this.form.password = data;
            },
            senddata:function(form){
                var that = this;
                this.$refs[form].validate(
                    (valid) => {
                        if(valid){
                            // alert("登录成功");
                            //http://47.99.242.48:8080
                            // let url="http://127.0.0.1:60/login";
                            let url="http://47.99.242.48:60/login";
                            this.$ajax.post(url,{username:this.form.username,password:this.form.password})
                            .then(function(response){
                                console.log(response.data);
                                if(response.data['flag']==true){
                                    that.show=true;
                                    var uid = '00001';
                                    var username = that.form.username;
                                    that.get_user_data({
                                        uid:uid,
                                        name:username
                                    });
                                } else {
                                    that.errshow=true;
                                }
                            },function(err){
                                console.log(err);
                            })
                            
                            // this.$router.replace('/home');
                        } else {
                            console.log('error');
                            return false;
                        }
                    }
                );
            },
            back:function(){
                this.$router.go(-1);
            },
            find:function(){
                this.$router.replace('/findpassword');
            },
            tohome:function(){
                this.show=false;
                this.$router.replace('/home')
            },
            relogin:function(){
                this.errshow=false;
            }
        }
    }
</script>

<style>
.main_login_box {
    width: 500px;
    margin: 0 0 0 490px;
}
.top {
    height: 150px;
}
.loginform {
    
    height: 200px;
    width: 450px;
    display: flex;
    justify-content: center;
    align-items: center;
   
}
.formtitle{
    width: 500px;
    text-align: center;
    font-size: 50px;

}

</style>

