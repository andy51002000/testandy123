import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'home',
    component: () => import('../components/HelloWorld.vue')
  },
  {
    path: '/test',
    name: 'test',
    component: () => import('../components/TestView.vue')
  },
  {
    path: '/transactions',
    name: 'transaction',
    component: () => import('../components/TransactionView.vue'), 
  },
  {
    path: '/member-update',
    name: 'member-update',
    props: (route) => ({
      memberId: route.params.memberId,
    }),
    component: () => import('../components/MemberUpdate.vue'), 
  },
  {
    path: '/member-get',
    name: 'member-get',
    props: (route) => ({
      memberId: route.params.memberId,
    }),
    component: () => import('../components/MemberGet.vue'), 
  },
  {
    path: '/member-create',
    name: 'member-create',
    component: () => import('../components/MemberCreate.vue'), 
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router