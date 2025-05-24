<script setup lang="ts">
import { pubsubService } from '@/entities/pubsub/services/pubsubService';
import { onMounted } from 'vue';


onMounted( async () => {
    console.log('IndexPage mounted');
    const res = await pubsubService.AddTopic("test")

    console.log('AddTopic response:', res);
})
</script>

<template>
    INDEX
</template>