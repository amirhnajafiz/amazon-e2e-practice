<template>
  <div class="auth-container">
    <h1 style="text-align: center; margin-bottom: 40px;">Authenticator</h1>
    <div class="form-toggle">
      <button :class="{ active: isLogin }" @click="isLogin = true">Login</button>
      <button :class="{ active: !isLogin }" @click="isLogin = false">Register</button>
    </div>
    <form @submit.prevent="handleSubmit">
      <div class="form-group">
        <label for="username">Username</label>
        <input
          id="username"
          v-model="form.username"
          type="text"
          :placeholder="isLogin ? 'Enter your username ...' : 'Choose a username ...'"
          required
          autocomplete="username"
        />
      </div>
      <div class="form-group">
        <label for="password">Password</label>
        <input
          id="password"
          v-model="form.password"
          :placeholder="isLogin ? 'Enter your password ...' : 'Create a password ...'"
          type="password"
          required
          autocomplete="current-password"
        />
      </div>
      <button type="submit">{{ isLogin ? 'Login' : 'Register' }}</button>
    </form>
    <div v-if="message" class="message">{{ message }}</div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const isLogin = ref(true)
const form = ref({
  username: '',
  password: ''
})
const message = ref('')

function handleSubmit() {
  if (isLogin.value) {
    // HTTP POST request to login endpoint
    fetch('/api', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        username: form.value.username,
        password: form.value.password
      })
    }).then(async response => {
      if (response.ok) {
        return response.json()
      } else {
        const err = await response.json();
        throw new Error(err.error || 'Login failed');
      }
    }).then(data => {
      let token = data.token;
      if (token) {
        localStorage.setItem('token', token);
        // redirect to home page or another view
        window.location.href = '/home';
      } else {
        throw new Error('No token received');
      }
    }).catch(error => {
      alert(error.message);
    })
  } else {
    // HTTP Put request to register endpoint
    fetch('/api', {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        username: form.value.username,
        password: form.value.password
      })
    }).then(async response => {
      if (response.ok) {
        return response.json()
      } else {
        const err = await response.json();
        throw new Error(err.error || 'Registration failed');
      }
    }).then(_ => {
      alert('Registration successful! You can now login.');
    }).catch(error => {
      alert(error.message);
    })
  }
}
</script>

<style scoped>
.auth-container {
  width: 320px;
  margin: 200px auto;
  padding: 2rem;
  border-radius: 8px;
  background: #fff;
  box-shadow: 0 2px 12px rgba(0,0,0,0.08);
}
.form-toggle {
  display: flex;
  justify-content: space-between;
  margin-bottom: 1.5rem;
}
.form-toggle button {
  flex: 1;
  padding: 0.5rem 0;
  border: none;
  background: #f0f0f0;
  cursor: pointer;
  font-weight: bold;
  transition: background 0.2s;
}
.form-toggle .active {
  background: #e53935;
  color: #fff;
}
.form-group {
  margin-bottom: 1rem;
}
label {
  display: block;
  margin-bottom: 0.25rem;
  font-weight: 500;
}
input {
  width: 100%;
  padding: 0.5rem;
  border: 1px solid #ddd;
  border-radius: 4px;
}
button[type="submit"] {
  width: 100%;
  padding: 0.75rem;
  background: #e53935;
  color: #fff;
  border: none;
  border-radius: 4px;
  font-size: 1rem;
  font-weight: bold;
  cursor: pointer;
  margin-top: 1rem;
  transition: background 0.2s;
}
button[type="submit"]:hover {
  background: #b71c1c;
}
.message {
  margin-top: 1rem;
  color: #388e3c;
  font-weight: 500;
  text-align: center;
}
</style>
