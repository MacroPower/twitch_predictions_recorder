import { createApp } from "vue";
import App from "./App.vue";
import axios from "axios";
import router from "./router";
import { Model } from "vue-api-query";

Model.$http = axios;

createApp(App).use(router).mount("#app");
