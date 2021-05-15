<template>
<div>
    <el-row>
      <el-col>
        <h1>欢迎来到聊天室</h1>
      </el-col>
      <el-col :offset="11" :span="6">
        
        <router-link to="/login">登录</router-link>
        <span> | </span>
        <router-link to="/register">注册</router-link>
      </el-col>
    </el-row>
    <!-- <JwChat-index :taleList="list" @enter="binEnter" v-model="message" :showRightBox="true" scrollType="noscroll"/> -->
  </div>
</template>

<script>
export default{
    name:'welcome',
    data(){
      return {
      }
    }
}
</script>

<style>
.wel_img {
  height: 500px;
  
}
h1 {
  margin: 0 0 0 625px;
  color: rgb(11, 145, 145);
}
a {
  text-decoration: none;
}
.router-link-active {
  text-decoration: none;
}
.el-carousel__item h3 {
    color: #475669;
    font-size: 14px;
    opacity: 0.75;
    line-height: 200px;
    margin: 0;
  }
  
  /* .el-carousel__item:nth-child(2n) {
    background-color: #99a9bf;
  }
  
  .el-carousel__item:nth-child(2n+1) {
    background-color: #d3dce6;
  } */
</style>