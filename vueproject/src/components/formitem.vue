<template>
    <el-form-item :label="label" :prop="prop">
        <el-input style="width:300px" @change="send" v-model="msg"></el-input>
    </el-form-item>
</template>

<script>
export default {
    name:'formitem',
    props:['label','prop'],
    data(){
        return{
            msg:''
        }
    },
    methods:{
        send:function(){
            this.$emit('child-event',this.msg);
        }
    }
}
</script>

<style>

</style>