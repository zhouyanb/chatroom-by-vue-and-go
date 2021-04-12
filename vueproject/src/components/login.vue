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
                <div class="loginform">
                    <el-form ref="form" :model="form" label-width="80px" :rules="rules">
                        <el-form-item label="用户名" prop="username">
                            <el-input style="width:300px;" v-model="form.username"></el-input>
                        </el-form-item>
                        <el-form-item label="密码" prop="password">
                            <el-input style="width:300px" v-model="form.password" show-password></el-input>
                        </el-form-item>
                        <el-form-item>
                            <el-button type="primary" @click.native="senddata('form')">登录</el-button>
                            <!-- <el-button @click.native='find'>找回密码</el-button> -->
                            <el-button @click.native="back">返回</el-button>
                        </el-form-item>
                    </el-form>
                </div>
            </el-col>
        </el-row>
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
    export default{
        name:'login',
        data(){
            return{
                show:false,
                errshow:false,
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
        methods:{
            senddata:function(form){
                var that = this;
                this.$refs[form].validate(
                    (valid) => {
                        if(valid){
                            // alert("登录成功");
                            //http://47.99.242.48:8080
                            let urllocal="http://127.0.0.1:60/login";
                            // let url="http://47.99.242.48:60/login";
                            this.$ajax.post(urllocal,{username:this.form.username,password:this.form.password})
                            .then(function(response){
                                console.log(response.data);
                                if(response.data['flag']==true){
                                    that.show=true;
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

