<template>
  <nav>
    <router-link to="/">Home</router-link> |
    <router-link to="/about">About</router-link>
  </nav>
  <router-view />
</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #efeff1;
}

body {
  background-color: #0e0e10;
}

nav {
  padding: 30px;
}

nav a {
  font-weight: bold;
  color: #efeff1;
}

nav a.router-link-exact-active {
  color: #bf94ff;
}

a {
  color: #bf94ff;
}
</style>
