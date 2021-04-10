<template>
  <div>
    <router-view></router-view>
  </div>
  <!-- <test></test> -->
  <!-- <websocket></websocket> -->
</template>

<script>
// import websocket from './components/websocket.vue'
// import test from './components/test.vue'
export default{
  name:'app',
  components:{
    // websocket
  },
  created () {
  setTimeout(() => {
      window.L2Dwidget.init({
        pluginRootPath: 'live2dw/',
        pluginJsPath: 'lib/',
        pluginModelPath: 'live2d-widget-model-koharu/assets/',
        tagMode: false,
        debug: false,
        model: { jsonPath: 'live2dw/live2d-widget-model-koharu/assets/koharu.model.json' },
        display: { position: 'left', width: 150, height: 300 },
        mobile: { show: true }
      })
    }, 1000)
  }
}
</script>


<style>

</style>
