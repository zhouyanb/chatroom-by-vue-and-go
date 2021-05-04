<template>
    <div>
        <div style="height:200px;"></div>
        <el-col :offset="9">
            <el-container>
                <el-form ref="form" :model="form" label-width="80px" :rules="rules"> 
                    <formitem :label="label" :prop="prop" @child-event='get_email_child_data'></formitem>
                    <el-form-item>
                        <el-button type="primary" @click.native="createcode('form')">确定</el-button>
                    </el-form-item>
                </el-form>
            </el-container>
        </el-col>
        <el-dialog title="验证" :visible.sync="show" width="30%">
            <span style="font-size:20px; color:rgb(11, 145, 145);">请输入验证码:{{ code }}</span>
            <br>
            <el-input style="width:200px;" v-model="inputcode" placeholder="不区分大小写"></el-input>
            <el-button type="primary" @click.native="send">确定</el-button>
        </el-dialog>
        <el-dialog title="提示" :visible.sync="errshow" width="30%" >
            <span>该邮箱未被注册!</span>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click.native="close">确定</el-button>
            </span>
    </el-dialog>
    </div>
</template>

<script>
// document.oncontextmenu = function(){
//     return false;
// }
// document.onkeydown = function(){
//     if (event.ctrlKey && window.event.keyCode==67){ 
// 		return false; 
// 	} 
// 	if (event.ctrlKey && window.event.keyCode==86){ 
// 		return false; 
// 	} 

// }
import formitem from './formitem.vue'
export default{
    name:'findpassword',
    data(){
        return {
            show:false,
            errshow:false,
            code:'',
            inputcode:'',
            label:'找回密码',
            prop:'email',
            form:{
                email:''
            },
            rules:{
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
        get_email_child_data:function(data){
            this.form.email = data;
        },
        createcode:function(form){
            // var that = this;
            this.$refs[form].validate(
                (valid) => {
                    if(valid){
                        var code = "";
                        var codelen = 4;
                        var codes = new Array(0,1,2,3,4,5,6,7,8,9,'A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R', 
                                            'S','T','U','V','W','X','Y','Z');
                        for(var i=0;i<codelen;i++){
                            //floor为向下取整 random返回0到1的随机数
                            var index = Math.floor(Math.random()*36);
                            code += codes[index];
                        }
                        this.code = code;
                        this.show = true;
                    } else {
                        return false;
                    }
                }
            )
        },
        open:function(){
            this.$notify.error({
                title:'错误',
                message:'验证码错误'
            });
        },
        close:function () {
            this.errshow=false;
        },
        send:function(){
            var that=this;
            var input = this.inputcode.toUpperCase();
            if(input == this.code){
                // let url="http://127.0.0.1:60/findpassword";
                let url="http://47.99.242.48:60/findpassword";
                this.$ajax.post(url,{email:this.form.email})
                .then(function(response){
                    // console.log(response);
                    if(response.data["flag"]=="no"){
                        that.errshow=true;
                    }else{
                        that.show = false;
                        that.$router.replace('/checkemail');
                    }
                },function(err){
                    console.log(err);
                })
            } else {
                this.open();
            }
        }
    }
}
</script>

<style>

</style>