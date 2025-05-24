<script setup lang="ts">
import { ref } from 'vue';
import { Settings } from '../../types/settings';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { settingsService } from '../../services/settingsService';

const handleSave = async (settings: Settings) => {
    try {
        await settingsService.SetSettings(settings);
    } catch (error) {
        console.error('Error saving settings:', error);
    }
};

const settings = ref<Settings>({} as Settings);

</script>

<template>
    <div class="flex flex-col items-center justify-center p-4">
        <label class="w-full align-left" for="serviceaccountpath">Service account file path</label>
        <Input
            id="serviceaccountpath"
            type="text"
            v-model="settings.ServiceAccountPathFile"
            placeholder="./service-account.json"
            class="border border-gray-300 rounded p-2 mb-4 w-full"
        />
        <label class="w-full align-left" for="limitmessages">Limit messages</label>
        <Input
            id="limitmessages"
            type="number"
            v-model="settings.LimitMessages"
            placeholder="10"
            class="border border-gray-300 rounded p-2 mb-4 w-full"
        />
        <Button
            type="button"
            class="bg-blue-500 text-white hover:bg-blue-600 rounded p-2 w-full"
            @click="handleSave(settings)"
        >
            Guardar
        </Button>
    </div>
</template>