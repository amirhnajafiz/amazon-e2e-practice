<template>
    <div class="login-container">
        <form @submit.prevent="handleLogin" class="login-form">
            <h2>Login</h2>
            <div class="form-group">
                <label for="user">User</label>
                <input
                    id="user"
                    v-model="user"
                    type="text"
                    required
                    autocomplete="admin"
                />
            </div>
            <div class="form-group">
                <label for="password">Password</label>
                <input
                    id="password"
                    v-model="password"
                    type="password"
                    required
                    autocomplete="current-password"
                />
            </div>
            <button type="submit" :disabled="loading">
                {{ loading ? 'Logging in...' : 'Login' }}
            </button>
            <p v-if="error" class="error">{{ error }}</p>
        </form>
    </div>
</template>

<script setup>
import { ref } from 'vue'

const user = ref('')
const password = ref('')
const loading = ref(false)
const error = ref('')

const handleLogin = async () => {
    loading.value = true
    error.value = ''
    try {
        // Replace with your login API call
        await fakeLogin(user.value, password.value)
        // Redirect or emit success event here
    } catch (e) {
        error.value = 'Invalid username or password.'
    } finally {
        loading.value = false
    }
}

// Mock login function, replace with real API call
function fakeLogin(user, password) {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            if (user === 'test@example.com' && password === 'password') {
                resolve()
            } else {
                reject()
            }
        }, 1000)
    })
}
</script>

<style scoped>
.login-container {
    max-width: 400px;
    margin: 60px auto;
    padding: 2rem;
    border: 1px solid #eee;
    border-radius: 8px;
    background: #fff;
    box-shadow: 0 2px 8px rgba(0,0,0,0.05);
}
.login-form h2 {
    margin-bottom: 1.5rem;
    text-align: center;
}
.form-group {
    margin-bottom: 1rem;
}
.form-group label {
    display: block;
    margin-bottom: 0.5rem;
}
.form-group input {
    width: 100%;
    padding: 0.5rem;
    border: 1px solid #ccc;
    border-radius: 4px;
}
button[type="submit"] {
    width: 100%;
    padding: 0.75rem;
    background: #232f3e;
    color: #fff;
    border: none;
    border-radius: 4px;
    font-weight: bold;
    cursor: pointer;
}
button[disabled] {
    opacity: 0.6;
    cursor: not-allowed;
}
.error {
    color: #d32f2f;
    margin-top: 1rem;
    text-align: center;
}
</style>
