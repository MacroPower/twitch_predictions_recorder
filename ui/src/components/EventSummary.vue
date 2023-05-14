<template>
  <div class="summary">
    <h1>Hello</h1>
    <p>
      <button @click="getSummary">Get Summary</button>
    </p>
    <p>
      For a guide and recipes on how to configure / customize this project,<br />
      check out the
      <a href="https://cli.vuejs.org" target="_blank" rel="noopener"
        >vue-cli documentation</a
      >.
    </p>
    <h3>Installed CLI Plugins</h3>
    <ul>
      <li v-for="summary in summaries" :key="summary.id">
        <p>{{ summary.title }}</p>
      </li>
    </ul>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, toRefs } from "vue";
import Summary from "@/models/Summary";

export default defineComponent({
  name: "EventSummary",
  props: {},
  setup() {
    const state = reactive({
      summaries: [] as Summary[],
    });

    return { ...toRefs(state) };
  },
  methods: {
    async getSummary() {
      const summaries = (await Summary.get()) as Summary[];
      this.summaries = summaries;
      console.log(summaries);
    },
  },
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
