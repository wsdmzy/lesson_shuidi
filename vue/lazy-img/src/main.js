import Vue from 'vue'
import App from './App.vue'

// 两个性能优化  lazyload  click
import VueLazyload from 'vue-lazyload'
Vue.use(VueLazyload, {
  //src 默认图  data-src
  // commonjs 图片 webpack 静态资源打包器
  // js css img 
  error: require('./assets/loading.svg'),  //error
  loading: require('./assets/loading.svg'), //加载中
  attemp: 1 //viewport
})

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')
