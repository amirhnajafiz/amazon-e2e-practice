<script setup>
import { RouterView, RouterLink } from 'vue-router'
</script>

<template>
  <nav>
    <ul>
      <li><RouterLink to="/" exact-active-class="active">Home</RouterLink></li>
      <li><RouterLink to="/about" exact-active-class="active">About</RouterLink></li>
    </ul>

    <a href="https://github.com/amirhnajafiz/amazon-e2e-practice" target="_blank">
      <i class="fa-brands fa-github" style="font-size: 1.5em;"></i>
    </a>
  </nav>
  <RouterView />
</template>

<style>
nav {
  background-color: #f8f9fa;
  padding: 20px;
}

nav ul {
  list-style-type: none;
  padding: 0;
}

nav li {
  display: inline;
  margin-right: 30px;
}

nav a {
  text-decoration: none;
  color: #007bff;
}

nav a:hover {
  text-decoration: underline;
}

nav a.active {
  font-weight: bold;
  color: #0056b3;
}
</style>
