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

    <!-- Range input for limit -->
    <h2 style="margin-bottom: 20px;">Top URLs (number of visits)</h2>
    <div class="slider-container">
      <label for="limitRange">Show top: <b>{{ limit }}</b> URLs</label>
      <input
        id="limitRange"
        type="range"
        min="1"
        max="9"
        v-model="limit"
        @input="fetchTopUrls"
        class="custom-range"
      />
    </div>
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
      limit: 3,
    };
  },
  mounted() {
    this.fetchTopUrls();
  },
  methods: {
    fetchTopUrls() {
      fetch(`/api/stats?limit=${this.limit}`)
        .then(response => response.json())
        .then(data => {
          // Fetch all URLs in parallel and wait for all to finish
          return Promise.all(
            data.map(element =>
              fetch(`/api/urls/${element.ID}`)
                .then(response => response.json())
                .then(urlData => ({
                  title: element.ShortenedID,
                  count: urlData.count,
                }))
            )
          );
        })
        .then(urls => {
          // Sort and set once to avoid duplicates
          this.topUrls = urls.sort((a, b) => b.count - a.count);
        })
        .catch(error => {
          console.error('Error fetching URLs:', error);
        });
    }
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

.slider-container {
  margin-bottom: 1.2rem;
  display: flex;
  align-items: center;
  gap: 1rem;
}

.custom-range {
  -webkit-appearance: none;
  width: 160px;
  height: 24px;              /* Match thumb height for alignment */
  background: transparent;   /* Transparent so only track shows */
  margin: 0;
  padding: 0;
  vertical-align: middle;
  display: flex;
  align-items: center;
}

/* Track styles */
.custom-range::-webkit-slider-runnable-track {
  height: 8px;
  margin-top: 8px;           /* Center track in input */
  border-radius: 6px;
  background: #e0e7ff;
}
.custom-range::-moz-range-track {
  height: 8px;
  border-radius: 6px;
  background: #e0e7ff;
}
.custom-range::-ms-fill-lower,
.custom-range::-ms-fill-upper {
  height: 8px;
  border-radius: 6px;
  background: #e0e7ff;
}

/* Thumb styles */
.custom-range::-webkit-slider-thumb {
  -webkit-appearance: none;
  appearance: none;
  width: 24px;
  height: 24px;
  margin-top: -8px; /* Half of (thumb height - track height) to center */
  border-radius: 50%;
  background: #6366f1;
  cursor: pointer;
  box-shadow: 0 2px 6px #0002;
  border: 2px solid #fff;
  transition: background 0.3s;
}
.custom-range::-moz-range-thumb {
  width: 24px;
  height: 24px;
  border-radius: 50%;
  background: #6366f1;
  cursor: pointer;
  box-shadow: 0 2px 6px #0002;
  border: 2px solid #fff;
  transition: background 0.3s;
}
.custom-range::-ms-thumb {
  width: 24px;
  height: 24px;
  border-radius: 50%;
  background: #6366f1;
  cursor: pointer;
  box-shadow: 0 2px 6px #0002;
  border: 2px solid #fff;
  transition: background 0.3s;
}

/* Remove default outline and background for all browsers */
.custom-range:focus {
  outline: none;
  box-shadow: none;
  background: transparent;
}

/* Hide outline for Firefox */
.custom-range::-moz-focus-outer {
  border: 0;
}
</style>
