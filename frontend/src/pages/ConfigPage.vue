<script setup lang="ts">
import SettingsForm from '@/entities/users/components/forms/SettingsForm.vue';
import { useSettingsStore } from '@/entities/users/store/settings.store';
import { Settings } from '@/entities/users/types/settings';
import { ref } from 'vue';

const settingsStore = useSettingsStore();

const settingsState = ref<Settings>(settingsStore.$state);
</script>

<template>
    <div class="flex w-full gap-4">
        <div class="w-1/2 p-4">
            <SettingsForm />
        </div>
        <div class="w-1/2 p-4">
            {{ settingsState }}
        </div>
    </div>
    
</template>