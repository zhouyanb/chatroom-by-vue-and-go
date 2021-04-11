<template>
<div>
    <div style="height: 250px"></div>
        <el-container>
            <el-col :offset="14">
                <el-form ref="form" :model="form" label-width="100px" :rules="rules">
                    <el-form-item label="输入新密码" prop="new_passwd"> 
                        <el-input style="width:300px;" v-model="form.new_passwd" show-password></el-input>
                    </el-form-item>
                    <el-form-item label="再次输入" prop="again_passwd">
                        <el-input style="width:300px;" v-model="form.again_passwd" show-password></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button type="primary" @click.native="send('form')">确定</el-button>
                        <el-button @click.native="goback">返回</el-button>
                    </el-form-item>
                </el-form>

            </el-col>
        </el-container>
        <el-dialog title="提示" :visible.sync="show" width="30%" >
            <span>修改密码成功</span>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click.native="tologin">确定</el-button>
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
            form:{
                new_passwd:'',
                again_passwd:''
            },
            rules:{
                new_passwd:[
                    {required:true, message:'输入不能为空'},
                    {min:6, max:16, message:'密码长度在6到16之间'}
                ],
                again_passwd:[
                    {required:true, message:'输入不能为空'},
                    {min:6, max:16, message:'密码长度在6到16之间'}
                ]
            }
        }
    },
    methods:{
        open:function(){
            this.$notify.error({
                title:'错误',
                message:'两次输入不一样'
            });
        },
        send:function(form){
            var that = this;
            this.$refs[form].validate(
                (valid) => {
                    if(valid){
                        if(this.form.new_passwd === this.form.again_passwd){
                            // this.$ajax.post("",{newpasswd:this.form.new_passwd})
                            // .then(function(response){
                            //     console.log(response);
                            //     that.show = true;
                            // },function(err){
                            //     console.log(err);
                            // })
                            that.show = true;
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
        }
    }
}
</script>

<style>

</style>