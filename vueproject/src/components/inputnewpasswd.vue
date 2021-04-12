<template>
<div>
    <div style="height: 250px"></div>
        <el-container>
            <el-col :offset="14">
                <el-form ref="form" :model="form" label-width="100px" >
                    <el-form-item label="输入验证码" prop="yzm">
                        <el-input style="width:300px;" v-model="form.yzm" show-password></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button type="primary" @click.native="send('form')">确定</el-button>
                        <el-button @click.native="goback">返回</el-button>
                    </el-form-item>
                </el-form>

            </el-col>
        </el-container>
        <el-dialog title="提示" :visible.sync="show" width="30%" >
            <span>密码已发送至邮箱</span>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click.native="tologin">确定</el-button>
            </span>
        </el-dialog>
        <el-dialog title="提示" :visible.sync="errshow" width="30%" >
            <span>验证码错误(区分大小写)</span>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click.native="close">确定</el-button>
            </span>
        </el-dialog>
</div>
</template>

<script>
export default{
    name:'inputnewpasswd',
    data(){
        return{
            show:false,
            errshow:false,
            form:{
                yzm:'',
            },
           
        }
    },
    methods:{
        
        send:function(form){
            var that = this;
            this.$refs[form].validate(
                (valid) => {
                    if(valid){
                        if(this.form.new_passwd === this.form.again_passwd){
                            let urllocal="http://127.0.0.1:60/getpassword"
                            // let url="http://47.99.242.48:60/getpassword"
                            this.$ajax.post(urllocal,{flag:this.form.yzm})
                            .then(function(response){
                                // console.log(response);
                                if(response.data["flag"]=="yes"){
                                    that.show=true;
                                }else if(response.data["flag"]=="no"){
                                    that.errshow=true;
                                }
                                // that.show = true;
                            },function(err){
                                console.log(err);
                            })
                            
                        } else {
                            this.open();
                        }
                    }
                }
            )
        },
        goback:function(){
            this.$router.go(-1);
        },
        tologin:function(){
            this.$router.replace('/login');
        },
        close:function(){
            this.errshow=false;
        }
    }
}
</script>

<style>

</style>