import api from "./requestUtil";

export async function getLogs(params) {
    const response = await api.get("/api/logs", {
      params: params,
      timeout: 30000,
    });
    return response.data;
}
