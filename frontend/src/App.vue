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
  padding: 5px;
  background: linear-gradient(90deg, #ff6a00 0%, #ffb347 50%, #4fd1c5 100%);
  border-bottom: 2px solid #fff;
  box-shadow: 0 2px 8px rgba(0,0,0,0.07);
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
  color: #fff;
  font-weight: 500;
  padding: 6px 14px;
  border-radius: 8px;
  transition: background 0.2s, color 0.2s;
}

nav a:hover,
nav a.active {
  background: rgba(255,255,255,0.2);
  color: #222;
}

.github-link {
  color: #fff;
  text-decoration: none;
  transition: color 0.2s;
}

.github-link:hover {
  color: #222;
}
</style>
