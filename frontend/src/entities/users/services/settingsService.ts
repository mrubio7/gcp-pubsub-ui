import { useSettingsStore } from "../store/settings.store";
import { Settings } from "../types/settings";

export const settingsService = {
    SetSettings(stgs: Settings) {
        const settingsStore = useSettingsStore();
        settingsStore.Set(stgs);
    }
}