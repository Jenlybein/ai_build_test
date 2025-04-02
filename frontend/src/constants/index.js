// API 相关常量
export const API_BASE_URL =
  import.meta.env.VITE_API_BASE_URL || "http://localhost:8080";

// 本地存储键名
export const STORAGE_KEYS = {
  SELECTED_FIELDS: "selected_fields",
};

// 默认选中的字段
export const DEFAULT_SELECTED_FIELDS = ["data_time", "app_id", "user_id", "os"];

// 平台选项
export const PLATFORM_OPTIONS = [
  { text: "全部平台", value: "all" },
  { text: "Web", value: "web" },
  { text: "iOS", value: "ios" },
  { text: "Android", value: "android" },
];

export const ALL_FIELDS = [
  { key: "data_time", label: "data_time" },
  { key: "write_time", label: "write_time" },
  { key: "time_hour", label: "time_hour" },
  { key: "id", label: "id" },
  { key: "time", label: "time" },
  { key: "extra", label: "extra" },
  { key: "entrance_time", label: "entrance_time" },
  { key: "entrance_id", label: "entrance_id" },
  { key: "stamp", label: "stamp" },
  { key: "app_id", label: "app_id" },
  { key: "platform", label: "platform" },
  { key: "user_id", label: "user_id" },
  { key: "version", label: "version" },
  { key: "build_id", label: "build_id" },
  { key: "device_id", label: "device_id" },
  { key: "model", label: "model" },
  { key: "os", label: "os" },
  { key: "os_ver", label: "os_ver" },
  { key: "sdk_ver", label: "sdk_ver" },
  { key: "category", label: "category" },
  { key: "action", label: "action" },
  { key: "label", label: "label" },
  { key: "state", label: "state" },
  { key: "value", label: "value" },
  // D1-D40
  ...Array.from({ length: 40 }, (_, i) => {
    const key = `d${i + 1}`;
    return { key, label: key };
  }),
  // V1-V40
  ...Array.from({ length: 40 }, (_, i) => {
    const key = `v${i + 1}`;
    return { key, label: key };
  }),
  // Info1-Info10
  ...Array.from({ length: 10 }, (_, i) => {
    const key = `info${i + 1}`;
    return { key, label: key };
  }),
  // UD1-UD20
  ...Array.from({ length: 20 }, (_, i) => {
    const key = `ud${i + 1}`;
    return { key, label: key };
  }),
  // UV1-UV10
  ...Array.from({ length: 10 }, (_, i) => {
    const key = `uv${i + 1}`;
    return { key, label: key };
  }),
  // SD1-SD20
  ...Array.from({ length: 20 }, (_, i) => {
    const key = `sd${i + 1}`;
    return { key, label: key };
  }),
  // SV1-SV10
  ...Array.from({ length: 10 }, (_, i) => {
    const key = `sv${i + 1}`;
    return { key, label: key };
  }),
];
