<template lang="pug">
  div.hello
    h1 {{ msg }}
    h2 Essential Links
    button(@click="apiPublic") Public
    button(@click="apiPrivate") Private
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import axios from 'axios'

@Component
export default class HelloWorld extends Vue {
  private msg: string = 'Welcome to Your Vue.js App'
  private readonly URL: string = 'http://localhost:8080/'

  public async apiPublic(): Promise<unknown> {
    const res = await axios.get(`${URL}/public`)
    this.msg = res.data
  }
  public async apiPrivate(): Promise<unknown> {
    const res = await axios.get(`${URL}/private`)
    this.msg = res.data
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
