<template lang="pug">
  .hello
    h1 {{ msg }}
    h2 Essential Links
    button(@click="signOut") Sign out
    button(@click="apiPublic") Public
    button(@click="apiPrivate") Private
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import axios from 'axios'
import firebase from 'firebase'

@Component
export default class HelloWorld extends Vue {
  private msg: unknown = 'Welcome to Your Vue.js App'
  // private readonly URL: string = 'http://localhost:8888/'

  async apiPublic(): Promise<unknown> {
    const res = await axios.get('http://localhost:8888/public')
    this.msg = res.data
  }
  async apiPrivate(): Promise<unknown> {
    const res = await axios.get('http://localhost:8888/private', {
      headers: { Authorization: `Bearer ${sessionStorage.getItem('jwt')}` }
    })
    console.log(res.data)
    this.msg = res.data
  }
  async signOut() {
    const res = await firebase.auth().signOut()
    sessionStorage.removeItem('jwt')
    this.$router.push('/signin')
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
h1,
h2 {
  font-weight: 500;
}
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
