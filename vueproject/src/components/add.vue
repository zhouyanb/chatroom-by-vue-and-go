<template>
    <div>
        <h1>{{$store.state.count}}</h1>
        <el-button @click.native='push'>add</el-button>
        <el-button @click.native='sb'>sub</el-button>
    </div>    
</template>

<script>
import {mapMutations} from 'vuex'
export default {
    name:'add',
    data(){
        return {}
    },
    methods:{
         // ...js的展开语法
        ...mapMutations(['sub']),
        push:function(){
            // this.$store.commit('add',5);
            this.$store.dispatch('addasync',5);
        },
        sb:function(){
            this.sub();
        }
    }
}
</script>

<style>

</style>