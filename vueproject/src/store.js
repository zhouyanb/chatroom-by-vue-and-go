import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state:{
        count: 0,
        uid:'',
        username:''
    },
    mutations:{
        // 不能在mutations中执行异步操作，比如setTimeout
        add:function(state,num) {
            console.log(num);
            state.count++;
        },
        sub:function(state) {
            state.count--;
        },
        get_user_data:function(state,user){
            state.uid = user.uid;
            state.username = user.name;
        }
    },
    // actions用来执行异步任务
    actions:{
        addasync:function(context,num){
            setTimeout(() => {
                context.commit('add',num);
            },1000)
        }
    }
})