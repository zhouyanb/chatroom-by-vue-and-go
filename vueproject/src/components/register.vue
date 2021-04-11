<template>
<div>
    <el-row>
            <el-col  :offset="8">
                <div class="top"></div>
            </el-col>
    </el-row>
    <el-row :gutter="20">
        <el-col  :offset="8">
            <div class="formtitle">
                        <i class="el-icon-s-custom"></i>
            </div>
            <div class="registerform">
                <el-form ref="form" :model="form" label-width="80px" :rules="rules">
                    <el-form-item label="用户名" prop="username">
                        <el-input style="width:300px;" v-model="form.username"></el-input>
                    </el-form-item>
                    <el-form-item label="密码" prop="password">
                        <el-input style="width:300px" v-model="form.password" show-password></el-input>
                    </el-form-item>
                    <el-form-item label="邮箱地址" prop="email">
                        <el-input style="width:300px;" v-model="form.email"></el-input>
                    </el-form-item>
                    <el-form-item> 
                        <el-button type="primary" @click.native="senddata('form')">注册</el-button>
                        <el-button @click.native="back">返回</el-button>
                    </el-form-item>
                </el-form>
            </div>
        </el-col>
    </el-row>
    <el-dialog title="提示" :visible.sync="show" width="30%" >
            <span>注册成功,请完善个人信息</span>
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
    export default{
        name:'register',
        data(){
            return {
                show:false,
                errshow:false,
                errmessage:'',
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
        methods:{
            senddata:function(form){
                var that = this;
                this.$refs[form].validate(
                    (valid) => {
                        if(valid){
                            urllocal="http://127.0.0.1:60/register";
                            url="http://47.99.242.48:60/register";
                           this.$ajax.post(url,{username:this.form.username,password:this.form.password,email:this.form.email})
                           .then(function(response){
                               if(response){
                                   console.log(response);
                                   that.show=true;
                               } else {
                                   that.errshow=true;
                               }
                           })
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
.top {
    height: 150px;
}
.registerform {
    
    height: 315px;
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

