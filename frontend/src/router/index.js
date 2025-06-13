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
        requiresAuth: false,
      },
    },
  ],
});

// Navigation guard to check for authentication
router.beforeEach((to, _, next) => {
  const isAuthenticated = !!localStorage.getItem('token'); // check if token exists in localStorage

  if (to.matched.some(record => record.meta.requiresAuth) && !isAuthenticated) {
    alert('You must be logged in to access this page.'); // alert user if not authenticated
    next({ name: 'root' }); // redirect to login if not authenticated
  } else {
    next(); // proceed to the route
  }
});

export default router
