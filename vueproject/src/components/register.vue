<template>
<div>
    <el-row>
            <el-col>
                <div class="top"></div>
            </el-col>
    </el-row>
    <el-row>
        <div class="main_register_box">
            <div class="formtitle">
                        <i class="el-icon-s-custom"></i>
            </div>
            <div class="registerform">
                <el-form ref="form" :model="form" label-width="80px" :rules="rules">
                    <formitem :label="label.user" :prop="prop.user" @child-event='get_user_child_data'></formitem>
                    <formitem :label="label.passwd" :prop="prop.passwd" @child-event='get_passwd_child_data'></formitem>
                    <formitem :label="label.email" :prop="prop.email" @child-event='get_email_child_data'></formitem>
                    <el-form-item> 
                        <el-button type="primary" @click.native="senddata('form')">注册</el-button>
                        <el-button @click.native="back">返回</el-button>
                    </el-form-item>
                </el-form>
            </div>
        </div>
    </el-row>
    <el-dialog title="提示" :visible.sync="show" width="30%" >
            <span>注册成功,请验证邮箱</span>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click.native="tologin">确定</el-button>
            </span>
    </el-dialog>
    <el-dialog title="注册失败" :visible.sync="errshow" width="30%">
        <span>{{ errmessage }}</span>
        <span slot="footer" class="dialog-footer">
            <el-button type="primary" @click.native="reregister">确定</el-button>
        </span>
    </el-dialog>
</div>
</template>

<script>
import formitem from './formitem.vue'
    export default{
        name:'register',
        data(){
            return {
                show:false,
                errshow:false,
                errmessage:'',
                label:{
                    user:'用户名',
                    passwd:'密码',
                    email:'邮箱地址'
                },
                prop:{
                    user:'username',
                    passwd:'password',
                    email:'email'
                },
                form:{
                    username:'',
                    password:'',
                    email:'',
                    font_emailvalid:'',
                    back_emailvalid:''
                },
                rules:{
                    username:[
                        {required:true, message:'输入不能为空'},
                        {min:3, max:10, message:'用户名长度在3到10之间'}
                            ],
                    password:[
                        {required:true, message:'输入不能为空'},
                        {min:6, max:16, message:'密码长度在6到16之间'}
                    ],
                    email:[
                        {required:true, message:'输入不能为空'},
                        {type:'email', message:'请输入正确的邮箱地址'}
                    ]
                   
                }
            }
        },
        components:{
            formitem
        },
        methods:{
            get_user_child_data:function(data){
                this.form.username = data;
            },
            get_passwd_child_data:function(data){
                this.form.password = data;
            },
            get_email_child_data:function(data){
                this.form.email = data;
            },
            senddata:function(form){
                var that = this;
                this.$refs[form].validate(
                    (valid) => {
                        if(valid){
                            // let url="http://127.0.0.1:60/register";
                            let url="http://47.99.242.48:60/register";
                           this.$ajax.post(url,{username:this.form.username,password:this.form.password,email:this.form.email})
                           .then(function(response){
                               if(response.data["flagmail"]=="yes"){
                                //    console.log(response);
                                    that.errshow=true;
                                    that.errmessage="该邮箱已被注册！";
                                //    that.show=true;
                               } else if(response.data["flagname"]=="yes"){
                                    that.errshow=true;
                                    that.errmessage="该ID已被注册！";

                               }else {
                                   that.show=true;
                               }
                           })
                        // that.show =true;
                        } else {
                            console.log('error submit');
                            return false;
                        }
                    }
                );
            },
            back:function(){
                this.$router.go(-1);
            },
            tologin:function(){
                this.show=false;
                this.$router.replace('/checkemail')
            },
            reregister:function(){
                this.errshow=false;
            }
        }
    }
</script>

<style>
.main_register_box {
    margin: 0 0 0 490px;
}
.top {
    height: 150px;
}
.registerform {
    
    height: 250px;
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

