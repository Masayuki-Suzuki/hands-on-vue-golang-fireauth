import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import firebase from 'firebase'

Vue.config.productionTip = false

// Initialize Firebase
const config = {
  apiKey: 'AIzaSyBvqtrjU7fV7dumG4SZJI53FMPIavZtmJo',
  authDomain: 'vuego-auth.firebaseapp.com',
  databaseURL: 'https://vuego-auth.firebaseio.com',
  projectId: 'vuego-auth',
  storageBucket: 'vuego-auth.appspot.com',
  messagingSenderId: '89314316097'
}
firebase.initializeApp(config)

let app: Vue

firebase.auth().onAuthStateChanged(user => {
  if (!app) {
    new Vue({
      router,
      store,
      render: h => h(App)
    }).$mount('#app')
  }
})
