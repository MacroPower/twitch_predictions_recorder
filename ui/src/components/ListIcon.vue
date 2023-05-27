<template>
  <n-tag :bordered="false" :color="{ color: 'rgba(0,0,0,0)' }">
    {{ text }}
    <template #icon>
      <n-icon :component="icon" />
    </template>
  </n-tag>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "ListIcon",
  props: {
    icon: Object,
    text: String,
  },
});
</script>
