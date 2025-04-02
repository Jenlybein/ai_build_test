import axios from "axios";
import Message from "../components/Message";
import { API_BASE_URL } from "../constants";

// 创建axios实例
const api = axios.create({
  baseURL: API_BASE_URL,
  timeout: 15000,
  headers: {
    "Content-Type": "application/json",
    Accept: "application/json",
  },
});

// 统一的错误处理函数
const handleError = (error) => {
  let errorMessage = "获取数据失败，请稍后重试";

  if (error.response?.data?.error) {
    errorMessage = error.response.data.error;
  } else if (error.message) {
    errorMessage = error.message;
  }

  // 根据错误类型设置不同的提示
  if (error.status === 400) {
    errorMessage = "请求参数错误：" + errorMessage;
  } else if (error.status === 404) {
    errorMessage = "未找到请求的资源";
  } else if (error.status === 500) {
    errorMessage = "服务器内部错误：" + errorMessage;
  } else if (error.message.includes("timeout")) {
    errorMessage = "请求超时，请稍后重试";
  } else if (error.message.includes("Network Error")) {
    errorMessage = "网络连接错误，请检查网络设置";
  }

  return {
    message: errorMessage,
    type: "error",
    data: null,
    total: 0,
  };
};

// 请求拦截器
api.interceptors.request.use(
  (config) => {
    // 验证和处理参数，移除空值参数
    const validParams = { ...config.params };
    Object.keys(validParams).forEach((key) => {
      if (!validParams[key] === "") {
        delete validParams[key];
      }
    });
    config.params = validParams;

    console.log("处理后发送请求:", {
      url: config.url,
      method: config.method,
      params: config.params,
      headers: config.headers,
    });
    return config;
  },
  (error) => {
    console.error("请求配置错误:", error);
    return Promise.reject(error);
  }
);

// 响应拦截器
api.interceptors.response.use(
  (response) => {
    console.log("收到响应:", {
      url: response.config.url,
      status: response.status,
      data: response.data,
    });

    // 检查响应格式
    const data = response.data;
    if (!data) {
      return Promise.reject({
        message: "服务器返回数据为空",
        status: 500,
      });
    }

    // 处理成功响应
    return response;
  },
  (error) => {
    console.error("API请求错误:", {
      url: error.config?.url,
      method: error.config?.method,
      params: error.config?.params,
      status: error.response?.status,
      statusText: error.response?.statusText,
      data: error.response?.data,
      message: error.message,
    });

    return Promise.reject(handleError(error));
  }
);

export default api;
