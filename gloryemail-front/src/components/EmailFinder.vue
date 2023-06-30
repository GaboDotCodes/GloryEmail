<script setup>
  import { ref } from 'vue';

  import SearchBar from './SearchBar.vue';
  import EmailsViewer from './EmailsViewer.vue';

  const queryRef = ref("");

  const handleQuery = (query) => {
    queryRef.value = query;
  }
</script>
<template>
  <div>
    <SearchBar @change-query="handleQuery"></SearchBar>
    <EmailsViewer :query="queryRef"></EmailsViewer>
  </div>
</template>