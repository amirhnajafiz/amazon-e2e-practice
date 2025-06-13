<template>
  <main class="home-view">
    <h1>
      {{ welcomeMessage }}
    </h1>
    <p style="margin: 50px;">
      This is a simple D3.js chart example using Vue 3. Click the button below to randomize the data.
    </p>
    <div ref="d3chart" class="d3chart"></div>
    <button class="btn" @click="randomizeData">Randomize Data</button>
  </main>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import * as d3 from 'd3'

const d3chart = ref(null)
let svg = null
let data = ref([10, 15, 20, 25, 30, 35, 40]) // initial data for the chart
let tooltip = null

const width = 400
const height = 200
const margin = { top: 10, right: 20, bottom: 40, left: 40 } // Add margin for axes

// --- JWT parsing and welcome message ---
const welcomeMessage = ref('Welcome')

function parseJwt(token) {
  try {
    const base64Url = token.split('.')[1]
    const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/')
    const jsonPayload = decodeURIComponent(
      atob(base64)
        .split('')
        .map(c => '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2))
        .join('')
    )
    return JSON.parse(jsonPayload)
  } catch (e) {
    return null
  }
}

onMounted(() => {
  // Read token from localStorage and update welcome message
  const token = localStorage.getItem('token')
  if (token) {
    const payload = parseJwt(token)
    if (payload && payload.username) {
      welcomeMessage.value = `Welcome, ${payload.username}!`
    }
  }
  drawChart()
})
// --- end JWT parsing and welcome message ---

function drawChart() {
  if (!d3chart.value) return

  // Remove old svg and tooltip if exists
  d3.select(d3chart.value).selectAll('*').remove()

  svg = d3
    .select(d3chart.value)
    .append('svg')
    .attr('width', width + margin.left + margin.right)
    .attr('height', height + margin.top + margin.bottom)
    .append('g')
    .attr('transform', `translate(${margin.left},${margin.top})`)

  // Add light gray background inside axes
  svg.append('rect')
    .attr('x', 0)
    .attr('y', 0)
    .attr('width', width)
    .attr('height', height)
    .attr('fill', '#f5f5f5') // light gray

  // X scale
  const x = d3
    .scaleBand()
    .domain(data.value.map((d, i) => i))
    .range([0, width])
    .padding(0.2)

  // Y scale
  const y = d3
    .scaleLinear()
    .domain([0, d3.max(data.value) || 0])
    .range([height, 0]).nice()

  // X axis
  svg
    .append('g')
    .attr('transform', `translate(0, ${height})`)
    .call(d3.axisBottom(x).tickFormat(i => `Day ${i + 1}`))
    .selectAll('text')
    .attr('transform', 'rotate(-30)')
    .style('text-anchor', 'end')

  // Y axis
  svg
    .append('g')
    .call(d3.axisLeft(y).ticks(5).tickFormat(d => `${d}`))

  // Tooltip (create only once)
  tooltip = d3
    .select(d3chart.value)
    .append('div')
    .attr('class', 'd3-tooltip')
    .style('position', 'fixed')
    .style('background', '#fff')
    .style('border', '1px solid #ccc')
    .style('padding', '4px 8px')
    .style('border-radius', '4px')
    .style('pointer-events', 'none')
    .style('opacity', 0)
    .style('z-index', 10)

  // Bars
  svg
    .selectAll('rect.bar')
    .data(data.value)
    .enter()
    .append('rect')
    .attr('class', 'bar')
    .attr('x', (d, i) => x(i))
    .attr('y', d => y(d))
    .attr('width', x.bandwidth())
    .attr('height', d => height - y(d))
    .attr('fill', '#e53935')
    .on('mousemove', function (event, d) {
      d3.select(this).attr('fill', '#ff9800')
      tooltip
        .style('opacity', 1)
        .html(`Value: ${d}`)
        .style('left', event.clientX + 15 + 'px')
        .style('top', event.clientY - 20 + 'px')
    })
    .on('mouseleave', function () {
      d3.select(this).attr('fill', '#e53935')
      tooltip.style('opacity', 0)
    })
}

function randomizeData() {
  data.value = Array.from({ length: 7 }, () => Math.floor(Math.random() * 100) + 5)
  drawChart()
}
</script>

<style scoped>
.home-view {
  padding: 2rem;
  text-align: center;
}

.d3chart {
  position: relative;
  width: 400px;
  height: 200px;
  margin: 0 auto;
}

.d3-tooltip {
  min-width: 60px;
  text-align: left;
  pointer-events: none;
}

.btn {
  margin-top: 100px;
  padding: 0.5rem 1rem;
  background-color: #ff3f3b;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.btn:hover {
  background-color: #e53935;
}
</style>
