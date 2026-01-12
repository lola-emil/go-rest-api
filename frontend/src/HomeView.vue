<script setup lang="ts">
  import { RouterLink } from 'vue-router';
</script>

<template>
  <div class="w-full h-screen flex flex-col justify-center items-center">
    <div class="flex items-center gap-5">
      <RouterLink to="/users-list" class="btn btn-primary btn-outline">Users List</RouterLink>
      <RouterLink to="/contacts-list" class="btn btn-secondary btn-outline">Contacts List</RouterLink>
    </div>
  </div>  
</template>