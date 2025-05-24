import { defineStore } from "pinia";
import { Settings } from "../types/settings";

export const useSettingsStore = defineStore('settings', {
    persist: true,
    state: (): Settings => ({
        ServiceAccountPathFile: "",
        LimitMessages: ""
    }),

    actions: {
        Set(user: Settings) {
            this.ServiceAccountPathFile = user.ServiceAccountPathFile;
            this.LimitMessages = user.LimitMessages;
        },
    }
});

