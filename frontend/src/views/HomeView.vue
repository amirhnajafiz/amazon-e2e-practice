<template>
  <main class="main-center">
    <div class="url-list">
      <div
        v-for="(url, index) in urls"
        :key="index"
        class="url-item"
        :style="getBoxStyle(index)"
      >
        <h2>{{ url.title }}</h2>
        <p>{{ url.description }}</p>
        <a :href="url.address" target="_blank">{{ url.address }}</a>
      </div>
    </div>
  </main>
</template>

<script>
export default {
  name: 'HomeView',
  data() {
    return {
      urls: [
        {
          address: 'https://example.com',
          title: 'Example',
          description: 'This is an example URL.',
        },
        {
          address: 'https://vuejs.org',
          title: 'Vue.js',
          description: 'The Progressive JavaScript Framework.',
        },
        {
          address: 'https://github.com',
          title: 'GitHub',
          description: 'Where the world builds software.',
        },
      ],
      boxColors: [
        'rgba(255, 106, 0, 0.2)',    // #ff6a00
        'rgba(255, 179, 71, 0.2)',   // #ffb347
        'rgba(79, 209, 197, 0.2)',   // #4fd1c5
        'rgba(255, 140, 40, 0.2)',   // blend between #ff6a00 and #ffb347
        'rgba(255, 192, 120, 0.2)',  // lighter #ffb347
        'rgba(120, 220, 210, 0.2)',  // lighter #4fd1c5
        'rgba(255, 120, 30, 0.2)',   // another orange in range
      ],
      borderColors: [
        '#ff6a00',                   // #ff6a00
        '#ffb347',                   // #ffb347
        '#4fd1c5',                   // #4fd1c5
        '#ff8c28',                   // blend between #ff6a00 and #ffb347
        '#ffc078',                   // lighter #ffb347
        '#78dcd2',                   // lighter #4fd1c5
        '#ff781e',                   // another orange in range
      ]
    };
  },
  methods: {
    getBoxStyle(index) {
      return {
        background: this.boxColors[index % this.boxColors.length],
        border: `2px solid ${this.borderColors[index % this.borderColors.length]}`,
        borderRadius: '18px',
        margin: '16px 0',
        padding: '24px',
        color: '#222',
        transition: 'box-shadow 0.2s',
      };
    }
  }
}
</script>

<style scoped>
.main-center {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
}

.url-list {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 500px;
  max-width: 90vw;
}

.url-item {
  width: 100%;
  box-sizing: border-box;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06); /* match the original inline style */
  transition: box-shadow 0.2s ease-in-out;
}

.url-item:hover {
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.2);
}
</style>
