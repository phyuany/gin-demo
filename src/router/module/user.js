const userRoutes = [
  {
    path: '/register',
    name: 'register',
    component: () => import('@/views/register/Register.vue'),
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('@/views/login/Login.vue'),
  },
  {
    path: '/profile',
    name: 'profile',
    meta: {
      auth: true,
    },
    component: () => import('@/views/profile/Profile.vue'),
  },
];

export default userRoutes;
