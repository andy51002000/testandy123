<template>
  <nav>
    <router-link to="/">Home</router-link> |
    <router-link to="/test">Test</router-link> |
    <router-link to="/transactions">Transactions</router-link> |
    <router-link to="/member-create">Create User</router-link> |
    <router-link to="/member-update">Update User</router-link> |
    <router-link to="/member-get">Get User</router-link> |




    </nav>
  <router-view />
</template>

<script>
export default {
  name: 'App',
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
