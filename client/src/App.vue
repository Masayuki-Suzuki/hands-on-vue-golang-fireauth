<template lang="pug">
  #app
    nav#nav
      router-link(to="/") Home
      | &nbsp;|&nbsp;
      router-link(to="/about") About
      | &nbsp;|&nbsp;
      router-link(to="/signup") Sign up
      | &nbsp;|&nbsp;
      router-link(to="/signin") Sign in
    router-view
</template>

<style lang="scss">
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
#nav {
  padding: 30px;
  a {
    font-weight: bold;
    color: #2c3e50;
    text-decoration: none;
    &.router-link-exact-active {
      color: #42b983;
      text-decoration: underline;
    }
  }
}
button {
  background: #42b983;
  border: solid 1px (#42b983 - #222);
  border-radius: 4px;
  color: #fff;
  font-size: 14px;
  font-weight: 600;
  text-transform: uppercase;
  display: inline-block;
  margin: 10px;
  padding: 10px 40px;
}
</style>
