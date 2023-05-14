import { Model as BaseModel, HTTPRequestConfig } from "vue-api-query";

export default class Model extends BaseModel {
  // Define a base url for a REST API
  baseURL() {
    return window.location.origin;
  }

  // Implement a default request method
  request(config: HTTPRequestConfig) {
    return this.$http.request(config);
  }
}
