<template>
    <div>
        <div style="height: 50px"></div>
        <el-col :offset="5">
        <el-container>
            <el-main>
                <div>
                    <p style="color:rgb(11, 145, 145);">个人信息</p>
                </div>
                <div class="inform">
                <el-form ref="form" :model="form" label-width="80px" :rules="rules">
                    
                    <el-form-item label="出生年月" prop="date">
                        <el-date-picker type="date" style="width: 150px;" v-model="form.date"></el-date-picker>
                    </el-form-item>
                    <el-form-item label="星座" prop="star">
                        <el-select v-model="form.star">
                            <el-option v-for="item in constellation" :key="item.index" :label="item.name" :value="item.id"></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="性别" prop="sex">
                        <el-select v-model="form.sex">
                            <el-option label="男" value="man"></el-option>
                            <el-option label="女" value="woman"></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="地区" prop="adder">
                        <el-input v-model="form.adder" style="width:300px;"></el-input>
                    </el-form-item>
                    <el-form-item>
                    <el-button type="primary" @click.native="sendmessage('form')">确定</el-button>
                    <el-button @click.native="back">返回</el-button>
                    </el-form-item>
                </el-form>
                </div>
            </el-main>
        </el-container>
        </el-col>
        <el-dialog title="提示" :visible.sync="show" width="30%" >
            <span>已完善信息，请登录</span>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click.native="tologin">确定</el-button>
            </span>
        </el-dialog>
    </div>
</template>

<script>
export default{
    name:'userinformation',
    data(){
        return{
            show:false,
            constellation:[{name:'白羊座', id:1},{name:'金牛座', id:2},{name:'双子座', id:3},{name:'巨蟹座', id:4},{name:'狮子座', id:5},{name:'处女座', id:6},{name:'天秤座', id:7},{name:'天蝎座', id:8},{name:'射手座', id:9},{name:'摩羯座', id:10},{name:'水瓶座', id:11},{name:'双鱼座', id:12}],
            form:{
                date:'',
                star:'',
                sex:'',
                adder:''
            },
            rules:{
                date:[
                    {type: 'date', required:true, message:'不能为空', trigger: 'change'}
                ],
                star:[
                    {required:true, message:'不能为空', trigger: 'change'}
                ],
                sex:[
                    {required:true, message:'不能为空', trigger: 'change'}
                ],
                adder:[
                    {required:true, message:'不能为空'}
                ]
            }
        }
    },
    methods:{
        sendmessage:function(form){
            this.$refs[form].validate(
                (valid) => {
                    if(valid){
                        
                        this.show=true;
                    }
                }
            )
        },
        back:function(){
            this.$router.go(-1);
        },
        tologin:function(){
                this.show=false;
                this.$router.replace('/login')
            }
    }
}
</script>

<style>
.el-main {
    background-color: #FFF;
    text-align: center;
}
.el-form-item__content{
    text-align: left;
}
.inform {
    display: flex;
    justify-content: center;
    align-items: center;
}
</style>