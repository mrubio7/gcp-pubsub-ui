<script setup lang="ts">
import { PubSubTopic } from '../../types/pubsub';

const props = defineProps<{
    topics: PubSubTopic[];
}>();
</script>

<template>
    {{ props.topics }}
</template>