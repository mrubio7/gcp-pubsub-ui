import { AddTopic } from "../../../../wailsjs/go/main/App";

export const pubsubService = {
    AddTopic: async (topic: string) => {
        const isAdded = await AddTopic(topic);
        console.log("pubsubStore", isAdded);
    }
}