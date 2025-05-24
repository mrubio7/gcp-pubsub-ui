import { createMemoryHistory, createRouter } from 'vue-router'
import IndexPage from './pages/IndexPage.vue'
import Layout from './pages/layout/Layout.vue'
import ConfigPage from './pages/ConfigPage.vue'
import NewTaskPage from './pages/NewTaskPage.vue'


export const PathRoutes = {
  Home: "/",
  Config: "/config",
}

const routes = [
  { path: PathRoutes.Home, component: Layout, children: [
    { path: '', component: IndexPage },
    { path: PathRoutes.Config, component: ConfigPage },
  ] },
]
  
const router = createRouter({
    history: createMemoryHistory(),
    routes,
})

export default router
