<script setup lang="ts">
import { Button } from '@/components/ui/button';
import { PathRoutes } from '@/router';
import { Settings } from 'lucide-vue-next';
import { Toaster } from 'vue-sonner';

</script>

<template>
    <Toaster />
    <header class="flex justify-between gap-1 bg-[#00a0df] text-white p-2">
        <div class="flex items-center gap-2 w-1/2 px-3">
            <h1 class="text-md font-bold">PubSubUI</h1>
        </div>
        <div class="flex items-center gap-2 w-1/2 justify-end px-6">
            <router-link :to="PathRoutes.Home">
                <Button class="bg-[#00a0df] text-white hover:bg-[#008bbd] cursor-pointer shadow-none">
                    <span>Home</span>
                </Button> 
            </router-link>
            <router-link :to="PathRoutes.Config">
                <Button class="bg-[#00a0df] text-white hover:bg-[#008bbd] cursor-pointer shadow-none">
                    <Settings class="w-5 h-5" />
                </Button> 
            </router-link>
        </div>
    </header>
    <main>
        <RouterView />
    </main>
</template>