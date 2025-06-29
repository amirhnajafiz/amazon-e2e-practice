<template>
  <main class="main-center">
    <div class="controls">
      <label>
        Columns:
        <input
          type="range"
          v-model.number="columns"
          min="1"
          max="5"
          step="1"
          class="slider"
        />
        <span class="slider-value">{{ columns }}</span>
      </label>
    </div>
    <div class="url-list" :style="gridStyle">
      <div
        v-for="(url, index) in urls"
        :key="index"
        class="url-item"
        :style="getBoxStyle(index)"
      >
        <h2>{{ url.title }}</h2>
        <p>{{ url.description }}</p>
        <a :href="url.address" target="_blank" @click="linkClicked(url.id)">{{ url.address }}</a>
      </div>
    </div>
  </main>
</template>

<script>
export default {
  name: 'HomeView',
  data() {
    return {
      urls: [],
      columns: 2, // Default columns (M)
      boxColors: [
        'rgba(255, 106, 0, 0.2)',
        'rgba(255, 179, 71, 0.2)',
        'rgba(79, 209, 197, 0.2)',
        'rgba(255, 140, 40, 0.2)',
        'rgba(255, 192, 120, 0.2)',
        'rgba(120, 220, 210, 0.2)',
        'rgba(255, 120, 30, 0.2)',
      ],
      borderColors: [
        '#ff6a00',
        '#ffb347',
        '#4fd1c5',
        '#ff8c28',
        '#ffc078',
        '#78dcd2',
        '#ff781e',
      ]
    };
  },
  computed: {
    gridStyle() {
      return {
        display: 'grid',
        gridTemplateColumns: `repeat(${this.columns}, 1fr)`,
        gap: '24px',
        alignItems: 'stretch',
        width: '100%',
        maxWidth: '90vw',
      };
    }
  },
  methods: {
    getBoxStyle(index) {
      return {
        background: this.boxColors[index % this.boxColors.length],
        border: `2px solid ${this.borderColors[index % this.borderColors.length]}`,
        borderRadius: '18px',
        padding: '24px',
        color: '#222',
        transition: 'box-shadow 0.2s',
        boxSizing: 'border-box',
        width: '100%',
        margin: 0,
      };
    },
    linkClicked(id) {
      fetch(`/api/urls`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ url_id: id })
      }).then(response => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
      }).catch(error => {
        console.error('Error updating URL visit count:', error);
      });
    }
  },
  mounted() {
    fetch('/api/urls')
      .then(response => response.json())
      .then(data => {
        data.forEach(element => {
          this.urls.push({
            id: element.ID,
            address: element.Address,
            title: element.ShortenedID,
            description: element.Description,
          });
        });
      })
      .catch(error => {
        console.error('Error fetching URLs:', error);
      });
  }
}
</script>

<style scoped>
.main-center {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
  margin-bottom: 70px;
}

.controls {
  margin-top: 50px;
  margin-bottom: 50px;
  display: flex;
  align-items: center;
}

.slider {
  margin: 0 12px;
  width: 120px;
  accent-color: #ff6a00;
  /* Remove default styles for better appearance */
  -webkit-appearance: none;
  appearance: none;
  height: 6px;
  border-radius: 3px;
  background: #eee;
  outline: none;
}

.slider::-webkit-slider-thumb {
  -webkit-appearance: none;
  appearance: none;
  width: 22px;
  height: 22px;
  border-radius: 50%;
  background: #fff;
  border: 2px solid #ff6a00;
  cursor: pointer;
  box-shadow: 0 2px 6px rgba(0,0,0,0.15);
  transition: background 0.2s;
}

.slider::-moz-range-thumb {
  width: 22px;
  height: 22px;
  border-radius: 50%;
  background: #fff;
  border: 2px solid #ff6a00;
  cursor: pointer;
  box-shadow: 0 2px 6px rgba(0,0,0,0.15);
  transition: background 0.2s;
}

.slider::-ms-thumb {
  width: 22px;
  height: 22px;
  border-radius: 50%;
  background: #fff;
  border: 2px solid #ff6a00;
  cursor: pointer;
  box-shadow: 0 2px 6px rgba(0,0,0,0.15);
  transition: background 0.2s;
}

.slider-value {
  font-weight: bold;
  min-width: 24px;
  display: inline-block;
  text-align: center;
}

.url-list {
  /* grid styles are set dynamically via :style binding */
}

.url-item {
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
  transition: box-shadow 0.2s ease-in-out;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.url-item:hover {
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.2);
}
</style>
