<template lang="pug">
  .signIn
    img(alt="Vue logo" src="../assets/logo.png")
    h2 Sign in
    .inputGrp
      input(type="text" placeholder="Username" v-model="email")
      input(type="password" placeholder="Password" v-model="password")
      button(@click="signIn") register
    p
      | You don't have an account?&nbsp;
      router-link(to="/signup") create account now!!
</template>

<script lang="ts">
  import {Component, Vue} from 'vue-property-decorator'
  import firebase from 'firebase'

  @Component
  export default class SignIn extends Vue {
    private email: string = ''
    private password: string = ''

    async signIn() {
      const res: any = await firebase.auth().signInWithEmailAndPassword(this.email, this.password).catch(err => {
        alert(err.message)
        console.error(err)
        throw new Error(err)
      })
      console.log(res.user)
      sessionStorage.setItem('jwt', res.user.qa)
      this.$router.push('/')
    }
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
  .signIn {
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