import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import { createPinia } from 'pinia'
import router from './router';
import { createStatePersistence } from 'pinia-plugin-state-persistence'

const pinia = createPinia();
pinia.use(createStatePersistence())
const app = createApp(App);

app.use(router);
app.use(pinia);

app.mount('#app');