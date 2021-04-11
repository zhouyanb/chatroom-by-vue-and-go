import Vue from 'vue'
import VueRouter from 'vue-router'
import login from "../components/login.vue"
import register from "../components/register.vue"
import home from "../components/home.vue"
import welcome from "../components/welcome.vue"
import userhome from "../components/userhome.vue"
import userinformation from "../components/userinformation.vue"
import checkemail from "../components/checkemail.vue"
import findpassword from "../components/findpassword.vue"
import inputnewpasswd from "../components/inputnewpasswd.vue"

Vue.use(VueRouter)

const routes = [
  {
    path:'/',
    redirect:'/welcome'
  },
  {
    path:'/welcome',
    name:'welcome',
    component:welcome
  },
  {
    path:'/login',
    name:'login',
    component:login
  },
  {
    path:'/register',
    name:'register',
    component:register
  },
  {
    path:'/checkemail',
    name:'checkemail',
    component:checkemail
  },
  {
    path:'/findpassword',
    name:'findpassword',
    component:findpassword
  },
  {
    path:'/inputnewpasswd',
    name:'inputnewpasswd',
    component:inputnewpasswd
  },
  {
    path:'/userinformation',
    name:'userinformation',
    component:userinformation
  },
  {
    path:'/home',
    name:'home',
    component:home
  },
  {
    path:'/home/:username',
    name:'userhome',
    component:userhome,
    props: route => ({username:route.params.username}) 
  }
  
]

const router = new VueRouter({
  routes
})

export default router
