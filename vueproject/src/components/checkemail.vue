<template>
    <div>
        <div style="height: 250px"></div>
        <el-container>
            <el-col :offset="14">
                <el-form ref="form" :model="form" label-width="80px" :rules="rules">
                    <el-form-item label="邮箱验证" prop="font_email"> 
                        <el-input style="width:300px;" v-model="form.font_email"></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button type="primary" @click.native="check('form')">确定</el-button>
                        <el-button @click.native="goback">返回</el-button>
                    </el-form-item>
                </el-form>

            </el-col>
        </el-container>
        <el-dialog title="成功" :visible.sync="show" width="30%">
            <span>验证成功</span>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click.native="toinformation">确定</el-button>
            </span>
        </el-dialog>
        <el-dialog title="失败" :visible.sync="errshow" width="30%">
            <span>验证失败</span>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click.native="reinput">确定</el-button>
            </span>
        </el-dialog>
    </div>
</template>

<script>
export default{
    name:'checkemail',
    data(){
        var checkemail  = (rule,value,callback) => {
            if(!value){
                return callback(new Error('输入不能为空'));
             } else{
                 callback();
             }
            //else if(value!=this.back_email) {
            //     return callback(new Error('验证码错误'));
            // }
            // else{
            //    this.$ajax.post('http://127.0.0.1:8080/yes',{"flag":"yes"});
            // }
        }
        return{
            show:false,
            url:'',
            errshow:false,
            frompath:'',
            form:{
                font_email:'',
                back_email:'',
            },
            rules:{
                font_email:[
                    {validator:checkemail, trigger:'blur'}
                ]
            }
        }
    },
    methods:{
        check:function(form){
            var that = this;
            this.$refs[form].validate(
                (valid) => {
                    if(valid){
                        console.log(that.form.font_email);
                        console.log(typeof(that.form.font_email));
                        
                        this.$ajax.post(this.url,{flag:that.form.font_email})
                        .then(function(response){
                           if(response.data["flag"]=='yes'){
                               that.show=true;
                           }else{
                                that.errshow=true;
                           }
                        }
                        ,function(err){
                                console.log(err);
                            })
                        that.show = true;
                    } else {
                        console.log("no");
                    // return false;
                }
                }
            );
        },
        
        toinformation:function(){
            this.show = false;
            this.$router.replace('/login');     
        },
        goback:function(){
            this.$router.go(-1);
        },
        reinput:function(){
            this.errshow = false;
        }
    },
    //vue-router的钩子函数 在渲染该组件对应的路由被confirm前调用
    //不能获取this 因为在执行时组件实例还没有创建
    // to 即将要进入目标的路由对象
    // from 当前导航正要离开的路由
    // next 为一个function
    beforeRouteEnter(to , from , next){
        next((vm) => {
            // 通过vm来访问组件实例
            vm.frompath = from.path;
        });
    },
    mounted(){
        if(this.frompath === '/register'){
            this.url = 'http://47.99.242.48:60/checkmail';
        } else {
            this.url = 'http://47.99.242.48:60/getpassword';
        }
    }
}
</script>

<style>

</style>