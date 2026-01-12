import ContactListView from '@/ContactListView.vue'
import HomeView from '@/HomeView.vue'
import UserListView from '@/UserListView.vue'
import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      component: HomeView
    },
    {
      path: "/users-list",
      component: UserListView
    },
    {
      path: "/contacts-list",
      component: ContactListView
    }
  ],
})

export default router
