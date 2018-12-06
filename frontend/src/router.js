import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import NewBlog from '../src/components/NewBlog.vue'
import Blog from "./views/Blog.vue"

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/blog/:id',
      name: 'blog',
      component: Blog
    },
    {
      path: '/new',
      name: 'new',
      component: NewBlog
    }
  ]
})
