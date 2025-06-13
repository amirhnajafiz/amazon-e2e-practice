import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'root',
      component: () => import('../views/LoginView.vue'),
    },
    {
      path: '/home',
      name: 'home',
      component: () => import('../views/HomeView.vue'),
      meta: {
        requiresAuth: true,
      },
    },
  ],
});

// Navigation guard to check for authentication
router.beforeEach((to, _, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    const token = localStorage.getItem('token');
    const isAuthenticated = !!token;

    if (token) {
      try {
        const tokenPayload = JSON.parse(atob(token.split('.')[1]));
        const expiryTime = tokenPayload.exp * 1000; // convert to milliseconds
        if (Date.now() > expiryTime) {
          localStorage.removeItem('token');
          alert('Your session has expired. Please log in again.');
          return next({ name: 'root' });
        }
      } catch (e) {
        localStorage.removeItem('token');
        alert('Invalid session. Please log in again.');
        return next({ name: 'root' });
      }
    }

    if (!isAuthenticated) {
      alert('You must be logged in to access this page.');
      return next({ name: 'root' });
    }
  }
  next();
});

export default router
