<template>
  <main>
    <h1>About This Project</h1>
    <p style="text-align: justify;">
      <span class="highlight">Amazon E2E practice project</span> featuring a database, backend, and frontend.
      Includes a <span class="highlight">URLs panel</span> for accessing useful IT news and blogs,
      along with <span class="highlight">statistics tracking</span> for these URLs.
      This project is designed to help you practice end-to-end testing with a real-world application.
      Along with infrastructure as code (IaC) using <span class="highlight">Terraform</span> and <span class="highlight">Ansible</span> for deployment.
      The project is built with a focus on simplicity and ease of use, making it a great
      starting point for those looking to learn about end-to-end testing and web development.
    </p>
    <div class="about-divider"></div>
    <ul>
      <li>Database: <b>PostgreSQL</b></li>
      <li>Backend: <b>Golang</b></li>
      <li>Frontend: <b>Vue.js</b></li>
    </ul>
    <div class="about-divider"></div>
    <h2>Top URLs (number of visits)</h2>
    <ol>
      <li v-for="(url, index) in topUrls" :key="index">
        {{ url.title }} ({{ url.count }})
      </li>
    </ol>
  </main>
</template>

<script>
export default {
  name: 'AboutView',
  data() {
    return {
      topUrls: [],
    };
  },
  mounted() {
    fetch('/api/stats?limit=3')
      .then(response => response.json())
      .then(data => {
        data.forEach(element => {
          fetch(`/api/urls/${element.ID}`)
            .then(response => response.json())
            .then(urlData => {
              this.topUrls.push({
                title: element.ShortenedID,
                count: urlData.count,
              });
            });
        });

        this.topUrls.sort((a, b) => b.count - a.count);
      })
      .catch(error => {
        console.error('Error fetching URLs:', error);
      });
  }
};
</script>

<style scoped>
main {
  max-width: 70%;
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
