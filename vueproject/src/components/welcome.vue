<template>
<div>
    <el-row>
      <el-col :offset="10">
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
    name:'welcome'
}
</script>

<style>
a {
  text-decoration: none;
}
.router-link-active {
  text-decoration: none;
}
h1 {
  color: rgb(11, 145, 145);
}
</style>