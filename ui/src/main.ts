import { createApp } from "vue";
import App from "./App.vue";
import axios from "axios";
import router from "./router";
import { Model } from "vue-api-query";
import naive from "naive-ui";
import "vfonts/Inter.css";
import "vfonts/FiraCode.css";

Model.$http = axios;

const app = createApp(App);
app.use(router);
app.use(naive);
app.mount("#app");
