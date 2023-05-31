import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/', name: 'home',
      component: () => import(/* webpackChunkName: "home" */ '../views/HomeView.vue')
    },
    {
      path: '/about', name: 'about',
      component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
    },
	{
		path: '/login', name: 'login',
		component: () => import(/* webpackChunkName: "login" */ '../views/LoginView.vue'),
	},
	{
		path: '/logout', name: 'logout',
		component: () => import(/* webpackChunkName: "logout" */ '../views/LogoutView.vue'),
	},
  ]
})

router.beforeEach((to, from, next) => {
	// if (to.name == 'logout') {
	// 	cookie.clear('token');
	// 	next();
	// } else {
	// 	const token = cookie.get('token');
	// 	if (!token && to.name != 'login') {
	// 		next({ name: 'login' });
	// 	} else if (token && to.name == 'login') {
	// 		next({ name: 'home' });
	// 	} else {
	// 		next();
	// 	}
	// }
    next();
});

export default router
