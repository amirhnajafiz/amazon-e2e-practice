<script setup>
import { RouterView, RouterLink } from 'vue-router'
</script>

<template>
  <nav>
    <div class="nav-content">
      <ul>
        <li><RouterLink to="/" exact-active-class="active">Home</RouterLink></li>
        <li><RouterLink to="/about" exact-active-class="active">About</RouterLink></li>
      </ul>
      <a href="https://github.com/amirhnajafiz/amazon-e2e-practice" target="_blank" class="github-link">
        <i class="fa-brands fa-github" style="font-size: 1.5em;"></i>
      </a>
    </div>
  </nav>
  <RouterView />
</template>

<style>
nav {
  padding: 20px;
  background-color: #f8f9fa;
  border-bottom: 1px solid #dee2e6;
}

.nav-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

nav ul {
  list-style-type: none;
  padding: 0;
  margin: 0;
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

.github-link {
  color: #333;
  text-decoration: none;
}
</style>
