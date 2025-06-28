<template>
  <main>
    <h1>About This Project</h1>
    <p style="text-align: justify;">
      <span class="highlight">Amazon E2E practice project</span> featuring a database, backend, and frontend.
      Includes a <span class="highlight">URLs panel</span> for accessing useful IT news and blogs,
      along with <span class="highlight">statistics tracking</span> for these URLs.
    </p>
    <div class="about-divider"></div>
    <ul>
      <li>Database: <b>PostgreSQL</b></li>
      <li>Backend: <b>Golang</b></li>
      <li>Frontend: <b>Vue.js</b></li>
    </ul>
    <div class="about-divider"></div>
    <h2>Top URLs</h2>
    <ol>
      <li v-for="(url, index) in topUrls" :key="index">
        {{ url.title }}
      </li>
    </ol>
  </main>
</template>

<script>
export default {
  name: 'AboutView',
  data() {
    return {
      topUrls: [
        {
          title: 'Hacker News',
        },
        {
          title: 'TechCrunch',
        },
        {
          title: 'The Verge',
        },
        {
          title: 'Ars Technica',
        },
        {
          title: 'Wired',
        },
        {
          title: 'Engadget',
        },
      ],
    };
  },
};
</script>

<style scoped>
main {
  max-width: 800px;
  margin: 40px auto;
  padding: 0 16px;
}

h1 {
  color: #2d3748;
  margin-bottom: 1.2rem;
  font-size: 2rem;
  font-weight: 700;
}

p {
  color: #4a5568;
  font-size: 1.1rem;
  line-height: 1.7;
}

.highlight {
  color: #2563eb;
  font-weight: 600;
  background: #e0e7ff;
  padding: 0.1em 0.4em;
  border-radius: 6px;
}

.about-divider {
  margin: 1.5rem 0 1.2rem 0;
  width: 60px;
  height: 4px;
  border-radius: 2px;
  background: #6366f1;
}

ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

li {
  margin: 0.7rem 0;
  font-size: 1.08rem;
  color: #374151;
}

ol {
  padding-left: 1.5rem;
  margin: 0;
}

ol li {
  margin: 0.5rem 0;
  font-size: 1.08rem;
  color: #374151;
}

ol li::marker {
  color: #2563eb;
  font-weight: bold;
}
</style>
