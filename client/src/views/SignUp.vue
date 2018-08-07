<template lang="pug">
  .signUp
    img(alt="Vue logo" src="../assets/logo.png")
    h2 Sign up
    .inputGrp
      input(type="text" placeholder="Username" v-model="email")
      input(type="password" placeholder="Password" v-model="password")
      button register
    p
      | Do you have an account?&nbsp;
      router-link(to="/signin") sign in now!!
</template>

<script lang="ts">
  import {Component, Vue} from 'vue-property-decorator'

  @Component
  export default class SignUp extends Vue {
    private email: string = '';
    private password: string = '';
  }
</script>

<style lang="scss" scoped>
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
  .signUp {
    margin-top: 20px;
    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
    align-items: center
  }
  input {
    border: solid 1px #ccc;
    font-size: 14px;
    display: block;
    margin: 15px 0;
    padding: 10px;
    width: 200px;
  }
</style>